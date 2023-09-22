// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hospital/internal/modules/db/ent/treatment"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Treatment is the model entity for the Treatment schema.
type Treatment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Tablets holds the value of the "tablets" field.
	Tablets string `json:"tablets,omitempty"`
	// PsychologicalTreatment holds the value of the "psychologicalTreatment" field.
	PsychologicalTreatment string `json:"psychologicalTreatment,omitempty"`
	// Survey holds the value of the "survey" field.
	Survey string `json:"survey,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TreatmentQuery when eager-loading is set.
	Edges        TreatmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TreatmentEdges holds the relations/edges for other nodes in the graph.
type TreatmentEdges struct {
	// Cured holds the value of the cured edge.
	Cured []*Patient `json:"cured,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CuredOrErr returns the Cured value or an error if the edge
// was not loaded in eager-loading.
func (e TreatmentEdges) CuredOrErr() ([]*Patient, error) {
	if e.loadedTypes[0] {
		return e.Cured, nil
	}
	return nil, &NotLoadedError{edge: "cured"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Treatment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case treatment.FieldID:
			values[i] = new(sql.NullInt64)
		case treatment.FieldTablets, treatment.FieldPsychologicalTreatment, treatment.FieldSurvey:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Treatment fields.
func (t *Treatment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case treatment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case treatment.FieldTablets:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tablets", values[i])
			} else if value.Valid {
				t.Tablets = value.String
			}
		case treatment.FieldPsychologicalTreatment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field psychologicalTreatment", values[i])
			} else if value.Valid {
				t.PsychologicalTreatment = value.String
			}
		case treatment.FieldSurvey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field survey", values[i])
			} else if value.Valid {
				t.Survey = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Treatment.
// This includes values selected through modifiers, order, etc.
func (t *Treatment) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryCured queries the "cured" edge of the Treatment entity.
func (t *Treatment) QueryCured() *PatientQuery {
	return NewTreatmentClient(t.config).QueryCured(t)
}

// Update returns a builder for updating this Treatment.
// Note that you need to call Treatment.Unwrap() before calling this method if this Treatment
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Treatment) Update() *TreatmentUpdateOne {
	return NewTreatmentClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Treatment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Treatment) Unwrap() *Treatment {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Treatment is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Treatment) String() string {
	var builder strings.Builder
	builder.WriteString("Treatment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("tablets=")
	builder.WriteString(t.Tablets)
	builder.WriteString(", ")
	builder.WriteString("psychologicalTreatment=")
	builder.WriteString(t.PsychologicalTreatment)
	builder.WriteString(", ")
	builder.WriteString("survey=")
	builder.WriteString(t.Survey)
	builder.WriteByte(')')
	return builder.String()
}

// Treatments is a parsable slice of Treatment.
type Treatments []*Treatment