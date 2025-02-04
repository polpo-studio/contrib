// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"context"

	"entgo.io/contrib/entoas/internal/oastypes/oastypes"
	"entgo.io/contrib/entoas/internal/oastypes/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OASTypesDelete is the builder for deleting a OASTypes entity.
type OASTypesDelete struct {
	config
	hooks    []Hook
	mutation *OASTypesMutation
}

// Where appends a list predicates to the OASTypesDelete builder.
func (otd *OASTypesDelete) Where(ps ...predicate.OASTypes) *OASTypesDelete {
	otd.mutation.Where(ps...)
	return otd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (otd *OASTypesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OASTypesMutation](ctx, otd.sqlExec, otd.mutation, otd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (otd *OASTypesDelete) ExecX(ctx context.Context) int {
	n, err := otd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (otd *OASTypesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: oastypes.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oastypes.FieldID,
			},
		},
	}
	if ps := otd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, otd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	otd.mutation.done = true
	return affected, err
}

// OASTypesDeleteOne is the builder for deleting a single OASTypes entity.
type OASTypesDeleteOne struct {
	otd *OASTypesDelete
}

// Exec executes the deletion query.
func (otdo *OASTypesDeleteOne) Exec(ctx context.Context) error {
	n, err := otdo.otd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oastypes.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (otdo *OASTypesDeleteOne) ExecX(ctx context.Context) {
	otdo.otd.ExecX(ctx)
}
