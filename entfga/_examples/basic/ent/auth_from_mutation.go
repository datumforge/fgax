// Code generated by entfga, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga"
)

func (m *OrgMembershipMutation) CreateTuplesFromCreate(ctx context.Context) error {

	// Get fields for tuple creation
	userID, _ := m.UserID()
	objectID, _ := m.OrganizationID()
	role, _ := m.Role()

	// get tuple key
	req := TupleRequest{
		SubjectID:   userID,
		SubjectType: "user",
		ObjectID:    objectID,
		ObjectType:  "organization",
		Relation:    role.String(),
	}
	tuple := fgax.GetTupleKey(req)

	if _, err := m.Authz.WriteTupleKeys(ctx, []fgax.TupleKey{tuple}, nil); err != nil {
		m.Logger.Errorw("failed to create relationship tuple", "error", err)

		return err
	}

	m.Logger.Debugw("created relationship tuples", "relation", role, "object", tuple.Object)

	return nil
}

func (m *OrgMembershipMutation) CreateTuplesFromUpdate(ctx context.Context) error {

	// get ids that will be updated
	ids, err := m.IDs(ctx)
	if err != nil {
		return err
	}

	var (
		writes  []fgax.TupleKey
		deletes []fgax.TupleKey
	)

	oldRole, err := m.OldRole(ctx)
	if err != nil {
		return err
	}

	newRole, exists := m.Role()
	if !exists {
		return entfga.ErrMissingRole
	}

	if oldRole == newRole {
		m.Logger.Debugw("nothing to update, roles are the same", "old_role", oldRole, "new_role", newRole)

		return nil
	}

	// User the IDs of the memberships and delete all related tuples
	for _, id := range ids {
		member, err := m.Client().OrgMembership.Get(ctx, id)
		if err != nil {
			return err
		}

		req := TupleRequest{
			SubjectID:   member.UserID,
			SubjectType: "user",
			ObjectID:    member.OrganizationID,
			ObjectType:  "organization",
			Relation:    oldRole.String(),
		}
		d := fgax.GetTupleKey(req)

		deletes = append(deletes, d)

		req.Relation = newRole.String()
		w := fgax.GetTupleKey(req)

		writes = append(writes, w)

		if len(writes) == 0 && len(deletes) == 0 {
			m.Logger.Debugw("no relationships to create or delete")

			return nil
		}

		if _, err := m.Authz.WriteTupleKeys(ctx, writes, deletes); err != nil {
			m.Logger.Errorw("failed to update relationship tuple", "error", err)

			return err
		}
	}

	return nil
}

func (m *OrgMembershipMutation) CreateTuplesFromDelete(ctx context.Context) error {

	// get ids that will be deleted
	ids, err := m.IDs(ctx)
	if err != nil {
		return err
	}

	tuples := []fgax.TupleKey{}

	// User the IDs of the memberships and delete all related tuples
	for _, id := range ids {
		// this wont work with soft deletes
		members, err := m.Client().OrgMembership.Get(ctx, id)
		if err != nil {
			return err
		}

		req := TupleRequest{
			SubjectID:   members.UserID,
			SubjectType: "user",
			ObjectID:    members.OrganizationID,
			ObjectType:  "organization",
			Relation:    members.Role.String(),
		}
		t := fgax.GetTupleKey(req)

		tuples = append(tuples, t)
	}

	if len(tuples) > 0 {
		if _, err := m.Authz.WriteTupleKeys(ctx, nil, tuples); err != nil {
			m.Logger.Errorw("failed to delete relationship tuple", "error", err)

			return err
		}

		m.Logger.Debugw("deleted relationship tuples")
	}

	return nil
}
