// Code generated by ent, DO NOT EDIT.

package scriptstats

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the scriptstats type in the database.
	Label = "script_stats"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldScriptID holds the string denoting the script_id field in the database.
	FieldScriptID = "script_id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldTotalRuns holds the string denoting the total_runs field in the database.
	FieldTotalRuns = "total_runs"
	// FieldTotalErrors holds the string denoting the total_errors field in the database.
	FieldTotalErrors = "total_errors"
	// FieldTotalSuccess holds the string denoting the total_success field in the database.
	FieldTotalSuccess = "total_success"
	// FieldAverageRuntime holds the string denoting the average_runtime field in the database.
	FieldAverageRuntime = "average_runtime"
	// FieldMinRuntime holds the string denoting the min_runtime field in the database.
	FieldMinRuntime = "min_runtime"
	// FieldMaxRuntime holds the string denoting the max_runtime field in the database.
	FieldMaxRuntime = "max_runtime"
	// FieldLastRun holds the string denoting the last_run field in the database.
	FieldLastRun = "last_run"
	// FieldTotalRuntime holds the string denoting the total_runtime field in the database.
	FieldTotalRuntime = "total_runtime"
	// FieldSuccessRate holds the string denoting the success_rate field in the database.
	FieldSuccessRate = "success_rate"
	// Table holds the table name of the scriptstats in the database.
	Table = "script_stats"
)

// Columns holds all SQL columns for scriptstats fields.
var Columns = []string{
	FieldID,
	FieldScriptID,
	FieldProjectID,
	FieldTotalRuns,
	FieldTotalErrors,
	FieldTotalSuccess,
	FieldAverageRuntime,
	FieldMinRuntime,
	FieldMaxRuntime,
	FieldLastRun,
	FieldTotalRuntime,
	FieldSuccessRate,
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

// OrderOption defines the ordering options for the ScriptStats queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByScriptID orders the results by the script_id field.
func ByScriptID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScriptID, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByTotalRuns orders the results by the total_runs field.
func ByTotalRuns(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalRuns, opts...).ToFunc()
}

// ByTotalErrors orders the results by the total_errors field.
func ByTotalErrors(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalErrors, opts...).ToFunc()
}

// ByTotalSuccess orders the results by the total_success field.
func ByTotalSuccess(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalSuccess, opts...).ToFunc()
}

// ByAverageRuntime orders the results by the average_runtime field.
func ByAverageRuntime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAverageRuntime, opts...).ToFunc()
}

// ByMinRuntime orders the results by the min_runtime field.
func ByMinRuntime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinRuntime, opts...).ToFunc()
}

// ByMaxRuntime orders the results by the max_runtime field.
func ByMaxRuntime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxRuntime, opts...).ToFunc()
}

// ByLastRun orders the results by the last_run field.
func ByLastRun(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastRun, opts...).ToFunc()
}

// ByTotalRuntime orders the results by the total_runtime field.
func ByTotalRuntime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalRuntime, opts...).ToFunc()
}

// BySuccessRate orders the results by the success_rate field.
func BySuccessRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuccessRate, opts...).ToFunc()
}
