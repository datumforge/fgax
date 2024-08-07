package entfga

import (
	"context"
	"fmt"

	"entgo.io/ent"
)

// Mutation interface that all generated Mutation types must implement
// These functions (with the exception of Op() which is already created) are generated by the ent extension for every schema
// that includes the `entfga.NewFGAExtension“ extension to satisfy the interface
// If hooks are skipped by the mutation, the functions are created to satisfy the interface but always return nil and are not added to the client
type Mutation interface {
	// Op is the ent operation being taken on the Mutation (Create, Update, UpdateOne, Delete, DeleteOne)
	Op() ent.Op
	// CreateTuplesFromCreate creates tuple relationships for the user/object type on Create Mutations
	CreateTuplesFromCreate(ctx context.Context) error
	// CreateTuplesFromUpdate creates new and deletes old tuple relationships for the user/object type on Update Mutations
	CreateTuplesFromUpdate(ctx context.Context) error
	// CreateTuplesFromDelete deletes tuple relationships for the user/object type on Delete Mutations
	CreateTuplesFromDelete(ctx context.Context) error
	// CheckAccessForEdit checks if the user has access to edit the object type
	CheckAccessForEdit(ctx context.Context) error
	// CheckAccessForDelete checks if the user has access to delete the object type
	CheckAccessForDelete(ctx context.Context) error
}

// Mutator is an interface thats defines a method for mutating a generic ent value based on a given mutation.
// This is used as a generic interface that ent generated Mutations will implement
type Mutator interface {
	Mutate(context.Context, Mutation) (ent.Value, error)
}

// Query interface that all generated Query types must implement
type Query interface {
	// Op is the ent operation being taken on the Mutation (Create, Update, UpdateOne, Delete, DeleteOne)
	Op() ent.Op

	// CheckAccess checks if the user has read access to the object type
	CheckAccess(ctx context.Context) error
}

// Querier is an interface thats defines a method for querying a generic ent value based on a given query.
// This is used as a generic interface that ent generated Query will implement
type Querier interface {
	Query(context.Context, Query) (ent.Value, error)
}

// getTypedMutation determines the specific mutation type
func getTypedMutation[T Mutation](m ent.Mutation) (T, error) {
	f, ok := any(m).(T)
	if !ok {
		return f, fmt.Errorf("expected appropriately typed mutation, got: %+v", m) //nolint:err113
	}

	return f, nil
}

// On will execute the appropriate hook based on the ent operation
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			hasOp := m.Op().Is(op)

			if hasOp {
				return hk(next).Mutate(ctx, m)
			}

			return next.Mutate(ctx, m)
		})
	}
}
