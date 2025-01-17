// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationevent"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationeventdeliverystatus"
	"github.com/openmeterio/openmeter/internal/ent/db/predicate"
)

// NotificationEventUpdate is the builder for updating NotificationEvent entities.
type NotificationEventUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationEventMutation
}

// Where appends a list predicates to the NotificationEventUpdate builder.
func (neu *NotificationEventUpdate) Where(ps ...predicate.NotificationEvent) *NotificationEventUpdate {
	neu.mutation.Where(ps...)
	return neu
}

// SetPayload sets the "payload" field.
func (neu *NotificationEventUpdate) SetPayload(s string) *NotificationEventUpdate {
	neu.mutation.SetPayload(s)
	return neu
}

// SetNillablePayload sets the "payload" field if the given value is not nil.
func (neu *NotificationEventUpdate) SetNillablePayload(s *string) *NotificationEventUpdate {
	if s != nil {
		neu.SetPayload(*s)
	}
	return neu
}

// AddDeliveryStatusIDs adds the "delivery_statuses" edge to the NotificationEventDeliveryStatus entity by IDs.
func (neu *NotificationEventUpdate) AddDeliveryStatusIDs(ids ...string) *NotificationEventUpdate {
	neu.mutation.AddDeliveryStatusIDs(ids...)
	return neu
}

// AddDeliveryStatuses adds the "delivery_statuses" edges to the NotificationEventDeliveryStatus entity.
func (neu *NotificationEventUpdate) AddDeliveryStatuses(n ...*NotificationEventDeliveryStatus) *NotificationEventUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return neu.AddDeliveryStatusIDs(ids...)
}

// Mutation returns the NotificationEventMutation object of the builder.
func (neu *NotificationEventUpdate) Mutation() *NotificationEventMutation {
	return neu.mutation
}

// ClearDeliveryStatuses clears all "delivery_statuses" edges to the NotificationEventDeliveryStatus entity.
func (neu *NotificationEventUpdate) ClearDeliveryStatuses() *NotificationEventUpdate {
	neu.mutation.ClearDeliveryStatuses()
	return neu
}

// RemoveDeliveryStatusIDs removes the "delivery_statuses" edge to NotificationEventDeliveryStatus entities by IDs.
func (neu *NotificationEventUpdate) RemoveDeliveryStatusIDs(ids ...string) *NotificationEventUpdate {
	neu.mutation.RemoveDeliveryStatusIDs(ids...)
	return neu
}

// RemoveDeliveryStatuses removes "delivery_statuses" edges to NotificationEventDeliveryStatus entities.
func (neu *NotificationEventUpdate) RemoveDeliveryStatuses(n ...*NotificationEventDeliveryStatus) *NotificationEventUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return neu.RemoveDeliveryStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (neu *NotificationEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, neu.sqlSave, neu.mutation, neu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (neu *NotificationEventUpdate) SaveX(ctx context.Context) int {
	affected, err := neu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (neu *NotificationEventUpdate) Exec(ctx context.Context) error {
	_, err := neu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (neu *NotificationEventUpdate) ExecX(ctx context.Context) {
	if err := neu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (neu *NotificationEventUpdate) check() error {
	if neu.mutation.RulesCleared() && len(neu.mutation.RulesIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "NotificationEvent.rules"`)
	}
	return nil
}

func (neu *NotificationEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := neu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationevent.Table, notificationevent.Columns, sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString))
	if ps := neu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := neu.mutation.Payload(); ok {
		_spec.SetField(notificationevent.FieldPayload, field.TypeString, value)
	}
	if neu.mutation.DeliveryStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationevent.DeliveryStatusesTable,
			Columns: notificationevent.DeliveryStatusesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationeventdeliverystatus.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := neu.mutation.RemovedDeliveryStatusesIDs(); len(nodes) > 0 && !neu.mutation.DeliveryStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationevent.DeliveryStatusesTable,
			Columns: notificationevent.DeliveryStatusesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationeventdeliverystatus.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := neu.mutation.DeliveryStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationevent.DeliveryStatusesTable,
			Columns: notificationevent.DeliveryStatusesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationeventdeliverystatus.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, neu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	neu.mutation.done = true
	return n, nil
}

// NotificationEventUpdateOne is the builder for updating a single NotificationEvent entity.
type NotificationEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationEventMutation
}

// SetPayload sets the "payload" field.
func (neuo *NotificationEventUpdateOne) SetPayload(s string) *NotificationEventUpdateOne {
	neuo.mutation.SetPayload(s)
	return neuo
}

// SetNillablePayload sets the "payload" field if the given value is not nil.
func (neuo *NotificationEventUpdateOne) SetNillablePayload(s *string) *NotificationEventUpdateOne {
	if s != nil {
		neuo.SetPayload(*s)
	}
	return neuo
}

// AddDeliveryStatusIDs adds the "delivery_statuses" edge to the NotificationEventDeliveryStatus entity by IDs.
func (neuo *NotificationEventUpdateOne) AddDeliveryStatusIDs(ids ...string) *NotificationEventUpdateOne {
	neuo.mutation.AddDeliveryStatusIDs(ids...)
	return neuo
}

// AddDeliveryStatuses adds the "delivery_statuses" edges to the NotificationEventDeliveryStatus entity.
func (neuo *NotificationEventUpdateOne) AddDeliveryStatuses(n ...*NotificationEventDeliveryStatus) *NotificationEventUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return neuo.AddDeliveryStatusIDs(ids...)
}

// Mutation returns the NotificationEventMutation object of the builder.
func (neuo *NotificationEventUpdateOne) Mutation() *NotificationEventMutation {
	return neuo.mutation
}

// ClearDeliveryStatuses clears all "delivery_statuses" edges to the NotificationEventDeliveryStatus entity.
func (neuo *NotificationEventUpdateOne) ClearDeliveryStatuses() *NotificationEventUpdateOne {
	neuo.mutation.ClearDeliveryStatuses()
	return neuo
}

// RemoveDeliveryStatusIDs removes the "delivery_statuses" edge to NotificationEventDeliveryStatus entities by IDs.
func (neuo *NotificationEventUpdateOne) RemoveDeliveryStatusIDs(ids ...string) *NotificationEventUpdateOne {
	neuo.mutation.RemoveDeliveryStatusIDs(ids...)
	return neuo
}

// RemoveDeliveryStatuses removes "delivery_statuses" edges to NotificationEventDeliveryStatus entities.
func (neuo *NotificationEventUpdateOne) RemoveDeliveryStatuses(n ...*NotificationEventDeliveryStatus) *NotificationEventUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return neuo.RemoveDeliveryStatusIDs(ids...)
}

// Where appends a list predicates to the NotificationEventUpdate builder.
func (neuo *NotificationEventUpdateOne) Where(ps ...predicate.NotificationEvent) *NotificationEventUpdateOne {
	neuo.mutation.Where(ps...)
	return neuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (neuo *NotificationEventUpdateOne) Select(field string, fields ...string) *NotificationEventUpdateOne {
	neuo.fields = append([]string{field}, fields...)
	return neuo
}

// Save executes the query and returns the updated NotificationEvent entity.
func (neuo *NotificationEventUpdateOne) Save(ctx context.Context) (*NotificationEvent, error) {
	return withHooks(ctx, neuo.sqlSave, neuo.mutation, neuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (neuo *NotificationEventUpdateOne) SaveX(ctx context.Context) *NotificationEvent {
	node, err := neuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (neuo *NotificationEventUpdateOne) Exec(ctx context.Context) error {
	_, err := neuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (neuo *NotificationEventUpdateOne) ExecX(ctx context.Context) {
	if err := neuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (neuo *NotificationEventUpdateOne) check() error {
	if neuo.mutation.RulesCleared() && len(neuo.mutation.RulesIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "NotificationEvent.rules"`)
	}
	return nil
}

func (neuo *NotificationEventUpdateOne) sqlSave(ctx context.Context) (_node *NotificationEvent, err error) {
	if err := neuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationevent.Table, notificationevent.Columns, sqlgraph.NewFieldSpec(notificationevent.FieldID, field.TypeString))
	id, ok := neuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "NotificationEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := neuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationevent.FieldID)
		for _, f := range fields {
			if !notificationevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != notificationevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := neuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := neuo.mutation.Payload(); ok {
		_spec.SetField(notificationevent.FieldPayload, field.TypeString, value)
	}
	if neuo.mutation.DeliveryStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationevent.DeliveryStatusesTable,
			Columns: notificationevent.DeliveryStatusesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationeventdeliverystatus.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := neuo.mutation.RemovedDeliveryStatusesIDs(); len(nodes) > 0 && !neuo.mutation.DeliveryStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationevent.DeliveryStatusesTable,
			Columns: notificationevent.DeliveryStatusesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationeventdeliverystatus.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := neuo.mutation.DeliveryStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notificationevent.DeliveryStatusesTable,
			Columns: notificationevent.DeliveryStatusesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationeventdeliverystatus.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotificationEvent{config: neuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, neuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	neuo.mutation.done = true
	return _node, nil
}
