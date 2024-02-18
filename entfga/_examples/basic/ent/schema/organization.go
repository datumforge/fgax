package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/datumforge/fgax/entfga"
)

// Organization holds the schema definition for the Organization entity
type Organization struct {
	ent.Schema
}

// Fields of the Organization
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable(),
		field.String("name").
			Comment("the name of the organization").
			NotEmpty(),
		field.String("description").
			Comment("An optional description of the organization").
			Optional(),
	}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}

// Annotations of the Organization
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entfga.Annotations{
			ObjectType:   "organization",
			IncludeHooks: false,
		},
	}
}

// // Policy defines the privacy policy of the Organization.
// func (Organization) Policy() ent.Policy {
// 	return privacy.Policy{
// 		Query: privacy.QueryPolicy{
// 			privacy.OrganizationQueryRuleFunc(func(ctx context.Context, q *generated.OrganizationQuery) error {
// 				return q.CheckAccess(ctx)
// 			}),
// 			privacy.AlwaysDenyRule(), // Deny all other users
// 		},
// 	}
// }
