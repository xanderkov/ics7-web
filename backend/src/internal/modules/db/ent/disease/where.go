// Code generated by ent, DO NOT EDIT.

package disease

import (
	"hospital/internal/modules/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Disease {
	return predicate.Disease(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Disease {
	return predicate.Disease(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Disease {
	return predicate.Disease(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Disease {
	return predicate.Disease(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Disease {
	return predicate.Disease(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Disease {
	return predicate.Disease(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Disease {
	return predicate.Disease(sql.FieldLTE(FieldID, id))
}

// Threat applies equality check predicate on the "threat" field. It's identical to ThreatEQ.
func Threat(v string) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldThreat, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldName, v))
}

// DegreeOfDanger applies equality check predicate on the "degreeOfDanger" field. It's identical to DegreeOfDangerEQ.
func DegreeOfDanger(v int) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldDegreeOfDanger, v))
}

// ThreatEQ applies the EQ predicate on the "threat" field.
func ThreatEQ(v string) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldThreat, v))
}

// ThreatNEQ applies the NEQ predicate on the "threat" field.
func ThreatNEQ(v string) predicate.Disease {
	return predicate.Disease(sql.FieldNEQ(FieldThreat, v))
}

// ThreatIn applies the In predicate on the "threat" field.
func ThreatIn(vs ...string) predicate.Disease {
	return predicate.Disease(sql.FieldIn(FieldThreat, vs...))
}

// ThreatNotIn applies the NotIn predicate on the "threat" field.
func ThreatNotIn(vs ...string) predicate.Disease {
	return predicate.Disease(sql.FieldNotIn(FieldThreat, vs...))
}

// ThreatGT applies the GT predicate on the "threat" field.
func ThreatGT(v string) predicate.Disease {
	return predicate.Disease(sql.FieldGT(FieldThreat, v))
}

// ThreatGTE applies the GTE predicate on the "threat" field.
func ThreatGTE(v string) predicate.Disease {
	return predicate.Disease(sql.FieldGTE(FieldThreat, v))
}

// ThreatLT applies the LT predicate on the "threat" field.
func ThreatLT(v string) predicate.Disease {
	return predicate.Disease(sql.FieldLT(FieldThreat, v))
}

// ThreatLTE applies the LTE predicate on the "threat" field.
func ThreatLTE(v string) predicate.Disease {
	return predicate.Disease(sql.FieldLTE(FieldThreat, v))
}

// ThreatContains applies the Contains predicate on the "threat" field.
func ThreatContains(v string) predicate.Disease {
	return predicate.Disease(sql.FieldContains(FieldThreat, v))
}

// ThreatHasPrefix applies the HasPrefix predicate on the "threat" field.
func ThreatHasPrefix(v string) predicate.Disease {
	return predicate.Disease(sql.FieldHasPrefix(FieldThreat, v))
}

// ThreatHasSuffix applies the HasSuffix predicate on the "threat" field.
func ThreatHasSuffix(v string) predicate.Disease {
	return predicate.Disease(sql.FieldHasSuffix(FieldThreat, v))
}

// ThreatEqualFold applies the EqualFold predicate on the "threat" field.
func ThreatEqualFold(v string) predicate.Disease {
	return predicate.Disease(sql.FieldEqualFold(FieldThreat, v))
}

// ThreatContainsFold applies the ContainsFold predicate on the "threat" field.
func ThreatContainsFold(v string) predicate.Disease {
	return predicate.Disease(sql.FieldContainsFold(FieldThreat, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Disease {
	return predicate.Disease(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Disease {
	return predicate.Disease(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Disease {
	return predicate.Disease(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Disease {
	return predicate.Disease(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Disease {
	return predicate.Disease(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Disease {
	return predicate.Disease(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Disease {
	return predicate.Disease(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Disease {
	return predicate.Disease(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Disease {
	return predicate.Disease(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Disease {
	return predicate.Disease(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Disease {
	return predicate.Disease(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Disease {
	return predicate.Disease(sql.FieldContainsFold(FieldName, v))
}

// DegreeOfDangerEQ applies the EQ predicate on the "degreeOfDanger" field.
func DegreeOfDangerEQ(v int) predicate.Disease {
	return predicate.Disease(sql.FieldEQ(FieldDegreeOfDanger, v))
}

// DegreeOfDangerNEQ applies the NEQ predicate on the "degreeOfDanger" field.
func DegreeOfDangerNEQ(v int) predicate.Disease {
	return predicate.Disease(sql.FieldNEQ(FieldDegreeOfDanger, v))
}

// DegreeOfDangerIn applies the In predicate on the "degreeOfDanger" field.
func DegreeOfDangerIn(vs ...int) predicate.Disease {
	return predicate.Disease(sql.FieldIn(FieldDegreeOfDanger, vs...))
}

// DegreeOfDangerNotIn applies the NotIn predicate on the "degreeOfDanger" field.
func DegreeOfDangerNotIn(vs ...int) predicate.Disease {
	return predicate.Disease(sql.FieldNotIn(FieldDegreeOfDanger, vs...))
}

// DegreeOfDangerGT applies the GT predicate on the "degreeOfDanger" field.
func DegreeOfDangerGT(v int) predicate.Disease {
	return predicate.Disease(sql.FieldGT(FieldDegreeOfDanger, v))
}

// DegreeOfDangerGTE applies the GTE predicate on the "degreeOfDanger" field.
func DegreeOfDangerGTE(v int) predicate.Disease {
	return predicate.Disease(sql.FieldGTE(FieldDegreeOfDanger, v))
}

// DegreeOfDangerLT applies the LT predicate on the "degreeOfDanger" field.
func DegreeOfDangerLT(v int) predicate.Disease {
	return predicate.Disease(sql.FieldLT(FieldDegreeOfDanger, v))
}

// DegreeOfDangerLTE applies the LTE predicate on the "degreeOfDanger" field.
func DegreeOfDangerLTE(v int) predicate.Disease {
	return predicate.Disease(sql.FieldLTE(FieldDegreeOfDanger, v))
}

// HasHas applies the HasEdge predicate on the "has" edge.
func HasHas() predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HasTable, HasColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasWith applies the HasEdge predicate on the "has" edge with a given conditions (other predicates).
func HasHasWith(preds ...predicate.Patient) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		step := newHasStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Disease) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Disease) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
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
func Not(p predicate.Disease) predicate.Disease {
	return predicate.Disease(func(s *sql.Selector) {
		p(s.Not())
	})
}
