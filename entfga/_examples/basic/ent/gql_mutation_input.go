// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/enums"
)

// CreateOrgMembershipInput represents a mutation input for creating orgmemberships.
type CreateOrgMembershipInput struct {
	Role           *enums.Role
	UserID         string
	OrganizationID string
}

// Mutate applies the CreateOrgMembershipInput on the OrgMembershipMutation builder.
func (i *CreateOrgMembershipInput) Mutate(m *OrgMembershipMutation) {
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	m.SetUserID(i.UserID)
	m.SetOrganizationID(i.OrganizationID)
}

// SetInput applies the change-set in the CreateOrgMembershipInput on the OrgMembershipCreate builder.
func (c *OrgMembershipCreate) SetInput(i CreateOrgMembershipInput) *OrgMembershipCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrgMembershipInput represents a mutation input for updating orgmemberships.
type UpdateOrgMembershipInput struct {
	Role *enums.Role
}

// Mutate applies the UpdateOrgMembershipInput on the OrgMembershipMutation builder.
func (i *UpdateOrgMembershipInput) Mutate(m *OrgMembershipMutation) {
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
}

// SetInput applies the change-set in the UpdateOrgMembershipInput on the OrgMembershipUpdate builder.
func (c *OrgMembershipUpdate) SetInput(i UpdateOrgMembershipInput) *OrgMembershipUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrgMembershipInput on the OrgMembershipUpdateOne builder.
func (c *OrgMembershipUpdateOne) SetInput(i UpdateOrgMembershipInput) *OrgMembershipUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrganizationInput represents a mutation input for creating organizations.
type CreateOrganizationInput struct {
	Name        string
	Description *string
}

// Mutate applies the CreateOrganizationInput on the OrganizationMutation builder.
func (i *CreateOrganizationInput) Mutate(m *OrganizationMutation) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
}

// SetInput applies the change-set in the CreateOrganizationInput on the OrganizationCreate builder.
func (c *OrganizationCreate) SetInput(i CreateOrganizationInput) *OrganizationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrganizationInput represents a mutation input for updating organizations.
type UpdateOrganizationInput struct {
	Name             *string
	ClearDescription bool
	Description      *string
}

// Mutate applies the UpdateOrganizationInput on the OrganizationMutation builder.
func (i *UpdateOrganizationInput) Mutate(m *OrganizationMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdate builder.
func (c *OrganizationUpdate) SetInput(i UpdateOrganizationInput) *OrganizationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdateOne builder.
func (c *OrganizationUpdateOne) SetInput(i UpdateOrganizationInput) *OrganizationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
