// Code generated by entc, DO NOT EDIT.

package activity

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ACTIVITYNAME applies equality check predicate on the "ACTIVITYNAME" field. It's identical to ACTIVITYNAMEEQ.
func ACTIVITYNAME(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldACTIVITYNAME), v))
	})
}

// ADDED applies equality check predicate on the "ADDED" field. It's identical to ADDEDEQ.
func ADDED(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldADDED), v))
	})
}

// HOURS applies equality check predicate on the "HOURS" field. It's identical to HOURSEQ.
func HOURS(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHOURS), v))
	})
}

// ACTIVITYNAMEEQ applies the EQ predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEEQ(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMENEQ applies the NEQ predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMENEQ(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEIn applies the In predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEIn(vs ...string) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldACTIVITYNAME), v...))
	})
}

// ACTIVITYNAMENotIn applies the NotIn predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMENotIn(vs ...string) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldACTIVITYNAME), v...))
	})
}

// ACTIVITYNAMEGT applies the GT predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEGT(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEGTE applies the GTE predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEGTE(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMELT applies the LT predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMELT(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMELTE applies the LTE predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMELTE(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEContains applies the Contains predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEContains(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEHasPrefix applies the HasPrefix predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEHasPrefix(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEHasSuffix applies the HasSuffix predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEHasSuffix(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEEqualFold applies the EqualFold predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEEqualFold(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldACTIVITYNAME), v))
	})
}

// ACTIVITYNAMEContainsFold applies the ContainsFold predicate on the "ACTIVITYNAME" field.
func ACTIVITYNAMEContainsFold(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldACTIVITYNAME), v))
	})
}

// ADDEDEQ applies the EQ predicate on the "ADDED" field.
func ADDEDEQ(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldADDED), v))
	})
}

// ADDEDNEQ applies the NEQ predicate on the "ADDED" field.
func ADDEDNEQ(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldADDED), v))
	})
}

// ADDEDIn applies the In predicate on the "ADDED" field.
func ADDEDIn(vs ...time.Time) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldADDED), v...))
	})
}

// ADDEDNotIn applies the NotIn predicate on the "ADDED" field.
func ADDEDNotIn(vs ...time.Time) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldADDED), v...))
	})
}

// ADDEDGT applies the GT predicate on the "ADDED" field.
func ADDEDGT(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldADDED), v))
	})
}

// ADDEDGTE applies the GTE predicate on the "ADDED" field.
func ADDEDGTE(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldADDED), v))
	})
}

// ADDEDLT applies the LT predicate on the "ADDED" field.
func ADDEDLT(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldADDED), v))
	})
}

// ADDEDLTE applies the LTE predicate on the "ADDED" field.
func ADDEDLTE(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldADDED), v))
	})
}

// HOURSEQ applies the EQ predicate on the "HOURS" field.
func HOURSEQ(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHOURS), v))
	})
}

// HOURSNEQ applies the NEQ predicate on the "HOURS" field.
func HOURSNEQ(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHOURS), v))
	})
}

// HOURSIn applies the In predicate on the "HOURS" field.
func HOURSIn(vs ...int) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHOURS), v...))
	})
}

// HOURSNotIn applies the NotIn predicate on the "HOURS" field.
func HOURSNotIn(vs ...int) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHOURS), v...))
	})
}

// HOURSGT applies the GT predicate on the "HOURS" field.
func HOURSGT(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHOURS), v))
	})
}

// HOURSGTE applies the GTE predicate on the "HOURS" field.
func HOURSGTE(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHOURS), v))
	})
}

// HOURSLT applies the LT predicate on the "HOURS" field.
func HOURSLT(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHOURS), v))
	})
}

// HOURSLTE applies the LTE predicate on the "HOURS" field.
func HOURSLTE(v int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHOURS), v))
	})
}

// HasActiStud applies the HasEdge predicate on the "acti_stud" edge.
func HasActiStud() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiStudTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiStudTable, ActiStudColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActiStudWith applies the HasEdge predicate on the "acti_stud" edge with a given conditions (other predicates).
func HasActiStudWith(preds ...predicate.Student) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiStudInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiStudTable, ActiStudColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActiPlace applies the HasEdge predicate on the "acti_place" edge.
func HasActiPlace() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiPlaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiPlaceTable, ActiPlaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActiPlaceWith applies the HasEdge predicate on the "acti_place" edge with a given conditions (other predicates).
func HasActiPlaceWith(preds ...predicate.Place) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiPlaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiPlaceTable, ActiPlaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActiAgen applies the HasEdge predicate on the "acti_agen" edge.
func HasActiAgen() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiAgenTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiAgenTable, ActiAgenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActiAgenWith applies the HasEdge predicate on the "acti_agen" edge with a given conditions (other predicates).
func HasActiAgenWith(preds ...predicate.Agency) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiAgenInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiAgenTable, ActiAgenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActiYear applies the HasEdge predicate on the "acti_year" edge.
func HasActiYear() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiYearTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiYearTable, ActiYearColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActiYearWith applies the HasEdge predicate on the "acti_year" edge with a given conditions (other predicates).
func HasActiYearWith(preds ...predicate.Year) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActiYearInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiYearTable, ActiYearColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Activity) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Activity) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
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
func Not(p predicate.Activity) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		p(s.Not())
	})
}
