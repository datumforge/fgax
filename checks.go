package fgax

import (
	"context"

	ofgaclient "github.com/openfga/go-sdk/client"
)

// AccessCheck is a struct to hold the information needed to check access
type AccessCheck struct {
	// ObjectType is the type of object being checked
	ObjectType Kind
	// ObjectID is the ID of the object being checked
	ObjectID string
	// Relation is the relationship being checked (e.g. "view", "edit", "delete")
	Relation string
	// UserID is the ID of the user making the request
	UserID string
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

	sub := Entity{
		Kind:       "user",
		Identifier: ac.UserID,
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
func (c *Client) CheckOrgAccess(ctx context.Context, userID, orgID, relation string) (bool, error) {
	ac := AccessCheck{
		ObjectType: "organization",
		ObjectID:   orgID,
		Relation:   relation,
		UserID:     userID,
	}

	return c.CheckAccess(ctx, ac)
}

// CheckGroupAccess checks if the user has access to the group with the given relation
func (c *Client) CheckGroupAccess(ctx context.Context, userID, groupID, relation string) (bool, error) {
	ac := AccessCheck{
		ObjectType: "group",
		ObjectID:   groupID,
		Relation:   relation,
		UserID:     userID,
	}

	return c.CheckAccess(ctx, ac)
}

// validateAccessCheck checks if the AccessCheck struct is valid
func validateAccessCheck(ac AccessCheck) error {
	if ac.UserID == "" {
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
