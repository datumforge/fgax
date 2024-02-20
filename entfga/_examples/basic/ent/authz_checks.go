// Code generated by entfga, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/privacy"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga/_examples/basic/auth"
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/organization"
	"github.com/datumforge/fgax/entfga/_examples/basic/ent/orgmembership"
)

func (q *OrgMembershipQuery) CheckAccess(ctx context.Context) error {
	gCtx := graphql.GetFieldContext(ctx)

	if gCtx != nil {
		ac := fgax.AccessCheck{
			Relation:   fgax.CanView,
			ObjectType: "organization",
		}

		// check id from graphql arg context
		// when all objects are requested, the interceptor will check object access
		// check the where input first
		whereArg := gCtx.Args["where"]
		if whereArg != nil {
			where, ok := whereArg.(*OrgMembershipWhereInput)
			if ok && where != nil && where.OrganizationID != nil {
				ac.ObjectID = *where.OrganizationID
			}
		}

		// if that doesn't work, check for the id in the args
		if ac.ObjectID == "" {
			ac.ObjectID, _ = gCtx.Args["organizationid"].(string)
		}

		// if we still don't have an object id, run the query and grab the object ID
		// from the result
		// this happens on join tables where we have the join ID (for updates and deletes)
		// and not the actual object id
		if ac.ObjectID == "" && "id" != "organizationid" {
			// allow this query to run
			reqCtx := privacy.DecisionContext(ctx, privacy.Allow)
			ob, err := q.Only(reqCtx)
			if err != nil {
				return privacy.Allowf("nil request, bypassing auth check")
			}

			ac.ObjectID = ob.OrganizationID
		}

		// request is for a list objects, will get filtered in interceptors
		if ac.ObjectID == "" {
			return privacy.Allowf("nil request, bypassing auth check")
		}

		var err error
		ac.UserID, err = auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		access, err := q.Authz.CheckAccess(ctx, ac)
		if err != nil {
			return privacy.Skipf("unable to check access, %s", err.Error())
		}

		if access {
			return privacy.Allow
		}
	}

	// Skip to the next privacy rule (equivalent to return nil)
	return privacy.Skip
}

func (m *OrgMembershipMutation) CheckAccessForEdit(ctx context.Context) error {
	ac := fgax.AccessCheck{
		Relation:   fgax.CanEdit,
		ObjectType: "organization",
	}

	gCtx := graphql.GetFieldContext(ctx)

	// get the input from the context
	gInput := gCtx.Args["input"]

	// check if the input is a CreateOrgMembershipInput
	input, ok := gInput.(CreateOrgMembershipInput)
	if ok {
		ac.ObjectID = input.OrganizationID
	}

	// check the id from the args
	if ac.ObjectID == "" {
		ac.ObjectID, _ = gCtx.Args["organizationid"].(string)
	}

	// if this is still empty, we need to query the object to get the object id
	// this happens on join tables where we have the join ID (for updates and deletes)
	if ac.ObjectID == "" {
		id, ok := gCtx.Args["id"].(string)
		if ok {
			// allow this query to run
			reqCtx := privacy.DecisionContext(ctx, privacy.Allow)
			ob, err := m.Client().OrgMembership.Query().Where(orgmembership.ID(id)).Only(reqCtx)
			if err != nil {
				return privacy.Allowf("nil request, bypassing auth check")
			}

			ac.ObjectID = ob.OrganizationID
		}
	}

	// request is for a list objects, will get filtered in interceptors
	if ac.ObjectID == "" {
		return privacy.Allowf("nil request, bypassing auth check")
	}

	m.Logger.Debugw("checking mutation access")

	var err error
	ac.UserID, err = auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	m.Logger.Infow("checking relationship tuples", "relation", ac.Relation, "object_id", ac.ObjectID)

	access, err := m.Authz.CheckAccess(ctx, ac)
	if err != nil {
		return privacy.Skipf("unable to check access, %s", err.Error())
	}

	if access {
		m.Logger.Debugw("access allowed", "relation", ac.Relation, "object_id", ac.ObjectID)

		return privacy.Allow
	}

	// deny if it was a mutation is not allowed
	return privacy.Deny
}

func (m *OrgMembershipMutation) CheckAccessForDelete(ctx context.Context) error {
	ac := fgax.AccessCheck{
		Relation:   fgax.CanDelete,
		ObjectType: "organization",
	}

	gCtx := graphql.GetFieldContext(ctx)

	var ok bool
	ac.ObjectID, ok = gCtx.Args["id"].(string)
	if !ok {
		return privacy.Allowf("nil request, bypassing auth check")
	}

	m.Logger.Debugw("checking mutation access")

	var err error
	ac.UserID, err = auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	m.Logger.Infow("checking relationship tuples", "relation", ac.Relation, "object_id", ac.ObjectID)

	access, err := m.Authz.CheckAccess(ctx, ac)
	if err != nil {
		return privacy.Skipf("unable to check access, %s", err.Error())
	}

	if access {
		m.Logger.Debugw("access allowed", "relation", ac.Relation, "object_id", ac.ObjectID)

		return privacy.Allow
	}

	// deny if it was a mutation is not allowed
	return privacy.Deny
}

func (q *OrganizationQuery) CheckAccess(ctx context.Context) error {
	gCtx := graphql.GetFieldContext(ctx)

	if gCtx != nil {
		ac := fgax.AccessCheck{
			Relation:   fgax.CanView,
			ObjectType: "organization",
		}

		// check id from graphql arg context
		// when all objects are requested, the interceptor will check object access
		// check the where input first
		whereArg := gCtx.Args["where"]
		if whereArg != nil {
			where, ok := whereArg.(*OrganizationWhereInput)
			if ok && where != nil && where.ID != nil {
				ac.ObjectID = *where.ID
			}
		}

		// if that doesn't work, check for the id in the args
		if ac.ObjectID == "" {
			ac.ObjectID, _ = gCtx.Args["id"].(string)
		}

		// if we still don't have an object id, run the query and grab the object ID
		// from the result
		// this happens on join tables where we have the join ID (for updates and deletes)
		// and not the actual object id
		if ac.ObjectID == "" && "id" != "id" {
			// allow this query to run
			reqCtx := privacy.DecisionContext(ctx, privacy.Allow)
			ob, err := q.Only(reqCtx)
			if err != nil {
				return privacy.Allowf("nil request, bypassing auth check")
			}

			ac.ObjectID = ob.ID
		}

		// request is for a list objects, will get filtered in interceptors
		if ac.ObjectID == "" {
			return privacy.Allowf("nil request, bypassing auth check")
		}

		var err error
		ac.UserID, err = auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		access, err := q.Authz.CheckAccess(ctx, ac)
		if err != nil {
			return privacy.Skipf("unable to check access, %s", err.Error())
		}

		if access {
			return privacy.Allow
		}
	}

	// Skip to the next privacy rule (equivalent to return nil)
	return privacy.Skip
}

func (m *OrganizationMutation) CheckAccessForEdit(ctx context.Context) error {
	ac := fgax.AccessCheck{
		Relation:   fgax.CanEdit,
		ObjectType: "organization",
	}

	gCtx := graphql.GetFieldContext(ctx)

	// check the id from the args
	if ac.ObjectID == "" {
		ac.ObjectID, _ = gCtx.Args["id"].(string)
	}

	// if this is still empty, we need to query the object to get the object id
	// this happens on join tables where we have the join ID (for updates and deletes)
	if ac.ObjectID == "" {
		id, ok := gCtx.Args["id"].(string)
		if ok {
			// allow this query to run
			reqCtx := privacy.DecisionContext(ctx, privacy.Allow)
			ob, err := m.Client().Organization.Query().Where(organization.ID(id)).Only(reqCtx)
			if err != nil {
				return privacy.Allowf("nil request, bypassing auth check")
			}

			ac.ObjectID = ob.ID
		}
	}

	// request is for a list objects, will get filtered in interceptors
	if ac.ObjectID == "" {
		return privacy.Allowf("nil request, bypassing auth check")
	}

	m.Logger.Debugw("checking mutation access")

	var err error
	ac.UserID, err = auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	m.Logger.Infow("checking relationship tuples", "relation", ac.Relation, "object_id", ac.ObjectID)

	access, err := m.Authz.CheckAccess(ctx, ac)
	if err != nil {
		return privacy.Skipf("unable to check access, %s", err.Error())
	}

	if access {
		m.Logger.Debugw("access allowed", "relation", ac.Relation, "object_id", ac.ObjectID)

		return privacy.Allow
	}

	// deny if it was a mutation is not allowed
	return privacy.Deny
}

func (m *OrganizationMutation) CheckAccessForDelete(ctx context.Context) error {
	ac := fgax.AccessCheck{
		Relation:   fgax.CanDelete,
		ObjectType: "organization",
	}

	gCtx := graphql.GetFieldContext(ctx)

	var ok bool
	ac.ObjectID, ok = gCtx.Args["id"].(string)
	if !ok {
		return privacy.Allowf("nil request, bypassing auth check")
	}

	m.Logger.Debugw("checking mutation access")

	var err error
	ac.UserID, err = auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	m.Logger.Infow("checking relationship tuples", "relation", ac.Relation, "object_id", ac.ObjectID)

	access, err := m.Authz.CheckAccess(ctx, ac)
	if err != nil {
		return privacy.Skipf("unable to check access, %s", err.Error())
	}

	if access {
		m.Logger.Debugw("access allowed", "relation", ac.Relation, "object_id", ac.ObjectID)

		return privacy.Allow
	}

	// deny if it was a mutation is not allowed
	return privacy.Deny
}
