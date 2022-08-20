// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithenum"
	"entgo.io/ent/dialect/sql"
)

// MessageWithEnum is the model entity for the MessageWithEnum schema.
type MessageWithEnum struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EnumType holds the value of the "enum_type" field.
	EnumType messagewithenum.EnumType `json:"enum_type,omitempty"`
	// EnumWithoutDefault holds the value of the "enum_without_default" field.
	EnumWithoutDefault messagewithenum.EnumWithoutDefault `json:"enum_without_default,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageWithEnum) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagewithenum.FieldID:
			values[i] = new(sql.NullInt64)
		case messagewithenum.FieldEnumType, messagewithenum.FieldEnumWithoutDefault:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MessageWithEnum", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageWithEnum fields.
func (mwe *MessageWithEnum) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagewithenum.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mwe.ID = int(value.Int64)
		case messagewithenum.FieldEnumType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field enum_type", values[i])
			} else if value.Valid {
				mwe.EnumType = messagewithenum.EnumType(value.String)
			}
		case messagewithenum.FieldEnumWithoutDefault:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field enum_without_default", values[i])
			} else if value.Valid {
				mwe.EnumWithoutDefault = messagewithenum.EnumWithoutDefault(value.String)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MessageWithEnum.
// Note that you need to call MessageWithEnum.Unwrap() before calling this method if this MessageWithEnum
// was returned from a transaction, and the transaction was committed or rolled back.
func (mwe *MessageWithEnum) Update() *MessageWithEnumUpdateOne {
	return (&MessageWithEnumClient{config: mwe.config}).UpdateOne(mwe)
}

// Unwrap unwraps the MessageWithEnum entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mwe *MessageWithEnum) Unwrap() *MessageWithEnum {
	_tx, ok := mwe.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageWithEnum is not a transactional entity")
	}
	mwe.config.driver = _tx.drv
	return mwe
}

// String implements the fmt.Stringer.
func (mwe *MessageWithEnum) String() string {
	var builder strings.Builder
	builder.WriteString("MessageWithEnum(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mwe.ID))
	builder.WriteString("enum_type=")
	builder.WriteString(fmt.Sprintf("%v", mwe.EnumType))
	builder.WriteString(", ")
	builder.WriteString("enum_without_default=")
	builder.WriteString(fmt.Sprintf("%v", mwe.EnumWithoutDefault))
	builder.WriteByte(')')
	return builder.String()
}

// MessageWithEnums is a parsable slice of MessageWithEnum.
type MessageWithEnums []*MessageWithEnum

func (mwe MessageWithEnums) config(cfg config) {
	for _i := range mwe {
		mwe[_i].config = cfg
	}
}
