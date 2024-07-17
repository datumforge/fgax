package fgax

import (
	"context"

	fgasdk "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
)

// listObjects checks the openFGA store for all objects associated with a user+relation
func (c *Client) listObjects(ctx context.Context, req ofgaclient.ClientListObjectsRequest) (*ofgaclient.ClientListObjectsResponse, error) {
	list, err := c.Ofga.ListObjects(ctx).Body(req).Execute()
	if err != nil {
		c.Logger.Errorw("error listing objects",
			"user", req.User,
			"relation", req.Relation,
			"type", req.Type,
			"error", err.Error())

		return nil, err
	}

	return list, nil
}

// ListObjectsRequest creates the ClientListObjectsRequest and queries the FGA store for all objects with the user+relation
func (c *Client) ListObjectsRequest(ctx context.Context, subjectID, subjectType, objectType, relation string) (*ofgaclient.ClientListObjectsResponse, error) {
	// default to user if no subjectType is provided
	if subjectType == "" {
		subjectType = "user"
	}

	sub := Entity{
		Kind:       Kind(subjectType),
		Identifier: subjectID,
	}

	listReq := ofgaclient.ClientListObjectsRequest{
		User:     sub.String(),
		Relation: relation,
		Type:     objectType,
		// TODO: Support contextual tuples
	}

	c.Logger.Debugw("listing objects", "relation", subjectType, sub.String(), relation, "type", objectType)

	return c.listObjects(ctx, listReq)
}

// listUsers checks the openFGA store for all users associated with a object+relation
func (c *Client) listUsers(ctx context.Context, req ofgaclient.ClientListUsersRequest) (*ofgaclient.ClientListUsersResponse, error) {
	list, err := c.Ofga.ListUsers(ctx).Body(req).Execute()
	if err != nil {
		c.Logger.Errorw("error listing users",
			"object", req.Object.Id,
			"type", req.Object.Type,
			"relation", req.Relation,
			"error", err.Error())

		return nil, err
	}

	return list, nil
}

// ListUserRequest is the fields needed to list users
type ListUserRequest struct {
	// ObjectID is the identifier of the object that the subject is related to
	ObjectID string
	// ObjectType is the type of object that the subject is related to
	ObjectType string
	// Relation is the relationship between the subject and object
	Relation string
	// UserTypeFilter is the type of user to filter by, by default it is "user"
	UserTypeFilter string
}

// ListUserRequest creates the ClientListUserRequest and queries the FGA store for all users with the object+relation
func (c *Client) ListUserRequest(ctx context.Context, l ListUserRequest) (*ofgaclient.ClientListUsersResponse, error) {
	// create the fga object
	obj := fgasdk.FgaObject{
		Type: l.ObjectType,
		Id:   l.ObjectID,
	}

	// default to user if no subjectType is provided
	if l.UserTypeFilter == "" {
		l.UserTypeFilter = defaultSubjectType
	}

	// compose the list request
	listReq := ofgaclient.ClientListUsersRequest{
		Object:      obj,
		Relation:    l.Relation,
		UserFilters: []fgasdk.UserTypeFilter{{Type: l.UserTypeFilter}},
		// TODO: Support contextual tuples
	}

	c.Logger.Debugw("listing users", "relation", l.Relation, "object", obj.Id, "type", obj.Type)

	return c.listUsers(ctx, listReq)
}

// ListContains checks the results of an fga ListObjects and parses the entities
// to get the identifier to compare to another identifier based on entity type
func ListContains(entityType string, l []string, i string) bool {
	for _, o := range l {
		e, _ := ParseEntity(o)

		// make sure its the correct entity type
		if e.Kind.String() != entityType {
			continue
		}

		if i == e.Identifier {
			return true
		}
	}

	return false
}

// GetEntityIDs returns a list of identifiers from a list of objects
func GetEntityIDs(l *ofgaclient.ClientListObjectsResponse) ([]string, error) {
	ids := make([]string, 0, len(l.Objects))

	for _, o := range l.Objects {
		e, err := ParseEntity(o)
		if err != nil {
			return nil, err
		}

		ids = append(ids, e.Identifier)
	}

	return ids, nil
}
