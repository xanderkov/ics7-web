// Code generated by ent, DO NOT EDIT.

package treatment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the treatment type in the database.
	Label = "treatment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTablets holds the string denoting the tablets field in the database.
	FieldTablets = "tablets"
	// FieldPsychologicalTreatment holds the string denoting the psychologicaltreatment field in the database.
	FieldPsychologicalTreatment = "psychological_treatment"
	// FieldSurvey holds the string denoting the survey field in the database.
	FieldSurvey = "survey"
	// EdgeCured holds the string denoting the cured edge name in mutations.
	EdgeCured = "cured"
	// Table holds the table name of the treatment in the database.
	Table = "treatments"
	// CuredTable is the table that holds the cured relation/edge.
	CuredTable = "patients"
	// CuredInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	CuredInverseTable = "patients"
	// CuredColumn is the table column denoting the cured relation/edge.
	CuredColumn = "treatment_cured"
)

// Columns holds all SQL columns for treatment fields.
var Columns = []string{
	FieldID,
	FieldTablets,
	FieldPsychologicalTreatment,
	FieldSurvey,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the Treatment queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTablets orders the results by the tablets field.
func ByTablets(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTablets, opts...).ToFunc()
}

// ByPsychologicalTreatment orders the results by the psychologicalTreatment field.
func ByPsychologicalTreatment(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPsychologicalTreatment, opts...).ToFunc()
}

// BySurvey orders the results by the survey field.
func BySurvey(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSurvey, opts...).ToFunc()
}

// ByCuredCount orders the results by cured count.
func ByCuredCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCuredStep(), opts...)
	}
}

// ByCured orders the results by cured terms.
func ByCured(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCuredStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCuredStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CuredInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CuredTable, CuredColumn),
	)
}