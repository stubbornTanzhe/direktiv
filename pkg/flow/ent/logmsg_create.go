// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// LogMsgCreate is the builder for creating a LogMsg entity.
type LogMsgCreate struct {
	config
	mutation *LogMsgMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetT sets the "t" field.
func (lmc *LogMsgCreate) SetT(t time.Time) *LogMsgCreate {
	lmc.mutation.SetT(t)
	return lmc
}

// SetMsg sets the "msg" field.
func (lmc *LogMsgCreate) SetMsg(s string) *LogMsgCreate {
	lmc.mutation.SetMsg(s)
	return lmc
}

// SetID sets the "id" field.
func (lmc *LogMsgCreate) SetID(u uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetID(u)
	return lmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableID(u *uuid.UUID) *LogMsgCreate {
	if u != nil {
		lmc.SetID(*u)
	}
	return lmc
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (lmc *LogMsgCreate) SetNamespaceID(id uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetNamespaceID(id)
	return lmc
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableNamespaceID(id *uuid.UUID) *LogMsgCreate {
	if id != nil {
		lmc = lmc.SetNamespaceID(*id)
	}
	return lmc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (lmc *LogMsgCreate) SetNamespace(n *Namespace) *LogMsgCreate {
	return lmc.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (lmc *LogMsgCreate) SetWorkflowID(id uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetWorkflowID(id)
	return lmc
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableWorkflowID(id *uuid.UUID) *LogMsgCreate {
	if id != nil {
		lmc = lmc.SetWorkflowID(*id)
	}
	return lmc
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (lmc *LogMsgCreate) SetWorkflow(w *Workflow) *LogMsgCreate {
	return lmc.SetWorkflowID(w.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (lmc *LogMsgCreate) SetInstanceID(id uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetInstanceID(id)
	return lmc
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableInstanceID(id *uuid.UUID) *LogMsgCreate {
	if id != nil {
		lmc = lmc.SetInstanceID(*id)
	}
	return lmc
}

// SetInstance sets the "instance" edge to the Instance entity.
func (lmc *LogMsgCreate) SetInstance(i *Instance) *LogMsgCreate {
	return lmc.SetInstanceID(i.ID)
}

// SetActivityID sets the "activity" edge to the MirrorActivity entity by ID.
func (lmc *LogMsgCreate) SetActivityID(id uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetActivityID(id)
	return lmc
}

// SetNillableActivityID sets the "activity" edge to the MirrorActivity entity by ID if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableActivityID(id *uuid.UUID) *LogMsgCreate {
	if id != nil {
		lmc = lmc.SetActivityID(*id)
	}
	return lmc
}

// SetActivity sets the "activity" edge to the MirrorActivity entity.
func (lmc *LogMsgCreate) SetActivity(m *MirrorActivity) *LogMsgCreate {
	return lmc.SetActivityID(m.ID)
}

// Mutation returns the LogMsgMutation object of the builder.
func (lmc *LogMsgCreate) Mutation() *LogMsgMutation {
	return lmc.mutation
}

// Save creates the LogMsg in the database.
func (lmc *LogMsgCreate) Save(ctx context.Context) (*LogMsg, error) {
	var (
		err  error
		node *LogMsg
	)
	lmc.defaults()
	if len(lmc.hooks) == 0 {
		if err = lmc.check(); err != nil {
			return nil, err
		}
		node, err = lmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogMsgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmc.check(); err != nil {
				return nil, err
			}
			lmc.mutation = mutation
			if node, err = lmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lmc.hooks) - 1; i >= 0; i-- {
			if lmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LogMsg)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LogMsgMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lmc *LogMsgCreate) SaveX(ctx context.Context) *LogMsg {
	v, err := lmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lmc *LogMsgCreate) Exec(ctx context.Context) error {
	_, err := lmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmc *LogMsgCreate) ExecX(ctx context.Context) {
	if err := lmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lmc *LogMsgCreate) defaults() {
	if _, ok := lmc.mutation.ID(); !ok {
		v := logmsg.DefaultID()
		lmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmc *LogMsgCreate) check() error {
	if _, ok := lmc.mutation.T(); !ok {
		return &ValidationError{Name: "t", err: errors.New(`ent: missing required field "LogMsg.t"`)}
	}
	if _, ok := lmc.mutation.Msg(); !ok {
		return &ValidationError{Name: "msg", err: errors.New(`ent: missing required field "LogMsg.msg"`)}
	}
	return nil
}

func (lmc *LogMsgCreate) sqlSave(ctx context.Context) (*LogMsg, error) {
	_node, _spec := lmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (lmc *LogMsgCreate) createSpec() (*LogMsg, *sqlgraph.CreateSpec) {
	var (
		_node = &LogMsg{config: lmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: logmsg.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: logmsg.FieldID,
			},
		}
	)
	_spec.OnConflict = lmc.conflict
	if id, ok := lmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lmc.mutation.T(); ok {
		_spec.SetField(logmsg.FieldT, field.TypeTime, value)
		_node.T = value
	}
	if value, ok := lmc.mutation.Msg(); ok {
		_spec.SetField(logmsg.FieldMsg, field.TypeString, value)
		_node.Msg = value
	}
	if nodes := lmc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lmc.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.WorkflowTable,
			Columns: []string{logmsg.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lmc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lmc.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.ActivityTable,
			Columns: []string{logmsg.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirroractivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.mirror_activity_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LogMsg.Create().
//		SetT(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LogMsgUpsert) {
//			SetT(v+v).
//		}).
//		Exec(ctx)
func (lmc *LogMsgCreate) OnConflict(opts ...sql.ConflictOption) *LogMsgUpsertOne {
	lmc.conflict = opts
	return &LogMsgUpsertOne{
		create: lmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lmc *LogMsgCreate) OnConflictColumns(columns ...string) *LogMsgUpsertOne {
	lmc.conflict = append(lmc.conflict, sql.ConflictColumns(columns...))
	return &LogMsgUpsertOne{
		create: lmc,
	}
}

type (
	// LogMsgUpsertOne is the builder for "upsert"-ing
	//  one LogMsg node.
	LogMsgUpsertOne struct {
		create *LogMsgCreate
	}

	// LogMsgUpsert is the "OnConflict" setter.
	LogMsgUpsert struct {
		*sql.UpdateSet
	}
)

// SetT sets the "t" field.
func (u *LogMsgUpsert) SetT(v time.Time) *LogMsgUpsert {
	u.Set(logmsg.FieldT, v)
	return u
}

// UpdateT sets the "t" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateT() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldT)
	return u
}

// SetMsg sets the "msg" field.
func (u *LogMsgUpsert) SetMsg(v string) *LogMsgUpsert {
	u.Set(logmsg.FieldMsg, v)
	return u
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateMsg() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldMsg)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(logmsg.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LogMsgUpsertOne) UpdateNewValues() *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(logmsg.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LogMsgUpsertOne) Ignore() *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LogMsgUpsertOne) DoNothing() *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LogMsgCreate.OnConflict
// documentation for more info.
func (u *LogMsgUpsertOne) Update(set func(*LogMsgUpsert)) *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LogMsgUpsert{UpdateSet: update})
	}))
	return u
}

// SetT sets the "t" field.
func (u *LogMsgUpsertOne) SetT(v time.Time) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetT(v)
	})
}

// UpdateT sets the "t" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateT() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateT()
	})
}

// SetMsg sets the "msg" field.
func (u *LogMsgUpsertOne) SetMsg(v string) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetMsg(v)
	})
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateMsg() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateMsg()
	})
}

// Exec executes the query.
func (u *LogMsgUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LogMsgCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LogMsgUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LogMsgUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LogMsgUpsertOne.ID is not supported by MySQL driver. Use LogMsgUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LogMsgUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LogMsgCreateBulk is the builder for creating many LogMsg entities in bulk.
type LogMsgCreateBulk struct {
	config
	builders []*LogMsgCreate
	conflict []sql.ConflictOption
}

// Save creates the LogMsg entities in the database.
func (lmcb *LogMsgCreateBulk) Save(ctx context.Context) ([]*LogMsg, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lmcb.builders))
	nodes := make([]*LogMsg, len(lmcb.builders))
	mutators := make([]Mutator, len(lmcb.builders))
	for i := range lmcb.builders {
		func(i int, root context.Context) {
			builder := lmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LogMsgMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lmcb *LogMsgCreateBulk) SaveX(ctx context.Context) []*LogMsg {
	v, err := lmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lmcb *LogMsgCreateBulk) Exec(ctx context.Context) error {
	_, err := lmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmcb *LogMsgCreateBulk) ExecX(ctx context.Context) {
	if err := lmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LogMsg.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LogMsgUpsert) {
//			SetT(v+v).
//		}).
//		Exec(ctx)
func (lmcb *LogMsgCreateBulk) OnConflict(opts ...sql.ConflictOption) *LogMsgUpsertBulk {
	lmcb.conflict = opts
	return &LogMsgUpsertBulk{
		create: lmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lmcb *LogMsgCreateBulk) OnConflictColumns(columns ...string) *LogMsgUpsertBulk {
	lmcb.conflict = append(lmcb.conflict, sql.ConflictColumns(columns...))
	return &LogMsgUpsertBulk{
		create: lmcb,
	}
}

// LogMsgUpsertBulk is the builder for "upsert"-ing
// a bulk of LogMsg nodes.
type LogMsgUpsertBulk struct {
	create *LogMsgCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(logmsg.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LogMsgUpsertBulk) UpdateNewValues() *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(logmsg.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LogMsgUpsertBulk) Ignore() *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LogMsgUpsertBulk) DoNothing() *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LogMsgCreateBulk.OnConflict
// documentation for more info.
func (u *LogMsgUpsertBulk) Update(set func(*LogMsgUpsert)) *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LogMsgUpsert{UpdateSet: update})
	}))
	return u
}

// SetT sets the "t" field.
func (u *LogMsgUpsertBulk) SetT(v time.Time) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetT(v)
	})
}

// UpdateT sets the "t" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateT() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateT()
	})
}

// SetMsg sets the "msg" field.
func (u *LogMsgUpsertBulk) SetMsg(v string) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetMsg(v)
	})
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateMsg() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateMsg()
	})
}

// Exec executes the query.
func (u *LogMsgUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LogMsgCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LogMsgCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LogMsgUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
