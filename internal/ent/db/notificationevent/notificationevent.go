// Code generated by ent, DO NOT EDIT.

package notificationevent

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/internal/notification"
)

const (
	// Label holds the string label denoting the notificationevent type in the database.
	Label = "notification_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldRuleID holds the string denoting the rule_id field in the database.
	FieldRuleID = "rule_id"
	// FieldPayload holds the string denoting the payload field in the database.
	FieldPayload = "payload"
	// FieldHandlerDeduplicationHash holds the string denoting the handler_deduplication_hash field in the database.
	FieldHandlerDeduplicationHash = "handler_deduplication_hash"
	// EdgeDeliveryStatuses holds the string denoting the delivery_statuses edge name in mutations.
	EdgeDeliveryStatuses = "delivery_statuses"
	// EdgeRules holds the string denoting the rules edge name in mutations.
	EdgeRules = "rules"
	// Table holds the table name of the notificationevent in the database.
	Table = "notification_events"
	// DeliveryStatusesTable is the table that holds the delivery_statuses relation/edge. The primary key declared below.
	DeliveryStatusesTable = "notification_event_delivery_status_events"
	// DeliveryStatusesInverseTable is the table name for the NotificationEventDeliveryStatus entity.
	// It exists in this package in order to avoid circular dependency with the "notificationeventdeliverystatus" package.
	DeliveryStatusesInverseTable = "notification_event_delivery_status"
	// RulesTable is the table that holds the rules relation/edge.
	RulesTable = "notification_events"
	// RulesInverseTable is the table name for the NotificationRule entity.
	// It exists in this package in order to avoid circular dependency with the "notificationrule" package.
	RulesInverseTable = "notification_rules"
	// RulesColumn is the table column denoting the rules relation/edge.
	RulesColumn = "rule_id"
)

// Columns holds all SQL columns for notificationevent fields.
var Columns = []string{
	FieldID,
	FieldNamespace,
	FieldCreatedAt,
	FieldType,
	FieldRuleID,
	FieldPayload,
	FieldHandlerDeduplicationHash,
}

var (
	// DeliveryStatusesPrimaryKey and DeliveryStatusesColumn2 are the table columns denoting the
	// primary key for the delivery_statuses relation (M2M).
	DeliveryStatusesPrimaryKey = []string{"notification_event_delivery_status_id", "notification_event_id"}
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

var (
	// NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	NamespaceValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type notification.EventType) error {
	switch _type {
	case "entitlements.balance.threshold":
		return nil
	default:
		return fmt.Errorf("notificationevent: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the NotificationEvent queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByRuleID orders the results by the rule_id field.
func ByRuleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRuleID, opts...).ToFunc()
}

// ByPayload orders the results by the payload field.
func ByPayload(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayload, opts...).ToFunc()
}

// ByHandlerDeduplicationHash orders the results by the handler_deduplication_hash field.
func ByHandlerDeduplicationHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHandlerDeduplicationHash, opts...).ToFunc()
}

// ByDeliveryStatusesCount orders the results by delivery_statuses count.
func ByDeliveryStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeliveryStatusesStep(), opts...)
	}
}

// ByDeliveryStatuses orders the results by delivery_statuses terms.
func ByDeliveryStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeliveryStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRulesField orders the results by rules field.
func ByRulesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRulesStep(), sql.OrderByField(field, opts...))
	}
}
func newDeliveryStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeliveryStatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DeliveryStatusesTable, DeliveryStatusesPrimaryKey...),
	)
}
func newRulesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RulesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RulesTable, RulesColumn),
	)
}
