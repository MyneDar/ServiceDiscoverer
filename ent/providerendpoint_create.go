// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"servicediscoverer/ent/endpointdata"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProviderEndpointCreate is the builder for creating a ProviderEndpoint entity.
type ProviderEndpointCreate struct {
	config
	mutation *ProviderEndpointMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pec *ProviderEndpointCreate) SetName(s string) *ProviderEndpointCreate {
	pec.mutation.SetName(s)
	return pec
}

// SetPath sets the "path" field.
func (pec *ProviderEndpointCreate) SetPath(s string) *ProviderEndpointCreate {
	pec.mutation.SetPath(s)
	return pec
}

// SetType sets the "type" field.
func (pec *ProviderEndpointCreate) SetType(s string) *ProviderEndpointCreate {
	pec.mutation.SetType(s)
	return pec
}

// SetDescription sets the "description" field.
func (pec *ProviderEndpointCreate) SetDescription(s string) *ProviderEndpointCreate {
	pec.mutation.SetDescription(s)
	return pec
}

// SetID sets the "id" field.
func (pec *ProviderEndpointCreate) SetID(i int) *ProviderEndpointCreate {
	pec.mutation.SetID(i)
	return pec
}

// AddRequiredDatumIDs adds the "required_data" edge to the EndpointData entity by IDs.
func (pec *ProviderEndpointCreate) AddRequiredDatumIDs(ids ...int) *ProviderEndpointCreate {
	pec.mutation.AddRequiredDatumIDs(ids...)
	return pec
}

// AddRequiredData adds the "required_data" edges to the EndpointData entity.
func (pec *ProviderEndpointCreate) AddRequiredData(e ...*EndpointData) *ProviderEndpointCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pec.AddRequiredDatumIDs(ids...)
}

// AddProvidedDatumIDs adds the "provided_data" edge to the EndpointData entity by IDs.
func (pec *ProviderEndpointCreate) AddProvidedDatumIDs(ids ...int) *ProviderEndpointCreate {
	pec.mutation.AddProvidedDatumIDs(ids...)
	return pec
}

// AddProvidedData adds the "provided_data" edges to the EndpointData entity.
func (pec *ProviderEndpointCreate) AddProvidedData(e ...*EndpointData) *ProviderEndpointCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pec.AddProvidedDatumIDs(ids...)
}

// SetProviderID sets the "provider" edge to the ProviderRegisterData entity by ID.
func (pec *ProviderEndpointCreate) SetProviderID(id int) *ProviderEndpointCreate {
	pec.mutation.SetProviderID(id)
	return pec
}

// SetNillableProviderID sets the "provider" edge to the ProviderRegisterData entity by ID if the given value is not nil.
func (pec *ProviderEndpointCreate) SetNillableProviderID(id *int) *ProviderEndpointCreate {
	if id != nil {
		pec = pec.SetProviderID(*id)
	}
	return pec
}

// SetProvider sets the "provider" edge to the ProviderRegisterData entity.
func (pec *ProviderEndpointCreate) SetProvider(p *ProviderRegisterData) *ProviderEndpointCreate {
	return pec.SetProviderID(p.ID)
}

// Mutation returns the ProviderEndpointMutation object of the builder.
func (pec *ProviderEndpointCreate) Mutation() *ProviderEndpointMutation {
	return pec.mutation
}

// Save creates the ProviderEndpoint in the database.
func (pec *ProviderEndpointCreate) Save(ctx context.Context) (*ProviderEndpoint, error) {
	return withHooks[*ProviderEndpoint, ProviderEndpointMutation](ctx, pec.sqlSave, pec.mutation, pec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pec *ProviderEndpointCreate) SaveX(ctx context.Context) *ProviderEndpoint {
	v, err := pec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pec *ProviderEndpointCreate) Exec(ctx context.Context) error {
	_, err := pec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pec *ProviderEndpointCreate) ExecX(ctx context.Context) {
	if err := pec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pec *ProviderEndpointCreate) check() error {
	if _, ok := pec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProviderEndpoint.name"`)}
	}
	if _, ok := pec.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "ProviderEndpoint.path"`)}
	}
	if _, ok := pec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ProviderEndpoint.type"`)}
	}
	if _, ok := pec.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ProviderEndpoint.description"`)}
	}
	return nil
}

func (pec *ProviderEndpointCreate) sqlSave(ctx context.Context) (*ProviderEndpoint, error) {
	if err := pec.check(); err != nil {
		return nil, err
	}
	_node, _spec := pec.createSpec()
	if err := sqlgraph.CreateNode(ctx, pec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pec.mutation.id = &_node.ID
	pec.mutation.done = true
	return _node, nil
}

func (pec *ProviderEndpointCreate) createSpec() (*ProviderEndpoint, *sqlgraph.CreateSpec) {
	var (
		_node = &ProviderEndpoint{config: pec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: providerendpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: providerendpoint.FieldID,
			},
		}
	)
	if id, ok := pec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pec.mutation.Name(); ok {
		_spec.SetField(providerendpoint.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pec.mutation.Path(); ok {
		_spec.SetField(providerendpoint.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := pec.mutation.GetType(); ok {
		_spec.SetField(providerendpoint.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := pec.mutation.Description(); ok {
		_spec.SetField(providerendpoint.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := pec.mutation.RequiredDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerendpoint.RequiredDataTable,
			Columns: []string{providerendpoint.RequiredDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: endpointdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pec.mutation.ProvidedDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerendpoint.ProvidedDataTable,
			Columns: []string{providerendpoint.ProvidedDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: endpointdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pec.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerendpoint.ProviderTable,
			Columns: []string{providerendpoint.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: providerregisterdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provider_register_data_endpoints = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProviderEndpointCreateBulk is the builder for creating many ProviderEndpoint entities in bulk.
type ProviderEndpointCreateBulk struct {
	config
	builders []*ProviderEndpointCreate
}

// Save creates the ProviderEndpoint entities in the database.
func (pecb *ProviderEndpointCreateBulk) Save(ctx context.Context) ([]*ProviderEndpoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pecb.builders))
	nodes := make([]*ProviderEndpoint, len(pecb.builders))
	mutators := make([]Mutator, len(pecb.builders))
	for i := range pecb.builders {
		func(i int, root context.Context) {
			builder := pecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderEndpointMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pecb *ProviderEndpointCreateBulk) SaveX(ctx context.Context) []*ProviderEndpoint {
	v, err := pecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pecb *ProviderEndpointCreateBulk) Exec(ctx context.Context) error {
	_, err := pecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pecb *ProviderEndpointCreateBulk) ExecX(ctx context.Context) {
	if err := pecb.Exec(ctx); err != nil {
		panic(err)
	}
}
