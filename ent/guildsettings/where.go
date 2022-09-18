// Code generated by ent, DO NOT EDIT.

package guildsettings

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/ent-relay-conn-bug/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnabled), v))
	})
}

// DefaultMaxClones applies equality check predicate on the "default_max_clones" field. It's identical to DefaultMaxClonesEQ.
func DefaultMaxClones(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDefaultMaxClones), v))
	})
}

// RegexMatch applies equality check predicate on the "regex_match" field. It's identical to RegexMatchEQ.
func RegexMatch(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegexMatch), v))
	})
}

// ContactEmail applies equality check predicate on the "contact_email" field. It's identical to ContactEmailEQ.
func ContactEmail(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactEmail), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnabled), v))
	})
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnabled), v))
	})
}

// DefaultMaxClonesEQ applies the EQ predicate on the "default_max_clones" field.
func DefaultMaxClonesEQ(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDefaultMaxClones), v))
	})
}

// DefaultMaxClonesNEQ applies the NEQ predicate on the "default_max_clones" field.
func DefaultMaxClonesNEQ(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDefaultMaxClones), v))
	})
}

// DefaultMaxClonesIn applies the In predicate on the "default_max_clones" field.
func DefaultMaxClonesIn(vs ...int) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDefaultMaxClones), v...))
	})
}

// DefaultMaxClonesNotIn applies the NotIn predicate on the "default_max_clones" field.
func DefaultMaxClonesNotIn(vs ...int) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDefaultMaxClones), v...))
	})
}

// DefaultMaxClonesGT applies the GT predicate on the "default_max_clones" field.
func DefaultMaxClonesGT(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDefaultMaxClones), v))
	})
}

// DefaultMaxClonesGTE applies the GTE predicate on the "default_max_clones" field.
func DefaultMaxClonesGTE(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDefaultMaxClones), v))
	})
}

// DefaultMaxClonesLT applies the LT predicate on the "default_max_clones" field.
func DefaultMaxClonesLT(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDefaultMaxClones), v))
	})
}

// DefaultMaxClonesLTE applies the LTE predicate on the "default_max_clones" field.
func DefaultMaxClonesLTE(v int) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDefaultMaxClones), v))
	})
}

// DefaultMaxClonesIsNil applies the IsNil predicate on the "default_max_clones" field.
func DefaultMaxClonesIsNil() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDefaultMaxClones)))
	})
}

// DefaultMaxClonesNotNil applies the NotNil predicate on the "default_max_clones" field.
func DefaultMaxClonesNotNil() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDefaultMaxClones)))
	})
}

// RegexMatchEQ applies the EQ predicate on the "regex_match" field.
func RegexMatchEQ(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchNEQ applies the NEQ predicate on the "regex_match" field.
func RegexMatchNEQ(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchIn applies the In predicate on the "regex_match" field.
func RegexMatchIn(vs ...string) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRegexMatch), v...))
	})
}

// RegexMatchNotIn applies the NotIn predicate on the "regex_match" field.
func RegexMatchNotIn(vs ...string) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRegexMatch), v...))
	})
}

// RegexMatchGT applies the GT predicate on the "regex_match" field.
func RegexMatchGT(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchGTE applies the GTE predicate on the "regex_match" field.
func RegexMatchGTE(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchLT applies the LT predicate on the "regex_match" field.
func RegexMatchLT(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchLTE applies the LTE predicate on the "regex_match" field.
func RegexMatchLTE(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchContains applies the Contains predicate on the "regex_match" field.
func RegexMatchContains(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchHasPrefix applies the HasPrefix predicate on the "regex_match" field.
func RegexMatchHasPrefix(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchHasSuffix applies the HasSuffix predicate on the "regex_match" field.
func RegexMatchHasSuffix(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchIsNil applies the IsNil predicate on the "regex_match" field.
func RegexMatchIsNil() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRegexMatch)))
	})
}

// RegexMatchNotNil applies the NotNil predicate on the "regex_match" field.
func RegexMatchNotNil() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRegexMatch)))
	})
}

// RegexMatchEqualFold applies the EqualFold predicate on the "regex_match" field.
func RegexMatchEqualFold(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegexMatch), v))
	})
}

// RegexMatchContainsFold applies the ContainsFold predicate on the "regex_match" field.
func RegexMatchContainsFold(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegexMatch), v))
	})
}

// ContactEmailEQ applies the EQ predicate on the "contact_email" field.
func ContactEmailEQ(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactEmail), v))
	})
}

// ContactEmailNEQ applies the NEQ predicate on the "contact_email" field.
func ContactEmailNEQ(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContactEmail), v))
	})
}

// ContactEmailIn applies the In predicate on the "contact_email" field.
func ContactEmailIn(vs ...string) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContactEmail), v...))
	})
}

// ContactEmailNotIn applies the NotIn predicate on the "contact_email" field.
func ContactEmailNotIn(vs ...string) predicate.GuildSettings {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContactEmail), v...))
	})
}

// ContactEmailGT applies the GT predicate on the "contact_email" field.
func ContactEmailGT(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContactEmail), v))
	})
}

// ContactEmailGTE applies the GTE predicate on the "contact_email" field.
func ContactEmailGTE(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContactEmail), v))
	})
}

// ContactEmailLT applies the LT predicate on the "contact_email" field.
func ContactEmailLT(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContactEmail), v))
	})
}

// ContactEmailLTE applies the LTE predicate on the "contact_email" field.
func ContactEmailLTE(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContactEmail), v))
	})
}

// ContactEmailContains applies the Contains predicate on the "contact_email" field.
func ContactEmailContains(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContactEmail), v))
	})
}

// ContactEmailHasPrefix applies the HasPrefix predicate on the "contact_email" field.
func ContactEmailHasPrefix(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContactEmail), v))
	})
}

// ContactEmailHasSuffix applies the HasSuffix predicate on the "contact_email" field.
func ContactEmailHasSuffix(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContactEmail), v))
	})
}

// ContactEmailIsNil applies the IsNil predicate on the "contact_email" field.
func ContactEmailIsNil() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContactEmail)))
	})
}

// ContactEmailNotNil applies the NotNil predicate on the "contact_email" field.
func ContactEmailNotNil() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContactEmail)))
	})
}

// ContactEmailEqualFold applies the EqualFold predicate on the "contact_email" field.
func ContactEmailEqualFold(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContactEmail), v))
	})
}

// ContactEmailContainsFold applies the ContainsFold predicate on the "contact_email" field.
func ContactEmailContainsFold(v string) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContactEmail), v))
	})
}

// HasGuild applies the HasEdge predicate on the "guild" edge.
func HasGuild() predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GuildTable, GuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGuildWith applies the HasEdge predicate on the "guild" edge with a given conditions (other predicates).
func HasGuildWith(preds ...predicate.Guild) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GuildTable, GuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GuildSettings) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GuildSettings) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GuildSettings) predicate.GuildSettings {
	return predicate.GuildSettings(func(s *sql.Selector) {
		p(s.Not())
	})
}