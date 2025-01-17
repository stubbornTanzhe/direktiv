// Code generated by ent, DO NOT EDIT.

package services

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Services {
	return predicate.Services(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldUpdatedAt, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldURL, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldName, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldData, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Services {
	return predicate.Services(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Services {
	return predicate.Services(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Services {
	return predicate.Services(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Services {
	return predicate.Services(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Services {
	return predicate.Services(sql.FieldLTE(FieldUpdatedAt, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Services {
	return predicate.Services(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Services {
	return predicate.Services(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Services {
	return predicate.Services(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Services {
	return predicate.Services(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Services {
	return predicate.Services(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Services {
	return predicate.Services(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Services {
	return predicate.Services(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Services {
	return predicate.Services(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Services {
	return predicate.Services(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Services {
	return predicate.Services(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Services {
	return predicate.Services(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Services {
	return predicate.Services(sql.FieldContainsFold(FieldURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Services {
	return predicate.Services(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Services {
	return predicate.Services(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Services {
	return predicate.Services(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Services {
	return predicate.Services(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Services {
	return predicate.Services(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Services {
	return predicate.Services(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Services {
	return predicate.Services(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Services {
	return predicate.Services(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Services {
	return predicate.Services(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Services {
	return predicate.Services(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Services {
	return predicate.Services(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Services {
	return predicate.Services(sql.FieldContainsFold(FieldName, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Services {
	return predicate.Services(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Services {
	return predicate.Services(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Services {
	return predicate.Services(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Services {
	return predicate.Services(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Services {
	return predicate.Services(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Services {
	return predicate.Services(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Services {
	return predicate.Services(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Services {
	return predicate.Services(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Services {
	return predicate.Services(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Services {
	return predicate.Services(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Services {
	return predicate.Services(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Services {
	return predicate.Services(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Services {
	return predicate.Services(sql.FieldContainsFold(FieldData, v))
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Services {
	return predicate.Services(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Services {
	return predicate.Services(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Services) predicate.Services {
	return predicate.Services(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Services) predicate.Services {
	return predicate.Services(func(s *sql.Selector) {
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
func Not(p predicate.Services) predicate.Services {
	return predicate.Services(func(s *sql.Selector) {
		p(s.Not())
	})
}
