// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carbonable/carbonable-launchpad-backend/ent/launchpad"
	"github.com/carbonable/carbonable-launchpad-backend/ent/project"
	"github.com/carbonable/carbonable-launchpad-backend/ent/schema"
)

// Launchpad is the model entity for the Launchpad schema.
type Launchpad struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IsReady holds the value of the "is_ready" field.
	IsReady bool `json:"is_ready,omitempty"`
	// MinterContract holds the value of the "minter_contract" field.
	MinterContract schema.MinterContract `json:"minter_contract,omitempty"`
	// WhitelistedSaleOpen holds the value of the "whitelisted_sale_open" field.
	WhitelistedSaleOpen bool `json:"whitelisted_sale_open,omitempty"`
	// PublicSaleOpen holds the value of the "public_sale_open" field.
	PublicSaleOpen bool `json:"public_sale_open,omitempty"`
	// IsSoldOut holds the value of the "is_sold_out" field.
	IsSoldOut bool `json:"is_sold_out,omitempty"`
	// IsCanceled holds the value of the "is_canceled" field.
	IsCanceled bool `json:"is_canceled,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LaunchpadQuery when eager-loading is set.
	Edges             LaunchpadEdges `json:"edges"`
	project_launchpad *int
	selectValues      sql.SelectValues
}

// LaunchpadEdges holds the relations/edges for other nodes in the graph.
type LaunchpadEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LaunchpadEdges) ProjectOrErr() (*Project, error) {
	if e.Project != nil {
		return e.Project, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: project.Label}
	}
	return nil, &NotLoadedError{edge: "project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Launchpad) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case launchpad.FieldMinterContract:
			values[i] = new([]byte)
		case launchpad.FieldIsReady, launchpad.FieldWhitelistedSaleOpen, launchpad.FieldPublicSaleOpen, launchpad.FieldIsSoldOut, launchpad.FieldIsCanceled:
			values[i] = new(sql.NullBool)
		case launchpad.FieldID:
			values[i] = new(sql.NullInt64)
		case launchpad.ForeignKeys[0]: // project_launchpad
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Launchpad fields.
func (l *Launchpad) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case launchpad.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case launchpad.FieldIsReady:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_ready", values[i])
			} else if value.Valid {
				l.IsReady = value.Bool
			}
		case launchpad.FieldMinterContract:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field minter_contract", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &l.MinterContract); err != nil {
					return fmt.Errorf("unmarshal field minter_contract: %w", err)
				}
			}
		case launchpad.FieldWhitelistedSaleOpen:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field whitelisted_sale_open", values[i])
			} else if value.Valid {
				l.WhitelistedSaleOpen = value.Bool
			}
		case launchpad.FieldPublicSaleOpen:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public_sale_open", values[i])
			} else if value.Valid {
				l.PublicSaleOpen = value.Bool
			}
		case launchpad.FieldIsSoldOut:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_sold_out", values[i])
			} else if value.Valid {
				l.IsSoldOut = value.Bool
			}
		case launchpad.FieldIsCanceled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_canceled", values[i])
			} else if value.Valid {
				l.IsCanceled = value.Bool
			}
		case launchpad.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_launchpad", value)
			} else if value.Valid {
				l.project_launchpad = new(int)
				*l.project_launchpad = int(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Launchpad.
// This includes values selected through modifiers, order, etc.
func (l *Launchpad) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the Launchpad entity.
func (l *Launchpad) QueryProject() *ProjectQuery {
	return NewLaunchpadClient(l.config).QueryProject(l)
}

// Update returns a builder for updating this Launchpad.
// Note that you need to call Launchpad.Unwrap() before calling this method if this Launchpad
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Launchpad) Update() *LaunchpadUpdateOne {
	return NewLaunchpadClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Launchpad entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Launchpad) Unwrap() *Launchpad {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Launchpad is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Launchpad) String() string {
	var builder strings.Builder
	builder.WriteString("Launchpad(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("is_ready=")
	builder.WriteString(fmt.Sprintf("%v", l.IsReady))
	builder.WriteString(", ")
	builder.WriteString("minter_contract=")
	builder.WriteString(fmt.Sprintf("%v", l.MinterContract))
	builder.WriteString(", ")
	builder.WriteString("whitelisted_sale_open=")
	builder.WriteString(fmt.Sprintf("%v", l.WhitelistedSaleOpen))
	builder.WriteString(", ")
	builder.WriteString("public_sale_open=")
	builder.WriteString(fmt.Sprintf("%v", l.PublicSaleOpen))
	builder.WriteString(", ")
	builder.WriteString("is_sold_out=")
	builder.WriteString(fmt.Sprintf("%v", l.IsSoldOut))
	builder.WriteString(", ")
	builder.WriteString("is_canceled=")
	builder.WriteString(fmt.Sprintf("%v", l.IsCanceled))
	builder.WriteByte(')')
	return builder.String()
}

// Launchpads is a parsable slice of Launchpad.
type Launchpads []*Launchpad
