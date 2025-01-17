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
	"github.com/direktiv/direktiv/pkg/flow/ent/annotation"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudeventfilters"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/services"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/google/uuid"
)

// NamespaceCreate is the builder for creating a Namespace entity.
type NamespaceCreate struct {
	config
	mutation *NamespaceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (nc *NamespaceCreate) SetCreatedAt(t time.Time) *NamespaceCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableCreatedAt(t *time.Time) *NamespaceCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NamespaceCreate) SetUpdatedAt(t time.Time) *NamespaceCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableUpdatedAt(t *time.Time) *NamespaceCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetConfig sets the "config" field.
func (nc *NamespaceCreate) SetConfig(s string) *NamespaceCreate {
	nc.mutation.SetConfig(s)
	return nc
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableConfig(s *string) *NamespaceCreate {
	if s != nil {
		nc.SetConfig(*s)
	}
	return nc
}

// SetName sets the "name" field.
func (nc *NamespaceCreate) SetName(s string) *NamespaceCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetID sets the "id" field.
func (nc *NamespaceCreate) SetID(u uuid.UUID) *NamespaceCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableID(u *uuid.UUID) *NamespaceCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (nc *NamespaceCreate) AddInstanceIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddInstanceIDs(ids...)
	return nc
}

// AddInstances adds the "instances" edges to the Instance entity.
func (nc *NamespaceCreate) AddInstances(i ...*Instance) *NamespaceCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nc.AddInstanceIDs(ids...)
}

// AddLogIDs adds the "logs" edge to the LogMsg entity by IDs.
func (nc *NamespaceCreate) AddLogIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddLogIDs(ids...)
	return nc
}

// AddLogs adds the "logs" edges to the LogMsg entity.
func (nc *NamespaceCreate) AddLogs(l ...*LogMsg) *NamespaceCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return nc.AddLogIDs(ids...)
}

// AddVarIDs adds the "vars" edge to the VarRef entity by IDs.
func (nc *NamespaceCreate) AddVarIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddVarIDs(ids...)
	return nc
}

// AddVars adds the "vars" edges to the VarRef entity.
func (nc *NamespaceCreate) AddVars(v ...*VarRef) *NamespaceCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return nc.AddVarIDs(ids...)
}

// AddCloudeventIDs adds the "cloudevents" edge to the CloudEvents entity by IDs.
func (nc *NamespaceCreate) AddCloudeventIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddCloudeventIDs(ids...)
	return nc
}

// AddCloudevents adds the "cloudevents" edges to the CloudEvents entity.
func (nc *NamespaceCreate) AddCloudevents(c ...*CloudEvents) *NamespaceCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nc.AddCloudeventIDs(ids...)
}

// AddNamespacelistenerIDs adds the "namespacelisteners" edge to the Events entity by IDs.
func (nc *NamespaceCreate) AddNamespacelistenerIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddNamespacelistenerIDs(ids...)
	return nc
}

// AddNamespacelisteners adds the "namespacelisteners" edges to the Events entity.
func (nc *NamespaceCreate) AddNamespacelisteners(e ...*Events) *NamespaceCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nc.AddNamespacelistenerIDs(ids...)
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (nc *NamespaceCreate) AddAnnotationIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddAnnotationIDs(ids...)
	return nc
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (nc *NamespaceCreate) AddAnnotations(a ...*Annotation) *NamespaceCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nc.AddAnnotationIDs(ids...)
}

// AddCloudeventfilterIDs adds the "cloudeventfilters" edge to the CloudEventFilters entity by IDs.
func (nc *NamespaceCreate) AddCloudeventfilterIDs(ids ...int) *NamespaceCreate {
	nc.mutation.AddCloudeventfilterIDs(ids...)
	return nc
}

// AddCloudeventfilters adds the "cloudeventfilters" edges to the CloudEventFilters entity.
func (nc *NamespaceCreate) AddCloudeventfilters(c ...*CloudEventFilters) *NamespaceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nc.AddCloudeventfilterIDs(ids...)
}

// AddServiceIDs adds the "services" edge to the Services entity by IDs.
func (nc *NamespaceCreate) AddServiceIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddServiceIDs(ids...)
	return nc
}

// AddServices adds the "services" edges to the Services entity.
func (nc *NamespaceCreate) AddServices(s ...*Services) *NamespaceCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nc.AddServiceIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nc *NamespaceCreate) Mutation() *NamespaceMutation {
	return nc.mutation
}

// Save creates the Namespace in the database.
func (nc *NamespaceCreate) Save(ctx context.Context) (*Namespace, error) {
	nc.defaults()
	return withHooks[*Namespace, NamespaceMutation](ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NamespaceCreate) SaveX(ctx context.Context) *Namespace {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NamespaceCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NamespaceCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NamespaceCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := namespace.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := namespace.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.Config(); !ok {
		v := namespace.DefaultConfig
		nc.mutation.SetConfig(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := namespace.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NamespaceCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Namespace.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Namespace.updated_at"`)}
	}
	if _, ok := nc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "Namespace.config"`)}
	}
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Namespace.name"`)}
	}
	if v, ok := nc.mutation.Name(); ok {
		if err := namespace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Namespace.name": %w`, err)}
		}
	}
	return nil
}

func (nc *NamespaceCreate) sqlSave(ctx context.Context) (*Namespace, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NamespaceCreate) createSpec() (*Namespace, *sqlgraph.CreateSpec) {
	var (
		_node = &Namespace{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(namespace.Table, sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(namespace.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(namespace.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Config(); ok {
		_spec.SetField(namespace.FieldConfig, field.TypeString, value)
		_node.Config = value
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.SetField(namespace.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := nc.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.InstancesTable,
			Columns: []string{namespace.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.LogsTable,
			Columns: []string{namespace.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(logmsg.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.VarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.VarsTable,
			Columns: []string{namespace.VarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(varref.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.CloudeventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.CloudeventsTable,
			Columns: []string{namespace.CloudeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudevents.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NamespacelistenersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.NamespacelistenersTable,
			Columns: []string{namespace.NamespacelistenersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.AnnotationsTable,
			Columns: []string{namespace.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.CloudeventfiltersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.CloudeventfiltersTable,
			Columns: []string{namespace.CloudeventfiltersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudeventfilters.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ServicesTable,
			Columns: []string{namespace.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(services.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Namespace.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NamespaceUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (nc *NamespaceCreate) OnConflict(opts ...sql.ConflictOption) *NamespaceUpsertOne {
	nc.conflict = opts
	return &NamespaceUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Namespace.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nc *NamespaceCreate) OnConflictColumns(columns ...string) *NamespaceUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NamespaceUpsertOne{
		create: nc,
	}
}

type (
	// NamespaceUpsertOne is the builder for "upsert"-ing
	//  one Namespace node.
	NamespaceUpsertOne struct {
		create *NamespaceCreate
	}

	// NamespaceUpsert is the "OnConflict" setter.
	NamespaceUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *NamespaceUpsert) SetUpdatedAt(v time.Time) *NamespaceUpsert {
	u.Set(namespace.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NamespaceUpsert) UpdateUpdatedAt() *NamespaceUpsert {
	u.SetExcluded(namespace.FieldUpdatedAt)
	return u
}

// SetConfig sets the "config" field.
func (u *NamespaceUpsert) SetConfig(v string) *NamespaceUpsert {
	u.Set(namespace.FieldConfig, v)
	return u
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *NamespaceUpsert) UpdateConfig() *NamespaceUpsert {
	u.SetExcluded(namespace.FieldConfig)
	return u
}

// SetName sets the "name" field.
func (u *NamespaceUpsert) SetName(v string) *NamespaceUpsert {
	u.Set(namespace.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NamespaceUpsert) UpdateName() *NamespaceUpsert {
	u.SetExcluded(namespace.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Namespace.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(namespace.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NamespaceUpsertOne) UpdateNewValues() *NamespaceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(namespace.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(namespace.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Namespace.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NamespaceUpsertOne) Ignore() *NamespaceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NamespaceUpsertOne) DoNothing() *NamespaceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NamespaceCreate.OnConflict
// documentation for more info.
func (u *NamespaceUpsertOne) Update(set func(*NamespaceUpsert)) *NamespaceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NamespaceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NamespaceUpsertOne) SetUpdatedAt(v time.Time) *NamespaceUpsertOne {
	return u.Update(func(s *NamespaceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NamespaceUpsertOne) UpdateUpdatedAt() *NamespaceUpsertOne {
	return u.Update(func(s *NamespaceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetConfig sets the "config" field.
func (u *NamespaceUpsertOne) SetConfig(v string) *NamespaceUpsertOne {
	return u.Update(func(s *NamespaceUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *NamespaceUpsertOne) UpdateConfig() *NamespaceUpsertOne {
	return u.Update(func(s *NamespaceUpsert) {
		s.UpdateConfig()
	})
}

// SetName sets the "name" field.
func (u *NamespaceUpsertOne) SetName(v string) *NamespaceUpsertOne {
	return u.Update(func(s *NamespaceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NamespaceUpsertOne) UpdateName() *NamespaceUpsertOne {
	return u.Update(func(s *NamespaceUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *NamespaceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NamespaceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NamespaceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NamespaceUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NamespaceUpsertOne.ID is not supported by MySQL driver. Use NamespaceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NamespaceUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NamespaceCreateBulk is the builder for creating many Namespace entities in bulk.
type NamespaceCreateBulk struct {
	config
	builders []*NamespaceCreate
	conflict []sql.ConflictOption
}

// Save creates the Namespace entities in the database.
func (ncb *NamespaceCreateBulk) Save(ctx context.Context) ([]*Namespace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Namespace, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NamespaceMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NamespaceCreateBulk) SaveX(ctx context.Context) []*Namespace {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NamespaceCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NamespaceCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Namespace.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NamespaceUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ncb *NamespaceCreateBulk) OnConflict(opts ...sql.ConflictOption) *NamespaceUpsertBulk {
	ncb.conflict = opts
	return &NamespaceUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Namespace.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ncb *NamespaceCreateBulk) OnConflictColumns(columns ...string) *NamespaceUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NamespaceUpsertBulk{
		create: ncb,
	}
}

// NamespaceUpsertBulk is the builder for "upsert"-ing
// a bulk of Namespace nodes.
type NamespaceUpsertBulk struct {
	create *NamespaceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Namespace.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(namespace.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NamespaceUpsertBulk) UpdateNewValues() *NamespaceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(namespace.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(namespace.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Namespace.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NamespaceUpsertBulk) Ignore() *NamespaceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NamespaceUpsertBulk) DoNothing() *NamespaceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NamespaceCreateBulk.OnConflict
// documentation for more info.
func (u *NamespaceUpsertBulk) Update(set func(*NamespaceUpsert)) *NamespaceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NamespaceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NamespaceUpsertBulk) SetUpdatedAt(v time.Time) *NamespaceUpsertBulk {
	return u.Update(func(s *NamespaceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NamespaceUpsertBulk) UpdateUpdatedAt() *NamespaceUpsertBulk {
	return u.Update(func(s *NamespaceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetConfig sets the "config" field.
func (u *NamespaceUpsertBulk) SetConfig(v string) *NamespaceUpsertBulk {
	return u.Update(func(s *NamespaceUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *NamespaceUpsertBulk) UpdateConfig() *NamespaceUpsertBulk {
	return u.Update(func(s *NamespaceUpsert) {
		s.UpdateConfig()
	})
}

// SetName sets the "name" field.
func (u *NamespaceUpsertBulk) SetName(v string) *NamespaceUpsertBulk {
	return u.Update(func(s *NamespaceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NamespaceUpsertBulk) UpdateName() *NamespaceUpsertBulk {
	return u.Update(func(s *NamespaceUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *NamespaceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NamespaceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NamespaceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NamespaceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
