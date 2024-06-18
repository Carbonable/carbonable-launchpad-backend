// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/carbonable/carbonable-launchpad-backend/ent/launchpad"
	"github.com/carbonable/carbonable-launchpad-backend/ent/mint"
	"github.com/carbonable/carbonable-launchpad-backend/ent/project"
	"github.com/carbonable/carbonable-launchpad-backend/ent/schema"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Slot holds the value of the "slot" field.
	Slot int `json:"slot,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// ValueDecimal holds the value of the "value_decimal" field.
	ValueDecimal int `json:"value_decimal,omitempty"`
	// ForecastedApr holds the value of the "forecasted_apr" field.
	ForecastedApr string `json:"forecasted_apr,omitempty"`
	// TotalValue holds the value of the "total_value" field.
	TotalValue string `json:"total_value,omitempty"`
	// PaymentToken holds the value of the "payment_token" field.
	PaymentToken schema.PaymentToken `json:"payment_token,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata schema.Metadata `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges        ProjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Mint holds the value of the mint edge.
	Mint *Mint `json:"mint,omitempty"`
	// Launchpad holds the value of the launchpad edge.
	Launchpad *Launchpad `json:"launchpad,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MintOrErr returns the Mint value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) MintOrErr() (*Mint, error) {
	if e.Mint != nil {
		return e.Mint, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: mint.Label}
	}
	return nil, &NotLoadedError{edge: "mint"}
}

// LaunchpadOrErr returns the Launchpad value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) LaunchpadOrErr() (*Launchpad, error) {
	if e.Launchpad != nil {
		return e.Launchpad, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: launchpad.Label}
	}
	return nil, &NotLoadedError{edge: "launchpad"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldPaymentToken, project.FieldMetadata:
			values[i] = new([]byte)
		case project.FieldID, project.FieldSlot, project.FieldValueDecimal:
			values[i] = new(sql.NullInt64)
		case project.FieldAddress, project.FieldName, project.FieldSlug, project.FieldForecastedApr, project.FieldTotalValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case project.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				pr.Address = value.String
			}
		case project.FieldSlot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field slot", values[i])
			} else if value.Valid {
				pr.Slot = int(value.Int64)
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				pr.Slug = value.String
			}
		case project.FieldValueDecimal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value_decimal", values[i])
			} else if value.Valid {
				pr.ValueDecimal = int(value.Int64)
			}
		case project.FieldForecastedApr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field forecasted_apr", values[i])
			} else if value.Valid {
				pr.ForecastedApr = value.String
			}
		case project.FieldTotalValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field total_value", values[i])
			} else if value.Valid {
				pr.TotalValue = value.String
			}
		case project.FieldPaymentToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field payment_token", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.PaymentToken); err != nil {
					return fmt.Errorf("unmarshal field payment_token: %w", err)
				}
			}
		case project.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryMint queries the "mint" edge of the Project entity.
func (pr *Project) QueryMint() *MintQuery {
	return NewProjectClient(pr.config).QueryMint(pr)
}

// QueryLaunchpad queries the "launchpad" edge of the Project entity.
func (pr *Project) QueryLaunchpad() *LaunchpadQuery {
	return NewProjectClient(pr.config).QueryLaunchpad(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("address=")
	builder.WriteString(pr.Address)
	builder.WriteString(", ")
	builder.WriteString("slot=")
	builder.WriteString(fmt.Sprintf("%v", pr.Slot))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(pr.Slug)
	builder.WriteString(", ")
	builder.WriteString("value_decimal=")
	builder.WriteString(fmt.Sprintf("%v", pr.ValueDecimal))
	builder.WriteString(", ")
	builder.WriteString("forecasted_apr=")
	builder.WriteString(pr.ForecastedApr)
	builder.WriteString(", ")
	builder.WriteString("total_value=")
	builder.WriteString(pr.TotalValue)
	builder.WriteString(", ")
	builder.WriteString("payment_token=")
	builder.WriteString(fmt.Sprintf("%v", pr.PaymentToken))
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", pr.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
