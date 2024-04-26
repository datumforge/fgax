package fgax

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	language "github.com/openfga/language/pkg/go/transformer"
	typesystem "github.com/openfga/openfga/pkg/typesystem"
	"github.com/samber/lo"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateModelFromFile creates a new fine grained authorization model and returns the model ID
func (c *Client) CreateModelFromFile(ctx context.Context, fn string, forceCreate bool) (string, error) {
	options := ofgaclient.ClientReadAuthorizationModelsOptions{}

	models, err := c.Ofga.ReadAuthorizationModels(context.Background()).Options(options).Execute()
	if err != nil {
		return "", err
	}

	// Only create a new test model if one does not exist and we aren't forcing a new model to be created
	if !forceCreate {
		if len(models.AuthorizationModels) > 0 {
			modelID := models.GetAuthorizationModels()[0].Id
			c.Logger.Infow("fga model exists", "model_id", modelID)

			return modelID, nil
		}
	}

	// Create new model
	dsl, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}

	return c.CreateModelFromDSL(ctx, dsl)
}

// CreateModelFromDSL creates a new fine grained authorization model from the DSL and returns the model ID
func (c *Client) CreateModelFromDSL(ctx context.Context, dsl []byte) (string, error) {
	// convert to json
	dslJSON, err := dslToJSON(dsl)
	if err != nil {
		return "", err
	}

	var body ofgaclient.ClientWriteAuthorizationModelRequest
	if err := json.Unmarshal(dslJSON, &body); err != nil {
		return "", err
	}

	return c.CreateModel(ctx, body)
}

// CreateModel updates the model and returns the new model ID
func (c *Client) CreateModel(ctx context.Context, model ofgaclient.ClientWriteAuthorizationModelRequest) (string, error) {
	resp, err := c.Ofga.WriteAuthorizationModel(ctx).Body(model).Execute()
	if err != nil {
		return "", err
	}

	modelID := resp.GetAuthorizationModelId()

	c.Logger.Infow("fga model created", "model_id", modelID)

	return modelID, nil
}

// dslToJSON converts fga model to JSON
func dslToJSON(dslString []byte) ([]byte, error) {
	parsedAuthModel, err := language.TransformDSLToProto(string(dslString))
	if err != nil {
		return []byte{}, errors.Wrap(err, ErrFailedToTransformModel.Error())
	}

	return protojson.Marshal(parsedAuthModel)
}

// RoleRequest is the request to add a role to the model for an existing object
type RoleRequest struct {
	// Role is the relation to add to the model
	Role string
	// Relation is the relation to the object
	Relations []RelationSetting
	// RelationCombination is the combination of the relation
	RelationCombination RelationCombination
	// ObjectType is the object type to add the role to
	ObjectType string
}

type RelationCombination struct {
	// IsUnion is an `or` relation
	IsUnion bool
	// IsIntersection is an `and` relation
	IsIntersection bool
	// IsDifference is a `not` relation
	IsDifference bool // TODO: implement
}

type RelationSetting struct {
	// Relation is the relation to the object
	Relation string
	// IsDirect is the direct relation to another fga object type
	IsDirect bool
	// FromRelation is the relation from another relation, leave empty if not a from relation
	FromRelation string
}

// AddRoleToModel adds a role to the model and returns the new model ID
func (c *Client) AddRoleToModel(ctx context.Context, r RoleRequest) error {
	// read the latest model
	model, err := c.Ofga.ReadLatestAuthorizationModel(ctx).Execute()
	if err != nil {
		return err
	}

	if r.RelationCombination.IsUnion && r.RelationCombination.IsIntersection {
		return errors.New("cannot have both union and intersection") // TODO: error definition
	}

	// get the model
	m := model.GetAuthorizationModel()

	// then add the role
	td := m.TypeDefinitions

	addedRole := false
	for i, t := range td {
		if t.Type == r.ObjectType {
			// initialize the relation map
			relations := t.GetRelations()

			// add the role to the relation map
			var metadata *openfga.Metadata
			relations[r.Role], metadata = createNewUserset(r)

			// set the relation map and metadata
			t.SetRelations(relations)
			// t.SetMetadata(metadata)

			metadataRelations := t.Metadata.GetRelations()
			for k, v := range *metadata.Relations {
				metadataRelations[k] = v
			}

			t.Metadata.SetRelations(metadataRelations)

			// set the updated type definition
			m.TypeDefinitions[i] = t

			// track that we added the role
			addedRole = true
		}
	}

	// if we didn't add the role, create a new type definition
	if !addedRole {
		m.TypeDefinitions = append(m.TypeDefinitions,
			createNewTypeDefinition(r))
	}

	// create the request to write the model
	request := ofgaclient.ClientWriteAuthorizationModelRequest{
		SchemaVersion:   m.SchemaVersion,
		Conditions:      m.Conditions,
		TypeDefinitions: m.TypeDefinitions,
	}

	// then write the model back
	resp, err := c.Ofga.WriteAuthorizationModel(ctx).Body(request).Execute()
	if err != nil {
		return err
	}

	// then update to the new model ID in the config
	c.Config.AuthorizationModelId = resp.GetAuthorizationModelId()

	return nil
}

// createNewTypeDefinition creates a new type definition for the model
func createNewTypeDefinition(r RoleRequest) openfga.TypeDefinition {
	rel := make(map[string]openfga.Userset)
	td := openfga.TypeDefinition{
		Type: r.ObjectType,
	}

	// get all the usersets
	us, metadata := createNewUserset(r)

	rel[r.Role] = us
	td.Relations = &rel
	td.Metadata = metadata
	fmt.Println(rel)
	return td
}

func createNewUserset(r RoleRequest) (openfga.Userset, *openfga.Metadata) {
	// get all the usersets
	us := openfga.Userset{}
	metadata := openfga.NewMetadataWithDefaults()

	// place holders to combine the usersets
	var (
		// ttus []openfga.TupleToUserset
		uses []openfga.Userset
	)

	hasDirect := false

	for _, relation := range r.Relations {
		if relation.IsDirect {
			// create direct relation
			relations, md := newDirectRelation(r.Role, relation)
			// TODO: technically you can have more than one direct, so we need to append later
			uses = append(uses, relations)
			metadata.SetRelations(md)
			hasDirect = true
		} else if relation.FromRelation != "" {
			// create tupleSet
			ts := openfga.TupleToUserset{
				Tupleset: openfga.ObjectRelation{
					Relation: &relation.FromRelation,
				},
				ComputedUserset: openfga.ObjectRelation{
					Object:   lo.ToPtr(""),
					Relation: &relation.Relation,
				},
			}

			userset := openfga.Userset{
				TupleToUserset: &ts,
			}

			uses = append(uses, userset)
		} else {
			// create computedUserset
			uses = append(uses, openfga.Userset{
				ComputedUserset: &openfga.ObjectRelation{
					Object:   lo.ToPtr(""),
					Relation: &relation.Relation,
				},
			})
		}

	}

	// don't overwrite if we have a direct relation
	if !hasDirect {
		md := createNewMetadata(r.Role, "")
		metadata.SetRelations(md)
	}

	// now combine the usersets
	if len(uses) == 1 {
		us = uses[0]
	} else if r.RelationCombination.IsIntersection {
		us.Intersection = &openfga.Usersets{
			Child: uses,
		}
	} else if r.RelationCombination.IsUnion {
		us.Union = &openfga.Usersets{
			Child: uses,
		}
	}

	return us, metadata
}

func createNewMetadata(r string, userType string) map[string]openfga.RelationMetadata {
	rd := make(map[string]openfga.RelationMetadata)
	rd[r] = openfga.RelationMetadata{
		DirectlyRelatedUserTypes: &[]openfga.RelationReference{},
	}

	if userType != "" {
		rd[r] = openfga.RelationMetadata{
			DirectlyRelatedUserTypes: &[]openfga.RelationReference{
				{
					Type: userType,
				},
			},
		}

	}

	return rd
}

// newDirectRelation creates a new relation to an existing object
func newDirectRelation(role string, r RelationSetting) (openfga.Userset, map[string]openfga.RelationMetadata) {
	// create the userset
	thisRelation := make(map[string]interface{})
	thisRelation[role] = typesystem.This()

	us := openfga.Userset{
		This: &thisRelation,
	}

	// create the relation metadata, this is required for the relation to be valid
	rds := createNewMetadata(role, r.Relation)

	return us, rds
}

func getParentUserType() string {
	return ""
}

// NOTES:

// computedUserset is a relation 1 relation
// intersection is a relation 2 relation with an and
// union is a relation 2 relation with an or
// direct is a relation to another object type
// direct requires metadata
// tupletoUserset is a from relation from another relation
// once we get all the usersets, we need to combine them into a single userset
