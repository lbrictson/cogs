// Code generated by ent, DO NOT EDIT.

package access

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/cogs/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldUserID, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldProjectID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldUserID, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v int) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v int) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...int) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...int) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v int) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v int) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v int) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v int) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldProjectID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldRole, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
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
func Not(p predicate.Access) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		p(s.Not())
	})
}
