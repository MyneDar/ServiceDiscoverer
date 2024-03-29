// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"servicediscoverer/ent/endpointdata"
	"servicediscoverer/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EndpointDataDelete is the builder for deleting a EndpointData entity.
type EndpointDataDelete struct {
	config
	hooks    []Hook
	mutation *EndpointDataMutation
}

// Where appends a list predicates to the EndpointDataDelete builder.
func (edd *EndpointDataDelete) Where(ps ...predicate.EndpointData) *EndpointDataDelete {
	edd.mutation.Where(ps...)
	return edd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (edd *EndpointDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, EndpointDataMutation](ctx, edd.sqlExec, edd.mutation, edd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (edd *EndpointDataDelete) ExecX(ctx context.Context) int {
	n, err := edd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (edd *EndpointDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: endpointdata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endpointdata.FieldID,
			},
		},
	}
	if ps := edd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, edd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	edd.mutation.done = true
	return affected, err
}

// EndpointDataDeleteOne is the builder for deleting a single EndpointData entity.
type EndpointDataDeleteOne struct {
	edd *EndpointDataDelete
}

// Exec executes the deletion query.
func (eddo *EndpointDataDeleteOne) Exec(ctx context.Context) error {
	n, err := eddo.edd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{endpointdata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eddo *EndpointDataDeleteOne) ExecX(ctx context.Context) {
	eddo.edd.ExecX(ctx)
}
