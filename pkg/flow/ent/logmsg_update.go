// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// LogMsgUpdate is the builder for updating LogMsg entities.
type LogMsgUpdate struct {
	config
	hooks     []Hook
	mutation  *LogMsgMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LogMsgUpdate builder.
func (lmu *LogMsgUpdate) Where(ps ...predicate.LogMsg) *LogMsgUpdate {
	lmu.mutation.Where(ps...)
	return lmu
}

// SetT sets the "t" field.
func (lmu *LogMsgUpdate) SetT(t time.Time) *LogMsgUpdate {
	lmu.mutation.SetT(t)
	return lmu
}

// SetMsg sets the "msg" field.
func (lmu *LogMsgUpdate) SetMsg(s string) *LogMsgUpdate {
	lmu.mutation.SetMsg(s)
	return lmu
}

// SetLevel sets the "level" field.
func (lmu *LogMsgUpdate) SetLevel(s string) *LogMsgUpdate {
	lmu.mutation.SetLevel(s)
	return lmu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableLevel(s *string) *LogMsgUpdate {
	if s != nil {
		lmu.SetLevel(*s)
	}
	return lmu
}

// SetRootInstanceId sets the "rootInstanceId" field.
func (lmu *LogMsgUpdate) SetRootInstanceId(s string) *LogMsgUpdate {
	lmu.mutation.SetRootInstanceId(s)
	return lmu
}

// SetNillableRootInstanceId sets the "rootInstanceId" field if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableRootInstanceId(s *string) *LogMsgUpdate {
	if s != nil {
		lmu.SetRootInstanceId(*s)
	}
	return lmu
}

// SetLogInstanceCallPath sets the "logInstanceCallPath" field.
func (lmu *LogMsgUpdate) SetLogInstanceCallPath(s string) *LogMsgUpdate {
	lmu.mutation.SetLogInstanceCallPath(s)
	return lmu
}

// SetNillableLogInstanceCallPath sets the "logInstanceCallPath" field if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableLogInstanceCallPath(s *string) *LogMsgUpdate {
	if s != nil {
		lmu.SetLogInstanceCallPath(*s)
	}
	return lmu
}

// SetTags sets the "tags" field.
func (lmu *LogMsgUpdate) SetTags(m map[string]string) *LogMsgUpdate {
	lmu.mutation.SetTags(m)
	return lmu
}

// ClearTags clears the value of the "tags" field.
func (lmu *LogMsgUpdate) ClearTags() *LogMsgUpdate {
	lmu.mutation.ClearTags()
	return lmu
}

// SetWorkflowID sets the "workflow_id" field.
func (lmu *LogMsgUpdate) SetWorkflowID(u uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetWorkflowID(u)
	return lmu
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableWorkflowID(u *uuid.UUID) *LogMsgUpdate {
	if u != nil {
		lmu.SetWorkflowID(*u)
	}
	return lmu
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (lmu *LogMsgUpdate) ClearWorkflowID() *LogMsgUpdate {
	lmu.mutation.ClearWorkflowID()
	return lmu
}

// SetMirrorActivityID sets the "mirror_activity_id" field.
func (lmu *LogMsgUpdate) SetMirrorActivityID(u uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetMirrorActivityID(u)
	return lmu
}

// SetNillableMirrorActivityID sets the "mirror_activity_id" field if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableMirrorActivityID(u *uuid.UUID) *LogMsgUpdate {
	if u != nil {
		lmu.SetMirrorActivityID(*u)
	}
	return lmu
}

// ClearMirrorActivityID clears the value of the "mirror_activity_id" field.
func (lmu *LogMsgUpdate) ClearMirrorActivityID() *LogMsgUpdate {
	lmu.mutation.ClearMirrorActivityID()
	return lmu
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (lmu *LogMsgUpdate) SetNamespaceID(id uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetNamespaceID(id)
	return lmu
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableNamespaceID(id *uuid.UUID) *LogMsgUpdate {
	if id != nil {
		lmu = lmu.SetNamespaceID(*id)
	}
	return lmu
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (lmu *LogMsgUpdate) SetNamespace(n *Namespace) *LogMsgUpdate {
	return lmu.SetNamespaceID(n.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (lmu *LogMsgUpdate) SetInstanceID(id uuid.UUID) *LogMsgUpdate {
	lmu.mutation.SetInstanceID(id)
	return lmu
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (lmu *LogMsgUpdate) SetNillableInstanceID(id *uuid.UUID) *LogMsgUpdate {
	if id != nil {
		lmu = lmu.SetInstanceID(*id)
	}
	return lmu
}

// SetInstance sets the "instance" edge to the Instance entity.
func (lmu *LogMsgUpdate) SetInstance(i *Instance) *LogMsgUpdate {
	return lmu.SetInstanceID(i.ID)
}

// Mutation returns the LogMsgMutation object of the builder.
func (lmu *LogMsgUpdate) Mutation() *LogMsgMutation {
	return lmu.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (lmu *LogMsgUpdate) ClearNamespace() *LogMsgUpdate {
	lmu.mutation.ClearNamespace()
	return lmu
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (lmu *LogMsgUpdate) ClearInstance() *LogMsgUpdate {
	lmu.mutation.ClearInstance()
	return lmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmu *LogMsgUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LogMsgMutation](ctx, lmu.sqlSave, lmu.mutation, lmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lmu *LogMsgUpdate) SaveX(ctx context.Context) int {
	affected, err := lmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmu *LogMsgUpdate) Exec(ctx context.Context) error {
	_, err := lmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmu *LogMsgUpdate) ExecX(ctx context.Context) {
	if err := lmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lmu *LogMsgUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LogMsgUpdate {
	lmu.modifiers = append(lmu.modifiers, modifiers...)
	return lmu
}

func (lmu *LogMsgUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(logmsg.Table, logmsg.Columns, sqlgraph.NewFieldSpec(logmsg.FieldID, field.TypeUUID))
	if ps := lmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmu.mutation.T(); ok {
		_spec.SetField(logmsg.FieldT, field.TypeTime, value)
	}
	if value, ok := lmu.mutation.Msg(); ok {
		_spec.SetField(logmsg.FieldMsg, field.TypeString, value)
	}
	if value, ok := lmu.mutation.Level(); ok {
		_spec.SetField(logmsg.FieldLevel, field.TypeString, value)
	}
	if value, ok := lmu.mutation.RootInstanceId(); ok {
		_spec.SetField(logmsg.FieldRootInstanceId, field.TypeString, value)
	}
	if value, ok := lmu.mutation.LogInstanceCallPath(); ok {
		_spec.SetField(logmsg.FieldLogInstanceCallPath, field.TypeString, value)
	}
	if value, ok := lmu.mutation.Tags(); ok {
		_spec.SetField(logmsg.FieldTags, field.TypeJSON, value)
	}
	if lmu.mutation.TagsCleared() {
		_spec.ClearField(logmsg.FieldTags, field.TypeJSON)
	}
	if value, ok := lmu.mutation.WorkflowID(); ok {
		_spec.SetField(logmsg.FieldWorkflowID, field.TypeUUID, value)
	}
	if lmu.mutation.WorkflowIDCleared() {
		_spec.ClearField(logmsg.FieldWorkflowID, field.TypeUUID)
	}
	if value, ok := lmu.mutation.MirrorActivityID(); ok {
		_spec.SetField(logmsg.FieldMirrorActivityID, field.TypeUUID, value)
	}
	if lmu.mutation.MirrorActivityIDCleared() {
		_spec.ClearField(logmsg.FieldMirrorActivityID, field.TypeUUID)
	}
	if lmu.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmu.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmu.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lmu.mutation.done = true
	return n, nil
}

// LogMsgUpdateOne is the builder for updating a single LogMsg entity.
type LogMsgUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LogMsgMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetT sets the "t" field.
func (lmuo *LogMsgUpdateOne) SetT(t time.Time) *LogMsgUpdateOne {
	lmuo.mutation.SetT(t)
	return lmuo
}

// SetMsg sets the "msg" field.
func (lmuo *LogMsgUpdateOne) SetMsg(s string) *LogMsgUpdateOne {
	lmuo.mutation.SetMsg(s)
	return lmuo
}

// SetLevel sets the "level" field.
func (lmuo *LogMsgUpdateOne) SetLevel(s string) *LogMsgUpdateOne {
	lmuo.mutation.SetLevel(s)
	return lmuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableLevel(s *string) *LogMsgUpdateOne {
	if s != nil {
		lmuo.SetLevel(*s)
	}
	return lmuo
}

// SetRootInstanceId sets the "rootInstanceId" field.
func (lmuo *LogMsgUpdateOne) SetRootInstanceId(s string) *LogMsgUpdateOne {
	lmuo.mutation.SetRootInstanceId(s)
	return lmuo
}

// SetNillableRootInstanceId sets the "rootInstanceId" field if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableRootInstanceId(s *string) *LogMsgUpdateOne {
	if s != nil {
		lmuo.SetRootInstanceId(*s)
	}
	return lmuo
}

// SetLogInstanceCallPath sets the "logInstanceCallPath" field.
func (lmuo *LogMsgUpdateOne) SetLogInstanceCallPath(s string) *LogMsgUpdateOne {
	lmuo.mutation.SetLogInstanceCallPath(s)
	return lmuo
}

// SetNillableLogInstanceCallPath sets the "logInstanceCallPath" field if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableLogInstanceCallPath(s *string) *LogMsgUpdateOne {
	if s != nil {
		lmuo.SetLogInstanceCallPath(*s)
	}
	return lmuo
}

// SetTags sets the "tags" field.
func (lmuo *LogMsgUpdateOne) SetTags(m map[string]string) *LogMsgUpdateOne {
	lmuo.mutation.SetTags(m)
	return lmuo
}

// ClearTags clears the value of the "tags" field.
func (lmuo *LogMsgUpdateOne) ClearTags() *LogMsgUpdateOne {
	lmuo.mutation.ClearTags()
	return lmuo
}

// SetWorkflowID sets the "workflow_id" field.
func (lmuo *LogMsgUpdateOne) SetWorkflowID(u uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetWorkflowID(u)
	return lmuo
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableWorkflowID(u *uuid.UUID) *LogMsgUpdateOne {
	if u != nil {
		lmuo.SetWorkflowID(*u)
	}
	return lmuo
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (lmuo *LogMsgUpdateOne) ClearWorkflowID() *LogMsgUpdateOne {
	lmuo.mutation.ClearWorkflowID()
	return lmuo
}

// SetMirrorActivityID sets the "mirror_activity_id" field.
func (lmuo *LogMsgUpdateOne) SetMirrorActivityID(u uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetMirrorActivityID(u)
	return lmuo
}

// SetNillableMirrorActivityID sets the "mirror_activity_id" field if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableMirrorActivityID(u *uuid.UUID) *LogMsgUpdateOne {
	if u != nil {
		lmuo.SetMirrorActivityID(*u)
	}
	return lmuo
}

// ClearMirrorActivityID clears the value of the "mirror_activity_id" field.
func (lmuo *LogMsgUpdateOne) ClearMirrorActivityID() *LogMsgUpdateOne {
	lmuo.mutation.ClearMirrorActivityID()
	return lmuo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (lmuo *LogMsgUpdateOne) SetNamespaceID(id uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetNamespaceID(id)
	return lmuo
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableNamespaceID(id *uuid.UUID) *LogMsgUpdateOne {
	if id != nil {
		lmuo = lmuo.SetNamespaceID(*id)
	}
	return lmuo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (lmuo *LogMsgUpdateOne) SetNamespace(n *Namespace) *LogMsgUpdateOne {
	return lmuo.SetNamespaceID(n.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (lmuo *LogMsgUpdateOne) SetInstanceID(id uuid.UUID) *LogMsgUpdateOne {
	lmuo.mutation.SetInstanceID(id)
	return lmuo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (lmuo *LogMsgUpdateOne) SetNillableInstanceID(id *uuid.UUID) *LogMsgUpdateOne {
	if id != nil {
		lmuo = lmuo.SetInstanceID(*id)
	}
	return lmuo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (lmuo *LogMsgUpdateOne) SetInstance(i *Instance) *LogMsgUpdateOne {
	return lmuo.SetInstanceID(i.ID)
}

// Mutation returns the LogMsgMutation object of the builder.
func (lmuo *LogMsgUpdateOne) Mutation() *LogMsgMutation {
	return lmuo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (lmuo *LogMsgUpdateOne) ClearNamespace() *LogMsgUpdateOne {
	lmuo.mutation.ClearNamespace()
	return lmuo
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (lmuo *LogMsgUpdateOne) ClearInstance() *LogMsgUpdateOne {
	lmuo.mutation.ClearInstance()
	return lmuo
}

// Where appends a list predicates to the LogMsgUpdate builder.
func (lmuo *LogMsgUpdateOne) Where(ps ...predicate.LogMsg) *LogMsgUpdateOne {
	lmuo.mutation.Where(ps...)
	return lmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmuo *LogMsgUpdateOne) Select(field string, fields ...string) *LogMsgUpdateOne {
	lmuo.fields = append([]string{field}, fields...)
	return lmuo
}

// Save executes the query and returns the updated LogMsg entity.
func (lmuo *LogMsgUpdateOne) Save(ctx context.Context) (*LogMsg, error) {
	return withHooks[*LogMsg, LogMsgMutation](ctx, lmuo.sqlSave, lmuo.mutation, lmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lmuo *LogMsgUpdateOne) SaveX(ctx context.Context) *LogMsg {
	node, err := lmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmuo *LogMsgUpdateOne) Exec(ctx context.Context) error {
	_, err := lmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmuo *LogMsgUpdateOne) ExecX(ctx context.Context) {
	if err := lmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lmuo *LogMsgUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LogMsgUpdateOne {
	lmuo.modifiers = append(lmuo.modifiers, modifiers...)
	return lmuo
}

func (lmuo *LogMsgUpdateOne) sqlSave(ctx context.Context) (_node *LogMsg, err error) {
	_spec := sqlgraph.NewUpdateSpec(logmsg.Table, logmsg.Columns, sqlgraph.NewFieldSpec(logmsg.FieldID, field.TypeUUID))
	id, ok := lmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LogMsg.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logmsg.FieldID)
		for _, f := range fields {
			if !logmsg.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logmsg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmuo.mutation.T(); ok {
		_spec.SetField(logmsg.FieldT, field.TypeTime, value)
	}
	if value, ok := lmuo.mutation.Msg(); ok {
		_spec.SetField(logmsg.FieldMsg, field.TypeString, value)
	}
	if value, ok := lmuo.mutation.Level(); ok {
		_spec.SetField(logmsg.FieldLevel, field.TypeString, value)
	}
	if value, ok := lmuo.mutation.RootInstanceId(); ok {
		_spec.SetField(logmsg.FieldRootInstanceId, field.TypeString, value)
	}
	if value, ok := lmuo.mutation.LogInstanceCallPath(); ok {
		_spec.SetField(logmsg.FieldLogInstanceCallPath, field.TypeString, value)
	}
	if value, ok := lmuo.mutation.Tags(); ok {
		_spec.SetField(logmsg.FieldTags, field.TypeJSON, value)
	}
	if lmuo.mutation.TagsCleared() {
		_spec.ClearField(logmsg.FieldTags, field.TypeJSON)
	}
	if value, ok := lmuo.mutation.WorkflowID(); ok {
		_spec.SetField(logmsg.FieldWorkflowID, field.TypeUUID, value)
	}
	if lmuo.mutation.WorkflowIDCleared() {
		_spec.ClearField(logmsg.FieldWorkflowID, field.TypeUUID)
	}
	if value, ok := lmuo.mutation.MirrorActivityID(); ok {
		_spec.SetField(logmsg.FieldMirrorActivityID, field.TypeUUID, value)
	}
	if lmuo.mutation.MirrorActivityIDCleared() {
		_spec.ClearField(logmsg.FieldMirrorActivityID, field.TypeUUID)
	}
	if lmuo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lmuo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmuo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(lmuo.modifiers...)
	_node = &LogMsg{config: lmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lmuo.mutation.done = true
	return _node, nil
}
