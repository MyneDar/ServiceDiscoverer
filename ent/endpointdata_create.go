// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"servicediscoverer/ent/endpointdata"
	"servicediscoverer/ent/providerendpoint"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EndpointDataCreate is the builder for creating a EndpointData entity.
type EndpointDataCreate struct {
	config
	mutation *EndpointDataMutation
	hooks    []Hook
}

// SetDataName sets the "dataName" field.
func (edc *EndpointDataCreate) SetDataName(s string) *EndpointDataCreate {
	edc.mutation.SetDataName(s)
	return edc
}

// SetDescription sets the "description" field.
func (edc *EndpointDataCreate) SetDescription(s string) *EndpointDataCreate {
	edc.mutation.SetDescription(s)
	return edc
}

// SetType sets the "type" field.
func (edc *EndpointDataCreate) SetType(s string) *EndpointDataCreate {
	edc.mutation.SetType(s)
	return edc
}

// SetID sets the "id" field.
func (edc *EndpointDataCreate) SetID(i int) *EndpointDataCreate {
	edc.mutation.SetID(i)
	return edc
}

// SetEndpointRequiredID sets the "endpoint_required" edge to the ProviderEndpoint entity by ID.
func (edc *EndpointDataCreate) SetEndpointRequiredID(id int) *EndpointDataCreate {
	edc.mutation.SetEndpointRequiredID(id)
	return edc
}

// SetNillableEndpointRequiredID sets the "endpoint_required" edge to the ProviderEndpoint entity by ID if the given value is not nil.
func (edc *EndpointDataCreate) SetNillableEndpointRequiredID(id *int) *EndpointDataCreate {
	if id != nil {
		edc = edc.SetEndpointRequiredID(*id)
	}
	return edc
}

// SetEndpointRequired sets the "endpoint_required" edge to the ProviderEndpoint entity.
func (edc *EndpointDataCreate) SetEndpointRequired(p *ProviderEndpoint) *EndpointDataCreate {
	return edc.SetEndpointRequiredID(p.ID)
}

// SetEndpointProvidedID sets the "endpoint_provided" edge to the ProviderEndpoint entity by ID.
func (edc *EndpointDataCreate) SetEndpointProvidedID(id int) *EndpointDataCreate {
	edc.mutation.SetEndpointProvidedID(id)
	return edc
}

// SetNillableEndpointProvidedID sets the "endpoint_provided" edge to the ProviderEndpoint entity by ID if the given value is not nil.
func (edc *EndpointDataCreate) SetNillableEndpointProvidedID(id *int) *EndpointDataCreate {
	if id != nil {
		edc = edc.SetEndpointProvidedID(*id)
	}
	return edc
}

// SetEndpointProvided sets the "endpoint_provided" edge to the ProviderEndpoint entity.
func (edc *EndpointDataCreate) SetEndpointProvided(p *ProviderEndpoint) *EndpointDataCreate {
	return edc.SetEndpointProvidedID(p.ID)
}

// Mutation returns the EndpointDataMutation object of the builder.
func (edc *EndpointDataCreate) Mutation() *EndpointDataMutation {
	return edc.mutation
}

// Save creates the EndpointData in the database.
func (edc *EndpointDataCreate) Save(ctx context.Context) (*EndpointData, error) {
	return withHooks[*EndpointData, EndpointDataMutation](ctx, edc.sqlSave, edc.mutation, edc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (edc *EndpointDataCreate) SaveX(ctx context.Context) *EndpointData {
	v, err := edc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (edc *EndpointDataCreate) Exec(ctx context.Context) error {
	_, err := edc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (edc *EndpointDataCreate) ExecX(ctx context.Context) {
	if err := edc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (edc *EndpointDataCreate) check() error {
	if _, ok := edc.mutation.DataName(); !ok {
		return &ValidationError{Name: "dataName", err: errors.New(`ent: missing required field "EndpointData.dataName"`)}
	}
	if _, ok := edc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "EndpointData.description"`)}
	}
	if _, ok := edc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "EndpointData.type"`)}
	}
	return nil
}

func (edc *EndpointDataCreate) sqlSave(ctx context.Context) (*EndpointData, error) {
	if err := edc.check(); err != nil {
		return nil, err
	}
	_node, _spec := edc.createSpec()
	if err := sqlgraph.CreateNode(ctx, edc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	edc.mutation.id = &_node.ID
	edc.mutation.done = true
	return _node, nil
}

func (edc *EndpointDataCreate) createSpec() (*EndpointData, *sqlgraph.CreateSpec) {
	var (
		_node = &EndpointData{config: edc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: endpointdata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endpointdata.FieldID,
			},
		}
	)
	if id, ok := edc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := edc.mutation.DataName(); ok {
		_spec.SetField(endpointdata.FieldDataName, field.TypeString, value)
		_node.DataName = value
	}
	if value, ok := edc.mutation.Description(); ok {
		_spec.SetField(endpointdata.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := edc.mutation.GetType(); ok {
		_spec.SetField(endpointdata.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := edc.mutation.EndpointRequiredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   endpointdata.EndpointRequiredTable,
			Columns: []string{endpointdata.EndpointRequiredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: providerendpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provider_endpoint_required_data = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := edc.mutation.EndpointProvidedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   endpointdata.EndpointProvidedTable,
			Columns: []string{endpointdata.EndpointProvidedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: providerendpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provider_endpoint_provided_data = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EndpointDataCreateBulk is the builder for creating many EndpointData entities in bulk.
type EndpointDataCreateBulk struct {
	config
	builders []*EndpointDataCreate
}

// Save creates the EndpointData entities in the database.
func (edcb *EndpointDataCreateBulk) Save(ctx context.Context) ([]*EndpointData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(edcb.builders))
	nodes := make([]*EndpointData, len(edcb.builders))
	mutators := make([]Mutator, len(edcb.builders))
	for i := range edcb.builders {
		func(i int, root context.Context) {
			builder := edcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EndpointDataMutation)
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
					_, err = mutators[i+1].Mutate(root, edcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, edcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, edcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (edcb *EndpointDataCreateBulk) SaveX(ctx context.Context) []*EndpointData {
	v, err := edcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (edcb *EndpointDataCreateBulk) Exec(ctx context.Context) error {
	_, err := edcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (edcb *EndpointDataCreateBulk) ExecX(ctx context.Context) {
	if err := edcb.Exec(ctx); err != nil {
		panic(err)
	}
}
