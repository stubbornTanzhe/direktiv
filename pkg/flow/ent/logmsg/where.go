// Code generated by ent, DO NOT EDIT.

package logmsg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// T applies equality check predicate on the "t" field. It's identical to TEQ.
func T(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldT), v))
	})
}

// Msg applies equality check predicate on the "msg" field. It's identical to MsgEQ.
func Msg(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsg), v))
	})
}

// TEQ applies the EQ predicate on the "t" field.
func TEQ(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldT), v))
	})
}

// TNEQ applies the NEQ predicate on the "t" field.
func TNEQ(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldT), v))
	})
}

// TIn applies the In predicate on the "t" field.
func TIn(vs ...time.Time) predicate.LogMsg {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldT), v...))
	})
}

// TNotIn applies the NotIn predicate on the "t" field.
func TNotIn(vs ...time.Time) predicate.LogMsg {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldT), v...))
	})
}

// TGT applies the GT predicate on the "t" field.
func TGT(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldT), v))
	})
}

// TGTE applies the GTE predicate on the "t" field.
func TGTE(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldT), v))
	})
}

// TLT applies the LT predicate on the "t" field.
func TLT(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldT), v))
	})
}

// TLTE applies the LTE predicate on the "t" field.
func TLTE(v time.Time) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldT), v))
	})
}

// MsgEQ applies the EQ predicate on the "msg" field.
func MsgEQ(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsg), v))
	})
}

// MsgNEQ applies the NEQ predicate on the "msg" field.
func MsgNEQ(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsg), v))
	})
}

// MsgIn applies the In predicate on the "msg" field.
func MsgIn(vs ...string) predicate.LogMsg {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsg), v...))
	})
}

// MsgNotIn applies the NotIn predicate on the "msg" field.
func MsgNotIn(vs ...string) predicate.LogMsg {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsg), v...))
	})
}

// MsgGT applies the GT predicate on the "msg" field.
func MsgGT(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsg), v))
	})
}

// MsgGTE applies the GTE predicate on the "msg" field.
func MsgGTE(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsg), v))
	})
}

// MsgLT applies the LT predicate on the "msg" field.
func MsgLT(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsg), v))
	})
}

// MsgLTE applies the LTE predicate on the "msg" field.
func MsgLTE(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsg), v))
	})
}

// MsgContains applies the Contains predicate on the "msg" field.
func MsgContains(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMsg), v))
	})
}

// MsgHasPrefix applies the HasPrefix predicate on the "msg" field.
func MsgHasPrefix(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMsg), v))
	})
}

// MsgHasSuffix applies the HasSuffix predicate on the "msg" field.
func MsgHasSuffix(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMsg), v))
	})
}

// MsgEqualFold applies the EqualFold predicate on the "msg" field.
func MsgEqualFold(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMsg), v))
	})
}

// MsgContainsFold applies the ContainsFold predicate on the "msg" field.
func MsgContainsFold(v string) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMsg), v))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
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

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstance applies the HasEdge predicate on the "instance" edge.
func HasInstance() predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstanceWith applies the HasEdge predicate on the "instance" edge with a given conditions (other predicates).
func HasInstanceWith(preds ...predicate.Instance) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActivity applies the HasEdge predicate on the "activity" edge.
func HasActivity() predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivityTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActivityWith applies the HasEdge predicate on the "activity" edge with a given conditions (other predicates).
func HasActivityWith(preds ...predicate.MirrorActivity) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivityInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LogMsg) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LogMsg) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
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
func Not(p predicate.LogMsg) predicate.LogMsg {
	return predicate.LogMsg(func(s *sql.Selector) {
		p(s.Not())
	})
}
