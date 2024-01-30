package fgax

import (
	"context"
	"regexp"
	"strings"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
)

const (
	// setup relations for use in creating tuples
	MemberRelation = "member"
	AdminRelation  = "admin"
	OwnerRelation  = "owner"
	ParentRelation = "parent"
	CanView        = "can_view"
	CanEdit        = "can_edit"
	CanDelete      = "can_delete"
)

type TupleKey struct {
	Subject  Entity
	Object   Entity
	Relation Relation `json:"relation"`
}

func NewTupleKey() TupleKey { return TupleKey{} }

// entityRegex is used to validate that a string represents an Entity/EntitySet
// and helps to convert from a string representation into an Entity struct.
var entityRegex = regexp.MustCompile(`([A-za-z0-9_][A-za-z0-9_-]*):([A-za-z0-9_][A-za-z0-9_@.+-]*)(#([A-za-z0-9_][A-za-z0-9_-]*))?`)

// Kind represents the type of the entity in OpenFGA.
type Kind string

// String implements the Stringer interface.
func (k Kind) String() string {
	return strings.ToLower(string(k))
}

// Relation represents the type of relation between entities in OpenFGA.
type Relation string

// String implements the Stringer interface.
func (r Relation) String() string {
	return strings.ToLower(string(r))
}

// Entity represents an entity/entity-set in OpenFGA.
// Example: `user:<user-id>`, `org:<org-id>#member`
type Entity struct {
	Kind       Kind
	Identifier string
	Relation   Relation
}

// String returns a string representation of the entity/entity-set.
func (e *Entity) String() string {
	if e.Relation == "" {
		return e.Kind.String() + ":" + e.Identifier
	}

	return e.Kind.String() + ":" + e.Identifier + "#" + e.Relation.String()
}

// ParseEntity will parse a string representation into an Entity. It expects to
// find entities of the form:
//   - <entityType>:<Identifier>
//     eg. organization:datum
//   - <entityType>:<Identifier>#<relationship-set>
//     eg. organization:datum#member
func ParseEntity(s string) (Entity, error) {
	// entities should only contain a single colon
	c := strings.Count(s, ":")
	if c != 1 {
		return Entity{}, newInvalidEntityError(s)
	}

	match := entityRegex.FindStringSubmatch(s)
	if match == nil {
		return Entity{}, newInvalidEntityError(s)
	}

	// Extract and return the relevant information from the sub-matches.
	return Entity{
		Kind:       Kind(match[1]),
		Identifier: match[2],
		Relation:   Relation(match[4]),
	}, nil
}

// tupleKeyToWriteRequest converts a TupleKey to a ClientTupleKey to send to FGA
func tupleKeyToWriteRequest(writes []TupleKey) (w []ofgaclient.ClientTupleKey) {
	for _, k := range writes {
		ctk := ofgaclient.ClientTupleKey{}
		ctk.SetObject(k.Object.String())
		ctk.SetUser(k.Subject.String())
		ctk.SetRelation(k.Relation.String())

		w = append(w, ctk)
	}

	return
}

// tupleKeyToDeleteRequest converts a TupleKey to a TupleKeyWithoutCondition to send to FGA
func tupleKeyToDeleteRequest(deletes []TupleKey) (d []openfga.TupleKeyWithoutCondition) {
	for _, k := range deletes {
		ctk := openfga.TupleKeyWithoutCondition{}
		ctk.SetObject(k.Object.String())
		ctk.SetUser(k.Subject.String())
		ctk.SetRelation(string(k.Relation))

		d = append(d, ctk)
	}

	return
}

// WriteTupleKeys takes a tuples keys, converts them to a client write request, which can contain up to 10 writes and deletes,
// and executes in a single transaction
func (c *Client) WriteTupleKeys(ctx context.Context, writes []TupleKey, deletes []TupleKey) (*ofgaclient.ClientWriteResponse, error) {
	opts := ofgaclient.ClientWriteOptions{AuthorizationModelId: openfga.PtrString(c.Config.AuthorizationModelId)}

	body := ofgaclient.ClientWriteRequest{
		Writes:  tupleKeyToWriteRequest(writes),
		Deletes: tupleKeyToDeleteRequest(deletes),
	}

	resp, err := c.Ofga.Write(ctx).Body(body).Options(opts).Execute()
	if err != nil {
		c.Logger.Infow("error writing relationship tuples", "error", err.Error(), "user", resp.Writes)

		return resp, err
	}

	for _, writes := range resp.Writes {
		if writes.Error != nil {
			c.Logger.Errorw("error creating relationship tuples", "user", writes.TupleKey.User, "relation", writes.TupleKey.Relation, "object", writes.TupleKey.Object)

			return resp, newWritingTuplesError(writes.TupleKey.User, writes.TupleKey.Relation, writes.TupleKey.Object, "writing", err)
		}
	}

	for _, deletes := range resp.Deletes {
		if deletes.Error != nil {
			c.Logger.Errorw("error deleting relationship tuples", "user", deletes.TupleKey.User, "relation", deletes.TupleKey.Relation, "object", deletes.TupleKey.Object)

			return resp, newWritingTuplesError(deletes.TupleKey.User, deletes.TupleKey.Relation, deletes.TupleKey.Object, "writing", err)
		}
	}

	return resp, nil
}

// DeleteRelationshipTuple deletes a relationship tuple in the openFGA store
func (c *Client) DeleteRelationshipTuple(ctx context.Context, tuples []openfga.TupleKeyWithoutCondition) (*ofgaclient.ClientWriteResponse, error) {
	if len(tuples) == 0 {
		return nil, nil
	}

	opts := ofgaclient.ClientWriteOptions{AuthorizationModelId: openfga.PtrString(c.Config.AuthorizationModelId)}

	resp, err := c.Ofga.DeleteTuples(ctx).Body(tuples).Options(opts).Execute()
	if err != nil {
		c.Logger.Errorw("error deleting relationship tuples", "error", err.Error())

		return resp, err
	}

	for _, del := range resp.Deletes {
		if del.Error != nil {
			c.Logger.Errorw("error deleting relationship tuples", "user", del.TupleKey.User, "relation", del.TupleKey.Relation, "object", del.TupleKey.Object)

			return resp, newWritingTuplesError(del.TupleKey.User, del.TupleKey.Relation, del.TupleKey.Object, "deleting", err)
		}
	}

	return resp, nil
}

func (c *Client) DeleteAllObjectRelations(ctx context.Context, object string) error {
	// validate object is not empty
	if object == "" {
		return ErrMissingObjectOnDeletion
	}

	match := entityRegex.FindStringSubmatch(object)
	if match == nil {
		return newInvalidEntityError(object)
	}

	// TODO: update page size for pagination
	opts := ofgaclient.ClientReadOptions{}

	resp, err := c.Ofga.Read(ctx).Options(opts).Execute()
	if err != nil {
		c.Logger.Errorw("error deleting relationship tuples", "error", err.Error())

		return err
	}

	var tuplesToDelete []openfga.TupleKeyWithoutCondition

	// check all the tuples for the object?
	for _, t := range resp.GetTuples() {
		if t.Key.Object == object {
			k := openfga.TupleKeyWithoutCondition{
				User:     t.Key.User,
				Relation: t.Key.Relation,
				Object:   t.Key.Object,
			}
			tuplesToDelete = append(tuplesToDelete, k)
		}
	}

	// Notes: Writes only allow 10 tuples per call, this will need to be fixed
	_, err = c.DeleteRelationshipTuple(ctx, tuplesToDelete)

	return err
}

// GetTupleKey creates a Tuple key with the provided subject, object, and role
func GetTupleKey(subjectID, subjectType, objectID, objectType, relation string) TupleKey {
	sub := Entity{
		Kind:       Kind(subjectType),
		Identifier: subjectID,
	}

	object := Entity{
		Kind:       Kind(objectType),
		Identifier: objectID,
	}

	return TupleKey{
		Subject:  sub,
		Object:   object,
		Relation: Relation(relation),
	}
}
