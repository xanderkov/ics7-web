// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hospital/internal/modules/db/ent/account"
	"hospital/internal/modules/db/ent/doctor"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Doctor is the model entity for the Doctor schema.
type Doctor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TokenId holds the value of the "tokenId" field.
	TokenId string `json:"tokenId,omitempty"`
	// Surname holds the value of the "surname" field.
	Surname string `json:"surname,omitempty"`
	// Speciality holds the value of the "speciality" field.
	Speciality string `json:"speciality,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DoctorQuery when eager-loading is set.
	Edges        DoctorEdges `json:"edges"`
	account_is   *int
	selectValues sql.SelectValues
}

// DoctorEdges holds the relations/edges for other nodes in the graph.
type DoctorEdges struct {
	// Treats holds the value of the treats edge.
	Treats []*Patient `json:"treats,omitempty"`
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TreatsOrErr returns the Treats value or an error if the edge
// was not loaded in eager-loading.
func (e DoctorEdges) TreatsOrErr() ([]*Patient, error) {
	if e.loadedTypes[0] {
		return e.Treats, nil
	}
	return nil, &NotLoadedError{edge: "treats"}
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DoctorEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[1] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Doctor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case doctor.FieldID:
			values[i] = new(sql.NullInt64)
		case doctor.FieldTokenId, doctor.FieldSurname, doctor.FieldSpeciality, doctor.FieldRole:
			values[i] = new(sql.NullString)
		case doctor.ForeignKeys[0]: // account_is
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Doctor fields.
func (d *Doctor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case doctor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case doctor.FieldTokenId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tokenId", values[i])
			} else if value.Valid {
				d.TokenId = value.String
			}
		case doctor.FieldSurname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field surname", values[i])
			} else if value.Valid {
				d.Surname = value.String
			}
		case doctor.FieldSpeciality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field speciality", values[i])
			} else if value.Valid {
				d.Speciality = value.String
			}
		case doctor.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				d.Role = value.String
			}
		case doctor.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_is", value)
			} else if value.Valid {
				d.account_is = new(int)
				*d.account_is = int(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Doctor.
// This includes values selected through modifiers, order, etc.
func (d *Doctor) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryTreats queries the "treats" edge of the Doctor entity.
func (d *Doctor) QueryTreats() *PatientQuery {
	return NewDoctorClient(d.config).QueryTreats(d)
}

// QueryAccount queries the "account" edge of the Doctor entity.
func (d *Doctor) QueryAccount() *AccountQuery {
	return NewDoctorClient(d.config).QueryAccount(d)
}

// Update returns a builder for updating this Doctor.
// Note that you need to call Doctor.Unwrap() before calling this method if this Doctor
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Doctor) Update() *DoctorUpdateOne {
	return NewDoctorClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Doctor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Doctor) Unwrap() *Doctor {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Doctor is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Doctor) String() string {
	var builder strings.Builder
	builder.WriteString("Doctor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("tokenId=")
	builder.WriteString(d.TokenId)
	builder.WriteString(", ")
	builder.WriteString("surname=")
	builder.WriteString(d.Surname)
	builder.WriteString(", ")
	builder.WriteString("speciality=")
	builder.WriteString(d.Speciality)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(d.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Doctors is a parsable slice of Doctor.
type Doctors []*Doctor
