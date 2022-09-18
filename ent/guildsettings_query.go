// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-relay-conn-bug/ent/guild"
	"github.com/lrstanley/ent-relay-conn-bug/ent/guildsettings"
	"github.com/lrstanley/ent-relay-conn-bug/ent/predicate"
)

// GuildSettingsQuery is the builder for querying GuildSettings entities.
type GuildSettingsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GuildSettings
	withGuild  *GuildQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GuildSettings) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GuildSettingsQuery builder.
func (gsq *GuildSettingsQuery) Where(ps ...predicate.GuildSettings) *GuildSettingsQuery {
	gsq.predicates = append(gsq.predicates, ps...)
	return gsq
}

// Limit adds a limit step to the query.
func (gsq *GuildSettingsQuery) Limit(limit int) *GuildSettingsQuery {
	gsq.limit = &limit
	return gsq
}

// Offset adds an offset step to the query.
func (gsq *GuildSettingsQuery) Offset(offset int) *GuildSettingsQuery {
	gsq.offset = &offset
	return gsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gsq *GuildSettingsQuery) Unique(unique bool) *GuildSettingsQuery {
	gsq.unique = &unique
	return gsq
}

// Order adds an order step to the query.
func (gsq *GuildSettingsQuery) Order(o ...OrderFunc) *GuildSettingsQuery {
	gsq.order = append(gsq.order, o...)
	return gsq
}

// QueryGuild chains the current query on the "guild" edge.
func (gsq *GuildSettingsQuery) QueryGuild() *GuildQuery {
	query := &GuildQuery{config: gsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(guildsettings.Table, guildsettings.FieldID, selector),
			sqlgraph.To(guild.Table, guild.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, guildsettings.GuildTable, guildsettings.GuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GuildSettings entity from the query.
// Returns a *NotFoundError when no GuildSettings was found.
func (gsq *GuildSettingsQuery) First(ctx context.Context) (*GuildSettings, error) {
	nodes, err := gsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{guildsettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gsq *GuildSettingsQuery) FirstX(ctx context.Context) *GuildSettings {
	node, err := gsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GuildSettings ID from the query.
// Returns a *NotFoundError when no GuildSettings ID was found.
func (gsq *GuildSettingsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{guildsettings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gsq *GuildSettingsQuery) FirstIDX(ctx context.Context) int {
	id, err := gsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GuildSettings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GuildSettings entity is found.
// Returns a *NotFoundError when no GuildSettings entities are found.
func (gsq *GuildSettingsQuery) Only(ctx context.Context) (*GuildSettings, error) {
	nodes, err := gsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{guildsettings.Label}
	default:
		return nil, &NotSingularError{guildsettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gsq *GuildSettingsQuery) OnlyX(ctx context.Context) *GuildSettings {
	node, err := gsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GuildSettings ID in the query.
// Returns a *NotSingularError when more than one GuildSettings ID is found.
// Returns a *NotFoundError when no entities are found.
func (gsq *GuildSettingsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{guildsettings.Label}
	default:
		err = &NotSingularError{guildsettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gsq *GuildSettingsQuery) OnlyIDX(ctx context.Context) int {
	id, err := gsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GuildSettingsSlice.
func (gsq *GuildSettingsQuery) All(ctx context.Context) ([]*GuildSettings, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gsq *GuildSettingsQuery) AllX(ctx context.Context) []*GuildSettings {
	nodes, err := gsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GuildSettings IDs.
func (gsq *GuildSettingsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gsq.Select(guildsettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gsq *GuildSettingsQuery) IDsX(ctx context.Context) []int {
	ids, err := gsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gsq *GuildSettingsQuery) Count(ctx context.Context) (int, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gsq *GuildSettingsQuery) CountX(ctx context.Context) int {
	count, err := gsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gsq *GuildSettingsQuery) Exist(ctx context.Context) (bool, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gsq *GuildSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := gsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GuildSettingsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gsq *GuildSettingsQuery) Clone() *GuildSettingsQuery {
	if gsq == nil {
		return nil
	}
	return &GuildSettingsQuery{
		config:     gsq.config,
		limit:      gsq.limit,
		offset:     gsq.offset,
		order:      append([]OrderFunc{}, gsq.order...),
		predicates: append([]predicate.GuildSettings{}, gsq.predicates...),
		withGuild:  gsq.withGuild.Clone(),
		// clone intermediate query.
		sql:    gsq.sql.Clone(),
		path:   gsq.path,
		unique: gsq.unique,
	}
}

// WithGuild tells the query-builder to eager-load the nodes that are connected to
// the "guild" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GuildSettingsQuery) WithGuild(opts ...func(*GuildQuery)) *GuildSettingsQuery {
	query := &GuildQuery{config: gsq.config}
	for _, opt := range opts {
		opt(query)
	}
	gsq.withGuild = query
	return gsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GuildSettings.Query().
//		GroupBy(guildsettings.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gsq *GuildSettingsQuery) GroupBy(field string, fields ...string) *GuildSettingsGroupBy {
	grbuild := &GuildSettingsGroupBy{config: gsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gsq.sqlQuery(ctx), nil
	}
	grbuild.label = guildsettings.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.GuildSettings.Query().
//		Select(guildsettings.FieldCreateTime).
//		Scan(ctx, &v)
func (gsq *GuildSettingsQuery) Select(fields ...string) *GuildSettingsSelect {
	gsq.fields = append(gsq.fields, fields...)
	selbuild := &GuildSettingsSelect{GuildSettingsQuery: gsq}
	selbuild.label = guildsettings.Label
	selbuild.flds, selbuild.scan = &gsq.fields, selbuild.Scan
	return selbuild
}

func (gsq *GuildSettingsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gsq.fields {
		if !guildsettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gsq.path != nil {
		prev, err := gsq.path(ctx)
		if err != nil {
			return err
		}
		gsq.sql = prev
	}
	if guildsettings.Policy == nil {
		return errors.New("ent: uninitialized guildsettings.Policy (forgotten import ent/runtime?)")
	}
	if err := guildsettings.Policy.EvalQuery(ctx, gsq); err != nil {
		return err
	}
	return nil
}

func (gsq *GuildSettingsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GuildSettings, error) {
	var (
		nodes       = []*GuildSettings{}
		withFKs     = gsq.withFKs
		_spec       = gsq.querySpec()
		loadedTypes = [1]bool{
			gsq.withGuild != nil,
		}
	)
	if gsq.withGuild != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, guildsettings.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GuildSettings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GuildSettings{config: gsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gsq.withGuild; query != nil {
		if err := gsq.loadGuild(ctx, query, nodes, nil,
			func(n *GuildSettings, e *Guild) { n.Edges.Guild = e }); err != nil {
			return nil, err
		}
	}
	for i := range gsq.loadTotal {
		if err := gsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gsq *GuildSettingsQuery) loadGuild(ctx context.Context, query *GuildQuery, nodes []*GuildSettings, init func(*GuildSettings), assign func(*GuildSettings, *Guild)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GuildSettings)
	for i := range nodes {
		if nodes[i].guild_guild_settings == nil {
			continue
		}
		fk := *nodes[i].guild_guild_settings
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(guild.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "guild_guild_settings" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gsq *GuildSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gsq.querySpec()
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	_spec.Node.Columns = gsq.fields
	if len(gsq.fields) > 0 {
		_spec.Unique = gsq.unique != nil && *gsq.unique
	}
	return sqlgraph.CountNodes(ctx, gsq.driver, _spec)
}

func (gsq *GuildSettingsQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := gsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (gsq *GuildSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guildsettings.Table,
			Columns: guildsettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: guildsettings.FieldID,
			},
		},
		From:   gsq.sql,
		Unique: true,
	}
	if unique := gsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guildsettings.FieldID)
		for i := range fields {
			if fields[i] != guildsettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gsq *GuildSettingsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gsq.driver.Dialect())
	t1 := builder.Table(guildsettings.Table)
	columns := gsq.fields
	if len(columns) == 0 {
		columns = guildsettings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gsq.sql != nil {
		selector = gsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gsq.unique != nil && *gsq.unique {
		selector.Distinct()
	}
	for _, p := range gsq.predicates {
		p(selector)
	}
	for _, p := range gsq.order {
		p(selector)
	}
	if offset := gsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GuildSettingsGroupBy is the group-by builder for GuildSettings entities.
type GuildSettingsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gsgb *GuildSettingsGroupBy) Aggregate(fns ...AggregateFunc) *GuildSettingsGroupBy {
	gsgb.fns = append(gsgb.fns, fns...)
	return gsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gsgb *GuildSettingsGroupBy) Scan(ctx context.Context, v any) error {
	query, err := gsgb.path(ctx)
	if err != nil {
		return err
	}
	gsgb.sql = query
	return gsgb.sqlScan(ctx, v)
}

func (gsgb *GuildSettingsGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range gsgb.fields {
		if !guildsettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gsgb *GuildSettingsGroupBy) sqlQuery() *sql.Selector {
	selector := gsgb.sql.Select()
	aggregation := make([]string, 0, len(gsgb.fns))
	for _, fn := range gsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gsgb.fields)+len(gsgb.fns))
		for _, f := range gsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gsgb.fields...)...)
}

// GuildSettingsSelect is the builder for selecting fields of GuildSettings entities.
type GuildSettingsSelect struct {
	*GuildSettingsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gss *GuildSettingsSelect) Scan(ctx context.Context, v any) error {
	if err := gss.prepareQuery(ctx); err != nil {
		return err
	}
	gss.sql = gss.GuildSettingsQuery.sqlQuery(ctx)
	return gss.sqlScan(ctx, v)
}

func (gss *GuildSettingsSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := gss.sql.Query()
	if err := gss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
