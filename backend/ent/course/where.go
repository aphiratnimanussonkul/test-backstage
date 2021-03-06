// Code generated by entc, DO NOT EDIT.

package course

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Course applies equality check predicate on the "course" field. It's identical to CourseEQ.
func Course(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCourse), v))
	})
}

// CourseEQ applies the EQ predicate on the "course" field.
func CourseEQ(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCourse), v))
	})
}

// CourseNEQ applies the NEQ predicate on the "course" field.
func CourseNEQ(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCourse), v))
	})
}

// CourseIn applies the In predicate on the "course" field.
func CourseIn(vs ...string) predicate.Course {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Course(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCourse), v...))
	})
}

// CourseNotIn applies the NotIn predicate on the "course" field.
func CourseNotIn(vs ...string) predicate.Course {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Course(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCourse), v...))
	})
}

// CourseGT applies the GT predicate on the "course" field.
func CourseGT(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCourse), v))
	})
}

// CourseGTE applies the GTE predicate on the "course" field.
func CourseGTE(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCourse), v))
	})
}

// CourseLT applies the LT predicate on the "course" field.
func CourseLT(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCourse), v))
	})
}

// CourseLTE applies the LTE predicate on the "course" field.
func CourseLTE(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCourse), v))
	})
}

// CourseContains applies the Contains predicate on the "course" field.
func CourseContains(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCourse), v))
	})
}

// CourseHasPrefix applies the HasPrefix predicate on the "course" field.
func CourseHasPrefix(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCourse), v))
	})
}

// CourseHasSuffix applies the HasSuffix predicate on the "course" field.
func CourseHasSuffix(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCourse), v))
	})
}

// CourseEqualFold applies the EqualFold predicate on the "course" field.
func CourseEqualFold(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCourse), v))
	})
}

// CourseContainsFold applies the ContainsFold predicate on the "course" field.
func CourseContainsFold(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCourse), v))
	})
}

// HasCourFacu applies the HasEdge predicate on the "cour_facu" edge.
func HasCourFacu() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourFacuTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourFacuTable, CourFacuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourFacuWith applies the HasEdge predicate on the "cour_facu" edge with a given conditions (other predicates).
func HasCourFacuWith(preds ...predicate.Faculty) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourFacuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourFacuTable, CourFacuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCourDegr applies the HasEdge predicate on the "cour_degr" edge.
func HasCourDegr() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourDegrTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourDegrTable, CourDegrColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourDegrWith applies the HasEdge predicate on the "cour_degr" edge with a given conditions (other predicates).
func HasCourDegrWith(preds ...predicate.Degree) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourDegrInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourDegrTable, CourDegrColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCourInst applies the HasEdge predicate on the "cour_inst" edge.
func HasCourInst() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourInstTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourInstTable, CourInstColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourInstWith applies the HasEdge predicate on the "cour_inst" edge with a given conditions (other predicates).
func HasCourInstWith(preds ...predicate.Institution) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourInstInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourInstTable, CourInstColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func Not(p predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		p(s.Not())
	})
}
