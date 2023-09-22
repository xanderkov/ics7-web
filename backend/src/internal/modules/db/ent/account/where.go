// Code generated by ent, DO NOT EDIT.

package account

import (
	"hospital/internal/modules/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// Login applies equality check predicate on the "login" field. It's identical to LoginEQ.
func Login(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldLogin, v))
}

// PasswordHash applies equality check predicate on the "password_hash" field. It's identical to PasswordHashEQ.
func PasswordHash(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPasswordHash, v))
}

// LoginEQ applies the EQ predicate on the "login" field.
func LoginEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldLogin, v))
}

// LoginNEQ applies the NEQ predicate on the "login" field.
func LoginNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldLogin, v))
}

// LoginIn applies the In predicate on the "login" field.
func LoginIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldLogin, vs...))
}

// LoginNotIn applies the NotIn predicate on the "login" field.
func LoginNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldLogin, vs...))
}

// LoginGT applies the GT predicate on the "login" field.
func LoginGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldLogin, v))
}

// LoginGTE applies the GTE predicate on the "login" field.
func LoginGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldLogin, v))
}

// LoginLT applies the LT predicate on the "login" field.
func LoginLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldLogin, v))
}

// LoginLTE applies the LTE predicate on the "login" field.
func LoginLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldLogin, v))
}

// LoginContains applies the Contains predicate on the "login" field.
func LoginContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldLogin, v))
}

// LoginHasPrefix applies the HasPrefix predicate on the "login" field.
func LoginHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldLogin, v))
}

// LoginHasSuffix applies the HasSuffix predicate on the "login" field.
func LoginHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldLogin, v))
}

// LoginEqualFold applies the EqualFold predicate on the "login" field.
func LoginEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldLogin, v))
}

// LoginContainsFold applies the ContainsFold predicate on the "login" field.
func LoginContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldLogin, v))
}

// PasswordHashEQ applies the EQ predicate on the "password_hash" field.
func PasswordHashEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPasswordHash, v))
}

// PasswordHashNEQ applies the NEQ predicate on the "password_hash" field.
func PasswordHashNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldPasswordHash, v))
}

// PasswordHashIn applies the In predicate on the "password_hash" field.
func PasswordHashIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldPasswordHash, vs...))
}

// PasswordHashNotIn applies the NotIn predicate on the "password_hash" field.
func PasswordHashNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldPasswordHash, vs...))
}

// PasswordHashGT applies the GT predicate on the "password_hash" field.
func PasswordHashGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldPasswordHash, v))
}

// PasswordHashGTE applies the GTE predicate on the "password_hash" field.
func PasswordHashGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldPasswordHash, v))
}

// PasswordHashLT applies the LT predicate on the "password_hash" field.
func PasswordHashLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldPasswordHash, v))
}

// PasswordHashLTE applies the LTE predicate on the "password_hash" field.
func PasswordHashLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldPasswordHash, v))
}

// PasswordHashContains applies the Contains predicate on the "password_hash" field.
func PasswordHashContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldPasswordHash, v))
}

// PasswordHashHasPrefix applies the HasPrefix predicate on the "password_hash" field.
func PasswordHashHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldPasswordHash, v))
}

// PasswordHashHasSuffix applies the HasSuffix predicate on the "password_hash" field.
func PasswordHashHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldPasswordHash, v))
}

// PasswordHashEqualFold applies the EqualFold predicate on the "password_hash" field.
func PasswordHashEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldPasswordHash, v))
}

// PasswordHashContainsFold applies the ContainsFold predicate on the "password_hash" field.
func PasswordHashContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldPasswordHash, v))
}

// HasIs applies the HasEdge predicate on the "is" edge.
func HasIs() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, IsTable, IsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIsWith applies the HasEdge predicate on the "is" edge with a given conditions (other predicates).
func HasIsWith(preds ...predicate.Doctor) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newIsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
