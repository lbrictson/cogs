// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/cogs/ent/schema"
	"github.com/lbrictson/cogs/ent/script"
)

// Script is the model entity for the Script schema.
type Script struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Script holds the value of the "script" field.
	Script string `json:"script,omitempty"`
	// TimeoutSeconds holds the value of the "timeout_seconds" field.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID int `json:"project_id,omitempty"`
	// Parameters holds the value of the "parameters" field.
	Parameters []schema.ScriptInputOptions `json:"parameters,omitempty"`
	// ScheduleEnabled holds the value of the "schedule_enabled" field.
	ScheduleEnabled bool `json:"schedule_enabled,omitempty"`
	// ScheduleCron holds the value of the "schedule_cron" field.
	ScheduleCron string `json:"schedule_cron,omitempty"`
	// SuccessNotificationChannelID holds the value of the "success_notification_channel_id" field.
	SuccessNotificationChannelID *int `json:"success_notification_channel_id,omitempty"`
	// FailureNotificationChannelID holds the value of the "failure_notification_channel_id" field.
	FailureNotificationChannelID *int `json:"failure_notification_channel_id,omitempty"`
	selectValues                 sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Script) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case script.FieldParameters:
			values[i] = new([]byte)
		case script.FieldScheduleEnabled:
			values[i] = new(sql.NullBool)
		case script.FieldID, script.FieldTimeoutSeconds, script.FieldProjectID, script.FieldSuccessNotificationChannelID, script.FieldFailureNotificationChannelID:
			values[i] = new(sql.NullInt64)
		case script.FieldName, script.FieldDescription, script.FieldScript, script.FieldScheduleCron:
			values[i] = new(sql.NullString)
		case script.FieldCreatedAt, script.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Script fields.
func (s *Script) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case script.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case script.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case script.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case script.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case script.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case script.FieldScript:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field script", values[i])
			} else if value.Valid {
				s.Script = value.String
			}
		case script.FieldTimeoutSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timeout_seconds", values[i])
			} else if value.Valid {
				s.TimeoutSeconds = int(value.Int64)
			}
		case script.FieldProjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value.Valid {
				s.ProjectID = int(value.Int64)
			}
		case script.FieldParameters:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field parameters", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Parameters); err != nil {
					return fmt.Errorf("unmarshal field parameters: %w", err)
				}
			}
		case script.FieldScheduleEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_enabled", values[i])
			} else if value.Valid {
				s.ScheduleEnabled = value.Bool
			}
		case script.FieldScheduleCron:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_cron", values[i])
			} else if value.Valid {
				s.ScheduleCron = value.String
			}
		case script.FieldSuccessNotificationChannelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field success_notification_channel_id", values[i])
			} else if value.Valid {
				s.SuccessNotificationChannelID = new(int)
				*s.SuccessNotificationChannelID = int(value.Int64)
			}
		case script.FieldFailureNotificationChannelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field failure_notification_channel_id", values[i])
			} else if value.Valid {
				s.FailureNotificationChannelID = new(int)
				*s.FailureNotificationChannelID = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Script.
// This includes values selected through modifiers, order, etc.
func (s *Script) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Script.
// Note that you need to call Script.Unwrap() before calling this method if this Script
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Script) Update() *ScriptUpdateOne {
	return NewScriptClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Script entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Script) Unwrap() *Script {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Script is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Script) String() string {
	var builder strings.Builder
	builder.WriteString("Script(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("script=")
	builder.WriteString(s.Script)
	builder.WriteString(", ")
	builder.WriteString("timeout_seconds=")
	builder.WriteString(fmt.Sprintf("%v", s.TimeoutSeconds))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("parameters=")
	builder.WriteString(fmt.Sprintf("%v", s.Parameters))
	builder.WriteString(", ")
	builder.WriteString("schedule_enabled=")
	builder.WriteString(fmt.Sprintf("%v", s.ScheduleEnabled))
	builder.WriteString(", ")
	builder.WriteString("schedule_cron=")
	builder.WriteString(s.ScheduleCron)
	builder.WriteString(", ")
	if v := s.SuccessNotificationChannelID; v != nil {
		builder.WriteString("success_notification_channel_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := s.FailureNotificationChannelID; v != nil {
		builder.WriteString("failure_notification_channel_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Scripts is a parsable slice of Script.
type Scripts []*Script
