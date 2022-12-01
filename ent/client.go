// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"servicediscoverer/ent/migrate"

	"servicediscoverer/ent/endpointdata"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// EndpointData is the client for interacting with the EndpointData builders.
	EndpointData *EndpointDataClient
	// ProviderEndpoint is the client for interacting with the ProviderEndpoint builders.
	ProviderEndpoint *ProviderEndpointClient
	// ProviderRegisterData is the client for interacting with the ProviderRegisterData builders.
	ProviderRegisterData *ProviderRegisterDataClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.EndpointData = NewEndpointDataClient(c.config)
	c.ProviderEndpoint = NewProviderEndpointClient(c.config)
	c.ProviderRegisterData = NewProviderRegisterDataClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                  ctx,
		config:               cfg,
		EndpointData:         NewEndpointDataClient(cfg),
		ProviderEndpoint:     NewProviderEndpointClient(cfg),
		ProviderRegisterData: NewProviderRegisterDataClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                  ctx,
		config:               cfg,
		EndpointData:         NewEndpointDataClient(cfg),
		ProviderEndpoint:     NewProviderEndpointClient(cfg),
		ProviderRegisterData: NewProviderRegisterDataClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		EndpointData.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.EndpointData.Use(hooks...)
	c.ProviderEndpoint.Use(hooks...)
	c.ProviderRegisterData.Use(hooks...)
}

// EndpointDataClient is a client for the EndpointData schema.
type EndpointDataClient struct {
	config
}

// NewEndpointDataClient returns a client for the EndpointData from the given config.
func NewEndpointDataClient(c config) *EndpointDataClient {
	return &EndpointDataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `endpointdata.Hooks(f(g(h())))`.
func (c *EndpointDataClient) Use(hooks ...Hook) {
	c.hooks.EndpointData = append(c.hooks.EndpointData, hooks...)
}

// Create returns a builder for creating a EndpointData entity.
func (c *EndpointDataClient) Create() *EndpointDataCreate {
	mutation := newEndpointDataMutation(c.config, OpCreate)
	return &EndpointDataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EndpointData entities.
func (c *EndpointDataClient) CreateBulk(builders ...*EndpointDataCreate) *EndpointDataCreateBulk {
	return &EndpointDataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EndpointData.
func (c *EndpointDataClient) Update() *EndpointDataUpdate {
	mutation := newEndpointDataMutation(c.config, OpUpdate)
	return &EndpointDataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EndpointDataClient) UpdateOne(ed *EndpointData) *EndpointDataUpdateOne {
	mutation := newEndpointDataMutation(c.config, OpUpdateOne, withEndpointData(ed))
	return &EndpointDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EndpointDataClient) UpdateOneID(id int) *EndpointDataUpdateOne {
	mutation := newEndpointDataMutation(c.config, OpUpdateOne, withEndpointDataID(id))
	return &EndpointDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EndpointData.
func (c *EndpointDataClient) Delete() *EndpointDataDelete {
	mutation := newEndpointDataMutation(c.config, OpDelete)
	return &EndpointDataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EndpointDataClient) DeleteOne(ed *EndpointData) *EndpointDataDeleteOne {
	return c.DeleteOneID(ed.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EndpointDataClient) DeleteOneID(id int) *EndpointDataDeleteOne {
	builder := c.Delete().Where(endpointdata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EndpointDataDeleteOne{builder}
}

// Query returns a query builder for EndpointData.
func (c *EndpointDataClient) Query() *EndpointDataQuery {
	return &EndpointDataQuery{
		config: c.config,
	}
}

// Get returns a EndpointData entity by its id.
func (c *EndpointDataClient) Get(ctx context.Context, id int) (*EndpointData, error) {
	return c.Query().Where(endpointdata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EndpointDataClient) GetX(ctx context.Context, id int) *EndpointData {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEndpointRequired queries the endpoint_required edge of a EndpointData.
func (c *EndpointDataClient) QueryEndpointRequired(ed *EndpointData) *ProviderEndpointQuery {
	query := &ProviderEndpointQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ed.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(endpointdata.Table, endpointdata.FieldID, id),
			sqlgraph.To(providerendpoint.Table, providerendpoint.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, endpointdata.EndpointRequiredTable, endpointdata.EndpointRequiredColumn),
		)
		fromV = sqlgraph.Neighbors(ed.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEndpointProvided queries the endpoint_provided edge of a EndpointData.
func (c *EndpointDataClient) QueryEndpointProvided(ed *EndpointData) *ProviderEndpointQuery {
	query := &ProviderEndpointQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ed.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(endpointdata.Table, endpointdata.FieldID, id),
			sqlgraph.To(providerendpoint.Table, providerendpoint.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, endpointdata.EndpointProvidedTable, endpointdata.EndpointProvidedColumn),
		)
		fromV = sqlgraph.Neighbors(ed.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EndpointDataClient) Hooks() []Hook {
	return c.hooks.EndpointData
}

// ProviderEndpointClient is a client for the ProviderEndpoint schema.
type ProviderEndpointClient struct {
	config
}

// NewProviderEndpointClient returns a client for the ProviderEndpoint from the given config.
func NewProviderEndpointClient(c config) *ProviderEndpointClient {
	return &ProviderEndpointClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `providerendpoint.Hooks(f(g(h())))`.
func (c *ProviderEndpointClient) Use(hooks ...Hook) {
	c.hooks.ProviderEndpoint = append(c.hooks.ProviderEndpoint, hooks...)
}

// Create returns a builder for creating a ProviderEndpoint entity.
func (c *ProviderEndpointClient) Create() *ProviderEndpointCreate {
	mutation := newProviderEndpointMutation(c.config, OpCreate)
	return &ProviderEndpointCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ProviderEndpoint entities.
func (c *ProviderEndpointClient) CreateBulk(builders ...*ProviderEndpointCreate) *ProviderEndpointCreateBulk {
	return &ProviderEndpointCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ProviderEndpoint.
func (c *ProviderEndpointClient) Update() *ProviderEndpointUpdate {
	mutation := newProviderEndpointMutation(c.config, OpUpdate)
	return &ProviderEndpointUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProviderEndpointClient) UpdateOne(pe *ProviderEndpoint) *ProviderEndpointUpdateOne {
	mutation := newProviderEndpointMutation(c.config, OpUpdateOne, withProviderEndpoint(pe))
	return &ProviderEndpointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProviderEndpointClient) UpdateOneID(id int) *ProviderEndpointUpdateOne {
	mutation := newProviderEndpointMutation(c.config, OpUpdateOne, withProviderEndpointID(id))
	return &ProviderEndpointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ProviderEndpoint.
func (c *ProviderEndpointClient) Delete() *ProviderEndpointDelete {
	mutation := newProviderEndpointMutation(c.config, OpDelete)
	return &ProviderEndpointDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProviderEndpointClient) DeleteOne(pe *ProviderEndpoint) *ProviderEndpointDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProviderEndpointClient) DeleteOneID(id int) *ProviderEndpointDeleteOne {
	builder := c.Delete().Where(providerendpoint.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProviderEndpointDeleteOne{builder}
}

// Query returns a query builder for ProviderEndpoint.
func (c *ProviderEndpointClient) Query() *ProviderEndpointQuery {
	return &ProviderEndpointQuery{
		config: c.config,
	}
}

// Get returns a ProviderEndpoint entity by its id.
func (c *ProviderEndpointClient) Get(ctx context.Context, id int) (*ProviderEndpoint, error) {
	return c.Query().Where(providerendpoint.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProviderEndpointClient) GetX(ctx context.Context, id int) *ProviderEndpoint {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRequiredData queries the required_data edge of a ProviderEndpoint.
func (c *ProviderEndpointClient) QueryRequiredData(pe *ProviderEndpoint) *EndpointDataQuery {
	query := &EndpointDataQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(providerendpoint.Table, providerendpoint.FieldID, id),
			sqlgraph.To(endpointdata.Table, endpointdata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, providerendpoint.RequiredDataTable, providerendpoint.RequiredDataColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvidedData queries the provided_data edge of a ProviderEndpoint.
func (c *ProviderEndpointClient) QueryProvidedData(pe *ProviderEndpoint) *EndpointDataQuery {
	query := &EndpointDataQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(providerendpoint.Table, providerendpoint.FieldID, id),
			sqlgraph.To(endpointdata.Table, endpointdata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, providerendpoint.ProvidedDataTable, providerendpoint.ProvidedDataColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvider queries the provider edge of a ProviderEndpoint.
func (c *ProviderEndpointClient) QueryProvider(pe *ProviderEndpoint) *ProviderRegisterDataQuery {
	query := &ProviderRegisterDataQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(providerendpoint.Table, providerendpoint.FieldID, id),
			sqlgraph.To(providerregisterdata.Table, providerregisterdata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, providerendpoint.ProviderTable, providerendpoint.ProviderColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProviderEndpointClient) Hooks() []Hook {
	return c.hooks.ProviderEndpoint
}

// ProviderRegisterDataClient is a client for the ProviderRegisterData schema.
type ProviderRegisterDataClient struct {
	config
}

// NewProviderRegisterDataClient returns a client for the ProviderRegisterData from the given config.
func NewProviderRegisterDataClient(c config) *ProviderRegisterDataClient {
	return &ProviderRegisterDataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `providerregisterdata.Hooks(f(g(h())))`.
func (c *ProviderRegisterDataClient) Use(hooks ...Hook) {
	c.hooks.ProviderRegisterData = append(c.hooks.ProviderRegisterData, hooks...)
}

// Create returns a builder for creating a ProviderRegisterData entity.
func (c *ProviderRegisterDataClient) Create() *ProviderRegisterDataCreate {
	mutation := newProviderRegisterDataMutation(c.config, OpCreate)
	return &ProviderRegisterDataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ProviderRegisterData entities.
func (c *ProviderRegisterDataClient) CreateBulk(builders ...*ProviderRegisterDataCreate) *ProviderRegisterDataCreateBulk {
	return &ProviderRegisterDataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ProviderRegisterData.
func (c *ProviderRegisterDataClient) Update() *ProviderRegisterDataUpdate {
	mutation := newProviderRegisterDataMutation(c.config, OpUpdate)
	return &ProviderRegisterDataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProviderRegisterDataClient) UpdateOne(prd *ProviderRegisterData) *ProviderRegisterDataUpdateOne {
	mutation := newProviderRegisterDataMutation(c.config, OpUpdateOne, withProviderRegisterData(prd))
	return &ProviderRegisterDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProviderRegisterDataClient) UpdateOneID(id int) *ProviderRegisterDataUpdateOne {
	mutation := newProviderRegisterDataMutation(c.config, OpUpdateOne, withProviderRegisterDataID(id))
	return &ProviderRegisterDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ProviderRegisterData.
func (c *ProviderRegisterDataClient) Delete() *ProviderRegisterDataDelete {
	mutation := newProviderRegisterDataMutation(c.config, OpDelete)
	return &ProviderRegisterDataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProviderRegisterDataClient) DeleteOne(prd *ProviderRegisterData) *ProviderRegisterDataDeleteOne {
	return c.DeleteOneID(prd.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProviderRegisterDataClient) DeleteOneID(id int) *ProviderRegisterDataDeleteOne {
	builder := c.Delete().Where(providerregisterdata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProviderRegisterDataDeleteOne{builder}
}

// Query returns a query builder for ProviderRegisterData.
func (c *ProviderRegisterDataClient) Query() *ProviderRegisterDataQuery {
	return &ProviderRegisterDataQuery{
		config: c.config,
	}
}

// Get returns a ProviderRegisterData entity by its id.
func (c *ProviderRegisterDataClient) Get(ctx context.Context, id int) (*ProviderRegisterData, error) {
	return c.Query().Where(providerregisterdata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProviderRegisterDataClient) GetX(ctx context.Context, id int) *ProviderRegisterData {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEndpoints queries the endpoints edge of a ProviderRegisterData.
func (c *ProviderRegisterDataClient) QueryEndpoints(prd *ProviderRegisterData) *ProviderEndpointQuery {
	query := &ProviderEndpointQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := prd.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(providerregisterdata.Table, providerregisterdata.FieldID, id),
			sqlgraph.To(providerendpoint.Table, providerendpoint.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, providerregisterdata.EndpointsTable, providerregisterdata.EndpointsColumn),
		)
		fromV = sqlgraph.Neighbors(prd.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProviderRegisterDataClient) Hooks() []Hook {
	return c.hooks.ProviderRegisterData
}
