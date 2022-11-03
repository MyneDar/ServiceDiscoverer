// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"servicediscoverer/ent/predicate"
	"servicediscoverer/ent/providerregisterdata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProviderRegisterDataUpdate is the builder for updating ProviderRegisterData entities.
type ProviderRegisterDataUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderRegisterDataMutation
}

// Where appends a list predicates to the ProviderRegisterDataUpdate builder.
func (prdu *ProviderRegisterDataUpdate) Where(ps ...predicate.ProviderRegisterData) *ProviderRegisterDataUpdate {
	prdu.mutation.Where(ps...)
	return prdu
}

// SetName sets the "name" field.
func (prdu *ProviderRegisterDataUpdate) SetName(s string) *ProviderRegisterDataUpdate {
	prdu.mutation.SetName(s)
	return prdu
}

// SetPort sets the "port" field.
func (prdu *ProviderRegisterDataUpdate) SetPort(s string) *ProviderRegisterDataUpdate {
	prdu.mutation.SetPort(s)
	return prdu
}

// SetAddress sets the "address" field.
func (prdu *ProviderRegisterDataUpdate) SetAddress(s string) *ProviderRegisterDataUpdate {
	prdu.mutation.SetAddress(s)
	return prdu
}

// SetDescription sets the "description" field.
func (prdu *ProviderRegisterDataUpdate) SetDescription(s string) *ProviderRegisterDataUpdate {
	prdu.mutation.SetDescription(s)
	return prdu
}

// SetLiveInterval sets the "liveInterval" field.
func (prdu *ProviderRegisterDataUpdate) SetLiveInterval(i int) *ProviderRegisterDataUpdate {
	prdu.mutation.ResetLiveInterval()
	prdu.mutation.SetLiveInterval(i)
	return prdu
}

// AddLiveInterval adds i to the "liveInterval" field.
func (prdu *ProviderRegisterDataUpdate) AddLiveInterval(i int) *ProviderRegisterDataUpdate {
	prdu.mutation.AddLiveInterval(i)
	return prdu
}

// SetLiveTimeout sets the "liveTimeout" field.
func (prdu *ProviderRegisterDataUpdate) SetLiveTimeout(i int) *ProviderRegisterDataUpdate {
	prdu.mutation.ResetLiveTimeout()
	prdu.mutation.SetLiveTimeout(i)
	return prdu
}

// AddLiveTimeout adds i to the "liveTimeout" field.
func (prdu *ProviderRegisterDataUpdate) AddLiveTimeout(i int) *ProviderRegisterDataUpdate {
	prdu.mutation.AddLiveTimeout(i)
	return prdu
}

// Mutation returns the ProviderRegisterDataMutation object of the builder.
func (prdu *ProviderRegisterDataUpdate) Mutation() *ProviderRegisterDataMutation {
	return prdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (prdu *ProviderRegisterDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(prdu.hooks) == 0 {
		affected, err = prdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProviderRegisterDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prdu.mutation = mutation
			affected, err = prdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(prdu.hooks) - 1; i >= 0; i-- {
			if prdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, prdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (prdu *ProviderRegisterDataUpdate) SaveX(ctx context.Context) int {
	affected, err := prdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (prdu *ProviderRegisterDataUpdate) Exec(ctx context.Context) error {
	_, err := prdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prdu *ProviderRegisterDataUpdate) ExecX(ctx context.Context) {
	if err := prdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (prdu *ProviderRegisterDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   providerregisterdata.Table,
			Columns: providerregisterdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: providerregisterdata.FieldID,
			},
		},
	}
	if ps := prdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := prdu.mutation.Name(); ok {
		_spec.SetField(providerregisterdata.FieldName, field.TypeString, value)
	}
	if value, ok := prdu.mutation.Port(); ok {
		_spec.SetField(providerregisterdata.FieldPort, field.TypeString, value)
	}
	if value, ok := prdu.mutation.Address(); ok {
		_spec.SetField(providerregisterdata.FieldAddress, field.TypeString, value)
	}
	if value, ok := prdu.mutation.Description(); ok {
		_spec.SetField(providerregisterdata.FieldDescription, field.TypeString, value)
	}
	if value, ok := prdu.mutation.LiveInterval(); ok {
		_spec.SetField(providerregisterdata.FieldLiveInterval, field.TypeInt, value)
	}
	if value, ok := prdu.mutation.AddedLiveInterval(); ok {
		_spec.AddField(providerregisterdata.FieldLiveInterval, field.TypeInt, value)
	}
	if value, ok := prdu.mutation.LiveTimeout(); ok {
		_spec.SetField(providerregisterdata.FieldLiveTimeout, field.TypeInt, value)
	}
	if value, ok := prdu.mutation.AddedLiveTimeout(); ok {
		_spec.AddField(providerregisterdata.FieldLiveTimeout, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, prdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerregisterdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ProviderRegisterDataUpdateOne is the builder for updating a single ProviderRegisterData entity.
type ProviderRegisterDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderRegisterDataMutation
}

// SetName sets the "name" field.
func (prduo *ProviderRegisterDataUpdateOne) SetName(s string) *ProviderRegisterDataUpdateOne {
	prduo.mutation.SetName(s)
	return prduo
}

// SetPort sets the "port" field.
func (prduo *ProviderRegisterDataUpdateOne) SetPort(s string) *ProviderRegisterDataUpdateOne {
	prduo.mutation.SetPort(s)
	return prduo
}

// SetAddress sets the "address" field.
func (prduo *ProviderRegisterDataUpdateOne) SetAddress(s string) *ProviderRegisterDataUpdateOne {
	prduo.mutation.SetAddress(s)
	return prduo
}

// SetDescription sets the "description" field.
func (prduo *ProviderRegisterDataUpdateOne) SetDescription(s string) *ProviderRegisterDataUpdateOne {
	prduo.mutation.SetDescription(s)
	return prduo
}

// SetLiveInterval sets the "liveInterval" field.
func (prduo *ProviderRegisterDataUpdateOne) SetLiveInterval(i int) *ProviderRegisterDataUpdateOne {
	prduo.mutation.ResetLiveInterval()
	prduo.mutation.SetLiveInterval(i)
	return prduo
}

// AddLiveInterval adds i to the "liveInterval" field.
func (prduo *ProviderRegisterDataUpdateOne) AddLiveInterval(i int) *ProviderRegisterDataUpdateOne {
	prduo.mutation.AddLiveInterval(i)
	return prduo
}

// SetLiveTimeout sets the "liveTimeout" field.
func (prduo *ProviderRegisterDataUpdateOne) SetLiveTimeout(i int) *ProviderRegisterDataUpdateOne {
	prduo.mutation.ResetLiveTimeout()
	prduo.mutation.SetLiveTimeout(i)
	return prduo
}

// AddLiveTimeout adds i to the "liveTimeout" field.
func (prduo *ProviderRegisterDataUpdateOne) AddLiveTimeout(i int) *ProviderRegisterDataUpdateOne {
	prduo.mutation.AddLiveTimeout(i)
	return prduo
}

// Mutation returns the ProviderRegisterDataMutation object of the builder.
func (prduo *ProviderRegisterDataUpdateOne) Mutation() *ProviderRegisterDataMutation {
	return prduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (prduo *ProviderRegisterDataUpdateOne) Select(field string, fields ...string) *ProviderRegisterDataUpdateOne {
	prduo.fields = append([]string{field}, fields...)
	return prduo
}

// Save executes the query and returns the updated ProviderRegisterData entity.
func (prduo *ProviderRegisterDataUpdateOne) Save(ctx context.Context) (*ProviderRegisterData, error) {
	var (
		err  error
		node *ProviderRegisterData
	)
	if len(prduo.hooks) == 0 {
		node, err = prduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProviderRegisterDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prduo.mutation = mutation
			node, err = prduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(prduo.hooks) - 1; i >= 0; i-- {
			if prduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, prduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProviderRegisterData)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProviderRegisterDataMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (prduo *ProviderRegisterDataUpdateOne) SaveX(ctx context.Context) *ProviderRegisterData {
	node, err := prduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (prduo *ProviderRegisterDataUpdateOne) Exec(ctx context.Context) error {
	_, err := prduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prduo *ProviderRegisterDataUpdateOne) ExecX(ctx context.Context) {
	if err := prduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (prduo *ProviderRegisterDataUpdateOne) sqlSave(ctx context.Context) (_node *ProviderRegisterData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   providerregisterdata.Table,
			Columns: providerregisterdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: providerregisterdata.FieldID,
			},
		},
	}
	id, ok := prduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProviderRegisterData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := prduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, providerregisterdata.FieldID)
		for _, f := range fields {
			if !providerregisterdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != providerregisterdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := prduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := prduo.mutation.Name(); ok {
		_spec.SetField(providerregisterdata.FieldName, field.TypeString, value)
	}
	if value, ok := prduo.mutation.Port(); ok {
		_spec.SetField(providerregisterdata.FieldPort, field.TypeString, value)
	}
	if value, ok := prduo.mutation.Address(); ok {
		_spec.SetField(providerregisterdata.FieldAddress, field.TypeString, value)
	}
	if value, ok := prduo.mutation.Description(); ok {
		_spec.SetField(providerregisterdata.FieldDescription, field.TypeString, value)
	}
	if value, ok := prduo.mutation.LiveInterval(); ok {
		_spec.SetField(providerregisterdata.FieldLiveInterval, field.TypeInt, value)
	}
	if value, ok := prduo.mutation.AddedLiveInterval(); ok {
		_spec.AddField(providerregisterdata.FieldLiveInterval, field.TypeInt, value)
	}
	if value, ok := prduo.mutation.LiveTimeout(); ok {
		_spec.SetField(providerregisterdata.FieldLiveTimeout, field.TypeInt, value)
	}
	if value, ok := prduo.mutation.AddedLiveTimeout(); ok {
		_spec.AddField(providerregisterdata.FieldLiveTimeout, field.TypeInt, value)
	}
	_node = &ProviderRegisterData{config: prduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, prduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerregisterdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
