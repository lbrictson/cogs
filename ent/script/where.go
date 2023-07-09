// Code generated by ent, DO NOT EDIT.

package script

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/cogs/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldDescription, v))
}

// Script applies equality check predicate on the "script" field. It's identical to ScriptEQ.
func Script(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldScript, v))
}

// TimeoutSeconds applies equality check predicate on the "timeout_seconds" field. It's identical to TimeoutSecondsEQ.
func TimeoutSeconds(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldTimeoutSeconds, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldProjectID, v))
}

// ScheduleEnabled applies equality check predicate on the "schedule_enabled" field. It's identical to ScheduleEnabledEQ.
func ScheduleEnabled(v bool) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldScheduleEnabled, v))
}

// ScheduleCron applies equality check predicate on the "schedule_cron" field. It's identical to ScheduleCronEQ.
func ScheduleCron(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldScheduleCron, v))
}

// SuccessNotificationChannelID applies equality check predicate on the "success_notification_channel_id" field. It's identical to SuccessNotificationChannelIDEQ.
func SuccessNotificationChannelID(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldSuccessNotificationChannelID, v))
}

// FailureNotificationChannelID applies equality check predicate on the "failure_notification_channel_id" field. It's identical to FailureNotificationChannelIDEQ.
func FailureNotificationChannelID(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldFailureNotificationChannelID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Script {
	return predicate.Script(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Script {
	return predicate.Script(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Script {
	return predicate.Script(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Script {
	return predicate.Script(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Script {
	return predicate.Script(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Script {
	return predicate.Script(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Script {
	return predicate.Script(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Script {
	return predicate.Script(sql.FieldContainsFold(FieldDescription, v))
}

// ScriptEQ applies the EQ predicate on the "script" field.
func ScriptEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldScript, v))
}

// ScriptNEQ applies the NEQ predicate on the "script" field.
func ScriptNEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldScript, v))
}

// ScriptIn applies the In predicate on the "script" field.
func ScriptIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldScript, vs...))
}

// ScriptNotIn applies the NotIn predicate on the "script" field.
func ScriptNotIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldScript, vs...))
}

// ScriptGT applies the GT predicate on the "script" field.
func ScriptGT(v string) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldScript, v))
}

// ScriptGTE applies the GTE predicate on the "script" field.
func ScriptGTE(v string) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldScript, v))
}

// ScriptLT applies the LT predicate on the "script" field.
func ScriptLT(v string) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldScript, v))
}

// ScriptLTE applies the LTE predicate on the "script" field.
func ScriptLTE(v string) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldScript, v))
}

// ScriptContains applies the Contains predicate on the "script" field.
func ScriptContains(v string) predicate.Script {
	return predicate.Script(sql.FieldContains(FieldScript, v))
}

// ScriptHasPrefix applies the HasPrefix predicate on the "script" field.
func ScriptHasPrefix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasPrefix(FieldScript, v))
}

// ScriptHasSuffix applies the HasSuffix predicate on the "script" field.
func ScriptHasSuffix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasSuffix(FieldScript, v))
}

// ScriptEqualFold applies the EqualFold predicate on the "script" field.
func ScriptEqualFold(v string) predicate.Script {
	return predicate.Script(sql.FieldEqualFold(FieldScript, v))
}

// ScriptContainsFold applies the ContainsFold predicate on the "script" field.
func ScriptContainsFold(v string) predicate.Script {
	return predicate.Script(sql.FieldContainsFold(FieldScript, v))
}

// TimeoutSecondsEQ applies the EQ predicate on the "timeout_seconds" field.
func TimeoutSecondsEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldTimeoutSeconds, v))
}

// TimeoutSecondsNEQ applies the NEQ predicate on the "timeout_seconds" field.
func TimeoutSecondsNEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldTimeoutSeconds, v))
}

// TimeoutSecondsIn applies the In predicate on the "timeout_seconds" field.
func TimeoutSecondsIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldTimeoutSeconds, vs...))
}

// TimeoutSecondsNotIn applies the NotIn predicate on the "timeout_seconds" field.
func TimeoutSecondsNotIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldTimeoutSeconds, vs...))
}

// TimeoutSecondsGT applies the GT predicate on the "timeout_seconds" field.
func TimeoutSecondsGT(v int) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldTimeoutSeconds, v))
}

// TimeoutSecondsGTE applies the GTE predicate on the "timeout_seconds" field.
func TimeoutSecondsGTE(v int) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldTimeoutSeconds, v))
}

// TimeoutSecondsLT applies the LT predicate on the "timeout_seconds" field.
func TimeoutSecondsLT(v int) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldTimeoutSeconds, v))
}

// TimeoutSecondsLTE applies the LTE predicate on the "timeout_seconds" field.
func TimeoutSecondsLTE(v int) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldTimeoutSeconds, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v int) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v int) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v int) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v int) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldProjectID, v))
}

// ParametersIsNil applies the IsNil predicate on the "parameters" field.
func ParametersIsNil() predicate.Script {
	return predicate.Script(sql.FieldIsNull(FieldParameters))
}

// ParametersNotNil applies the NotNil predicate on the "parameters" field.
func ParametersNotNil() predicate.Script {
	return predicate.Script(sql.FieldNotNull(FieldParameters))
}

// ScheduleEnabledEQ applies the EQ predicate on the "schedule_enabled" field.
func ScheduleEnabledEQ(v bool) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldScheduleEnabled, v))
}

// ScheduleEnabledNEQ applies the NEQ predicate on the "schedule_enabled" field.
func ScheduleEnabledNEQ(v bool) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldScheduleEnabled, v))
}

// ScheduleCronEQ applies the EQ predicate on the "schedule_cron" field.
func ScheduleCronEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldScheduleCron, v))
}

// ScheduleCronNEQ applies the NEQ predicate on the "schedule_cron" field.
func ScheduleCronNEQ(v string) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldScheduleCron, v))
}

// ScheduleCronIn applies the In predicate on the "schedule_cron" field.
func ScheduleCronIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldScheduleCron, vs...))
}

// ScheduleCronNotIn applies the NotIn predicate on the "schedule_cron" field.
func ScheduleCronNotIn(vs ...string) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldScheduleCron, vs...))
}

// ScheduleCronGT applies the GT predicate on the "schedule_cron" field.
func ScheduleCronGT(v string) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldScheduleCron, v))
}

// ScheduleCronGTE applies the GTE predicate on the "schedule_cron" field.
func ScheduleCronGTE(v string) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldScheduleCron, v))
}

// ScheduleCronLT applies the LT predicate on the "schedule_cron" field.
func ScheduleCronLT(v string) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldScheduleCron, v))
}

// ScheduleCronLTE applies the LTE predicate on the "schedule_cron" field.
func ScheduleCronLTE(v string) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldScheduleCron, v))
}

// ScheduleCronContains applies the Contains predicate on the "schedule_cron" field.
func ScheduleCronContains(v string) predicate.Script {
	return predicate.Script(sql.FieldContains(FieldScheduleCron, v))
}

// ScheduleCronHasPrefix applies the HasPrefix predicate on the "schedule_cron" field.
func ScheduleCronHasPrefix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasPrefix(FieldScheduleCron, v))
}

// ScheduleCronHasSuffix applies the HasSuffix predicate on the "schedule_cron" field.
func ScheduleCronHasSuffix(v string) predicate.Script {
	return predicate.Script(sql.FieldHasSuffix(FieldScheduleCron, v))
}

// ScheduleCronIsNil applies the IsNil predicate on the "schedule_cron" field.
func ScheduleCronIsNil() predicate.Script {
	return predicate.Script(sql.FieldIsNull(FieldScheduleCron))
}

// ScheduleCronNotNil applies the NotNil predicate on the "schedule_cron" field.
func ScheduleCronNotNil() predicate.Script {
	return predicate.Script(sql.FieldNotNull(FieldScheduleCron))
}

// ScheduleCronEqualFold applies the EqualFold predicate on the "schedule_cron" field.
func ScheduleCronEqualFold(v string) predicate.Script {
	return predicate.Script(sql.FieldEqualFold(FieldScheduleCron, v))
}

// ScheduleCronContainsFold applies the ContainsFold predicate on the "schedule_cron" field.
func ScheduleCronContainsFold(v string) predicate.Script {
	return predicate.Script(sql.FieldContainsFold(FieldScheduleCron, v))
}

// SuccessNotificationChannelIDEQ applies the EQ predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldSuccessNotificationChannelID, v))
}

// SuccessNotificationChannelIDNEQ applies the NEQ predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDNEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldSuccessNotificationChannelID, v))
}

// SuccessNotificationChannelIDIn applies the In predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldSuccessNotificationChannelID, vs...))
}

// SuccessNotificationChannelIDNotIn applies the NotIn predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDNotIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldSuccessNotificationChannelID, vs...))
}

// SuccessNotificationChannelIDGT applies the GT predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDGT(v int) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldSuccessNotificationChannelID, v))
}

// SuccessNotificationChannelIDGTE applies the GTE predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDGTE(v int) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldSuccessNotificationChannelID, v))
}

// SuccessNotificationChannelIDLT applies the LT predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDLT(v int) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldSuccessNotificationChannelID, v))
}

// SuccessNotificationChannelIDLTE applies the LTE predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDLTE(v int) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldSuccessNotificationChannelID, v))
}

// SuccessNotificationChannelIDIsNil applies the IsNil predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDIsNil() predicate.Script {
	return predicate.Script(sql.FieldIsNull(FieldSuccessNotificationChannelID))
}

// SuccessNotificationChannelIDNotNil applies the NotNil predicate on the "success_notification_channel_id" field.
func SuccessNotificationChannelIDNotNil() predicate.Script {
	return predicate.Script(sql.FieldNotNull(FieldSuccessNotificationChannelID))
}

// FailureNotificationChannelIDEQ applies the EQ predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldEQ(FieldFailureNotificationChannelID, v))
}

// FailureNotificationChannelIDNEQ applies the NEQ predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDNEQ(v int) predicate.Script {
	return predicate.Script(sql.FieldNEQ(FieldFailureNotificationChannelID, v))
}

// FailureNotificationChannelIDIn applies the In predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldIn(FieldFailureNotificationChannelID, vs...))
}

// FailureNotificationChannelIDNotIn applies the NotIn predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDNotIn(vs ...int) predicate.Script {
	return predicate.Script(sql.FieldNotIn(FieldFailureNotificationChannelID, vs...))
}

// FailureNotificationChannelIDGT applies the GT predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDGT(v int) predicate.Script {
	return predicate.Script(sql.FieldGT(FieldFailureNotificationChannelID, v))
}

// FailureNotificationChannelIDGTE applies the GTE predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDGTE(v int) predicate.Script {
	return predicate.Script(sql.FieldGTE(FieldFailureNotificationChannelID, v))
}

// FailureNotificationChannelIDLT applies the LT predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDLT(v int) predicate.Script {
	return predicate.Script(sql.FieldLT(FieldFailureNotificationChannelID, v))
}

// FailureNotificationChannelIDLTE applies the LTE predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDLTE(v int) predicate.Script {
	return predicate.Script(sql.FieldLTE(FieldFailureNotificationChannelID, v))
}

// FailureNotificationChannelIDIsNil applies the IsNil predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDIsNil() predicate.Script {
	return predicate.Script(sql.FieldIsNull(FieldFailureNotificationChannelID))
}

// FailureNotificationChannelIDNotNil applies the NotNil predicate on the "failure_notification_channel_id" field.
func FailureNotificationChannelIDNotNil() predicate.Script {
	return predicate.Script(sql.FieldNotNull(FieldFailureNotificationChannelID))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Script) predicate.Script {
	return predicate.Script(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Script) predicate.Script {
	return predicate.Script(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Script) predicate.Script {
	return predicate.Script(func(s *sql.Selector) {
		p(s.Not())
	})
}
