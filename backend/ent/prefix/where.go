// Code generated by entc, DO NOT EDIT.

package prefix

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Prefix applies equality check predicate on the "prefix" field. It's identical to PrefixEQ.
func Prefix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrefix), v))
	})
}

// PrefixEQ applies the EQ predicate on the "prefix" field.
func PrefixEQ(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrefix), v))
	})
}

// PrefixNEQ applies the NEQ predicate on the "prefix" field.
func PrefixNEQ(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrefix), v))
	})
}

// PrefixIn applies the In predicate on the "prefix" field.
func PrefixIn(vs ...string) predicate.Prefix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrefix), v...))
	})
}

// PrefixNotIn applies the NotIn predicate on the "prefix" field.
func PrefixNotIn(vs ...string) predicate.Prefix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrefix), v...))
	})
}

// PrefixGT applies the GT predicate on the "prefix" field.
func PrefixGT(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrefix), v))
	})
}

// PrefixGTE applies the GTE predicate on the "prefix" field.
func PrefixGTE(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrefix), v))
	})
}

// PrefixLT applies the LT predicate on the "prefix" field.
func PrefixLT(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrefix), v))
	})
}

// PrefixLTE applies the LTE predicate on the "prefix" field.
func PrefixLTE(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrefix), v))
	})
}

// PrefixContains applies the Contains predicate on the "prefix" field.
func PrefixContains(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrefix), v))
	})
}

// PrefixHasPrefix applies the HasPrefix predicate on the "prefix" field.
func PrefixHasPrefix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrefix), v))
	})
}

// PrefixHasSuffix applies the HasSuffix predicate on the "prefix" field.
func PrefixHasSuffix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrefix), v))
	})
}

// PrefixEqualFold applies the EqualFold predicate on the "prefix" field.
func PrefixEqualFold(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrefix), v))
	})
}

// PrefixContainsFold applies the ContainsFold predicate on the "prefix" field.
func PrefixContainsFold(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrefix), v))
	})
}

// HasPreProf applies the HasEdge predicate on the "pre_prof" edge.
func HasPreProf() predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PreProfTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PreProfTable, PreProfColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPreProfWith applies the HasEdge predicate on the "pre_prof" edge with a given conditions (other predicates).
func HasPreProfWith(preds ...predicate.Professor) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PreProfInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PreProfTable, PreProfColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrefStud applies the HasEdge predicate on the "pref_stud" edge.
func HasPrefStud() predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrefStudTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PrefStudTable, PrefStudColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrefStudWith applies the HasEdge predicate on the "pref_stud" edge with a given conditions (other predicates).
func HasPrefStudWith(preds ...predicate.Student) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrefStudInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PrefStudTable, PrefStudColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
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
func Not(p predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		p(s.Not())
	})
}
