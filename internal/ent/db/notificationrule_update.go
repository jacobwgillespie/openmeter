// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationchannel"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationevent"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/internal/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/notification"
)

// NotificationRuleUpdate is the builder for updating NotificationRule entities.
type NotificationRuleUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationRuleMutation
}

// Where appends a list predicates to the NotificationRuleUpdate builder.
func (nru *NotificationRuleUpdate) Where(ps ...predicate.NotificationRule) *NotificationRuleUpdate {
	nru.mutation.Where(ps...)
	return nru
}

// SetUpdatedAt sets the "updated_at" field.
func (nru *NotificationRuleUpdate) SetUpdatedAt(t time.Time) *NotificationRuleUpdate {
	nru.mutation.SetUpdatedAt(t)
	return nru
}

// SetDeletedAt sets the "deleted_at" field.
func (nru *NotificationRuleUpdate) SetDeletedAt(t time.Time) *NotificationRuleUpdate {
	nru.mutation.SetDeletedAt(t)
	return nru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nru *NotificationRuleUpdate) SetNillableDeletedAt(t *time.Time) *NotificationRuleUpdate {
	if t != nil {
		nru.SetDeletedAt(*t)
	}
	return nru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nru *NotificationRuleUpdate) ClearDeletedAt() *NotificationRuleUpdate {
	nru.mutation.ClearDeletedAt()
	return nru
}

// SetName sets the "name" field.
func (nru *NotificationRuleUpdate) SetName(s string) *NotificationRuleUpdate {
	nru.mutation.SetName(s)
	return nru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nru *NotificationRuleUpdate) SetNillableName(s *string) *NotificationRuleUpdate {
	if s != nil {
		nru.SetName(*s)
	}
	return nru
}

// SetDisabled sets the "disabled" field.
func (nru *NotificationRuleUpdate) SetDisabled(b bool) *NotificationRuleUpdate {
	nru.mutation.SetDisabled(b)
	return nru
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (nru *NotificationRuleUpdate) SetNillableDisabled(b *bool) *NotificationRuleUpdate {
	if b != nil {
		nru.SetDisabled(*b)
	}
	return nru
}

// ClearDisabled clears the value of the "disabled" field.
func (nru *NotificationRuleUpdate) ClearDisabled() *NotificationRuleUpdate {
	nru.mutation.ClearDisabled()
	return nru
}

// SetConfig sets the "config" field.
func (nru *NotificationRuleUpdate) SetConfig(nc notification.RuleConfig) *NotificationRuleUpdate {
	nru.mutation.SetConfig(nc)
	return nru
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (nru *NotificationRuleUpdate) SetNillableConfig(nc *notification.RuleConfig) *NotificationRuleUpdate {
	if nc != nil {
		nru.SetConfig(*nc)
	}
	return nru
}

// AddChannelIDs adds the "channels" edge to the NotificationChannel entity by IDs.
func (nru *NotificationRuleUpdate) AddChannelIDs(ids ...string) *NotificationRuleUpdate {
	nru.mutation.AddChannelIDs(ids...)
	return nru
}

// AddChannels adds the "channels" edges to the NotificationChannel entity.
func (nru *NotificationRuleUpdate) AddChannels(n ...*NotificationChannel) *NotificationRuleUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nru.AddChannelIDs(ids...)
}

// AddEventIDs adds the "events" edge to the NotificationEvent entity by IDs.
func (nru *NotificationRuleUpdate) AddEventIDs(ids ...string) *NotificationRuleUpdate {
	nru.mutation.AddEventIDs(ids...)
	return nru
}

// AddEvents adds the "events" edges to the NotificationEvent entity.
func (nru *NotificationRuleUpdate) AddEvents(n ...*NotificationEvent) *NotificationRuleUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nru.AddEventIDs(ids...)
}

// Mutation returns the NotificationRuleMutation object of the builder.
func (nru *NotificationRuleUpdate) Mutation() *NotificationRuleMutation {
	return nru.mutation
}

// ClearChannels clears all "channels" edges to the NotificationChannel entity.
func (nru *NotificationRuleUpdate) ClearChannels() *NotificationRuleUpdate {
	nru.mutation.ClearChannels()
	return nru
}

// RemoveChannelIDs removes the "channels" edge to NotificationChannel entities by IDs.
func (nru *NotificationRuleUpdate) RemoveChannelIDs(ids ...string) *NotificationRuleUpdate {
	nru.mutation.RemoveChannelIDs(ids...)
	return nru
}

// RemoveChannels removes "channels" edges to NotificationChannel entities.
func (nru *NotificationRuleUpdate) RemoveChannels(n ...*NotificationChannel) *NotificationRuleUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nru.RemoveChannelIDs(ids...)
}

// ClearEvents clears all "events" edges to the NotificationEvent entity.
func (nru *NotificationRuleUpdate) ClearEvents() *NotificationRuleUpdate {
	nru.mutation.ClearEvents()
	return nru
}

// RemoveEventIDs removes the "events" edge to NotificationEvent entities by IDs.
func (nru *NotificationRuleUpdate) RemoveEventIDs(ids ...string) *NotificationRuleUpdate {
	nru.mutation.RemoveEventIDs(ids...)
	return nru
}

// RemoveEvents removes "events" edges to NotificationEvent entities.
func (nru *NotificationRuleUpdate) RemoveEvents(n ...*NotificationEvent) *NotificationRuleUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nru.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nru *NotificationRuleUpdate) Save(ctx context.Context) (int, error) {
	nru.defaults()
	return withHooks(ctx, nru.sqlSave, nru.mutation, nru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nru *NotificationRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := nru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nru *NotificationRuleUpdate) Exec(ctx context.Context) error {
	_, err := nru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nru *NotificationRuleUpdate) ExecX(ctx context.Context) {
	if err := nru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nru *NotificationRuleUpdate) defaults() {
	if _, ok := nru.mutation.UpdatedAt(); !ok {
		v := notificationrule.UpdateDefaultUpdatedAt()
		nru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nru *NotificationRuleUpdate) check() error {
	if v, ok := nru.mutation.Name(); ok {
		if err := notificationrule.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "NotificationRule.name": %w`, err)}
		}
	}
	return nil
}

func (nru *NotificationRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationrule.Table, notificationrule.Columns, sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString))
	if ps := nru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nru.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationrule.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nru.mutation.DeletedAt(); ok {
		_spec.SetField(notificationrule.FieldDeletedAt, field.TypeTime, value)
	}
	if nru.mutation.DeletedAtCleared() {
		_spec.ClearField(notificationrule.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nru.mutation.Name(); ok {
		_spec.SetField(notificationrule.FieldName, field.TypeString, value)
	}
	if value, ok := nru.mutation.Disabled(); ok {
		_spec.SetField(notificationrule.FieldDisabled, field.TypeBool, value)
	}
	if nru.mutation.DisabledCleared() {
		_spec.ClearField(notificationrule.FieldDisabled, field.TypeBool)
	}
	if value, ok := nru.mutation.Config(); ok {
		vv, err := notificationrule.ValueScanner.Config.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(notificationrule.FieldConfig, field.TypeString, vv)
	}
	if nru.mutation.ChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationrule.ChannelsTable,
			Columns: notificationrule.ChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nru.mutation.RemovedChannelsIDs(); len(nodes) > 0 && !nru.mutation.ChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationrule.ChannelsTable,
			Columns: notificationrule.ChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nru.mutation.ChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationrule.ChannelsTable,
			Columns: notificationrule.ChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nru.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationrule.EventsTable,
			Columns: []string{notificationrule.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nru.mutation.RemovedEventsIDs(); len(nodes) > 0 && !nru.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationrule.EventsTable,
			Columns: []string{notificationrule.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nru.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationrule.EventsTable,
			Columns: []string{notificationrule.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nru.mutation.done = true
	return n, nil
}

// NotificationRuleUpdateOne is the builder for updating a single NotificationRule entity.
type NotificationRuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationRuleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nruo *NotificationRuleUpdateOne) SetUpdatedAt(t time.Time) *NotificationRuleUpdateOne {
	nruo.mutation.SetUpdatedAt(t)
	return nruo
}

// SetDeletedAt sets the "deleted_at" field.
func (nruo *NotificationRuleUpdateOne) SetDeletedAt(t time.Time) *NotificationRuleUpdateOne {
	nruo.mutation.SetDeletedAt(t)
	return nruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nruo *NotificationRuleUpdateOne) SetNillableDeletedAt(t *time.Time) *NotificationRuleUpdateOne {
	if t != nil {
		nruo.SetDeletedAt(*t)
	}
	return nruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nruo *NotificationRuleUpdateOne) ClearDeletedAt() *NotificationRuleUpdateOne {
	nruo.mutation.ClearDeletedAt()
	return nruo
}

// SetName sets the "name" field.
func (nruo *NotificationRuleUpdateOne) SetName(s string) *NotificationRuleUpdateOne {
	nruo.mutation.SetName(s)
	return nruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nruo *NotificationRuleUpdateOne) SetNillableName(s *string) *NotificationRuleUpdateOne {
	if s != nil {
		nruo.SetName(*s)
	}
	return nruo
}

// SetDisabled sets the "disabled" field.
func (nruo *NotificationRuleUpdateOne) SetDisabled(b bool) *NotificationRuleUpdateOne {
	nruo.mutation.SetDisabled(b)
	return nruo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (nruo *NotificationRuleUpdateOne) SetNillableDisabled(b *bool) *NotificationRuleUpdateOne {
	if b != nil {
		nruo.SetDisabled(*b)
	}
	return nruo
}

// ClearDisabled clears the value of the "disabled" field.
func (nruo *NotificationRuleUpdateOne) ClearDisabled() *NotificationRuleUpdateOne {
	nruo.mutation.ClearDisabled()
	return nruo
}

// SetConfig sets the "config" field.
func (nruo *NotificationRuleUpdateOne) SetConfig(nc notification.RuleConfig) *NotificationRuleUpdateOne {
	nruo.mutation.SetConfig(nc)
	return nruo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (nruo *NotificationRuleUpdateOne) SetNillableConfig(nc *notification.RuleConfig) *NotificationRuleUpdateOne {
	if nc != nil {
		nruo.SetConfig(*nc)
	}
	return nruo
}

// AddChannelIDs adds the "channels" edge to the NotificationChannel entity by IDs.
func (nruo *NotificationRuleUpdateOne) AddChannelIDs(ids ...string) *NotificationRuleUpdateOne {
	nruo.mutation.AddChannelIDs(ids...)
	return nruo
}

// AddChannels adds the "channels" edges to the NotificationChannel entity.
func (nruo *NotificationRuleUpdateOne) AddChannels(n ...*NotificationChannel) *NotificationRuleUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nruo.AddChannelIDs(ids...)
}

// AddEventIDs adds the "events" edge to the NotificationEvent entity by IDs.
func (nruo *NotificationRuleUpdateOne) AddEventIDs(ids ...string) *NotificationRuleUpdateOne {
	nruo.mutation.AddEventIDs(ids...)
	return nruo
}

// AddEvents adds the "events" edges to the NotificationEvent entity.
func (nruo *NotificationRuleUpdateOne) AddEvents(n ...*NotificationEvent) *NotificationRuleUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nruo.AddEventIDs(ids...)
}

// Mutation returns the NotificationRuleMutation object of the builder.
func (nruo *NotificationRuleUpdateOne) Mutation() *NotificationRuleMutation {
	return nruo.mutation
}

// ClearChannels clears all "channels" edges to the NotificationChannel entity.
func (nruo *NotificationRuleUpdateOne) ClearChannels() *NotificationRuleUpdateOne {
	nruo.mutation.ClearChannels()
	return nruo
}

// RemoveChannelIDs removes the "channels" edge to NotificationChannel entities by IDs.
func (nruo *NotificationRuleUpdateOne) RemoveChannelIDs(ids ...string) *NotificationRuleUpdateOne {
	nruo.mutation.RemoveChannelIDs(ids...)
	return nruo
}

// RemoveChannels removes "channels" edges to NotificationChannel entities.
func (nruo *NotificationRuleUpdateOne) RemoveChannels(n ...*NotificationChannel) *NotificationRuleUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nruo.RemoveChannelIDs(ids...)
}

// ClearEvents clears all "events" edges to the NotificationEvent entity.
func (nruo *NotificationRuleUpdateOne) ClearEvents() *NotificationRuleUpdateOne {
	nruo.mutation.ClearEvents()
	return nruo
}

// RemoveEventIDs removes the "events" edge to NotificationEvent entities by IDs.
func (nruo *NotificationRuleUpdateOne) RemoveEventIDs(ids ...string) *NotificationRuleUpdateOne {
	nruo.mutation.RemoveEventIDs(ids...)
	return nruo
}

// RemoveEvents removes "events" edges to NotificationEvent entities.
func (nruo *NotificationRuleUpdateOne) RemoveEvents(n ...*NotificationEvent) *NotificationRuleUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nruo.RemoveEventIDs(ids...)
}

// Where appends a list predicates to the NotificationRuleUpdate builder.
func (nruo *NotificationRuleUpdateOne) Where(ps ...predicate.NotificationRule) *NotificationRuleUpdateOne {
	nruo.mutation.Where(ps...)
	return nruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nruo *NotificationRuleUpdateOne) Select(field string, fields ...string) *NotificationRuleUpdateOne {
	nruo.fields = append([]string{field}, fields...)
	return nruo
}

// Save executes the query and returns the updated NotificationRule entity.
func (nruo *NotificationRuleUpdateOne) Save(ctx context.Context) (*NotificationRule, error) {
	nruo.defaults()
	return withHooks(ctx, nruo.sqlSave, nruo.mutation, nruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nruo *NotificationRuleUpdateOne) SaveX(ctx context.Context) *NotificationRule {
	node, err := nruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nruo *NotificationRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := nruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nruo *NotificationRuleUpdateOne) ExecX(ctx context.Context) {
	if err := nruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nruo *NotificationRuleUpdateOne) defaults() {
	if _, ok := nruo.mutation.UpdatedAt(); !ok {
		v := notificationrule.UpdateDefaultUpdatedAt()
		nruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nruo *NotificationRuleUpdateOne) check() error {
	if v, ok := nruo.mutation.Name(); ok {
		if err := notificationrule.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "NotificationRule.name": %w`, err)}
		}
	}
	return nil
}

func (nruo *NotificationRuleUpdateOne) sqlSave(ctx context.Context) (_node *NotificationRule, err error) {
	if err := nruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationrule.Table, notificationrule.Columns, sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString))
	id, ok := nruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "NotificationRule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationrule.FieldID)
		for _, f := range fields {
			if !notificationrule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != notificationrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nruo.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationrule.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nruo.mutation.DeletedAt(); ok {
		_spec.SetField(notificationrule.FieldDeletedAt, field.TypeTime, value)
	}
	if nruo.mutation.DeletedAtCleared() {
		_spec.ClearField(notificationrule.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nruo.mutation.Name(); ok {
		_spec.SetField(notificationrule.FieldName, field.TypeString, value)
	}
	if value, ok := nruo.mutation.Disabled(); ok {
		_spec.SetField(notificationrule.FieldDisabled, field.TypeBool, value)
	}
	if nruo.mutation.DisabledCleared() {
		_spec.ClearField(notificationrule.FieldDisabled, field.TypeBool)
	}
	if value, ok := nruo.mutation.Config(); ok {
		vv, err := notificationrule.ValueScanner.Config.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(notificationrule.FieldConfig, field.TypeString, vv)
	}
	if nruo.mutation.ChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationrule.ChannelsTable,
			Columns: notificationrule.ChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nruo.mutation.RemovedChannelsIDs(); len(nodes) > 0 && !nruo.mutation.ChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationrule.ChannelsTable,
			Columns: notificationrule.ChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nruo.mutation.ChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationrule.ChannelsTable,
			Columns: notificationrule.ChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nruo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationrule.EventsTable,
			Columns: []string{notificationrule.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nruo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !nruo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationrule.EventsTable,
			Columns: []string{notificationrule.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nruo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notificationrule.EventsTable,
			Columns: []string{notificationrule.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotificationRule{config: nruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nruo.mutation.done = true
	return _node, nil
}
