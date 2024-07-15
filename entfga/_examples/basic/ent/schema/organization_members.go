package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/datumforge/fgax/entfga"

	generated "github.com/datumforge/fgax/entfga/_examples/basic/ent"
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/enums"
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/privacy"
)

// OrgMembership holds the schema definition for the OrgMembership entity
type OrgMembership struct {
	ent.Schema
}

// Fields of the OrgMembership
func (OrgMembership) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable(),
		field.Enum("role").
			GoType(enums.Role("")).
			Default(string(enums.RoleMember)).
			Values(string(enums.RoleOwner)), // adds owner to possible values
		field.String("organization_id").Immutable(),
		field.String("user_id").Immutable(),
	}
}

// Edges of the OrgMembership
func (OrgMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization", Organization.Type).
			Field("organization_id").
			Required().
			Unique().
			Immutable(),
	}
}

// Annotations of the OrgMembership
func (OrgMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entfga.Annotations{
			ObjectType:    "organization",
			IncludeHooks:  true,
			OrgOwnedField: true,
			IDField:       "OrganizationID",
		},
	}
}

func (OrgMembership) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "organization_id").
			Unique(),
	}
}

// Policy of the OrgMembership
func (OrgMembership) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OrgMembershipMutationRuleFunc(func(ctx context.Context, m *generated.OrgMembershipMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.OrgMembershipQueryRuleFunc(func(ctx context.Context, q *generated.OrgMembershipQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
