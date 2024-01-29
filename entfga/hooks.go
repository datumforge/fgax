package entfga

import (
	"context"
	"fmt"

	"entgo.io/ent"
)

// Mutation interface that all generated Mutation types must implement
type Mutation interface {
	// Op is the ent operation being taken on the Mutation (Create, Update, UpdateOne, Delete, DeleteOne)
	Op() ent.Op
	// CreateTuplesFromCreate creates tuple relationships for the user/object type on Create Mutations
	CreateTuplesFromCreate(ctx context.Context) error
	// CreateTuplesFromUpdate creates new and deletes old tuple relationships for the user/object type on Update Mutations
	CreateTuplesFromUpdate(ctx context.Context) error
	// CreateTuplesFromDelete deletes tuple relationships for the user/object type on Delete Mutations
	CreateTuplesFromDelete(ctx context.Context) error
}

// Mutator is an interface thats defines a method for mutating a generic ent value based on a given mutation.
// This is used as a generic interface that ent generated Mutations will implement
type Mutator interface {
	Mutate(context.Context, Mutation) (ent.Value, error)
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

// AuthzHooks returns a list of authorization hooks for create, update, and delete
// operations on a specific type of mutation.
func AuthzHooks[T Mutation]() []ent.Hook {
	return []ent.Hook{
		On(authzHookCreate[T](), ent.OpCreate),
		On(authzHookUpdate[T](), ent.OpUpdate|ent.OpUpdateOne),
		On(authzHookDelete[T](), ent.OpDelete|ent.OpDeleteOne),
	}
}

// getTypedMutation determines the specific mutation type
func getTypedMutation[T Mutation](m ent.Mutation) (T, error) {
	f, ok := any(m).(T)
	if !ok {
		return f, fmt.Errorf("expected appropriately typed mutation in schema hook, got: %+v", m) //nolint:goerr113
	}

	return f, nil
}

// authzHookCreate creates tuple relations in FGA after the mutation is executed
func authzHookCreate[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			err = mutation.CreateTuplesFromCreate(ctx)
			if err != nil {
				return nil, err
			}

			return value, nil
		})
	}
}

// authzHookUpdate updates (involving a delete and create) tuple relations in FGA after the mutation is executed
func authzHookUpdate[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			if err = mutation.CreateTuplesFromUpdate(ctx); err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}

// authzHookDelete removes tuple relations in FGA after the mutation is executed
func authzHookDelete[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			if err = mutation.CreateTuplesFromDelete(ctx); err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}