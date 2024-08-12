// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/ent/db/balancesnapshot"
	"github.com/openmeterio/openmeter/internal/ent/db/entitlement"
	dbgrant "github.com/openmeterio/openmeter/internal/ent/db/grant"
	"github.com/openmeterio/openmeter/internal/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/ent/db/usagereset"
)

// EntitlementUpdate is the builder for updating Entitlement entities.
type EntitlementUpdate struct {
	config
	hooks    []Hook
	mutation *EntitlementMutation
}

// Where appends a list predicates to the EntitlementUpdate builder.
func (eu *EntitlementUpdate) Where(ps ...predicate.Entitlement) *EntitlementUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetMetadata sets the "metadata" field.
func (eu *EntitlementUpdate) SetMetadata(m map[string]string) *EntitlementUpdate {
	eu.mutation.SetMetadata(m)
	return eu
}

// ClearMetadata clears the value of the "metadata" field.
func (eu *EntitlementUpdate) ClearMetadata() *EntitlementUpdate {
	eu.mutation.ClearMetadata()
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EntitlementUpdate) SetUpdatedAt(t time.Time) *EntitlementUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetDeletedAt sets the "deleted_at" field.
func (eu *EntitlementUpdate) SetDeletedAt(t time.Time) *EntitlementUpdate {
	eu.mutation.SetDeletedAt(t)
	return eu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableDeletedAt(t *time.Time) *EntitlementUpdate {
	if t != nil {
		eu.SetDeletedAt(*t)
	}
	return eu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (eu *EntitlementUpdate) ClearDeletedAt() *EntitlementUpdate {
	eu.mutation.ClearDeletedAt()
	return eu
}

// SetConfig sets the "config" field.
func (eu *EntitlementUpdate) SetConfig(u []uint8) *EntitlementUpdate {
	eu.mutation.SetConfig(u)
	return eu
}

// AppendConfig appends u to the "config" field.
func (eu *EntitlementUpdate) AppendConfig(u []uint8) *EntitlementUpdate {
	eu.mutation.AppendConfig(u)
	return eu
}

// ClearConfig clears the value of the "config" field.
func (eu *EntitlementUpdate) ClearConfig() *EntitlementUpdate {
	eu.mutation.ClearConfig()
	return eu
}

// SetUsagePeriodAnchor sets the "usage_period_anchor" field.
func (eu *EntitlementUpdate) SetUsagePeriodAnchor(t time.Time) *EntitlementUpdate {
	eu.mutation.SetUsagePeriodAnchor(t)
	return eu
}

// SetNillableUsagePeriodAnchor sets the "usage_period_anchor" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableUsagePeriodAnchor(t *time.Time) *EntitlementUpdate {
	if t != nil {
		eu.SetUsagePeriodAnchor(*t)
	}
	return eu
}

// ClearUsagePeriodAnchor clears the value of the "usage_period_anchor" field.
func (eu *EntitlementUpdate) ClearUsagePeriodAnchor() *EntitlementUpdate {
	eu.mutation.ClearUsagePeriodAnchor()
	return eu
}

// SetCurrentUsagePeriodStart sets the "current_usage_period_start" field.
func (eu *EntitlementUpdate) SetCurrentUsagePeriodStart(t time.Time) *EntitlementUpdate {
	eu.mutation.SetCurrentUsagePeriodStart(t)
	return eu
}

// SetNillableCurrentUsagePeriodStart sets the "current_usage_period_start" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableCurrentUsagePeriodStart(t *time.Time) *EntitlementUpdate {
	if t != nil {
		eu.SetCurrentUsagePeriodStart(*t)
	}
	return eu
}

// ClearCurrentUsagePeriodStart clears the value of the "current_usage_period_start" field.
func (eu *EntitlementUpdate) ClearCurrentUsagePeriodStart() *EntitlementUpdate {
	eu.mutation.ClearCurrentUsagePeriodStart()
	return eu
}

// SetCurrentUsagePeriodEnd sets the "current_usage_period_end" field.
func (eu *EntitlementUpdate) SetCurrentUsagePeriodEnd(t time.Time) *EntitlementUpdate {
	eu.mutation.SetCurrentUsagePeriodEnd(t)
	return eu
}

// SetNillableCurrentUsagePeriodEnd sets the "current_usage_period_end" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableCurrentUsagePeriodEnd(t *time.Time) *EntitlementUpdate {
	if t != nil {
		eu.SetCurrentUsagePeriodEnd(*t)
	}
	return eu
}

// ClearCurrentUsagePeriodEnd clears the value of the "current_usage_period_end" field.
func (eu *EntitlementUpdate) ClearCurrentUsagePeriodEnd() *EntitlementUpdate {
	eu.mutation.ClearCurrentUsagePeriodEnd()
	return eu
}

// AddUsageResetIDs adds the "usage_reset" edge to the UsageReset entity by IDs.
func (eu *EntitlementUpdate) AddUsageResetIDs(ids ...string) *EntitlementUpdate {
	eu.mutation.AddUsageResetIDs(ids...)
	return eu
}

// AddUsageReset adds the "usage_reset" edges to the UsageReset entity.
func (eu *EntitlementUpdate) AddUsageReset(u ...*UsageReset) *EntitlementUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.AddUsageResetIDs(ids...)
}

// AddGrantIDs adds the "grant" edge to the Grant entity by IDs.
func (eu *EntitlementUpdate) AddGrantIDs(ids ...string) *EntitlementUpdate {
	eu.mutation.AddGrantIDs(ids...)
	return eu
}

// AddGrant adds the "grant" edges to the Grant entity.
func (eu *EntitlementUpdate) AddGrant(g ...*Grant) *EntitlementUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return eu.AddGrantIDs(ids...)
}

// AddBalanceSnapshotIDs adds the "balance_snapshot" edge to the BalanceSnapshot entity by IDs.
func (eu *EntitlementUpdate) AddBalanceSnapshotIDs(ids ...int) *EntitlementUpdate {
	eu.mutation.AddBalanceSnapshotIDs(ids...)
	return eu
}

// AddBalanceSnapshot adds the "balance_snapshot" edges to the BalanceSnapshot entity.
func (eu *EntitlementUpdate) AddBalanceSnapshot(b ...*BalanceSnapshot) *EntitlementUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return eu.AddBalanceSnapshotIDs(ids...)
}

// Mutation returns the EntitlementMutation object of the builder.
func (eu *EntitlementUpdate) Mutation() *EntitlementMutation {
	return eu.mutation
}

// ClearUsageReset clears all "usage_reset" edges to the UsageReset entity.
func (eu *EntitlementUpdate) ClearUsageReset() *EntitlementUpdate {
	eu.mutation.ClearUsageReset()
	return eu
}

// RemoveUsageResetIDs removes the "usage_reset" edge to UsageReset entities by IDs.
func (eu *EntitlementUpdate) RemoveUsageResetIDs(ids ...string) *EntitlementUpdate {
	eu.mutation.RemoveUsageResetIDs(ids...)
	return eu
}

// RemoveUsageReset removes "usage_reset" edges to UsageReset entities.
func (eu *EntitlementUpdate) RemoveUsageReset(u ...*UsageReset) *EntitlementUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.RemoveUsageResetIDs(ids...)
}

// ClearGrant clears all "grant" edges to the Grant entity.
func (eu *EntitlementUpdate) ClearGrant() *EntitlementUpdate {
	eu.mutation.ClearGrant()
	return eu
}

// RemoveGrantIDs removes the "grant" edge to Grant entities by IDs.
func (eu *EntitlementUpdate) RemoveGrantIDs(ids ...string) *EntitlementUpdate {
	eu.mutation.RemoveGrantIDs(ids...)
	return eu
}

// RemoveGrant removes "grant" edges to Grant entities.
func (eu *EntitlementUpdate) RemoveGrant(g ...*Grant) *EntitlementUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return eu.RemoveGrantIDs(ids...)
}

// ClearBalanceSnapshot clears all "balance_snapshot" edges to the BalanceSnapshot entity.
func (eu *EntitlementUpdate) ClearBalanceSnapshot() *EntitlementUpdate {
	eu.mutation.ClearBalanceSnapshot()
	return eu
}

// RemoveBalanceSnapshotIDs removes the "balance_snapshot" edge to BalanceSnapshot entities by IDs.
func (eu *EntitlementUpdate) RemoveBalanceSnapshotIDs(ids ...int) *EntitlementUpdate {
	eu.mutation.RemoveBalanceSnapshotIDs(ids...)
	return eu
}

// RemoveBalanceSnapshot removes "balance_snapshot" edges to BalanceSnapshot entities.
func (eu *EntitlementUpdate) RemoveBalanceSnapshot(b ...*BalanceSnapshot) *EntitlementUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return eu.RemoveBalanceSnapshotIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntitlementUpdate) Save(ctx context.Context) (int, error) {
	eu.defaults()
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntitlementUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntitlementUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntitlementUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EntitlementUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := entitlement.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EntitlementUpdate) check() error {
	if eu.mutation.FeatureCleared() && len(eu.mutation.FeatureIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "Entitlement.feature"`)
	}
	return nil
}

func (eu *EntitlementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(entitlement.Table, entitlement.Columns, sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Metadata(); ok {
		_spec.SetField(entitlement.FieldMetadata, field.TypeJSON, value)
	}
	if eu.mutation.MetadataCleared() {
		_spec.ClearField(entitlement.FieldMetadata, field.TypeJSON)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlement.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.SetField(entitlement.FieldDeletedAt, field.TypeTime, value)
	}
	if eu.mutation.DeletedAtCleared() {
		_spec.ClearField(entitlement.FieldDeletedAt, field.TypeTime)
	}
	if eu.mutation.MeasureUsageFromCleared() {
		_spec.ClearField(entitlement.FieldMeasureUsageFrom, field.TypeTime)
	}
	if eu.mutation.IssueAfterResetCleared() {
		_spec.ClearField(entitlement.FieldIssueAfterReset, field.TypeFloat64)
	}
	if eu.mutation.IssueAfterResetPriorityCleared() {
		_spec.ClearField(entitlement.FieldIssueAfterResetPriority, field.TypeUint8)
	}
	if eu.mutation.IsSoftLimitCleared() {
		_spec.ClearField(entitlement.FieldIsSoftLimit, field.TypeBool)
	}
	if eu.mutation.PreserveOverageAtResetCleared() {
		_spec.ClearField(entitlement.FieldPreserveOverageAtReset, field.TypeBool)
	}
	if value, ok := eu.mutation.Config(); ok {
		_spec.SetField(entitlement.FieldConfig, field.TypeJSON, value)
	}
	if value, ok := eu.mutation.AppendedConfig(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, entitlement.FieldConfig, value)
		})
	}
	if eu.mutation.ConfigCleared() {
		_spec.ClearField(entitlement.FieldConfig, field.TypeJSON)
	}
	if eu.mutation.UsagePeriodIntervalCleared() {
		_spec.ClearField(entitlement.FieldUsagePeriodInterval, field.TypeEnum)
	}
	if value, ok := eu.mutation.UsagePeriodAnchor(); ok {
		_spec.SetField(entitlement.FieldUsagePeriodAnchor, field.TypeTime, value)
	}
	if eu.mutation.UsagePeriodAnchorCleared() {
		_spec.ClearField(entitlement.FieldUsagePeriodAnchor, field.TypeTime)
	}
	if value, ok := eu.mutation.CurrentUsagePeriodStart(); ok {
		_spec.SetField(entitlement.FieldCurrentUsagePeriodStart, field.TypeTime, value)
	}
	if eu.mutation.CurrentUsagePeriodStartCleared() {
		_spec.ClearField(entitlement.FieldCurrentUsagePeriodStart, field.TypeTime)
	}
	if value, ok := eu.mutation.CurrentUsagePeriodEnd(); ok {
		_spec.SetField(entitlement.FieldCurrentUsagePeriodEnd, field.TypeTime, value)
	}
	if eu.mutation.CurrentUsagePeriodEndCleared() {
		_spec.ClearField(entitlement.FieldCurrentUsagePeriodEnd, field.TypeTime)
	}
	if eu.mutation.UsageResetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.UsageResetTable,
			Columns: []string{entitlement.UsageResetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedUsageResetIDs(); len(nodes) > 0 && !eu.mutation.UsageResetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.UsageResetTable,
			Columns: []string{entitlement.UsageResetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UsageResetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.UsageResetTable,
			Columns: []string{entitlement.UsageResetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.GrantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.GrantTable,
			Columns: []string{entitlement.GrantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dbgrant.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedGrantIDs(); len(nodes) > 0 && !eu.mutation.GrantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.GrantTable,
			Columns: []string{entitlement.GrantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dbgrant.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.GrantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.GrantTable,
			Columns: []string{entitlement.GrantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dbgrant.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.BalanceSnapshotCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.BalanceSnapshotTable,
			Columns: []string{entitlement.BalanceSnapshotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedBalanceSnapshotIDs(); len(nodes) > 0 && !eu.mutation.BalanceSnapshotCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.BalanceSnapshotTable,
			Columns: []string{entitlement.BalanceSnapshotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.BalanceSnapshotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.BalanceSnapshotTable,
			Columns: []string{entitlement.BalanceSnapshotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitlement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EntitlementUpdateOne is the builder for updating a single Entitlement entity.
type EntitlementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntitlementMutation
}

// SetMetadata sets the "metadata" field.
func (euo *EntitlementUpdateOne) SetMetadata(m map[string]string) *EntitlementUpdateOne {
	euo.mutation.SetMetadata(m)
	return euo
}

// ClearMetadata clears the value of the "metadata" field.
func (euo *EntitlementUpdateOne) ClearMetadata() *EntitlementUpdateOne {
	euo.mutation.ClearMetadata()
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EntitlementUpdateOne) SetUpdatedAt(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetDeletedAt sets the "deleted_at" field.
func (euo *EntitlementUpdateOne) SetDeletedAt(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetDeletedAt(t)
	return euo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableDeletedAt(t *time.Time) *EntitlementUpdateOne {
	if t != nil {
		euo.SetDeletedAt(*t)
	}
	return euo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (euo *EntitlementUpdateOne) ClearDeletedAt() *EntitlementUpdateOne {
	euo.mutation.ClearDeletedAt()
	return euo
}

// SetConfig sets the "config" field.
func (euo *EntitlementUpdateOne) SetConfig(u []uint8) *EntitlementUpdateOne {
	euo.mutation.SetConfig(u)
	return euo
}

// AppendConfig appends u to the "config" field.
func (euo *EntitlementUpdateOne) AppendConfig(u []uint8) *EntitlementUpdateOne {
	euo.mutation.AppendConfig(u)
	return euo
}

// ClearConfig clears the value of the "config" field.
func (euo *EntitlementUpdateOne) ClearConfig() *EntitlementUpdateOne {
	euo.mutation.ClearConfig()
	return euo
}

// SetUsagePeriodAnchor sets the "usage_period_anchor" field.
func (euo *EntitlementUpdateOne) SetUsagePeriodAnchor(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetUsagePeriodAnchor(t)
	return euo
}

// SetNillableUsagePeriodAnchor sets the "usage_period_anchor" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableUsagePeriodAnchor(t *time.Time) *EntitlementUpdateOne {
	if t != nil {
		euo.SetUsagePeriodAnchor(*t)
	}
	return euo
}

// ClearUsagePeriodAnchor clears the value of the "usage_period_anchor" field.
func (euo *EntitlementUpdateOne) ClearUsagePeriodAnchor() *EntitlementUpdateOne {
	euo.mutation.ClearUsagePeriodAnchor()
	return euo
}

// SetCurrentUsagePeriodStart sets the "current_usage_period_start" field.
func (euo *EntitlementUpdateOne) SetCurrentUsagePeriodStart(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetCurrentUsagePeriodStart(t)
	return euo
}

// SetNillableCurrentUsagePeriodStart sets the "current_usage_period_start" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableCurrentUsagePeriodStart(t *time.Time) *EntitlementUpdateOne {
	if t != nil {
		euo.SetCurrentUsagePeriodStart(*t)
	}
	return euo
}

// ClearCurrentUsagePeriodStart clears the value of the "current_usage_period_start" field.
func (euo *EntitlementUpdateOne) ClearCurrentUsagePeriodStart() *EntitlementUpdateOne {
	euo.mutation.ClearCurrentUsagePeriodStart()
	return euo
}

// SetCurrentUsagePeriodEnd sets the "current_usage_period_end" field.
func (euo *EntitlementUpdateOne) SetCurrentUsagePeriodEnd(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetCurrentUsagePeriodEnd(t)
	return euo
}

// SetNillableCurrentUsagePeriodEnd sets the "current_usage_period_end" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableCurrentUsagePeriodEnd(t *time.Time) *EntitlementUpdateOne {
	if t != nil {
		euo.SetCurrentUsagePeriodEnd(*t)
	}
	return euo
}

// ClearCurrentUsagePeriodEnd clears the value of the "current_usage_period_end" field.
func (euo *EntitlementUpdateOne) ClearCurrentUsagePeriodEnd() *EntitlementUpdateOne {
	euo.mutation.ClearCurrentUsagePeriodEnd()
	return euo
}

// AddUsageResetIDs adds the "usage_reset" edge to the UsageReset entity by IDs.
func (euo *EntitlementUpdateOne) AddUsageResetIDs(ids ...string) *EntitlementUpdateOne {
	euo.mutation.AddUsageResetIDs(ids...)
	return euo
}

// AddUsageReset adds the "usage_reset" edges to the UsageReset entity.
func (euo *EntitlementUpdateOne) AddUsageReset(u ...*UsageReset) *EntitlementUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.AddUsageResetIDs(ids...)
}

// AddGrantIDs adds the "grant" edge to the Grant entity by IDs.
func (euo *EntitlementUpdateOne) AddGrantIDs(ids ...string) *EntitlementUpdateOne {
	euo.mutation.AddGrantIDs(ids...)
	return euo
}

// AddGrant adds the "grant" edges to the Grant entity.
func (euo *EntitlementUpdateOne) AddGrant(g ...*Grant) *EntitlementUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return euo.AddGrantIDs(ids...)
}

// AddBalanceSnapshotIDs adds the "balance_snapshot" edge to the BalanceSnapshot entity by IDs.
func (euo *EntitlementUpdateOne) AddBalanceSnapshotIDs(ids ...int) *EntitlementUpdateOne {
	euo.mutation.AddBalanceSnapshotIDs(ids...)
	return euo
}

// AddBalanceSnapshot adds the "balance_snapshot" edges to the BalanceSnapshot entity.
func (euo *EntitlementUpdateOne) AddBalanceSnapshot(b ...*BalanceSnapshot) *EntitlementUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return euo.AddBalanceSnapshotIDs(ids...)
}

// Mutation returns the EntitlementMutation object of the builder.
func (euo *EntitlementUpdateOne) Mutation() *EntitlementMutation {
	return euo.mutation
}

// ClearUsageReset clears all "usage_reset" edges to the UsageReset entity.
func (euo *EntitlementUpdateOne) ClearUsageReset() *EntitlementUpdateOne {
	euo.mutation.ClearUsageReset()
	return euo
}

// RemoveUsageResetIDs removes the "usage_reset" edge to UsageReset entities by IDs.
func (euo *EntitlementUpdateOne) RemoveUsageResetIDs(ids ...string) *EntitlementUpdateOne {
	euo.mutation.RemoveUsageResetIDs(ids...)
	return euo
}

// RemoveUsageReset removes "usage_reset" edges to UsageReset entities.
func (euo *EntitlementUpdateOne) RemoveUsageReset(u ...*UsageReset) *EntitlementUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.RemoveUsageResetIDs(ids...)
}

// ClearGrant clears all "grant" edges to the Grant entity.
func (euo *EntitlementUpdateOne) ClearGrant() *EntitlementUpdateOne {
	euo.mutation.ClearGrant()
	return euo
}

// RemoveGrantIDs removes the "grant" edge to Grant entities by IDs.
func (euo *EntitlementUpdateOne) RemoveGrantIDs(ids ...string) *EntitlementUpdateOne {
	euo.mutation.RemoveGrantIDs(ids...)
	return euo
}

// RemoveGrant removes "grant" edges to Grant entities.
func (euo *EntitlementUpdateOne) RemoveGrant(g ...*Grant) *EntitlementUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return euo.RemoveGrantIDs(ids...)
}

// ClearBalanceSnapshot clears all "balance_snapshot" edges to the BalanceSnapshot entity.
func (euo *EntitlementUpdateOne) ClearBalanceSnapshot() *EntitlementUpdateOne {
	euo.mutation.ClearBalanceSnapshot()
	return euo
}

// RemoveBalanceSnapshotIDs removes the "balance_snapshot" edge to BalanceSnapshot entities by IDs.
func (euo *EntitlementUpdateOne) RemoveBalanceSnapshotIDs(ids ...int) *EntitlementUpdateOne {
	euo.mutation.RemoveBalanceSnapshotIDs(ids...)
	return euo
}

// RemoveBalanceSnapshot removes "balance_snapshot" edges to BalanceSnapshot entities.
func (euo *EntitlementUpdateOne) RemoveBalanceSnapshot(b ...*BalanceSnapshot) *EntitlementUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return euo.RemoveBalanceSnapshotIDs(ids...)
}

// Where appends a list predicates to the EntitlementUpdate builder.
func (euo *EntitlementUpdateOne) Where(ps ...predicate.Entitlement) *EntitlementUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntitlementUpdateOne) Select(field string, fields ...string) *EntitlementUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entitlement entity.
func (euo *EntitlementUpdateOne) Save(ctx context.Context) (*Entitlement, error) {
	euo.defaults()
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntitlementUpdateOne) SaveX(ctx context.Context) *Entitlement {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntitlementUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntitlementUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EntitlementUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := entitlement.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EntitlementUpdateOne) check() error {
	if euo.mutation.FeatureCleared() && len(euo.mutation.FeatureIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "Entitlement.feature"`)
	}
	return nil
}

func (euo *EntitlementUpdateOne) sqlSave(ctx context.Context) (_node *Entitlement, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(entitlement.Table, entitlement.Columns, sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Entitlement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitlement.FieldID)
		for _, f := range fields {
			if !entitlement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != entitlement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Metadata(); ok {
		_spec.SetField(entitlement.FieldMetadata, field.TypeJSON, value)
	}
	if euo.mutation.MetadataCleared() {
		_spec.ClearField(entitlement.FieldMetadata, field.TypeJSON)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlement.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.SetField(entitlement.FieldDeletedAt, field.TypeTime, value)
	}
	if euo.mutation.DeletedAtCleared() {
		_spec.ClearField(entitlement.FieldDeletedAt, field.TypeTime)
	}
	if euo.mutation.MeasureUsageFromCleared() {
		_spec.ClearField(entitlement.FieldMeasureUsageFrom, field.TypeTime)
	}
	if euo.mutation.IssueAfterResetCleared() {
		_spec.ClearField(entitlement.FieldIssueAfterReset, field.TypeFloat64)
	}
	if euo.mutation.IssueAfterResetPriorityCleared() {
		_spec.ClearField(entitlement.FieldIssueAfterResetPriority, field.TypeUint8)
	}
	if euo.mutation.IsSoftLimitCleared() {
		_spec.ClearField(entitlement.FieldIsSoftLimit, field.TypeBool)
	}
	if euo.mutation.PreserveOverageAtResetCleared() {
		_spec.ClearField(entitlement.FieldPreserveOverageAtReset, field.TypeBool)
	}
	if value, ok := euo.mutation.Config(); ok {
		_spec.SetField(entitlement.FieldConfig, field.TypeJSON, value)
	}
	if value, ok := euo.mutation.AppendedConfig(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, entitlement.FieldConfig, value)
		})
	}
	if euo.mutation.ConfigCleared() {
		_spec.ClearField(entitlement.FieldConfig, field.TypeJSON)
	}
	if euo.mutation.UsagePeriodIntervalCleared() {
		_spec.ClearField(entitlement.FieldUsagePeriodInterval, field.TypeEnum)
	}
	if value, ok := euo.mutation.UsagePeriodAnchor(); ok {
		_spec.SetField(entitlement.FieldUsagePeriodAnchor, field.TypeTime, value)
	}
	if euo.mutation.UsagePeriodAnchorCleared() {
		_spec.ClearField(entitlement.FieldUsagePeriodAnchor, field.TypeTime)
	}
	if value, ok := euo.mutation.CurrentUsagePeriodStart(); ok {
		_spec.SetField(entitlement.FieldCurrentUsagePeriodStart, field.TypeTime, value)
	}
	if euo.mutation.CurrentUsagePeriodStartCleared() {
		_spec.ClearField(entitlement.FieldCurrentUsagePeriodStart, field.TypeTime)
	}
	if value, ok := euo.mutation.CurrentUsagePeriodEnd(); ok {
		_spec.SetField(entitlement.FieldCurrentUsagePeriodEnd, field.TypeTime, value)
	}
	if euo.mutation.CurrentUsagePeriodEndCleared() {
		_spec.ClearField(entitlement.FieldCurrentUsagePeriodEnd, field.TypeTime)
	}
	if euo.mutation.UsageResetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.UsageResetTable,
			Columns: []string{entitlement.UsageResetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedUsageResetIDs(); len(nodes) > 0 && !euo.mutation.UsageResetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.UsageResetTable,
			Columns: []string{entitlement.UsageResetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UsageResetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.UsageResetTable,
			Columns: []string{entitlement.UsageResetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.GrantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.GrantTable,
			Columns: []string{entitlement.GrantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dbgrant.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedGrantIDs(); len(nodes) > 0 && !euo.mutation.GrantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.GrantTable,
			Columns: []string{entitlement.GrantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dbgrant.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.GrantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.GrantTable,
			Columns: []string{entitlement.GrantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dbgrant.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.BalanceSnapshotCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.BalanceSnapshotTable,
			Columns: []string{entitlement.BalanceSnapshotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedBalanceSnapshotIDs(); len(nodes) > 0 && !euo.mutation.BalanceSnapshotCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.BalanceSnapshotTable,
			Columns: []string{entitlement.BalanceSnapshotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.BalanceSnapshotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlement.BalanceSnapshotTable,
			Columns: []string{entitlement.BalanceSnapshotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Entitlement{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitlement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
