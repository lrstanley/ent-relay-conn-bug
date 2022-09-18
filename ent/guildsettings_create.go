// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-relay-conn-bug/ent/guild"
	"github.com/lrstanley/ent-relay-conn-bug/ent/guildsettings"
)

// GuildSettingsCreate is the builder for creating a GuildSettings entity.
type GuildSettingsCreate struct {
	config
	mutation *GuildSettingsMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (gsc *GuildSettingsCreate) SetCreateTime(t time.Time) *GuildSettingsCreate {
	gsc.mutation.SetCreateTime(t)
	return gsc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (gsc *GuildSettingsCreate) SetNillableCreateTime(t *time.Time) *GuildSettingsCreate {
	if t != nil {
		gsc.SetCreateTime(*t)
	}
	return gsc
}

// SetUpdateTime sets the "update_time" field.
func (gsc *GuildSettingsCreate) SetUpdateTime(t time.Time) *GuildSettingsCreate {
	gsc.mutation.SetUpdateTime(t)
	return gsc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (gsc *GuildSettingsCreate) SetNillableUpdateTime(t *time.Time) *GuildSettingsCreate {
	if t != nil {
		gsc.SetUpdateTime(*t)
	}
	return gsc
}

// SetEnabled sets the "enabled" field.
func (gsc *GuildSettingsCreate) SetEnabled(b bool) *GuildSettingsCreate {
	gsc.mutation.SetEnabled(b)
	return gsc
}

// SetDefaultMaxClones sets the "default_max_clones" field.
func (gsc *GuildSettingsCreate) SetDefaultMaxClones(i int) *GuildSettingsCreate {
	gsc.mutation.SetDefaultMaxClones(i)
	return gsc
}

// SetNillableDefaultMaxClones sets the "default_max_clones" field if the given value is not nil.
func (gsc *GuildSettingsCreate) SetNillableDefaultMaxClones(i *int) *GuildSettingsCreate {
	if i != nil {
		gsc.SetDefaultMaxClones(*i)
	}
	return gsc
}

// SetRegexMatch sets the "regex_match" field.
func (gsc *GuildSettingsCreate) SetRegexMatch(s string) *GuildSettingsCreate {
	gsc.mutation.SetRegexMatch(s)
	return gsc
}

// SetNillableRegexMatch sets the "regex_match" field if the given value is not nil.
func (gsc *GuildSettingsCreate) SetNillableRegexMatch(s *string) *GuildSettingsCreate {
	if s != nil {
		gsc.SetRegexMatch(*s)
	}
	return gsc
}

// SetContactEmail sets the "contact_email" field.
func (gsc *GuildSettingsCreate) SetContactEmail(s string) *GuildSettingsCreate {
	gsc.mutation.SetContactEmail(s)
	return gsc
}

// SetNillableContactEmail sets the "contact_email" field if the given value is not nil.
func (gsc *GuildSettingsCreate) SetNillableContactEmail(s *string) *GuildSettingsCreate {
	if s != nil {
		gsc.SetContactEmail(*s)
	}
	return gsc
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (gsc *GuildSettingsCreate) SetGuildID(id int) *GuildSettingsCreate {
	gsc.mutation.SetGuildID(id)
	return gsc
}

// SetGuild sets the "guild" edge to the Guild entity.
func (gsc *GuildSettingsCreate) SetGuild(g *Guild) *GuildSettingsCreate {
	return gsc.SetGuildID(g.ID)
}

// Mutation returns the GuildSettingsMutation object of the builder.
func (gsc *GuildSettingsCreate) Mutation() *GuildSettingsMutation {
	return gsc.mutation
}

// Save creates the GuildSettings in the database.
func (gsc *GuildSettingsCreate) Save(ctx context.Context) (*GuildSettings, error) {
	var (
		err  error
		node *GuildSettings
	)
	gsc.defaults()
	if len(gsc.hooks) == 0 {
		if err = gsc.check(); err != nil {
			return nil, err
		}
		node, err = gsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildSettingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gsc.check(); err != nil {
				return nil, err
			}
			gsc.mutation = mutation
			if node, err = gsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gsc.hooks) - 1; i >= 0; i-- {
			if gsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GuildSettings)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GuildSettingsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gsc *GuildSettingsCreate) SaveX(ctx context.Context) *GuildSettings {
	v, err := gsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gsc *GuildSettingsCreate) Exec(ctx context.Context) error {
	_, err := gsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsc *GuildSettingsCreate) ExecX(ctx context.Context) {
	if err := gsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsc *GuildSettingsCreate) defaults() {
	if _, ok := gsc.mutation.CreateTime(); !ok {
		v := guildsettings.DefaultCreateTime()
		gsc.mutation.SetCreateTime(v)
	}
	if _, ok := gsc.mutation.UpdateTime(); !ok {
		v := guildsettings.DefaultUpdateTime()
		gsc.mutation.SetUpdateTime(v)
	}
	if _, ok := gsc.mutation.DefaultMaxClones(); !ok {
		v := guildsettings.DefaultDefaultMaxClones
		gsc.mutation.SetDefaultMaxClones(v)
	}
	if _, ok := gsc.mutation.RegexMatch(); !ok {
		v := guildsettings.DefaultRegexMatch
		gsc.mutation.SetRegexMatch(v)
	}
	if _, ok := gsc.mutation.ContactEmail(); !ok {
		v := guildsettings.DefaultContactEmail
		gsc.mutation.SetContactEmail(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gsc *GuildSettingsCreate) check() error {
	if _, ok := gsc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "GuildSettings.create_time"`)}
	}
	if _, ok := gsc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "GuildSettings.update_time"`)}
	}
	if _, ok := gsc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`ent: missing required field "GuildSettings.enabled"`)}
	}
	if v, ok := gsc.mutation.RegexMatch(); ok {
		if err := guildsettings.RegexMatchValidator(v); err != nil {
			return &ValidationError{Name: "regex_match", err: fmt.Errorf(`ent: validator failed for field "GuildSettings.regex_match": %w`, err)}
		}
	}
	if v, ok := gsc.mutation.ContactEmail(); ok {
		if err := guildsettings.ContactEmailValidator(v); err != nil {
			return &ValidationError{Name: "contact_email", err: fmt.Errorf(`ent: validator failed for field "GuildSettings.contact_email": %w`, err)}
		}
	}
	if _, ok := gsc.mutation.GuildID(); !ok {
		return &ValidationError{Name: "guild", err: errors.New(`ent: missing required edge "GuildSettings.guild"`)}
	}
	return nil
}

func (gsc *GuildSettingsCreate) sqlSave(ctx context.Context) (*GuildSettings, error) {
	_node, _spec := gsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gsc *GuildSettingsCreate) createSpec() (*GuildSettings, *sqlgraph.CreateSpec) {
	var (
		_node = &GuildSettings{config: gsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: guildsettings.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: guildsettings.FieldID,
			},
		}
	)
	if value, ok := gsc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: guildsettings.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := gsc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: guildsettings.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := gsc.mutation.Enabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: guildsettings.FieldEnabled,
		})
		_node.Enabled = value
	}
	if value, ok := gsc.mutation.DefaultMaxClones(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: guildsettings.FieldDefaultMaxClones,
		})
		_node.DefaultMaxClones = value
	}
	if value, ok := gsc.mutation.RegexMatch(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: guildsettings.FieldRegexMatch,
		})
		_node.RegexMatch = value
	}
	if value, ok := gsc.mutation.ContactEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: guildsettings.FieldContactEmail,
		})
		_node.ContactEmail = value
	}
	if nodes := gsc.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   guildsettings.GuildTable,
			Columns: []string{guildsettings.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.guild_guild_settings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GuildSettingsCreateBulk is the builder for creating many GuildSettings entities in bulk.
type GuildSettingsCreateBulk struct {
	config
	builders []*GuildSettingsCreate
}

// Save creates the GuildSettings entities in the database.
func (gscb *GuildSettingsCreateBulk) Save(ctx context.Context) ([]*GuildSettings, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gscb.builders))
	nodes := make([]*GuildSettings, len(gscb.builders))
	mutators := make([]Mutator, len(gscb.builders))
	for i := range gscb.builders {
		func(i int, root context.Context) {
			builder := gscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GuildSettingsMutation)
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
					_, err = mutators[i+1].Mutate(root, gscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, gscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gscb *GuildSettingsCreateBulk) SaveX(ctx context.Context) []*GuildSettings {
	v, err := gscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gscb *GuildSettingsCreateBulk) Exec(ctx context.Context) error {
	_, err := gscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gscb *GuildSettingsCreateBulk) ExecX(ctx context.Context) {
	if err := gscb.Exec(ctx); err != nil {
		panic(err)
	}
}
