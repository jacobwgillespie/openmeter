// Code generated by ent, DO NOT EDIT.

package creditentry

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldUpdatedAt, v))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldNamespace, v))
}

// LedgerID applies equality check predicate on the "ledger_id" field. It's identical to LedgerIDEQ.
func LedgerID(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldLedgerID, v))
}

// FeatureID applies equality check predicate on the "feature_id" field. It's identical to FeatureIDEQ.
func FeatureID(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldFeatureID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldAmount, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldPriority, v))
}

// EffectiveAt applies equality check predicate on the "effective_at" field. It's identical to EffectiveAtEQ.
func EffectiveAt(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldEffectiveAt, v))
}

// ExpirationPeriodCount applies equality check predicate on the "expiration_period_count" field. It's identical to ExpirationPeriodCountEQ.
func ExpirationPeriodCount(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldExpirationPeriodCount, v))
}

// ExpirationAt applies equality check predicate on the "expiration_at" field. It's identical to ExpirationAtEQ.
func ExpirationAt(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldExpirationAt, v))
}

// RolloverMaxAmount applies equality check predicate on the "rollover_max_amount" field. It's identical to RolloverMaxAmountEQ.
func RolloverMaxAmount(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldRolloverMaxAmount, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldParentID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldUpdatedAt, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContainsFold(FieldNamespace, v))
}

// LedgerIDEQ applies the EQ predicate on the "ledger_id" field.
func LedgerIDEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldLedgerID, v))
}

// LedgerIDNEQ applies the NEQ predicate on the "ledger_id" field.
func LedgerIDNEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldLedgerID, v))
}

// LedgerIDIn applies the In predicate on the "ledger_id" field.
func LedgerIDIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldLedgerID, vs...))
}

// LedgerIDNotIn applies the NotIn predicate on the "ledger_id" field.
func LedgerIDNotIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldLedgerID, vs...))
}

// LedgerIDGT applies the GT predicate on the "ledger_id" field.
func LedgerIDGT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldLedgerID, v))
}

// LedgerIDGTE applies the GTE predicate on the "ledger_id" field.
func LedgerIDGTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldLedgerID, v))
}

// LedgerIDLT applies the LT predicate on the "ledger_id" field.
func LedgerIDLT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldLedgerID, v))
}

// LedgerIDLTE applies the LTE predicate on the "ledger_id" field.
func LedgerIDLTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldLedgerID, v))
}

// LedgerIDContains applies the Contains predicate on the "ledger_id" field.
func LedgerIDContains(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContains(FieldLedgerID, v))
}

// LedgerIDHasPrefix applies the HasPrefix predicate on the "ledger_id" field.
func LedgerIDHasPrefix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasPrefix(FieldLedgerID, v))
}

// LedgerIDHasSuffix applies the HasSuffix predicate on the "ledger_id" field.
func LedgerIDHasSuffix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasSuffix(FieldLedgerID, v))
}

// LedgerIDEqualFold applies the EqualFold predicate on the "ledger_id" field.
func LedgerIDEqualFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEqualFold(FieldLedgerID, v))
}

// LedgerIDContainsFold applies the ContainsFold predicate on the "ledger_id" field.
func LedgerIDContainsFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContainsFold(FieldLedgerID, v))
}

// EntryTypeEQ applies the EQ predicate on the "entry_type" field.
func EntryTypeEQ(v credit.EntryType) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldEQ(FieldEntryType, vc))
}

// EntryTypeNEQ applies the NEQ predicate on the "entry_type" field.
func EntryTypeNEQ(v credit.EntryType) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldNEQ(FieldEntryType, vc))
}

// EntryTypeIn applies the In predicate on the "entry_type" field.
func EntryTypeIn(vs ...credit.EntryType) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldIn(FieldEntryType, v...))
}

// EntryTypeNotIn applies the NotIn predicate on the "entry_type" field.
func EntryTypeNotIn(vs ...credit.EntryType) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldNotIn(FieldEntryType, v...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v credit.GrantType) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v credit.GrantType) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...credit.GrantType) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...credit.GrantType) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldNotIn(FieldType, v...))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldType))
}

// FeatureIDEQ applies the EQ predicate on the "feature_id" field.
func FeatureIDEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldFeatureID, v))
}

// FeatureIDNEQ applies the NEQ predicate on the "feature_id" field.
func FeatureIDNEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldFeatureID, v))
}

// FeatureIDIn applies the In predicate on the "feature_id" field.
func FeatureIDIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldFeatureID, vs...))
}

// FeatureIDNotIn applies the NotIn predicate on the "feature_id" field.
func FeatureIDNotIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldFeatureID, vs...))
}

// FeatureIDGT applies the GT predicate on the "feature_id" field.
func FeatureIDGT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldFeatureID, v))
}

// FeatureIDGTE applies the GTE predicate on the "feature_id" field.
func FeatureIDGTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldFeatureID, v))
}

// FeatureIDLT applies the LT predicate on the "feature_id" field.
func FeatureIDLT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldFeatureID, v))
}

// FeatureIDLTE applies the LTE predicate on the "feature_id" field.
func FeatureIDLTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldFeatureID, v))
}

// FeatureIDContains applies the Contains predicate on the "feature_id" field.
func FeatureIDContains(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContains(FieldFeatureID, v))
}

// FeatureIDHasPrefix applies the HasPrefix predicate on the "feature_id" field.
func FeatureIDHasPrefix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasPrefix(FieldFeatureID, v))
}

// FeatureIDHasSuffix applies the HasSuffix predicate on the "feature_id" field.
func FeatureIDHasSuffix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasSuffix(FieldFeatureID, v))
}

// FeatureIDIsNil applies the IsNil predicate on the "feature_id" field.
func FeatureIDIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldFeatureID))
}

// FeatureIDNotNil applies the NotNil predicate on the "feature_id" field.
func FeatureIDNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldFeatureID))
}

// FeatureIDEqualFold applies the EqualFold predicate on the "feature_id" field.
func FeatureIDEqualFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEqualFold(FieldFeatureID, v))
}

// FeatureIDContainsFold applies the ContainsFold predicate on the "feature_id" field.
func FeatureIDContainsFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContainsFold(FieldFeatureID, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldAmount, v))
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldAmount))
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldAmount))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldPriority, v))
}

// EffectiveAtEQ applies the EQ predicate on the "effective_at" field.
func EffectiveAtEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldEffectiveAt, v))
}

// EffectiveAtNEQ applies the NEQ predicate on the "effective_at" field.
func EffectiveAtNEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldEffectiveAt, v))
}

// EffectiveAtIn applies the In predicate on the "effective_at" field.
func EffectiveAtIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldEffectiveAt, vs...))
}

// EffectiveAtNotIn applies the NotIn predicate on the "effective_at" field.
func EffectiveAtNotIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldEffectiveAt, vs...))
}

// EffectiveAtGT applies the GT predicate on the "effective_at" field.
func EffectiveAtGT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldEffectiveAt, v))
}

// EffectiveAtGTE applies the GTE predicate on the "effective_at" field.
func EffectiveAtGTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldEffectiveAt, v))
}

// EffectiveAtLT applies the LT predicate on the "effective_at" field.
func EffectiveAtLT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldEffectiveAt, v))
}

// EffectiveAtLTE applies the LTE predicate on the "effective_at" field.
func EffectiveAtLTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldEffectiveAt, v))
}

// ExpirationPeriodDurationEQ applies the EQ predicate on the "expiration_period_duration" field.
func ExpirationPeriodDurationEQ(v credit.ExpirationPeriodDuration) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldEQ(FieldExpirationPeriodDuration, vc))
}

// ExpirationPeriodDurationNEQ applies the NEQ predicate on the "expiration_period_duration" field.
func ExpirationPeriodDurationNEQ(v credit.ExpirationPeriodDuration) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldNEQ(FieldExpirationPeriodDuration, vc))
}

// ExpirationPeriodDurationIn applies the In predicate on the "expiration_period_duration" field.
func ExpirationPeriodDurationIn(vs ...credit.ExpirationPeriodDuration) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldIn(FieldExpirationPeriodDuration, v...))
}

// ExpirationPeriodDurationNotIn applies the NotIn predicate on the "expiration_period_duration" field.
func ExpirationPeriodDurationNotIn(vs ...credit.ExpirationPeriodDuration) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldNotIn(FieldExpirationPeriodDuration, v...))
}

// ExpirationPeriodDurationIsNil applies the IsNil predicate on the "expiration_period_duration" field.
func ExpirationPeriodDurationIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldExpirationPeriodDuration))
}

// ExpirationPeriodDurationNotNil applies the NotNil predicate on the "expiration_period_duration" field.
func ExpirationPeriodDurationNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldExpirationPeriodDuration))
}

// ExpirationPeriodCountEQ applies the EQ predicate on the "expiration_period_count" field.
func ExpirationPeriodCountEQ(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldExpirationPeriodCount, v))
}

// ExpirationPeriodCountNEQ applies the NEQ predicate on the "expiration_period_count" field.
func ExpirationPeriodCountNEQ(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldExpirationPeriodCount, v))
}

// ExpirationPeriodCountIn applies the In predicate on the "expiration_period_count" field.
func ExpirationPeriodCountIn(vs ...uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldExpirationPeriodCount, vs...))
}

// ExpirationPeriodCountNotIn applies the NotIn predicate on the "expiration_period_count" field.
func ExpirationPeriodCountNotIn(vs ...uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldExpirationPeriodCount, vs...))
}

// ExpirationPeriodCountGT applies the GT predicate on the "expiration_period_count" field.
func ExpirationPeriodCountGT(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldExpirationPeriodCount, v))
}

// ExpirationPeriodCountGTE applies the GTE predicate on the "expiration_period_count" field.
func ExpirationPeriodCountGTE(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldExpirationPeriodCount, v))
}

// ExpirationPeriodCountLT applies the LT predicate on the "expiration_period_count" field.
func ExpirationPeriodCountLT(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldExpirationPeriodCount, v))
}

// ExpirationPeriodCountLTE applies the LTE predicate on the "expiration_period_count" field.
func ExpirationPeriodCountLTE(v uint8) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldExpirationPeriodCount, v))
}

// ExpirationPeriodCountIsNil applies the IsNil predicate on the "expiration_period_count" field.
func ExpirationPeriodCountIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldExpirationPeriodCount))
}

// ExpirationPeriodCountNotNil applies the NotNil predicate on the "expiration_period_count" field.
func ExpirationPeriodCountNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldExpirationPeriodCount))
}

// ExpirationAtEQ applies the EQ predicate on the "expiration_at" field.
func ExpirationAtEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldExpirationAt, v))
}

// ExpirationAtNEQ applies the NEQ predicate on the "expiration_at" field.
func ExpirationAtNEQ(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldExpirationAt, v))
}

// ExpirationAtIn applies the In predicate on the "expiration_at" field.
func ExpirationAtIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldExpirationAt, vs...))
}

// ExpirationAtNotIn applies the NotIn predicate on the "expiration_at" field.
func ExpirationAtNotIn(vs ...time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldExpirationAt, vs...))
}

// ExpirationAtGT applies the GT predicate on the "expiration_at" field.
func ExpirationAtGT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldExpirationAt, v))
}

// ExpirationAtGTE applies the GTE predicate on the "expiration_at" field.
func ExpirationAtGTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldExpirationAt, v))
}

// ExpirationAtLT applies the LT predicate on the "expiration_at" field.
func ExpirationAtLT(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldExpirationAt, v))
}

// ExpirationAtLTE applies the LTE predicate on the "expiration_at" field.
func ExpirationAtLTE(v time.Time) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldExpirationAt, v))
}

// ExpirationAtIsNil applies the IsNil predicate on the "expiration_at" field.
func ExpirationAtIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldExpirationAt))
}

// ExpirationAtNotNil applies the NotNil predicate on the "expiration_at" field.
func ExpirationAtNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldExpirationAt))
}

// RolloverTypeEQ applies the EQ predicate on the "rollover_type" field.
func RolloverTypeEQ(v credit.GrantRolloverType) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldEQ(FieldRolloverType, vc))
}

// RolloverTypeNEQ applies the NEQ predicate on the "rollover_type" field.
func RolloverTypeNEQ(v credit.GrantRolloverType) predicate.CreditEntry {
	vc := v
	return predicate.CreditEntry(sql.FieldNEQ(FieldRolloverType, vc))
}

// RolloverTypeIn applies the In predicate on the "rollover_type" field.
func RolloverTypeIn(vs ...credit.GrantRolloverType) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldIn(FieldRolloverType, v...))
}

// RolloverTypeNotIn applies the NotIn predicate on the "rollover_type" field.
func RolloverTypeNotIn(vs ...credit.GrantRolloverType) predicate.CreditEntry {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CreditEntry(sql.FieldNotIn(FieldRolloverType, v...))
}

// RolloverTypeIsNil applies the IsNil predicate on the "rollover_type" field.
func RolloverTypeIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldRolloverType))
}

// RolloverTypeNotNil applies the NotNil predicate on the "rollover_type" field.
func RolloverTypeNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldRolloverType))
}

// RolloverMaxAmountEQ applies the EQ predicate on the "rollover_max_amount" field.
func RolloverMaxAmountEQ(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldRolloverMaxAmount, v))
}

// RolloverMaxAmountNEQ applies the NEQ predicate on the "rollover_max_amount" field.
func RolloverMaxAmountNEQ(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldRolloverMaxAmount, v))
}

// RolloverMaxAmountIn applies the In predicate on the "rollover_max_amount" field.
func RolloverMaxAmountIn(vs ...float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldRolloverMaxAmount, vs...))
}

// RolloverMaxAmountNotIn applies the NotIn predicate on the "rollover_max_amount" field.
func RolloverMaxAmountNotIn(vs ...float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldRolloverMaxAmount, vs...))
}

// RolloverMaxAmountGT applies the GT predicate on the "rollover_max_amount" field.
func RolloverMaxAmountGT(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldRolloverMaxAmount, v))
}

// RolloverMaxAmountGTE applies the GTE predicate on the "rollover_max_amount" field.
func RolloverMaxAmountGTE(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldRolloverMaxAmount, v))
}

// RolloverMaxAmountLT applies the LT predicate on the "rollover_max_amount" field.
func RolloverMaxAmountLT(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldRolloverMaxAmount, v))
}

// RolloverMaxAmountLTE applies the LTE predicate on the "rollover_max_amount" field.
func RolloverMaxAmountLTE(v float64) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldRolloverMaxAmount, v))
}

// RolloverMaxAmountIsNil applies the IsNil predicate on the "rollover_max_amount" field.
func RolloverMaxAmountIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldRolloverMaxAmount))
}

// RolloverMaxAmountNotNil applies the NotNil predicate on the "rollover_max_amount" field.
func RolloverMaxAmountNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldRolloverMaxAmount))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldMetadata))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldLTE(FieldParentID, v))
}

// ParentIDContains applies the Contains predicate on the "parent_id" field.
func ParentIDContains(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContains(FieldParentID, v))
}

// ParentIDHasPrefix applies the HasPrefix predicate on the "parent_id" field.
func ParentIDHasPrefix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasPrefix(FieldParentID, v))
}

// ParentIDHasSuffix applies the HasSuffix predicate on the "parent_id" field.
func ParentIDHasSuffix(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldHasSuffix(FieldParentID, v))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldNotNull(FieldParentID))
}

// ParentIDEqualFold applies the EqualFold predicate on the "parent_id" field.
func ParentIDEqualFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldEqualFold(FieldParentID, v))
}

// ParentIDContainsFold applies the ContainsFold predicate on the "parent_id" field.
func ParentIDContainsFold(v string) predicate.CreditEntry {
	return predicate.CreditEntry(sql.FieldContainsFold(FieldParentID, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.CreditEntry) predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.CreditEntry) predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeature applies the HasEdge predicate on the "feature" edge.
func HasFeature() predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FeatureTable, FeatureColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeatureWith applies the HasEdge predicate on the "feature" edge with a given conditions (other predicates).
func HasFeatureWith(preds ...predicate.Feature) predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := newFeatureStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLedger applies the HasEdge predicate on the "ledger" edge.
func HasLedger() predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LedgerTable, LedgerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLedgerWith applies the HasEdge predicate on the "ledger" edge with a given conditions (other predicates).
func HasLedgerWith(preds ...predicate.Ledger) predicate.CreditEntry {
	return predicate.CreditEntry(func(s *sql.Selector) {
		step := newLedgerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CreditEntry) predicate.CreditEntry {
	return predicate.CreditEntry(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CreditEntry) predicate.CreditEntry {
	return predicate.CreditEntry(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CreditEntry) predicate.CreditEntry {
	return predicate.CreditEntry(sql.NotPredicates(p))
}
