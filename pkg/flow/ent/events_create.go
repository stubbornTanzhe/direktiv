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
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// EventsCreate is the builder for creating a Events entity.
type EventsCreate struct {
	config
	mutation *EventsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEvents sets the "events" field.
func (ec *EventsCreate) SetEvents(m []map[string]interface{}) *EventsCreate {
	ec.mutation.SetEvents(m)
	return ec
}

// SetCorrelations sets the "correlations" field.
func (ec *EventsCreate) SetCorrelations(s []string) *EventsCreate {
	ec.mutation.SetCorrelations(s)
	return ec
}

// SetSignature sets the "signature" field.
func (ec *EventsCreate) SetSignature(b []byte) *EventsCreate {
	ec.mutation.SetSignature(b)
	return ec
}

// SetCount sets the "count" field.
func (ec *EventsCreate) SetCount(i int) *EventsCreate {
	ec.mutation.SetCount(i)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventsCreate) SetCreatedAt(t time.Time) *EventsCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventsCreate) SetNillableCreatedAt(t *time.Time) *EventsCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EventsCreate) SetUpdatedAt(t time.Time) *EventsCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EventsCreate) SetNillableUpdatedAt(t *time.Time) *EventsCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetWorkflowID sets the "workflow_id" field.
func (ec *EventsCreate) SetWorkflowID(u uuid.UUID) *EventsCreate {
	ec.mutation.SetWorkflowID(u)
	return ec
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (ec *EventsCreate) SetNillableWorkflowID(u *uuid.UUID) *EventsCreate {
	if u != nil {
		ec.SetWorkflowID(*u)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventsCreate) SetID(u uuid.UUID) *EventsCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EventsCreate) SetNillableID(u *uuid.UUID) *EventsCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// AddWfeventswaitIDs adds the "wfeventswait" edge to the EventsWait entity by IDs.
func (ec *EventsCreate) AddWfeventswaitIDs(ids ...uuid.UUID) *EventsCreate {
	ec.mutation.AddWfeventswaitIDs(ids...)
	return ec
}

// AddWfeventswait adds the "wfeventswait" edges to the EventsWait entity.
func (ec *EventsCreate) AddWfeventswait(e ...*EventsWait) *EventsCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddWfeventswaitIDs(ids...)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (ec *EventsCreate) SetInstanceID(id uuid.UUID) *EventsCreate {
	ec.mutation.SetInstanceID(id)
	return ec
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (ec *EventsCreate) SetNillableInstanceID(id *uuid.UUID) *EventsCreate {
	if id != nil {
		ec = ec.SetInstanceID(*id)
	}
	return ec
}

// SetInstance sets the "instance" edge to the Instance entity.
func (ec *EventsCreate) SetInstance(i *Instance) *EventsCreate {
	return ec.SetInstanceID(i.ID)
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (ec *EventsCreate) SetNamespaceID(id uuid.UUID) *EventsCreate {
	ec.mutation.SetNamespaceID(id)
	return ec
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (ec *EventsCreate) SetNamespace(n *Namespace) *EventsCreate {
	return ec.SetNamespaceID(n.ID)
}

// Mutation returns the EventsMutation object of the builder.
func (ec *EventsCreate) Mutation() *EventsMutation {
	return ec.mutation
}

// Save creates the Events in the database.
func (ec *EventsCreate) Save(ctx context.Context) (*Events, error) {
	ec.defaults()
	return withHooks[*Events, EventsMutation](ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventsCreate) SaveX(ctx context.Context) *Events {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventsCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventsCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventsCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := events.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := events.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := events.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventsCreate) check() error {
	if _, ok := ec.mutation.Events(); !ok {
		return &ValidationError{Name: "events", err: errors.New(`ent: missing required field "Events.events"`)}
	}
	if _, ok := ec.mutation.Correlations(); !ok {
		return &ValidationError{Name: "correlations", err: errors.New(`ent: missing required field "Events.correlations"`)}
	}
	if _, ok := ec.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`ent: missing required field "Events.count"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Events.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Events.updated_at"`)}
	}
	if _, ok := ec.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Events.namespace"`)}
	}
	return nil
}

func (ec *EventsCreate) sqlSave(ctx context.Context) (*Events, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EventsCreate) createSpec() (*Events, *sqlgraph.CreateSpec) {
	var (
		_node = &Events{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(events.Table, sqlgraph.NewFieldSpec(events.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Events(); ok {
		_spec.SetField(events.FieldEvents, field.TypeJSON, value)
		_node.Events = value
	}
	if value, ok := ec.mutation.Correlations(); ok {
		_spec.SetField(events.FieldCorrelations, field.TypeJSON, value)
		_node.Correlations = value
	}
	if value, ok := ec.mutation.Signature(); ok {
		_spec.SetField(events.FieldSignature, field.TypeBytes, value)
		_node.Signature = value
	}
	if value, ok := ec.mutation.Count(); ok {
		_spec.SetField(events.FieldCount, field.TypeInt, value)
		_node.Count = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(events.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(events.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.WorkflowID(); ok {
		_spec.SetField(events.FieldWorkflowID, field.TypeUUID, value)
		_node.WorkflowID = value
	}
	if nodes := ec.mutation.WfeventswaitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventswait.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.InstanceTable,
			Columns: []string{events.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_eventlisteners = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.NamespaceTable,
			Columns: []string{events.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_namespacelisteners = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Events.Create().
//		SetEvents(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventsUpsert) {
//			SetEvents(v+v).
//		}).
//		Exec(ctx)
func (ec *EventsCreate) OnConflict(opts ...sql.ConflictOption) *EventsUpsertOne {
	ec.conflict = opts
	return &EventsUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Events.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EventsCreate) OnConflictColumns(columns ...string) *EventsUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EventsUpsertOne{
		create: ec,
	}
}

type (
	// EventsUpsertOne is the builder for "upsert"-ing
	//  one Events node.
	EventsUpsertOne struct {
		create *EventsCreate
	}

	// EventsUpsert is the "OnConflict" setter.
	EventsUpsert struct {
		*sql.UpdateSet
	}
)

// SetEvents sets the "events" field.
func (u *EventsUpsert) SetEvents(v []map[string]interface{}) *EventsUpsert {
	u.Set(events.FieldEvents, v)
	return u
}

// UpdateEvents sets the "events" field to the value that was provided on create.
func (u *EventsUpsert) UpdateEvents() *EventsUpsert {
	u.SetExcluded(events.FieldEvents)
	return u
}

// SetCorrelations sets the "correlations" field.
func (u *EventsUpsert) SetCorrelations(v []string) *EventsUpsert {
	u.Set(events.FieldCorrelations, v)
	return u
}

// UpdateCorrelations sets the "correlations" field to the value that was provided on create.
func (u *EventsUpsert) UpdateCorrelations() *EventsUpsert {
	u.SetExcluded(events.FieldCorrelations)
	return u
}

// SetSignature sets the "signature" field.
func (u *EventsUpsert) SetSignature(v []byte) *EventsUpsert {
	u.Set(events.FieldSignature, v)
	return u
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventsUpsert) UpdateSignature() *EventsUpsert {
	u.SetExcluded(events.FieldSignature)
	return u
}

// ClearSignature clears the value of the "signature" field.
func (u *EventsUpsert) ClearSignature() *EventsUpsert {
	u.SetNull(events.FieldSignature)
	return u
}

// SetCount sets the "count" field.
func (u *EventsUpsert) SetCount(v int) *EventsUpsert {
	u.Set(events.FieldCount, v)
	return u
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *EventsUpsert) UpdateCount() *EventsUpsert {
	u.SetExcluded(events.FieldCount)
	return u
}

// AddCount adds v to the "count" field.
func (u *EventsUpsert) AddCount(v int) *EventsUpsert {
	u.Add(events.FieldCount, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventsUpsert) SetUpdatedAt(v time.Time) *EventsUpsert {
	u.Set(events.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventsUpsert) UpdateUpdatedAt() *EventsUpsert {
	u.SetExcluded(events.FieldUpdatedAt)
	return u
}

// SetWorkflowID sets the "workflow_id" field.
func (u *EventsUpsert) SetWorkflowID(v uuid.UUID) *EventsUpsert {
	u.Set(events.FieldWorkflowID, v)
	return u
}

// UpdateWorkflowID sets the "workflow_id" field to the value that was provided on create.
func (u *EventsUpsert) UpdateWorkflowID() *EventsUpsert {
	u.SetExcluded(events.FieldWorkflowID)
	return u
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (u *EventsUpsert) ClearWorkflowID() *EventsUpsert {
	u.SetNull(events.FieldWorkflowID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Events.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(events.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EventsUpsertOne) UpdateNewValues() *EventsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(events.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(events.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Events.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EventsUpsertOne) Ignore() *EventsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventsUpsertOne) DoNothing() *EventsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventsCreate.OnConflict
// documentation for more info.
func (u *EventsUpsertOne) Update(set func(*EventsUpsert)) *EventsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventsUpsert{UpdateSet: update})
	}))
	return u
}

// SetEvents sets the "events" field.
func (u *EventsUpsertOne) SetEvents(v []map[string]interface{}) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.SetEvents(v)
	})
}

// UpdateEvents sets the "events" field to the value that was provided on create.
func (u *EventsUpsertOne) UpdateEvents() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateEvents()
	})
}

// SetCorrelations sets the "correlations" field.
func (u *EventsUpsertOne) SetCorrelations(v []string) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.SetCorrelations(v)
	})
}

// UpdateCorrelations sets the "correlations" field to the value that was provided on create.
func (u *EventsUpsertOne) UpdateCorrelations() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateCorrelations()
	})
}

// SetSignature sets the "signature" field.
func (u *EventsUpsertOne) SetSignature(v []byte) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventsUpsertOne) UpdateSignature() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateSignature()
	})
}

// ClearSignature clears the value of the "signature" field.
func (u *EventsUpsertOne) ClearSignature() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.ClearSignature()
	})
}

// SetCount sets the "count" field.
func (u *EventsUpsertOne) SetCount(v int) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *EventsUpsertOne) AddCount(v int) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *EventsUpsertOne) UpdateCount() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateCount()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventsUpsertOne) SetUpdatedAt(v time.Time) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventsUpsertOne) UpdateUpdatedAt() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetWorkflowID sets the "workflow_id" field.
func (u *EventsUpsertOne) SetWorkflowID(v uuid.UUID) *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.SetWorkflowID(v)
	})
}

// UpdateWorkflowID sets the "workflow_id" field to the value that was provided on create.
func (u *EventsUpsertOne) UpdateWorkflowID() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateWorkflowID()
	})
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (u *EventsUpsertOne) ClearWorkflowID() *EventsUpsertOne {
	return u.Update(func(s *EventsUpsert) {
		s.ClearWorkflowID()
	})
}

// Exec executes the query.
func (u *EventsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventsUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EventsUpsertOne.ID is not supported by MySQL driver. Use EventsUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventsUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventsCreateBulk is the builder for creating many Events entities in bulk.
type EventsCreateBulk struct {
	config
	builders []*EventsCreate
	conflict []sql.ConflictOption
}

// Save creates the Events entities in the database.
func (ecb *EventsCreateBulk) Save(ctx context.Context) ([]*Events, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Events, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventsMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventsCreateBulk) SaveX(ctx context.Context) []*Events {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventsCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventsCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Events.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventsUpsert) {
//			SetEvents(v+v).
//		}).
//		Exec(ctx)
func (ecb *EventsCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventsUpsertBulk {
	ecb.conflict = opts
	return &EventsUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Events.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EventsCreateBulk) OnConflictColumns(columns ...string) *EventsUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EventsUpsertBulk{
		create: ecb,
	}
}

// EventsUpsertBulk is the builder for "upsert"-ing
// a bulk of Events nodes.
type EventsUpsertBulk struct {
	create *EventsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Events.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(events.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EventsUpsertBulk) UpdateNewValues() *EventsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(events.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(events.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Events.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EventsUpsertBulk) Ignore() *EventsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventsUpsertBulk) DoNothing() *EventsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventsCreateBulk.OnConflict
// documentation for more info.
func (u *EventsUpsertBulk) Update(set func(*EventsUpsert)) *EventsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventsUpsert{UpdateSet: update})
	}))
	return u
}

// SetEvents sets the "events" field.
func (u *EventsUpsertBulk) SetEvents(v []map[string]interface{}) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.SetEvents(v)
	})
}

// UpdateEvents sets the "events" field to the value that was provided on create.
func (u *EventsUpsertBulk) UpdateEvents() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateEvents()
	})
}

// SetCorrelations sets the "correlations" field.
func (u *EventsUpsertBulk) SetCorrelations(v []string) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.SetCorrelations(v)
	})
}

// UpdateCorrelations sets the "correlations" field to the value that was provided on create.
func (u *EventsUpsertBulk) UpdateCorrelations() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateCorrelations()
	})
}

// SetSignature sets the "signature" field.
func (u *EventsUpsertBulk) SetSignature(v []byte) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventsUpsertBulk) UpdateSignature() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateSignature()
	})
}

// ClearSignature clears the value of the "signature" field.
func (u *EventsUpsertBulk) ClearSignature() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.ClearSignature()
	})
}

// SetCount sets the "count" field.
func (u *EventsUpsertBulk) SetCount(v int) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *EventsUpsertBulk) AddCount(v int) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *EventsUpsertBulk) UpdateCount() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateCount()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventsUpsertBulk) SetUpdatedAt(v time.Time) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventsUpsertBulk) UpdateUpdatedAt() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetWorkflowID sets the "workflow_id" field.
func (u *EventsUpsertBulk) SetWorkflowID(v uuid.UUID) *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.SetWorkflowID(v)
	})
}

// UpdateWorkflowID sets the "workflow_id" field to the value that was provided on create.
func (u *EventsUpsertBulk) UpdateWorkflowID() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.UpdateWorkflowID()
	})
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (u *EventsUpsertBulk) ClearWorkflowID() *EventsUpsertBulk {
	return u.Update(func(s *EventsUpsert) {
		s.ClearWorkflowID()
	})
}

// Exec executes the query.
func (u *EventsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
