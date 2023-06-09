// Code generated by ent, DO NOT EDIT.

package script

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the script type in the database.
	Label = "script"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldScript holds the string denoting the script field in the database.
	FieldScript = "script"
	// FieldTimeoutSeconds holds the string denoting the timeout_seconds field in the database.
	FieldTimeoutSeconds = "timeout_seconds"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldParameters holds the string denoting the parameters field in the database.
	FieldParameters = "parameters"
	// FieldScheduleEnabled holds the string denoting the schedule_enabled field in the database.
	FieldScheduleEnabled = "schedule_enabled"
	// FieldScheduleCron holds the string denoting the schedule_cron field in the database.
	FieldScheduleCron = "schedule_cron"
	// FieldSuccessNotificationChannelID holds the string denoting the success_notification_channel_id field in the database.
	FieldSuccessNotificationChannelID = "success_notification_channel_id"
	// FieldFailureNotificationChannelID holds the string denoting the failure_notification_channel_id field in the database.
	FieldFailureNotificationChannelID = "failure_notification_channel_id"
	// Table holds the table name of the script in the database.
	Table = "scripts"
)

// Columns holds all SQL columns for script fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldScript,
	FieldTimeoutSeconds,
	FieldProjectID,
	FieldParameters,
	FieldScheduleEnabled,
	FieldScheduleCron,
	FieldSuccessNotificationChannelID,
	FieldFailureNotificationChannelID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTimeoutSeconds holds the default value on creation for the "timeout_seconds" field.
	DefaultTimeoutSeconds int
	// DefaultScheduleEnabled holds the default value on creation for the "schedule_enabled" field.
	DefaultScheduleEnabled bool
)

// OrderOption defines the ordering options for the Script queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByScript orders the results by the script field.
func ByScript(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScript, opts...).ToFunc()
}

// ByTimeoutSeconds orders the results by the timeout_seconds field.
func ByTimeoutSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeoutSeconds, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByScheduleEnabled orders the results by the schedule_enabled field.
func ByScheduleEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduleEnabled, opts...).ToFunc()
}

// ByScheduleCron orders the results by the schedule_cron field.
func ByScheduleCron(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduleCron, opts...).ToFunc()
}

// BySuccessNotificationChannelID orders the results by the success_notification_channel_id field.
func BySuccessNotificationChannelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuccessNotificationChannelID, opts...).ToFunc()
}

// ByFailureNotificationChannelID orders the results by the failure_notification_channel_id field.
func ByFailureNotificationChannelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFailureNotificationChannelID, opts...).ToFunc()
}
