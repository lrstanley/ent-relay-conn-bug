// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/lrstanley/ent-relay-conn-bug/ent/guild"
	"github.com/lrstanley/ent-relay-conn-bug/ent/guildsettings"
	"github.com/lrstanley/ent-relay-conn-bug/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   guild.Table,
			Columns: guild.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: guild.FieldID,
			},
		},
		Type: "Guild",
		Fields: map[string]*sqlgraph.FieldSpec{
			guild.FieldCreateTime:         {Type: field.TypeTime, Column: guild.FieldCreateTime},
			guild.FieldUpdateTime:         {Type: field.TypeTime, Column: guild.FieldUpdateTime},
			guild.FieldGuildID:            {Type: field.TypeString, Column: guild.FieldGuildID},
			guild.FieldName:               {Type: field.TypeString, Column: guild.FieldName},
			guild.FieldFeatures:           {Type: field.TypeJSON, Column: guild.FieldFeatures},
			guild.FieldIconHash:           {Type: field.TypeString, Column: guild.FieldIconHash},
			guild.FieldIconURL:            {Type: field.TypeString, Column: guild.FieldIconURL},
			guild.FieldJoinedAt:           {Type: field.TypeTime, Column: guild.FieldJoinedAt},
			guild.FieldLarge:              {Type: field.TypeBool, Column: guild.FieldLarge},
			guild.FieldMemberCount:        {Type: field.TypeInt, Column: guild.FieldMemberCount},
			guild.FieldOwnerID:            {Type: field.TypeString, Column: guild.FieldOwnerID},
			guild.FieldPermissions:        {Type: field.TypeUint64, Column: guild.FieldPermissions},
			guild.FieldRegion:             {Type: field.TypeString, Column: guild.FieldRegion},
			guild.FieldSystemChannelFlags: {Type: field.TypeString, Column: guild.FieldSystemChannelFlags},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   guildsettings.Table,
			Columns: guildsettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: guildsettings.FieldID,
			},
		},
		Type: "GuildSettings",
		Fields: map[string]*sqlgraph.FieldSpec{
			guildsettings.FieldCreateTime:       {Type: field.TypeTime, Column: guildsettings.FieldCreateTime},
			guildsettings.FieldUpdateTime:       {Type: field.TypeTime, Column: guildsettings.FieldUpdateTime},
			guildsettings.FieldEnabled:          {Type: field.TypeBool, Column: guildsettings.FieldEnabled},
			guildsettings.FieldDefaultMaxClones: {Type: field.TypeInt, Column: guildsettings.FieldDefaultMaxClones},
			guildsettings.FieldRegexMatch:       {Type: field.TypeString, Column: guildsettings.FieldRegexMatch},
			guildsettings.FieldContactEmail:     {Type: field.TypeString, Column: guildsettings.FieldContactEmail},
		},
	}
	graph.MustAddE(
		"guild_settings",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   guild.GuildSettingsTable,
			Columns: []string{guild.GuildSettingsColumn},
			Bidi:    false,
		},
		"Guild",
		"GuildSettings",
	)
	graph.MustAddE(
		"guild",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   guildsettings.GuildTable,
			Columns: []string{guildsettings.GuildColumn},
			Bidi:    false,
		},
		"GuildSettings",
		"Guild",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (gq *GuildQuery) addPredicate(pred func(s *sql.Selector)) {
	gq.predicates = append(gq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GuildQuery builder.
func (gq *GuildQuery) Filter() *GuildFilter {
	return &GuildFilter{config: gq.config, predicateAdder: gq}
}

// addPredicate implements the predicateAdder interface.
func (m *GuildMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GuildMutation builder.
func (m *GuildMutation) Filter() *GuildFilter {
	return &GuildFilter{config: m.config, predicateAdder: m}
}

// GuildFilter provides a generic filtering capability at runtime for GuildQuery.
type GuildFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GuildFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *GuildFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(guild.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *GuildFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(guild.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *GuildFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(guild.FieldUpdateTime))
}

// WhereGuildID applies the entql string predicate on the guild_id field.
func (f *GuildFilter) WhereGuildID(p entql.StringP) {
	f.Where(p.Field(guild.FieldGuildID))
}

// WhereName applies the entql string predicate on the name field.
func (f *GuildFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(guild.FieldName))
}

// WhereFeatures applies the entql json.RawMessage predicate on the features field.
func (f *GuildFilter) WhereFeatures(p entql.BytesP) {
	f.Where(p.Field(guild.FieldFeatures))
}

// WhereIconHash applies the entql string predicate on the icon_hash field.
func (f *GuildFilter) WhereIconHash(p entql.StringP) {
	f.Where(p.Field(guild.FieldIconHash))
}

// WhereIconURL applies the entql string predicate on the icon_url field.
func (f *GuildFilter) WhereIconURL(p entql.StringP) {
	f.Where(p.Field(guild.FieldIconURL))
}

// WhereJoinedAt applies the entql time.Time predicate on the joined_at field.
func (f *GuildFilter) WhereJoinedAt(p entql.TimeP) {
	f.Where(p.Field(guild.FieldJoinedAt))
}

// WhereLarge applies the entql bool predicate on the large field.
func (f *GuildFilter) WhereLarge(p entql.BoolP) {
	f.Where(p.Field(guild.FieldLarge))
}

// WhereMemberCount applies the entql int predicate on the member_count field.
func (f *GuildFilter) WhereMemberCount(p entql.IntP) {
	f.Where(p.Field(guild.FieldMemberCount))
}

// WhereOwnerID applies the entql string predicate on the owner_id field.
func (f *GuildFilter) WhereOwnerID(p entql.StringP) {
	f.Where(p.Field(guild.FieldOwnerID))
}

// WherePermissions applies the entql uint64 predicate on the permissions field.
func (f *GuildFilter) WherePermissions(p entql.Uint64P) {
	f.Where(p.Field(guild.FieldPermissions))
}

// WhereRegion applies the entql string predicate on the region field.
func (f *GuildFilter) WhereRegion(p entql.StringP) {
	f.Where(p.Field(guild.FieldRegion))
}

// WhereSystemChannelFlags applies the entql string predicate on the system_channel_flags field.
func (f *GuildFilter) WhereSystemChannelFlags(p entql.StringP) {
	f.Where(p.Field(guild.FieldSystemChannelFlags))
}

// WhereHasGuildSettings applies a predicate to check if query has an edge guild_settings.
func (f *GuildFilter) WhereHasGuildSettings() {
	f.Where(entql.HasEdge("guild_settings"))
}

// WhereHasGuildSettingsWith applies a predicate to check if query has an edge guild_settings with a given conditions (other predicates).
func (f *GuildFilter) WhereHasGuildSettingsWith(preds ...predicate.GuildSettings) {
	f.Where(entql.HasEdgeWith("guild_settings", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (gsq *GuildSettingsQuery) addPredicate(pred func(s *sql.Selector)) {
	gsq.predicates = append(gsq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GuildSettingsQuery builder.
func (gsq *GuildSettingsQuery) Filter() *GuildSettingsFilter {
	return &GuildSettingsFilter{config: gsq.config, predicateAdder: gsq}
}

// addPredicate implements the predicateAdder interface.
func (m *GuildSettingsMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GuildSettingsMutation builder.
func (m *GuildSettingsMutation) Filter() *GuildSettingsFilter {
	return &GuildSettingsFilter{config: m.config, predicateAdder: m}
}

// GuildSettingsFilter provides a generic filtering capability at runtime for GuildSettingsQuery.
type GuildSettingsFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GuildSettingsFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *GuildSettingsFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(guildsettings.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *GuildSettingsFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(guildsettings.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *GuildSettingsFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(guildsettings.FieldUpdateTime))
}

// WhereEnabled applies the entql bool predicate on the enabled field.
func (f *GuildSettingsFilter) WhereEnabled(p entql.BoolP) {
	f.Where(p.Field(guildsettings.FieldEnabled))
}

// WhereDefaultMaxClones applies the entql int predicate on the default_max_clones field.
func (f *GuildSettingsFilter) WhereDefaultMaxClones(p entql.IntP) {
	f.Where(p.Field(guildsettings.FieldDefaultMaxClones))
}

// WhereRegexMatch applies the entql string predicate on the regex_match field.
func (f *GuildSettingsFilter) WhereRegexMatch(p entql.StringP) {
	f.Where(p.Field(guildsettings.FieldRegexMatch))
}

// WhereContactEmail applies the entql string predicate on the contact_email field.
func (f *GuildSettingsFilter) WhereContactEmail(p entql.StringP) {
	f.Where(p.Field(guildsettings.FieldContactEmail))
}

// WhereHasGuild applies a predicate to check if query has an edge guild.
func (f *GuildSettingsFilter) WhereHasGuild() {
	f.Where(entql.HasEdge("guild"))
}

// WhereHasGuildWith applies a predicate to check if query has an edge guild with a given conditions (other predicates).
func (f *GuildSettingsFilter) WhereHasGuildWith(preds ...predicate.Guild) {
	f.Where(entql.HasEdgeWith("guild", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
