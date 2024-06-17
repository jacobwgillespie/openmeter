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
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/creditentry"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/ledger"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/predicate"
)

// LedgerUpdate is the builder for updating Ledger entities.
type LedgerUpdate struct {
	config
	hooks    []Hook
	mutation *LedgerMutation
}

// Where appends a list predicates to the LedgerUpdate builder.
func (lu *LedgerUpdate) Where(ps ...predicate.Ledger) *LedgerUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LedgerUpdate) SetUpdatedAt(t time.Time) *LedgerUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetMetadata sets the "metadata" field.
func (lu *LedgerUpdate) SetMetadata(m map[string]string) *LedgerUpdate {
	lu.mutation.SetMetadata(m)
	return lu
}

// ClearMetadata clears the value of the "metadata" field.
func (lu *LedgerUpdate) ClearMetadata() *LedgerUpdate {
	lu.mutation.ClearMetadata()
	return lu
}

// SetHighwatermark sets the "highwatermark" field.
func (lu *LedgerUpdate) SetHighwatermark(t time.Time) *LedgerUpdate {
	lu.mutation.SetHighwatermark(t)
	return lu
}

// SetNillableHighwatermark sets the "highwatermark" field if the given value is not nil.
func (lu *LedgerUpdate) SetNillableHighwatermark(t *time.Time) *LedgerUpdate {
	if t != nil {
		lu.SetHighwatermark(*t)
	}
	return lu
}

// AddCreditGrantIDs adds the "credit_grants" edge to the CreditEntry entity by IDs.
func (lu *LedgerUpdate) AddCreditGrantIDs(ids ...string) *LedgerUpdate {
	lu.mutation.AddCreditGrantIDs(ids...)
	return lu
}

// AddCreditGrants adds the "credit_grants" edges to the CreditEntry entity.
func (lu *LedgerUpdate) AddCreditGrants(c ...*CreditEntry) *LedgerUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.AddCreditGrantIDs(ids...)
}

// Mutation returns the LedgerMutation object of the builder.
func (lu *LedgerUpdate) Mutation() *LedgerMutation {
	return lu.mutation
}

// ClearCreditGrants clears all "credit_grants" edges to the CreditEntry entity.
func (lu *LedgerUpdate) ClearCreditGrants() *LedgerUpdate {
	lu.mutation.ClearCreditGrants()
	return lu
}

// RemoveCreditGrantIDs removes the "credit_grants" edge to CreditEntry entities by IDs.
func (lu *LedgerUpdate) RemoveCreditGrantIDs(ids ...string) *LedgerUpdate {
	lu.mutation.RemoveCreditGrantIDs(ids...)
	return lu
}

// RemoveCreditGrants removes "credit_grants" edges to CreditEntry entities.
func (lu *LedgerUpdate) RemoveCreditGrants(c ...*CreditEntry) *LedgerUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.RemoveCreditGrantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LedgerUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LedgerUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LedgerUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LedgerUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LedgerUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := ledger.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

func (lu *LedgerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ledger.Table, ledger.Columns, sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeString))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(ledger.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.Metadata(); ok {
		_spec.SetField(ledger.FieldMetadata, field.TypeJSON, value)
	}
	if lu.mutation.MetadataCleared() {
		_spec.ClearField(ledger.FieldMetadata, field.TypeJSON)
	}
	if value, ok := lu.mutation.Highwatermark(); ok {
		_spec.SetField(ledger.FieldHighwatermark, field.TypeTime, value)
	}
	if lu.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ledger.CreditGrantsTable,
			Columns: []string{ledger.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedCreditGrantsIDs(); len(nodes) > 0 && !lu.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ledger.CreditGrantsTable,
			Columns: []string{ledger.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.CreditGrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ledger.CreditGrantsTable,
			Columns: []string{ledger.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ledger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LedgerUpdateOne is the builder for updating a single Ledger entity.
type LedgerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LedgerMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LedgerUpdateOne) SetUpdatedAt(t time.Time) *LedgerUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetMetadata sets the "metadata" field.
func (luo *LedgerUpdateOne) SetMetadata(m map[string]string) *LedgerUpdateOne {
	luo.mutation.SetMetadata(m)
	return luo
}

// ClearMetadata clears the value of the "metadata" field.
func (luo *LedgerUpdateOne) ClearMetadata() *LedgerUpdateOne {
	luo.mutation.ClearMetadata()
	return luo
}

// SetHighwatermark sets the "highwatermark" field.
func (luo *LedgerUpdateOne) SetHighwatermark(t time.Time) *LedgerUpdateOne {
	luo.mutation.SetHighwatermark(t)
	return luo
}

// SetNillableHighwatermark sets the "highwatermark" field if the given value is not nil.
func (luo *LedgerUpdateOne) SetNillableHighwatermark(t *time.Time) *LedgerUpdateOne {
	if t != nil {
		luo.SetHighwatermark(*t)
	}
	return luo
}

// AddCreditGrantIDs adds the "credit_grants" edge to the CreditEntry entity by IDs.
func (luo *LedgerUpdateOne) AddCreditGrantIDs(ids ...string) *LedgerUpdateOne {
	luo.mutation.AddCreditGrantIDs(ids...)
	return luo
}

// AddCreditGrants adds the "credit_grants" edges to the CreditEntry entity.
func (luo *LedgerUpdateOne) AddCreditGrants(c ...*CreditEntry) *LedgerUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.AddCreditGrantIDs(ids...)
}

// Mutation returns the LedgerMutation object of the builder.
func (luo *LedgerUpdateOne) Mutation() *LedgerMutation {
	return luo.mutation
}

// ClearCreditGrants clears all "credit_grants" edges to the CreditEntry entity.
func (luo *LedgerUpdateOne) ClearCreditGrants() *LedgerUpdateOne {
	luo.mutation.ClearCreditGrants()
	return luo
}

// RemoveCreditGrantIDs removes the "credit_grants" edge to CreditEntry entities by IDs.
func (luo *LedgerUpdateOne) RemoveCreditGrantIDs(ids ...string) *LedgerUpdateOne {
	luo.mutation.RemoveCreditGrantIDs(ids...)
	return luo
}

// RemoveCreditGrants removes "credit_grants" edges to CreditEntry entities.
func (luo *LedgerUpdateOne) RemoveCreditGrants(c ...*CreditEntry) *LedgerUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.RemoveCreditGrantIDs(ids...)
}

// Where appends a list predicates to the LedgerUpdate builder.
func (luo *LedgerUpdateOne) Where(ps ...predicate.Ledger) *LedgerUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LedgerUpdateOne) Select(field string, fields ...string) *LedgerUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Ledger entity.
func (luo *LedgerUpdateOne) Save(ctx context.Context) (*Ledger, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LedgerUpdateOne) SaveX(ctx context.Context) *Ledger {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LedgerUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LedgerUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LedgerUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := ledger.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

func (luo *LedgerUpdateOne) sqlSave(ctx context.Context) (_node *Ledger, err error) {
	_spec := sqlgraph.NewUpdateSpec(ledger.Table, ledger.Columns, sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeString))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Ledger.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ledger.FieldID)
		for _, f := range fields {
			if !ledger.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != ledger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(ledger.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.Metadata(); ok {
		_spec.SetField(ledger.FieldMetadata, field.TypeJSON, value)
	}
	if luo.mutation.MetadataCleared() {
		_spec.ClearField(ledger.FieldMetadata, field.TypeJSON)
	}
	if value, ok := luo.mutation.Highwatermark(); ok {
		_spec.SetField(ledger.FieldHighwatermark, field.TypeTime, value)
	}
	if luo.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ledger.CreditGrantsTable,
			Columns: []string{ledger.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedCreditGrantsIDs(); len(nodes) > 0 && !luo.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ledger.CreditGrantsTable,
			Columns: []string{ledger.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.CreditGrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ledger.CreditGrantsTable,
			Columns: []string{ledger.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ledger{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ledger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
