// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/ledger"
)

// Ledger is the model entity for the Ledger schema.
type Ledger struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Highwatermark holds the value of the "highwatermark" field.
	Highwatermark time.Time `json:"highwatermark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LedgerQuery when eager-loading is set.
	Edges        LedgerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LedgerEdges holds the relations/edges for other nodes in the graph.
type LedgerEdges struct {
	// CreditGrants holds the value of the credit_grants edge.
	CreditGrants []*CreditEntry `json:"credit_grants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CreditGrantsOrErr returns the CreditGrants value or an error if the edge
// was not loaded in eager-loading.
func (e LedgerEdges) CreditGrantsOrErr() ([]*CreditEntry, error) {
	if e.loadedTypes[0] {
		return e.CreditGrants, nil
	}
	return nil, &NotLoadedError{edge: "credit_grants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ledger) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ledger.FieldMetadata:
			values[i] = new([]byte)
		case ledger.FieldID, ledger.FieldNamespace, ledger.FieldSubject:
			values[i] = new(sql.NullString)
		case ledger.FieldCreatedAt, ledger.FieldUpdatedAt, ledger.FieldHighwatermark:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ledger fields.
func (l *Ledger) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ledger.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				l.ID = value.String
			}
		case ledger.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case ledger.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case ledger.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				l.Namespace = value.String
			}
		case ledger.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				l.Subject = value.String
			}
		case ledger.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &l.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case ledger.FieldHighwatermark:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field highwatermark", values[i])
			} else if value.Valid {
				l.Highwatermark = value.Time
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Ledger.
// This includes values selected through modifiers, order, etc.
func (l *Ledger) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryCreditGrants queries the "credit_grants" edge of the Ledger entity.
func (l *Ledger) QueryCreditGrants() *CreditEntryQuery {
	return NewLedgerClient(l.config).QueryCreditGrants(l)
}

// Update returns a builder for updating this Ledger.
// Note that you need to call Ledger.Unwrap() before calling this method if this Ledger
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Ledger) Update() *LedgerUpdateOne {
	return NewLedgerClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Ledger entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Ledger) Unwrap() *Ledger {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("db: Ledger is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Ledger) String() string {
	var builder strings.Builder
	builder.WriteString("Ledger(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("namespace=")
	builder.WriteString(l.Namespace)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(l.Subject)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", l.Metadata))
	builder.WriteString(", ")
	builder.WriteString("highwatermark=")
	builder.WriteString(l.Highwatermark.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Ledgers is a parsable slice of Ledger.
type Ledgers []*Ledger
