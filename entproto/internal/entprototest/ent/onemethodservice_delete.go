// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/onemethodservice"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OneMethodServiceDelete is the builder for deleting a OneMethodService entity.
type OneMethodServiceDelete struct {
	config
	hooks    []Hook
	mutation *OneMethodServiceMutation
}

// Where appends a list predicates to the OneMethodServiceDelete builder.
func (omsd *OneMethodServiceDelete) Where(ps ...predicate.OneMethodService) *OneMethodServiceDelete {
	omsd.mutation.Where(ps...)
	return omsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (omsd *OneMethodServiceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(omsd.hooks) == 0 {
		affected, err = omsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OneMethodServiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			omsd.mutation = mutation
			affected, err = omsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(omsd.hooks) - 1; i >= 0; i-- {
			if omsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = omsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, omsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (omsd *OneMethodServiceDelete) ExecX(ctx context.Context) int {
	n, err := omsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (omsd *OneMethodServiceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: onemethodservice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: onemethodservice.FieldID,
			},
		},
	}
	if ps := omsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, omsd.driver, _spec)
}

// OneMethodServiceDeleteOne is the builder for deleting a single OneMethodService entity.
type OneMethodServiceDeleteOne struct {
	omsd *OneMethodServiceDelete
}

// Exec executes the deletion query.
func (omsdo *OneMethodServiceDeleteOne) Exec(ctx context.Context) error {
	n, err := omsdo.omsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{onemethodservice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (omsdo *OneMethodServiceDeleteOne) ExecX(ctx context.Context) {
	omsdo.omsd.ExecX(ctx)
}
