// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/carbonable/carbonable-launchpad-backend/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/carbonable/carbonable-launchpad-backend/ent/launchpad"
	"github.com/carbonable/carbonable-launchpad-backend/ent/mint"
	"github.com/carbonable/carbonable-launchpad-backend/ent/project"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Launchpad is the client for interacting with the Launchpad builders.
	Launchpad *LaunchpadClient
	// Mint is the client for interacting with the Mint builders.
	Mint *MintClient
	// Project is the client for interacting with the Project builders.
	Project *ProjectClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Launchpad = NewLaunchpadClient(c.config)
	c.Mint = NewMintClient(c.config)
	c.Project = NewProjectClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Launchpad: NewLaunchpadClient(cfg),
		Mint:      NewMintClient(cfg),
		Project:   NewProjectClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Launchpad: NewLaunchpadClient(cfg),
		Mint:      NewMintClient(cfg),
		Project:   NewProjectClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Launchpad.
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
	c.Launchpad.Use(hooks...)
	c.Mint.Use(hooks...)
	c.Project.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Launchpad.Intercept(interceptors...)
	c.Mint.Intercept(interceptors...)
	c.Project.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LaunchpadMutation:
		return c.Launchpad.mutate(ctx, m)
	case *MintMutation:
		return c.Mint.mutate(ctx, m)
	case *ProjectMutation:
		return c.Project.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LaunchpadClient is a client for the Launchpad schema.
type LaunchpadClient struct {
	config
}

// NewLaunchpadClient returns a client for the Launchpad from the given config.
func NewLaunchpadClient(c config) *LaunchpadClient {
	return &LaunchpadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `launchpad.Hooks(f(g(h())))`.
func (c *LaunchpadClient) Use(hooks ...Hook) {
	c.hooks.Launchpad = append(c.hooks.Launchpad, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `launchpad.Intercept(f(g(h())))`.
func (c *LaunchpadClient) Intercept(interceptors ...Interceptor) {
	c.inters.Launchpad = append(c.inters.Launchpad, interceptors...)
}

// Create returns a builder for creating a Launchpad entity.
func (c *LaunchpadClient) Create() *LaunchpadCreate {
	mutation := newLaunchpadMutation(c.config, OpCreate)
	return &LaunchpadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Launchpad entities.
func (c *LaunchpadClient) CreateBulk(builders ...*LaunchpadCreate) *LaunchpadCreateBulk {
	return &LaunchpadCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LaunchpadClient) MapCreateBulk(slice any, setFunc func(*LaunchpadCreate, int)) *LaunchpadCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LaunchpadCreateBulk{err: fmt.Errorf("calling to LaunchpadClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LaunchpadCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LaunchpadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Launchpad.
func (c *LaunchpadClient) Update() *LaunchpadUpdate {
	mutation := newLaunchpadMutation(c.config, OpUpdate)
	return &LaunchpadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LaunchpadClient) UpdateOne(l *Launchpad) *LaunchpadUpdateOne {
	mutation := newLaunchpadMutation(c.config, OpUpdateOne, withLaunchpad(l))
	return &LaunchpadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LaunchpadClient) UpdateOneID(id int) *LaunchpadUpdateOne {
	mutation := newLaunchpadMutation(c.config, OpUpdateOne, withLaunchpadID(id))
	return &LaunchpadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Launchpad.
func (c *LaunchpadClient) Delete() *LaunchpadDelete {
	mutation := newLaunchpadMutation(c.config, OpDelete)
	return &LaunchpadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LaunchpadClient) DeleteOne(l *Launchpad) *LaunchpadDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LaunchpadClient) DeleteOneID(id int) *LaunchpadDeleteOne {
	builder := c.Delete().Where(launchpad.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LaunchpadDeleteOne{builder}
}

// Query returns a query builder for Launchpad.
func (c *LaunchpadClient) Query() *LaunchpadQuery {
	return &LaunchpadQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLaunchpad},
		inters: c.Interceptors(),
	}
}

// Get returns a Launchpad entity by its id.
func (c *LaunchpadClient) Get(ctx context.Context, id int) (*Launchpad, error) {
	return c.Query().Where(launchpad.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LaunchpadClient) GetX(ctx context.Context, id int) *Launchpad {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProject queries the project edge of a Launchpad.
func (c *LaunchpadClient) QueryProject(l *Launchpad) *ProjectQuery {
	query := (&ProjectClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(launchpad.Table, launchpad.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, launchpad.ProjectTable, launchpad.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LaunchpadClient) Hooks() []Hook {
	return c.hooks.Launchpad
}

// Interceptors returns the client interceptors.
func (c *LaunchpadClient) Interceptors() []Interceptor {
	return c.inters.Launchpad
}

func (c *LaunchpadClient) mutate(ctx context.Context, m *LaunchpadMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LaunchpadCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LaunchpadUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LaunchpadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LaunchpadDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Launchpad mutation op: %q", m.Op())
	}
}

// MintClient is a client for the Mint schema.
type MintClient struct {
	config
}

// NewMintClient returns a client for the Mint from the given config.
func NewMintClient(c config) *MintClient {
	return &MintClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mint.Hooks(f(g(h())))`.
func (c *MintClient) Use(hooks ...Hook) {
	c.hooks.Mint = append(c.hooks.Mint, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `mint.Intercept(f(g(h())))`.
func (c *MintClient) Intercept(interceptors ...Interceptor) {
	c.inters.Mint = append(c.inters.Mint, interceptors...)
}

// Create returns a builder for creating a Mint entity.
func (c *MintClient) Create() *MintCreate {
	mutation := newMintMutation(c.config, OpCreate)
	return &MintCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Mint entities.
func (c *MintClient) CreateBulk(builders ...*MintCreate) *MintCreateBulk {
	return &MintCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MintClient) MapCreateBulk(slice any, setFunc func(*MintCreate, int)) *MintCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MintCreateBulk{err: fmt.Errorf("calling to MintClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MintCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MintCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Mint.
func (c *MintClient) Update() *MintUpdate {
	mutation := newMintMutation(c.config, OpUpdate)
	return &MintUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MintClient) UpdateOne(m *Mint) *MintUpdateOne {
	mutation := newMintMutation(c.config, OpUpdateOne, withMint(m))
	return &MintUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MintClient) UpdateOneID(id int) *MintUpdateOne {
	mutation := newMintMutation(c.config, OpUpdateOne, withMintID(id))
	return &MintUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Mint.
func (c *MintClient) Delete() *MintDelete {
	mutation := newMintMutation(c.config, OpDelete)
	return &MintDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MintClient) DeleteOne(m *Mint) *MintDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MintClient) DeleteOneID(id int) *MintDeleteOne {
	builder := c.Delete().Where(mint.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MintDeleteOne{builder}
}

// Query returns a query builder for Mint.
func (c *MintClient) Query() *MintQuery {
	return &MintQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMint},
		inters: c.Interceptors(),
	}
}

// Get returns a Mint entity by its id.
func (c *MintClient) Get(ctx context.Context, id int) (*Mint, error) {
	return c.Query().Where(mint.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MintClient) GetX(ctx context.Context, id int) *Mint {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProject queries the project edge of a Mint.
func (c *MintClient) QueryProject(m *Mint) *ProjectQuery {
	query := (&ProjectClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mint.Table, mint.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, mint.ProjectTable, mint.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MintClient) Hooks() []Hook {
	return c.hooks.Mint
}

// Interceptors returns the client interceptors.
func (c *MintClient) Interceptors() []Interceptor {
	return c.inters.Mint
}

func (c *MintClient) mutate(ctx context.Context, m *MintMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MintCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MintUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MintUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MintDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Mint mutation op: %q", m.Op())
	}
}

// ProjectClient is a client for the Project schema.
type ProjectClient struct {
	config
}

// NewProjectClient returns a client for the Project from the given config.
func NewProjectClient(c config) *ProjectClient {
	return &ProjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `project.Hooks(f(g(h())))`.
func (c *ProjectClient) Use(hooks ...Hook) {
	c.hooks.Project = append(c.hooks.Project, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `project.Intercept(f(g(h())))`.
func (c *ProjectClient) Intercept(interceptors ...Interceptor) {
	c.inters.Project = append(c.inters.Project, interceptors...)
}

// Create returns a builder for creating a Project entity.
func (c *ProjectClient) Create() *ProjectCreate {
	mutation := newProjectMutation(c.config, OpCreate)
	return &ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Project entities.
func (c *ProjectClient) CreateBulk(builders ...*ProjectCreate) *ProjectCreateBulk {
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ProjectClient) MapCreateBulk(slice any, setFunc func(*ProjectCreate, int)) *ProjectCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ProjectCreateBulk{err: fmt.Errorf("calling to ProjectClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ProjectCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Project.
func (c *ProjectClient) Update() *ProjectUpdate {
	mutation := newProjectMutation(c.config, OpUpdate)
	return &ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProjectClient) UpdateOne(pr *Project) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProject(pr))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProjectClient) UpdateOneID(id int) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProjectID(id))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Project.
func (c *ProjectClient) Delete() *ProjectDelete {
	mutation := newProjectMutation(c.config, OpDelete)
	return &ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProjectClient) DeleteOne(pr *Project) *ProjectDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProjectClient) DeleteOneID(id int) *ProjectDeleteOne {
	builder := c.Delete().Where(project.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProjectDeleteOne{builder}
}

// Query returns a query builder for Project.
func (c *ProjectClient) Query() *ProjectQuery {
	return &ProjectQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProject},
		inters: c.Interceptors(),
	}
}

// Get returns a Project entity by its id.
func (c *ProjectClient) Get(ctx context.Context, id int) (*Project, error) {
	return c.Query().Where(project.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProjectClient) GetX(ctx context.Context, id int) *Project {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMint queries the mint edge of a Project.
func (c *ProjectClient) QueryMint(pr *Project) *MintQuery {
	query := (&MintClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(mint.Table, mint.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, project.MintTable, project.MintColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLaunchpad queries the launchpad edge of a Project.
func (c *ProjectClient) QueryLaunchpad(pr *Project) *LaunchpadQuery {
	query := (&LaunchpadClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(launchpad.Table, launchpad.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, project.LaunchpadTable, project.LaunchpadColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProjectClient) Hooks() []Hook {
	return c.hooks.Project
}

// Interceptors returns the client interceptors.
func (c *ProjectClient) Interceptors() []Interceptor {
	return c.inters.Project
}

func (c *ProjectClient) mutate(ctx context.Context, m *ProjectMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Project mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Launchpad, Mint, Project []ent.Hook
	}
	inters struct {
		Launchpad, Mint, Project []ent.Interceptor
	}
)