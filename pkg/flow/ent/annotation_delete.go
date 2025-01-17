// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/annotation"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// AnnotationDelete is the builder for deleting a Annotation entity.
type AnnotationDelete struct {
	config
	hooks    []Hook
	mutation *AnnotationMutation
}

// Where appends a list predicates to the AnnotationDelete builder.
func (ad *AnnotationDelete) Where(ps ...predicate.Annotation) *AnnotationDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AnnotationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AnnotationMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AnnotationDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AnnotationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(annotation.Table, sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeUUID))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AnnotationDeleteOne is the builder for deleting a single Annotation entity.
type AnnotationDeleteOne struct {
	ad *AnnotationDelete
}

// Where appends a list predicates to the AnnotationDelete builder.
func (ado *AnnotationDeleteOne) Where(ps ...predicate.Annotation) *AnnotationDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AnnotationDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{annotation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AnnotationDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
