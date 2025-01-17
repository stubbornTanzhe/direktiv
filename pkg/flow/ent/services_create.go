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
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/services"
	"github.com/google/uuid"
)

// ServicesCreate is the builder for creating a Services entity.
type ServicesCreate struct {
	config
	mutation *ServicesMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *ServicesCreate) SetCreatedAt(t time.Time) *ServicesCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ServicesCreate) SetNillableCreatedAt(t *time.Time) *ServicesCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ServicesCreate) SetUpdatedAt(t time.Time) *ServicesCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ServicesCreate) SetNillableUpdatedAt(t *time.Time) *ServicesCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetURL sets the "url" field.
func (sc *ServicesCreate) SetURL(s string) *ServicesCreate {
	sc.mutation.SetURL(s)
	return sc
}

// SetName sets the "name" field.
func (sc *ServicesCreate) SetName(s string) *ServicesCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetData sets the "data" field.
func (sc *ServicesCreate) SetData(s string) *ServicesCreate {
	sc.mutation.SetData(s)
	return sc
}

// SetID sets the "id" field.
func (sc *ServicesCreate) SetID(u uuid.UUID) *ServicesCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *ServicesCreate) SetNillableID(u *uuid.UUID) *ServicesCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (sc *ServicesCreate) SetNamespaceID(id uuid.UUID) *ServicesCreate {
	sc.mutation.SetNamespaceID(id)
	return sc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (sc *ServicesCreate) SetNamespace(n *Namespace) *ServicesCreate {
	return sc.SetNamespaceID(n.ID)
}

// Mutation returns the ServicesMutation object of the builder.
func (sc *ServicesCreate) Mutation() *ServicesMutation {
	return sc.mutation
}

// Save creates the Services in the database.
func (sc *ServicesCreate) Save(ctx context.Context) (*Services, error) {
	sc.defaults()
	return withHooks[*Services, ServicesMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServicesCreate) SaveX(ctx context.Context) *Services {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServicesCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServicesCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServicesCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := services.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := services.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := services.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServicesCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Services.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Services.updated_at"`)}
	}
	if _, ok := sc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Services.url"`)}
	}
	if v, ok := sc.mutation.URL(); ok {
		if err := services.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Services.url": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Services.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := services.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Services.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Services.data"`)}
	}
	if v, ok := sc.mutation.Data(); ok {
		if err := services.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "Services.data": %w`, err)}
		}
	}
	if _, ok := sc.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Services.namespace"`)}
	}
	return nil
}

func (sc *ServicesCreate) sqlSave(ctx context.Context) (*Services, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ServicesCreate) createSpec() (*Services, *sqlgraph.CreateSpec) {
	var (
		_node = &Services{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(services.Table, sqlgraph.NewFieldSpec(services.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(services.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(services.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.URL(); ok {
		_spec.SetField(services.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(services.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Data(); ok {
		_spec.SetField(services.FieldData, field.TypeString, value)
		_node.Data = value
	}
	if nodes := sc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   services.NamespaceTable,
			Columns: []string{services.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_services = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Services.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServicesUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sc *ServicesCreate) OnConflict(opts ...sql.ConflictOption) *ServicesUpsertOne {
	sc.conflict = opts
	return &ServicesUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Services.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *ServicesCreate) OnConflictColumns(columns ...string) *ServicesUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ServicesUpsertOne{
		create: sc,
	}
}

type (
	// ServicesUpsertOne is the builder for "upsert"-ing
	//  one Services node.
	ServicesUpsertOne struct {
		create *ServicesCreate
	}

	// ServicesUpsert is the "OnConflict" setter.
	ServicesUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ServicesUpsert) SetUpdatedAt(v time.Time) *ServicesUpsert {
	u.Set(services.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ServicesUpsert) UpdateUpdatedAt() *ServicesUpsert {
	u.SetExcluded(services.FieldUpdatedAt)
	return u
}

// SetURL sets the "url" field.
func (u *ServicesUpsert) SetURL(v string) *ServicesUpsert {
	u.Set(services.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ServicesUpsert) UpdateURL() *ServicesUpsert {
	u.SetExcluded(services.FieldURL)
	return u
}

// SetName sets the "name" field.
func (u *ServicesUpsert) SetName(v string) *ServicesUpsert {
	u.Set(services.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ServicesUpsert) UpdateName() *ServicesUpsert {
	u.SetExcluded(services.FieldName)
	return u
}

// SetData sets the "data" field.
func (u *ServicesUpsert) SetData(v string) *ServicesUpsert {
	u.Set(services.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *ServicesUpsert) UpdateData() *ServicesUpsert {
	u.SetExcluded(services.FieldData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Services.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(services.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServicesUpsertOne) UpdateNewValues() *ServicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(services.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(services.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Services.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ServicesUpsertOne) Ignore() *ServicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServicesUpsertOne) DoNothing() *ServicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServicesCreate.OnConflict
// documentation for more info.
func (u *ServicesUpsertOne) Update(set func(*ServicesUpsert)) *ServicesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServicesUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ServicesUpsertOne) SetUpdatedAt(v time.Time) *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ServicesUpsertOne) UpdateUpdatedAt() *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetURL sets the "url" field.
func (u *ServicesUpsertOne) SetURL(v string) *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ServicesUpsertOne) UpdateURL() *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateURL()
	})
}

// SetName sets the "name" field.
func (u *ServicesUpsertOne) SetName(v string) *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ServicesUpsertOne) UpdateName() *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateName()
	})
}

// SetData sets the "data" field.
func (u *ServicesUpsertOne) SetData(v string) *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *ServicesUpsertOne) UpdateData() *ServicesUpsertOne {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *ServicesUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ServicesCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServicesUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ServicesUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ServicesUpsertOne.ID is not supported by MySQL driver. Use ServicesUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ServicesUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ServicesCreateBulk is the builder for creating many Services entities in bulk.
type ServicesCreateBulk struct {
	config
	builders []*ServicesCreate
	conflict []sql.ConflictOption
}

// Save creates the Services entities in the database.
func (scb *ServicesCreateBulk) Save(ctx context.Context) ([]*Services, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Services, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServicesMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServicesCreateBulk) SaveX(ctx context.Context) []*Services {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServicesCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServicesCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Services.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServicesUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (scb *ServicesCreateBulk) OnConflict(opts ...sql.ConflictOption) *ServicesUpsertBulk {
	scb.conflict = opts
	return &ServicesUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Services.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *ServicesCreateBulk) OnConflictColumns(columns ...string) *ServicesUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ServicesUpsertBulk{
		create: scb,
	}
}

// ServicesUpsertBulk is the builder for "upsert"-ing
// a bulk of Services nodes.
type ServicesUpsertBulk struct {
	create *ServicesCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Services.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(services.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ServicesUpsertBulk) UpdateNewValues() *ServicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(services.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(services.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Services.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ServicesUpsertBulk) Ignore() *ServicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServicesUpsertBulk) DoNothing() *ServicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServicesCreateBulk.OnConflict
// documentation for more info.
func (u *ServicesUpsertBulk) Update(set func(*ServicesUpsert)) *ServicesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServicesUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ServicesUpsertBulk) SetUpdatedAt(v time.Time) *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ServicesUpsertBulk) UpdateUpdatedAt() *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetURL sets the "url" field.
func (u *ServicesUpsertBulk) SetURL(v string) *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ServicesUpsertBulk) UpdateURL() *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateURL()
	})
}

// SetName sets the "name" field.
func (u *ServicesUpsertBulk) SetName(v string) *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ServicesUpsertBulk) UpdateName() *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateName()
	})
}

// SetData sets the "data" field.
func (u *ServicesUpsertBulk) SetData(v string) *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *ServicesUpsertBulk) UpdateData() *ServicesUpsertBulk {
	return u.Update(func(s *ServicesUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *ServicesUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ServicesCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ServicesCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServicesUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
