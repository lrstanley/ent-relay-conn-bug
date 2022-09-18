// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (gu *Guild) GuildSettings(ctx context.Context) (*GuildSettings, error) {
	result, err := gu.Edges.GuildSettingsOrErr()
	if IsNotLoaded(err) {
		result, err = gu.QueryGuildSettings().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (gs *GuildSettings) Guild(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *GuildOrder, where *GuildWhereInput,
) (*GuildConnection, error) {
	opts := []GuildPaginateOption{
		WithGuildOrder(orderBy),
		WithGuildFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gs.Edges.totalCount[0][alias]
	if nodes, err := gs.NamedGuild(alias); err == nil || hasTotalCount {
		pager, err := newGuildPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &GuildConnection{Edges: []*GuildEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gs.QueryGuild().Paginate(ctx, after, first, before, last, opts...)
}
