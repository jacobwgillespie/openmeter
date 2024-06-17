// Code generated by ent, DO NOT EDIT.

package ledger

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the ledger type in the database.
	Label = "ledger"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldHighwatermark holds the string denoting the highwatermark field in the database.
	FieldHighwatermark = "highwatermark"
	// EdgeCreditGrants holds the string denoting the credit_grants edge name in mutations.
	EdgeCreditGrants = "credit_grants"
	// Table holds the table name of the ledger in the database.
	Table = "ledgers"
	// CreditGrantsTable is the table that holds the credit_grants relation/edge.
	CreditGrantsTable = "credit_entries"
	// CreditGrantsInverseTable is the table name for the CreditEntry entity.
	// It exists in this package in order to avoid circular dependency with the "creditentry" package.
	CreditGrantsInverseTable = "credit_entries"
	// CreditGrantsColumn is the table column denoting the credit_grants relation/edge.
	CreditGrantsColumn = "ledger_id"
)

// Columns holds all SQL columns for ledger fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldNamespace,
	FieldSubject,
	FieldMetadata,
	FieldHighwatermark,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	NamespaceValidator func(string) error
	// SubjectValidator is a validator for the "subject" field. It is called by the builders before save.
	SubjectValidator func(string) error
	// DefaultHighwatermark holds the default value on creation for the "highwatermark" field.
	DefaultHighwatermark func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Ledger queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByHighwatermark orders the results by the highwatermark field.
func ByHighwatermark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHighwatermark, opts...).ToFunc()
}

// ByCreditGrantsCount orders the results by credit_grants count.
func ByCreditGrantsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreditGrantsStep(), opts...)
	}
}

// ByCreditGrants orders the results by credit_grants terms.
func ByCreditGrants(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreditGrantsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCreditGrantsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreditGrantsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreditGrantsTable, CreditGrantsColumn),
	)
}
