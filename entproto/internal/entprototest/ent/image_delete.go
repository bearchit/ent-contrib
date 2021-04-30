// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/image"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageDelete is the builder for deleting a Image entity.
type ImageDelete struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where adds a new predicate to the ImageDelete builder.
func (id *ImageDelete) Where(ps ...predicate.Image) *ImageDelete {
	id.mutation.predicates = append(id.mutation.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *ImageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(id.hooks) - 1; i >= 0; i-- {
			mut = id.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *ImageDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *ImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: image.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: image.FieldID,
			},
		},
	}
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, id.driver, _spec)
}

// ImageDeleteOne is the builder for deleting a single Image entity.
type ImageDeleteOne struct {
	id *ImageDelete
}

// Exec executes the deletion query.
func (ido *ImageDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{image.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *ImageDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
