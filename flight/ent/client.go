// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"flookybooky/flight/ent/migrate"

	"flookybooky/flight/ent/flight"
	"flookybooky/flight/ent/place"
	"flookybooky/flight/ent/seat"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Flight is the client for interacting with the Flight builders.
	Flight *FlightClient
	// Place is the client for interacting with the Place builders.
	Place *PlaceClient
	// Seat is the client for interacting with the Seat builders.
	Seat *SeatClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Flight = NewFlightClient(c.config)
	c.Place = NewPlaceClient(c.config)
	c.Seat = NewSeatClient(c.config)
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
		ctx:    ctx,
		config: cfg,
		Flight: NewFlightClient(cfg),
		Place:  NewPlaceClient(cfg),
		Seat:   NewSeatClient(cfg),
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
		ctx:    ctx,
		config: cfg,
		Flight: NewFlightClient(cfg),
		Place:  NewPlaceClient(cfg),
		Seat:   NewSeatClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Flight.
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
	c.Flight.Use(hooks...)
	c.Place.Use(hooks...)
	c.Seat.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Flight.Intercept(interceptors...)
	c.Place.Intercept(interceptors...)
	c.Seat.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *FlightMutation:
		return c.Flight.mutate(ctx, m)
	case *PlaceMutation:
		return c.Place.mutate(ctx, m)
	case *SeatMutation:
		return c.Seat.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// FlightClient is a client for the Flight schema.
type FlightClient struct {
	config
}

// NewFlightClient returns a client for the Flight from the given config.
func NewFlightClient(c config) *FlightClient {
	return &FlightClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flight.Hooks(f(g(h())))`.
func (c *FlightClient) Use(hooks ...Hook) {
	c.hooks.Flight = append(c.hooks.Flight, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `flight.Intercept(f(g(h())))`.
func (c *FlightClient) Intercept(interceptors ...Interceptor) {
	c.inters.Flight = append(c.inters.Flight, interceptors...)
}

// Create returns a builder for creating a Flight entity.
func (c *FlightClient) Create() *FlightCreate {
	mutation := newFlightMutation(c.config, OpCreate)
	return &FlightCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Flight entities.
func (c *FlightClient) CreateBulk(builders ...*FlightCreate) *FlightCreateBulk {
	return &FlightCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Flight.
func (c *FlightClient) Update() *FlightUpdate {
	mutation := newFlightMutation(c.config, OpUpdate)
	return &FlightUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlightClient) UpdateOne(f *Flight) *FlightUpdateOne {
	mutation := newFlightMutation(c.config, OpUpdateOne, withFlight(f))
	return &FlightUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlightClient) UpdateOneID(id int) *FlightUpdateOne {
	mutation := newFlightMutation(c.config, OpUpdateOne, withFlightID(id))
	return &FlightUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Flight.
func (c *FlightClient) Delete() *FlightDelete {
	mutation := newFlightMutation(c.config, OpDelete)
	return &FlightDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FlightClient) DeleteOne(f *Flight) *FlightDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FlightClient) DeleteOneID(id int) *FlightDeleteOne {
	builder := c.Delete().Where(flight.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlightDeleteOne{builder}
}

// Query returns a query builder for Flight.
func (c *FlightClient) Query() *FlightQuery {
	return &FlightQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFlight},
		inters: c.Interceptors(),
	}
}

// Get returns a Flight entity by its id.
func (c *FlightClient) Get(ctx context.Context, id int) (*Flight, error) {
	return c.Query().Where(flight.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlightClient) GetX(ctx context.Context, id int) *Flight {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySeats queries the seats edge of a Flight.
func (c *FlightClient) QuerySeats(f *Flight) *SeatQuery {
	query := (&SeatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flight.Table, flight.FieldID, id),
			sqlgraph.To(seat.Table, seat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flight.SeatsTable, flight.SeatsColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlightClient) Hooks() []Hook {
	return c.hooks.Flight
}

// Interceptors returns the client interceptors.
func (c *FlightClient) Interceptors() []Interceptor {
	return c.inters.Flight
}

func (c *FlightClient) mutate(ctx context.Context, m *FlightMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FlightCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FlightUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FlightUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FlightDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Flight mutation op: %q", m.Op())
	}
}

// PlaceClient is a client for the Place schema.
type PlaceClient struct {
	config
}

// NewPlaceClient returns a client for the Place from the given config.
func NewPlaceClient(c config) *PlaceClient {
	return &PlaceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `place.Hooks(f(g(h())))`.
func (c *PlaceClient) Use(hooks ...Hook) {
	c.hooks.Place = append(c.hooks.Place, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `place.Intercept(f(g(h())))`.
func (c *PlaceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Place = append(c.inters.Place, interceptors...)
}

// Create returns a builder for creating a Place entity.
func (c *PlaceClient) Create() *PlaceCreate {
	mutation := newPlaceMutation(c.config, OpCreate)
	return &PlaceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Place entities.
func (c *PlaceClient) CreateBulk(builders ...*PlaceCreate) *PlaceCreateBulk {
	return &PlaceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Place.
func (c *PlaceClient) Update() *PlaceUpdate {
	mutation := newPlaceMutation(c.config, OpUpdate)
	return &PlaceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlaceClient) UpdateOne(pl *Place) *PlaceUpdateOne {
	mutation := newPlaceMutation(c.config, OpUpdateOne, withPlace(pl))
	return &PlaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlaceClient) UpdateOneID(id int) *PlaceUpdateOne {
	mutation := newPlaceMutation(c.config, OpUpdateOne, withPlaceID(id))
	return &PlaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Place.
func (c *PlaceClient) Delete() *PlaceDelete {
	mutation := newPlaceMutation(c.config, OpDelete)
	return &PlaceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PlaceClient) DeleteOne(pl *Place) *PlaceDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PlaceClient) DeleteOneID(id int) *PlaceDeleteOne {
	builder := c.Delete().Where(place.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlaceDeleteOne{builder}
}

// Query returns a query builder for Place.
func (c *PlaceClient) Query() *PlaceQuery {
	return &PlaceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePlace},
		inters: c.Interceptors(),
	}
}

// Get returns a Place entity by its id.
func (c *PlaceClient) Get(ctx context.Context, id int) (*Place, error) {
	return c.Query().Where(place.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlaceClient) GetX(ctx context.Context, id int) *Place {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlaceClient) Hooks() []Hook {
	return c.hooks.Place
}

// Interceptors returns the client interceptors.
func (c *PlaceClient) Interceptors() []Interceptor {
	return c.inters.Place
}

func (c *PlaceClient) mutate(ctx context.Context, m *PlaceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PlaceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PlaceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PlaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PlaceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Place mutation op: %q", m.Op())
	}
}

// SeatClient is a client for the Seat schema.
type SeatClient struct {
	config
}

// NewSeatClient returns a client for the Seat from the given config.
func NewSeatClient(c config) *SeatClient {
	return &SeatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `seat.Hooks(f(g(h())))`.
func (c *SeatClient) Use(hooks ...Hook) {
	c.hooks.Seat = append(c.hooks.Seat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `seat.Intercept(f(g(h())))`.
func (c *SeatClient) Intercept(interceptors ...Interceptor) {
	c.inters.Seat = append(c.inters.Seat, interceptors...)
}

// Create returns a builder for creating a Seat entity.
func (c *SeatClient) Create() *SeatCreate {
	mutation := newSeatMutation(c.config, OpCreate)
	return &SeatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Seat entities.
func (c *SeatClient) CreateBulk(builders ...*SeatCreate) *SeatCreateBulk {
	return &SeatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Seat.
func (c *SeatClient) Update() *SeatUpdate {
	mutation := newSeatMutation(c.config, OpUpdate)
	return &SeatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeatClient) UpdateOne(s *Seat) *SeatUpdateOne {
	mutation := newSeatMutation(c.config, OpUpdateOne, withSeat(s))
	return &SeatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeatClient) UpdateOneID(id int) *SeatUpdateOne {
	mutation := newSeatMutation(c.config, OpUpdateOne, withSeatID(id))
	return &SeatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Seat.
func (c *SeatClient) Delete() *SeatDelete {
	mutation := newSeatMutation(c.config, OpDelete)
	return &SeatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SeatClient) DeleteOne(s *Seat) *SeatDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SeatClient) DeleteOneID(id int) *SeatDeleteOne {
	builder := c.Delete().Where(seat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeatDeleteOne{builder}
}

// Query returns a query builder for Seat.
func (c *SeatClient) Query() *SeatQuery {
	return &SeatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSeat},
		inters: c.Interceptors(),
	}
}

// Get returns a Seat entity by its id.
func (c *SeatClient) Get(ctx context.Context, id int) (*Seat, error) {
	return c.Query().Where(seat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeatClient) GetX(ctx context.Context, id int) *Seat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SeatClient) Hooks() []Hook {
	return c.hooks.Seat
}

// Interceptors returns the client interceptors.
func (c *SeatClient) Interceptors() []Interceptor {
	return c.inters.Seat
}

func (c *SeatClient) mutate(ctx context.Context, m *SeatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SeatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SeatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SeatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SeatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Seat mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Flight, Place, Seat []ent.Hook
	}
	inters struct {
		Flight, Place, Seat []ent.Interceptor
	}
)
