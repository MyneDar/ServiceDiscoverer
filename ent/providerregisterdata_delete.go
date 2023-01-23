// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"servicediscoverer/ent/predicate"
	"servicediscoverer/ent/providerregisterdata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProviderRegisterDataDelete is the builder for deleting a ProviderRegisterData entity.
type ProviderRegisterDataDelete struct {
	config
	hooks    []Hook
	mutation *ProviderRegisterDataMutation
}

// Where appends a list predicates to the ProviderRegisterDataDelete builder.
func (prdd *ProviderRegisterDataDelete) Where(ps ...predicate.ProviderRegisterData) *ProviderRegisterDataDelete {
	prdd.mutation.Where(ps...)
	return prdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prdd *ProviderRegisterDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ProviderRegisterDataMutation](ctx, prdd.sqlExec, prdd.mutation, prdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prdd *ProviderRegisterDataDelete) ExecX(ctx context.Context) int {
	n, err := prdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prdd *ProviderRegisterDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: providerregisterdata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: providerregisterdata.FieldID,
			},
		},
	}
	if ps := prdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prdd.mutation.done = true
	return affected, err
}

// ProviderRegisterDataDeleteOne is the builder for deleting a single ProviderRegisterData entity.
type ProviderRegisterDataDeleteOne struct {
	prdd *ProviderRegisterDataDelete
}

// Exec executes the deletion query.
func (prddo *ProviderRegisterDataDeleteOne) Exec(ctx context.Context) error {
	n, err := prddo.prdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{providerregisterdata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prddo *ProviderRegisterDataDeleteOne) ExecX(ctx context.Context) {
	prddo.prdd.ExecX(ctx)
}
