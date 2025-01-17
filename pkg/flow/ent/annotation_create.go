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
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// AnnotationCreate is the builder for creating a Annotation entity.
type AnnotationCreate struct {
	config
	mutation *AnnotationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ac *AnnotationCreate) SetName(s string) *AnnotationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AnnotationCreate) SetCreatedAt(t time.Time) *AnnotationCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AnnotationCreate) SetNillableCreatedAt(t *time.Time) *AnnotationCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AnnotationCreate) SetUpdatedAt(t time.Time) *AnnotationCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AnnotationCreate) SetNillableUpdatedAt(t *time.Time) *AnnotationCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetSize sets the "size" field.
func (ac *AnnotationCreate) SetSize(i int) *AnnotationCreate {
	ac.mutation.SetSize(i)
	return ac
}

// SetHash sets the "hash" field.
func (ac *AnnotationCreate) SetHash(s string) *AnnotationCreate {
	ac.mutation.SetHash(s)
	return ac
}

// SetData sets the "data" field.
func (ac *AnnotationCreate) SetData(b []byte) *AnnotationCreate {
	ac.mutation.SetData(b)
	return ac
}

// SetMimeType sets the "mime_type" field.
func (ac *AnnotationCreate) SetMimeType(s string) *AnnotationCreate {
	ac.mutation.SetMimeType(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AnnotationCreate) SetID(u uuid.UUID) *AnnotationCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AnnotationCreate) SetNillableID(u *uuid.UUID) *AnnotationCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (ac *AnnotationCreate) SetNamespaceID(id uuid.UUID) *AnnotationCreate {
	ac.mutation.SetNamespaceID(id)
	return ac
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (ac *AnnotationCreate) SetNillableNamespaceID(id *uuid.UUID) *AnnotationCreate {
	if id != nil {
		ac = ac.SetNamespaceID(*id)
	}
	return ac
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (ac *AnnotationCreate) SetNamespace(n *Namespace) *AnnotationCreate {
	return ac.SetNamespaceID(n.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (ac *AnnotationCreate) SetInstanceID(id uuid.UUID) *AnnotationCreate {
	ac.mutation.SetInstanceID(id)
	return ac
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (ac *AnnotationCreate) SetNillableInstanceID(id *uuid.UUID) *AnnotationCreate {
	if id != nil {
		ac = ac.SetInstanceID(*id)
	}
	return ac
}

// SetInstance sets the "instance" edge to the Instance entity.
func (ac *AnnotationCreate) SetInstance(i *Instance) *AnnotationCreate {
	return ac.SetInstanceID(i.ID)
}

// Mutation returns the AnnotationMutation object of the builder.
func (ac *AnnotationCreate) Mutation() *AnnotationMutation {
	return ac.mutation
}

// Save creates the Annotation in the database.
func (ac *AnnotationCreate) Save(ctx context.Context) (*Annotation, error) {
	ac.defaults()
	return withHooks[*Annotation, AnnotationMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnnotationCreate) SaveX(ctx context.Context) *Annotation {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnnotationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnnotationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnnotationCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := annotation.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := annotation.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := annotation.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnnotationCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Annotation.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := annotation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Annotation.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Annotation.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Annotation.updated_at"`)}
	}
	if _, ok := ac.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Annotation.size"`)}
	}
	if _, ok := ac.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Annotation.hash"`)}
	}
	if _, ok := ac.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Annotation.data"`)}
	}
	if _, ok := ac.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mime_type", err: errors.New(`ent: missing required field "Annotation.mime_type"`)}
	}
	return nil
}

func (ac *AnnotationCreate) sqlSave(ctx context.Context) (*Annotation, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AnnotationCreate) createSpec() (*Annotation, *sqlgraph.CreateSpec) {
	var (
		_node = &Annotation{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(annotation.Table, sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(annotation.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(annotation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(annotation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Size(); ok {
		_spec.SetField(annotation.FieldSize, field.TypeInt, value)
		_node.Size = value
	}
	if value, ok := ac.mutation.Hash(); ok {
		_spec.SetField(annotation.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := ac.mutation.Data(); ok {
		_spec.SetField(annotation.FieldData, field.TypeBytes, value)
		_node.Data = value
	}
	if value, ok := ac.mutation.MimeType(); ok {
		_spec.SetField(annotation.FieldMimeType, field.TypeString, value)
		_node.MimeType = value
	}
	if nodes := ac.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.NamespaceTable,
			Columns: []string{annotation.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_annotations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InstanceTable,
			Columns: []string{annotation.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(instance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_annotations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Annotation.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnnotationUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ac *AnnotationCreate) OnConflict(opts ...sql.ConflictOption) *AnnotationUpsertOne {
	ac.conflict = opts
	return &AnnotationUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AnnotationCreate) OnConflictColumns(columns ...string) *AnnotationUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AnnotationUpsertOne{
		create: ac,
	}
}

type (
	// AnnotationUpsertOne is the builder for "upsert"-ing
	//  one Annotation node.
	AnnotationUpsertOne struct {
		create *AnnotationCreate
	}

	// AnnotationUpsert is the "OnConflict" setter.
	AnnotationUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *AnnotationUpsert) SetName(v string) *AnnotationUpsert {
	u.Set(annotation.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateName() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldName)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AnnotationUpsert) SetUpdatedAt(v time.Time) *AnnotationUpsert {
	u.Set(annotation.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateUpdatedAt() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldUpdatedAt)
	return u
}

// SetSize sets the "size" field.
func (u *AnnotationUpsert) SetSize(v int) *AnnotationUpsert {
	u.Set(annotation.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateSize() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *AnnotationUpsert) AddSize(v int) *AnnotationUpsert {
	u.Add(annotation.FieldSize, v)
	return u
}

// SetHash sets the "hash" field.
func (u *AnnotationUpsert) SetHash(v string) *AnnotationUpsert {
	u.Set(annotation.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateHash() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldHash)
	return u
}

// SetData sets the "data" field.
func (u *AnnotationUpsert) SetData(v []byte) *AnnotationUpsert {
	u.Set(annotation.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateData() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldData)
	return u
}

// SetMimeType sets the "mime_type" field.
func (u *AnnotationUpsert) SetMimeType(v string) *AnnotationUpsert {
	u.Set(annotation.FieldMimeType, v)
	return u
}

// UpdateMimeType sets the "mime_type" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateMimeType() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldMimeType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(annotation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AnnotationUpsertOne) UpdateNewValues() *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(annotation.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(annotation.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Annotation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AnnotationUpsertOne) Ignore() *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnnotationUpsertOne) DoNothing() *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnnotationCreate.OnConflict
// documentation for more info.
func (u *AnnotationUpsertOne) Update(set func(*AnnotationUpsert)) *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnnotationUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AnnotationUpsertOne) SetName(v string) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateName() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateName()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AnnotationUpsertOne) SetUpdatedAt(v time.Time) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateUpdatedAt() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSize sets the "size" field.
func (u *AnnotationUpsertOne) SetSize(v int) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *AnnotationUpsertOne) AddSize(v int) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateSize() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateSize()
	})
}

// SetHash sets the "hash" field.
func (u *AnnotationUpsertOne) SetHash(v string) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateHash() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateHash()
	})
}

// SetData sets the "data" field.
func (u *AnnotationUpsertOne) SetData(v []byte) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateData() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateData()
	})
}

// SetMimeType sets the "mime_type" field.
func (u *AnnotationUpsertOne) SetMimeType(v string) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetMimeType(v)
	})
}

// UpdateMimeType sets the "mime_type" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateMimeType() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateMimeType()
	})
}

// Exec executes the query.
func (u *AnnotationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AnnotationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnnotationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AnnotationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AnnotationUpsertOne.ID is not supported by MySQL driver. Use AnnotationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AnnotationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AnnotationCreateBulk is the builder for creating many Annotation entities in bulk.
type AnnotationCreateBulk struct {
	config
	builders []*AnnotationCreate
	conflict []sql.ConflictOption
}

// Save creates the Annotation entities in the database.
func (acb *AnnotationCreateBulk) Save(ctx context.Context) ([]*Annotation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Annotation, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnnotationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnnotationCreateBulk) SaveX(ctx context.Context) []*Annotation {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnnotationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnnotationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Annotation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnnotationUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (acb *AnnotationCreateBulk) OnConflict(opts ...sql.ConflictOption) *AnnotationUpsertBulk {
	acb.conflict = opts
	return &AnnotationUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AnnotationCreateBulk) OnConflictColumns(columns ...string) *AnnotationUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AnnotationUpsertBulk{
		create: acb,
	}
}

// AnnotationUpsertBulk is the builder for "upsert"-ing
// a bulk of Annotation nodes.
type AnnotationUpsertBulk struct {
	create *AnnotationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(annotation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AnnotationUpsertBulk) UpdateNewValues() *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(annotation.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(annotation.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AnnotationUpsertBulk) Ignore() *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnnotationUpsertBulk) DoNothing() *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnnotationCreateBulk.OnConflict
// documentation for more info.
func (u *AnnotationUpsertBulk) Update(set func(*AnnotationUpsert)) *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnnotationUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AnnotationUpsertBulk) SetName(v string) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateName() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateName()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AnnotationUpsertBulk) SetUpdatedAt(v time.Time) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateUpdatedAt() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSize sets the "size" field.
func (u *AnnotationUpsertBulk) SetSize(v int) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *AnnotationUpsertBulk) AddSize(v int) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateSize() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateSize()
	})
}

// SetHash sets the "hash" field.
func (u *AnnotationUpsertBulk) SetHash(v string) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateHash() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateHash()
	})
}

// SetData sets the "data" field.
func (u *AnnotationUpsertBulk) SetData(v []byte) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateData() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateData()
	})
}

// SetMimeType sets the "mime_type" field.
func (u *AnnotationUpsertBulk) SetMimeType(v string) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetMimeType(v)
	})
}

// UpdateMimeType sets the "mime_type" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateMimeType() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateMimeType()
	})
}

// Exec executes the query.
func (u *AnnotationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AnnotationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AnnotationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnnotationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
