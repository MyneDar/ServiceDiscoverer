// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"servicediscoverer/ent/providerregisterdata"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ProviderRegisterData is the model entity for the ProviderRegisterData schema.
type ProviderRegisterData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Port holds the value of the "port" field.
	Port string `json:"port,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// LiveInterval holds the value of the "liveInterval" field.
	LiveInterval int `json:"liveInterval,omitempty"`
	// LiveTimeout holds the value of the "liveTimeout" field.
	LiveTimeout int `json:"liveTimeout,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProviderRegisterData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case providerregisterdata.FieldID, providerregisterdata.FieldLiveInterval, providerregisterdata.FieldLiveTimeout:
			values[i] = new(sql.NullInt64)
		case providerregisterdata.FieldName, providerregisterdata.FieldPort, providerregisterdata.FieldAddress, providerregisterdata.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProviderRegisterData", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProviderRegisterData fields.
func (prd *ProviderRegisterData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case providerregisterdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			prd.ID = int(value.Int64)
		case providerregisterdata.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				prd.Name = value.String
			}
		case providerregisterdata.FieldPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				prd.Port = value.String
			}
		case providerregisterdata.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				prd.Address = value.String
			}
		case providerregisterdata.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				prd.Description = value.String
			}
		case providerregisterdata.FieldLiveInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field liveInterval", values[i])
			} else if value.Valid {
				prd.LiveInterval = int(value.Int64)
			}
		case providerregisterdata.FieldLiveTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field liveTimeout", values[i])
			} else if value.Valid {
				prd.LiveTimeout = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ProviderRegisterData.
// Note that you need to call ProviderRegisterData.Unwrap() before calling this method if this ProviderRegisterData
// was returned from a transaction, and the transaction was committed or rolled back.
func (prd *ProviderRegisterData) Update() *ProviderRegisterDataUpdateOne {
	return (&ProviderRegisterDataClient{config: prd.config}).UpdateOne(prd)
}

// Unwrap unwraps the ProviderRegisterData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (prd *ProviderRegisterData) Unwrap() *ProviderRegisterData {
	_tx, ok := prd.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProviderRegisterData is not a transactional entity")
	}
	prd.config.driver = _tx.drv
	return prd
}

// String implements the fmt.Stringer.
func (prd *ProviderRegisterData) String() string {
	var builder strings.Builder
	builder.WriteString("ProviderRegisterData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", prd.ID))
	builder.WriteString("name=")
	builder.WriteString(prd.Name)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(prd.Port)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(prd.Address)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(prd.Description)
	builder.WriteString(", ")
	builder.WriteString("liveInterval=")
	builder.WriteString(fmt.Sprintf("%v", prd.LiveInterval))
	builder.WriteString(", ")
	builder.WriteString("liveTimeout=")
	builder.WriteString(fmt.Sprintf("%v", prd.LiveTimeout))
	builder.WriteByte(')')
	return builder.String()
}

// ProviderRegisterDataSlice is a parsable slice of ProviderRegisterData.
type ProviderRegisterDataSlice []*ProviderRegisterData

func (prd ProviderRegisterDataSlice) config(cfg config) {
	for _i := range prd {
		prd[_i].config = cfg
	}
}
