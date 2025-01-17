// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// CloudEventsDelete is the builder for deleting a CloudEvents entity.
type CloudEventsDelete struct {
	config
	hooks    []Hook
	mutation *CloudEventsMutation
}

// Where appends a list predicates to the CloudEventsDelete builder.
func (ced *CloudEventsDelete) Where(ps ...predicate.CloudEvents) *CloudEventsDelete {
	ced.mutation.Where(ps...)
	return ced
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ced *CloudEventsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CloudEventsMutation](ctx, ced.sqlExec, ced.mutation, ced.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ced *CloudEventsDelete) ExecX(ctx context.Context) int {
	n, err := ced.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ced *CloudEventsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cloudevents.Table, sqlgraph.NewFieldSpec(cloudevents.FieldID, field.TypeUUID))
	if ps := ced.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ced.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ced.mutation.done = true
	return affected, err
}

// CloudEventsDeleteOne is the builder for deleting a single CloudEvents entity.
type CloudEventsDeleteOne struct {
	ced *CloudEventsDelete
}

// Where appends a list predicates to the CloudEventsDelete builder.
func (cedo *CloudEventsDeleteOne) Where(ps ...predicate.CloudEvents) *CloudEventsDeleteOne {
	cedo.ced.mutation.Where(ps...)
	return cedo
}

// Exec executes the deletion query.
func (cedo *CloudEventsDeleteOne) Exec(ctx context.Context) error {
	n, err := cedo.ced.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cloudevents.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cedo *CloudEventsDeleteOne) ExecX(ctx context.Context) {
	if err := cedo.Exec(ctx); err != nil {
		panic(err)
	}
}
