// Code generated by ent, DO NOT EDIT.

package doctor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the doctor type in the database.
	Label = "doctor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTokenId holds the string denoting the tokenid field in the database.
	FieldTokenId = "token_id"
	// FieldSurname holds the string denoting the surname field in the database.
	FieldSurname = "surname"
	// FieldSpeciality holds the string denoting the speciality field in the database.
	FieldSpeciality = "speciality"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldPhotoPath holds the string denoting the photopath field in the database.
	FieldPhotoPath = "photo_path"
	// EdgeTreats holds the string denoting the treats edge name in mutations.
	EdgeTreats = "treats"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the doctor in the database.
	Table = "doctors"
	// TreatsTable is the table that holds the treats relation/edge. The primary key declared below.
	TreatsTable = "doctor_patient"
	// TreatsInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	TreatsInverseTable = "patients"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "doctors"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_is"
)

// Columns holds all SQL columns for doctor fields.
var Columns = []string{
	FieldID,
	FieldTokenId,
	FieldSurname,
	FieldSpeciality,
	FieldRole,
	FieldPhotoPath,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "doctors"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_is",
}

var (
	// TreatsPrimaryKey and TreatsColumn2 are the table columns denoting the
	// primary key for the treats relation (M2M).
	TreatsPrimaryKey = []string{"doctor_id", "patient_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the Doctor queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTokenId orders the results by the tokenId field.
func ByTokenId(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTokenId, opts...).ToFunc()
}

// BySurname orders the results by the surname field.
func BySurname(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSurname, opts...).ToFunc()
}

// BySpeciality orders the results by the speciality field.
func BySpeciality(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSpeciality, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByPhotoPath orders the results by the photoPath field.
func ByPhotoPath(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPhotoPath, opts...).ToFunc()
}

// ByTreatsCount orders the results by treats count.
func ByTreatsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTreatsStep(), opts...)
	}
}

// ByTreats orders the results by treats terms.
func ByTreats(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTreatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}
func newTreatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TreatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TreatsTable, TreatsPrimaryKey...),
	)
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
