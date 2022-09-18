// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lrstanley/ent-relay-conn-bug/ent/privacy"
)

type Guild struct {
	ent.Schema
}

func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.String("guild_id").Immutable().Comment("Guild id."),
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		).MinLen(2).MaxLen(100).Comment("Guild name (2-100 chars, excl. trailing/leading spaces)."),
		field.Strings("features").Optional().Optional().Default([]string{}).Comment("Enabled guild features."),
		field.String("icon_hash").Optional().MaxLen(2048).Comment("Icon hash."),
		field.String("icon_url").MaxLen(2048),
		field.Time("joined_at").Immutable().Annotations(
			entgql.OrderField("JOINED_AT"),
		).Comment("When the bot joined the guild."),
		field.Bool("large").Optional().Default(false).Comment("True if the guild is considered large (according to Discord standards)."),
		field.Int("member_count").Optional().Default(0).Comment("Total number of members in the guild."),
		field.String("owner_id").Comment("Discord snowflake ID of the user that owns the guild."),
		field.Uint64("permissions").Optional().Default(0).Annotations(
			entgql.Type("Uint64"),
		).Comment("Permissions of the bot on this guild (excludes overrides)."),
		field.String("region").Optional().MaxLen(32).Comment("Region of the guild."),
		field.String("system_channel_flags").Optional().MaxLen(32).Comment("System channel flags."),
	}
}

func (Guild) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Guild) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			AllowRoles([]string{"admin"}, true),
		},
	}
}

func (Guild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("guild_settings", GuildSettings.Type).Unique(),
	}
}

func (Guild) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
