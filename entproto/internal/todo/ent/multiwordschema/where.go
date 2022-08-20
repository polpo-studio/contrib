// Code generated by ent, DO NOT EDIT.

package multiwordschema

import (
	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v Unit) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v Unit) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnit), v))
	})
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...Unit) predicate.MultiWordSchema {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnit), v...))
	})
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...Unit) predicate.MultiWordSchema {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnit), v...))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MultiWordSchema) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MultiWordSchema) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
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
func Not(p predicate.MultiWordSchema) predicate.MultiWordSchema {
	return predicate.MultiWordSchema(func(s *sql.Selector) {
		p(s.Not())
	})
}
