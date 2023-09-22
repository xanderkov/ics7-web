// Code generated by ent, DO NOT EDIT.

package account

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// EdgeIs holds the string denoting the is edge name in mutations.
	EdgeIs = "is"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// IsTable is the table that holds the is relation/edge. The primary key declared below.
	IsTable = "account_is"
	// IsInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	IsInverseTable = "doctors"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldPasswordHash,
}

var (
	// IsPrimaryKey and IsColumn2 are the table columns denoting the
	// primary key for the is relation (M2M).
	IsPrimaryKey = []string{"account_id", "doctor_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the Account queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLogin orders the results by the login field.
func ByLogin(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLogin, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByIsCount orders the results by is count.
func ByIsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIsStep(), opts...)
	}
}

// ByIs orders the results by is terms.
func ByIs(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newIsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, IsTable, IsPrimaryKey...),
	)
}