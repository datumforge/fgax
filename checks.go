package fgax

import (
	"context"

	ofgaclient "github.com/openfga/go-sdk/client"
)

const (
	defaultSubjectType = "user"
)

// AccessCheck is a struct to hold the information needed to check access
type AccessCheck struct {
	// ObjectType is the type of object being checked
	ObjectType Kind
	// ObjectID is the ID of the object being checked
	ObjectID string
	// SubjectID is the ID of the user making the request
	SubjectID string
	// SubjectType is the type of subject being checked
	SubjectType string
	// Relation is the relationship being checked (e.g. "view", "edit", "delete")
	Relation string
}

// CheckTuple checks the openFGA store for provided relationship tuple
func (c *Client) CheckTuple(ctx context.Context, check ofgaclient.ClientCheckRequest) (bool, error) {
	data, err := c.Ofga.Check(ctx).Body(check).Execute()
	if err != nil {
		c.Logger.Errorw("error checking tuple", "tuple", check, "error", err.Error())

		return false, err
	}

	return *data.Allowed, nil
}

// CheckAccess checks if the user has access to the object type with the given relation
func (c *Client) CheckAccess(ctx context.Context, ac AccessCheck) (bool, error) {
	if err := validateAccessCheck(ac); err != nil {
		return false, err
	}

	if ac.SubjectType == "" {
		ac.SubjectType = defaultSubjectType
	}

	sub := Entity{
		Kind:       Kind(ac.SubjectType),
		Identifier: ac.SubjectID,
	}

	obj := Entity{
		Kind:       ac.ObjectType,
		Identifier: ac.ObjectID,
	}

	c.Logger.Infow("checking relationship tuples", "relation", ac.Relation, "object", obj.String())

	checkReq := ofgaclient.ClientCheckRequest{
		User:     sub.String(),
		Relation: ac.Relation,
		Object:   obj.String(),
	}

	return c.CheckTuple(ctx, checkReq)
}

// CheckOrgAccess checks if the user has access to the organization with the given relation
func (c *Client) CheckOrgAccess(ctx context.Context, ac AccessCheck) (bool, error) {
	ac.ObjectType = "organization"

	return c.CheckAccess(ctx, ac)
}

// CheckGroupAccess checks if the user has access to the group with the given relation
func (c *Client) CheckGroupAccess(ctx context.Context, ac AccessCheck) (bool, error) {
	ac.ObjectType = "group"

	return c.CheckAccess(ctx, ac)
}

// CheckSystemAdminRole checks if the user has system admin access
func (c *Client) CheckSystemAdminRole(ctx context.Context, userID string) (bool, error) {
	ac := AccessCheck{
		ObjectType:  "role",
		ObjectID:    SystemAdminRole,
		Relation:    RoleRelation,
		SubjectID:   userID,
		SubjectType: defaultSubjectType, // admin roles are always user roles, never an API token
	}

	return c.CheckAccess(ctx, ac)
}

// validateAccessCheck checks if the AccessCheck struct is valid
func validateAccessCheck(ac AccessCheck) error {
	if ac.SubjectID == "" {
		return ErrInvalidAccessCheck
	}

	if ac.ObjectType == "" {
		return ErrInvalidAccessCheck
	}

	if ac.ObjectID == "" {
		return ErrInvalidAccessCheck
	}

	if ac.Relation == "" {
		return ErrInvalidAccessCheck
	}

	return nil
}
