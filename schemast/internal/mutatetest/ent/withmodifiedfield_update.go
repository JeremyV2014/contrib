// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/user"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withmodifiedfield"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithModifiedFieldUpdate is the builder for updating WithModifiedField entities.
type WithModifiedFieldUpdate struct {
	config
	hooks    []Hook
	mutation *WithModifiedFieldMutation
}

// Where adds a new predicate for the WithModifiedFieldUpdate builder.
func (wmfu *WithModifiedFieldUpdate) Where(ps ...predicate.WithModifiedField) *WithModifiedFieldUpdate {
	wmfu.mutation.predicates = append(wmfu.mutation.predicates, ps...)
	return wmfu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (wmfu *WithModifiedFieldUpdate) SetOwnerID(id int) *WithModifiedFieldUpdate {
	wmfu.mutation.SetOwnerID(id)
	return wmfu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (wmfu *WithModifiedFieldUpdate) SetNillableOwnerID(id *int) *WithModifiedFieldUpdate {
	if id != nil {
		wmfu = wmfu.SetOwnerID(*id)
	}
	return wmfu
}

// SetOwner sets the "owner" edge to the User entity.
func (wmfu *WithModifiedFieldUpdate) SetOwner(u *User) *WithModifiedFieldUpdate {
	return wmfu.SetOwnerID(u.ID)
}

// Mutation returns the WithModifiedFieldMutation object of the builder.
func (wmfu *WithModifiedFieldUpdate) Mutation() *WithModifiedFieldMutation {
	return wmfu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (wmfu *WithModifiedFieldUpdate) ClearOwner() *WithModifiedFieldUpdate {
	wmfu.mutation.ClearOwner()
	return wmfu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wmfu *WithModifiedFieldUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wmfu.hooks) == 0 {
		affected, err = wmfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithModifiedFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wmfu.mutation = mutation
			affected, err = wmfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wmfu.hooks) - 1; i >= 0; i-- {
			mut = wmfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wmfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wmfu *WithModifiedFieldUpdate) SaveX(ctx context.Context) int {
	affected, err := wmfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wmfu *WithModifiedFieldUpdate) Exec(ctx context.Context) error {
	_, err := wmfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wmfu *WithModifiedFieldUpdate) ExecX(ctx context.Context) {
	if err := wmfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wmfu *WithModifiedFieldUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withmodifiedfield.Table,
			Columns: withmodifiedfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withmodifiedfield.FieldID,
			},
		},
	}
	if ps := wmfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wmfu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   withmodifiedfield.OwnerTable,
			Columns: []string{withmodifiedfield.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wmfu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   withmodifiedfield.OwnerTable,
			Columns: []string{withmodifiedfield.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wmfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withmodifiedfield.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WithModifiedFieldUpdateOne is the builder for updating a single WithModifiedField entity.
type WithModifiedFieldUpdateOne struct {
	config
	hooks    []Hook
	mutation *WithModifiedFieldMutation
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (wmfuo *WithModifiedFieldUpdateOne) SetOwnerID(id int) *WithModifiedFieldUpdateOne {
	wmfuo.mutation.SetOwnerID(id)
	return wmfuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (wmfuo *WithModifiedFieldUpdateOne) SetNillableOwnerID(id *int) *WithModifiedFieldUpdateOne {
	if id != nil {
		wmfuo = wmfuo.SetOwnerID(*id)
	}
	return wmfuo
}

// SetOwner sets the "owner" edge to the User entity.
func (wmfuo *WithModifiedFieldUpdateOne) SetOwner(u *User) *WithModifiedFieldUpdateOne {
	return wmfuo.SetOwnerID(u.ID)
}

// Mutation returns the WithModifiedFieldMutation object of the builder.
func (wmfuo *WithModifiedFieldUpdateOne) Mutation() *WithModifiedFieldMutation {
	return wmfuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (wmfuo *WithModifiedFieldUpdateOne) ClearOwner() *WithModifiedFieldUpdateOne {
	wmfuo.mutation.ClearOwner()
	return wmfuo
}

// Save executes the query and returns the updated WithModifiedField entity.
func (wmfuo *WithModifiedFieldUpdateOne) Save(ctx context.Context) (*WithModifiedField, error) {
	var (
		err  error
		node *WithModifiedField
	)
	if len(wmfuo.hooks) == 0 {
		node, err = wmfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithModifiedFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wmfuo.mutation = mutation
			node, err = wmfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wmfuo.hooks) - 1; i >= 0; i-- {
			mut = wmfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wmfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wmfuo *WithModifiedFieldUpdateOne) SaveX(ctx context.Context) *WithModifiedField {
	node, err := wmfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wmfuo *WithModifiedFieldUpdateOne) Exec(ctx context.Context) error {
	_, err := wmfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wmfuo *WithModifiedFieldUpdateOne) ExecX(ctx context.Context) {
	if err := wmfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wmfuo *WithModifiedFieldUpdateOne) sqlSave(ctx context.Context) (_node *WithModifiedField, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withmodifiedfield.Table,
			Columns: withmodifiedfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withmodifiedfield.FieldID,
			},
		},
	}
	id, ok := wmfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WithModifiedField.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := wmfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wmfuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   withmodifiedfield.OwnerTable,
			Columns: []string{withmodifiedfield.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wmfuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   withmodifiedfield.OwnerTable,
			Columns: []string{withmodifiedfield.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WithModifiedField{config: wmfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wmfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withmodifiedfield.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}