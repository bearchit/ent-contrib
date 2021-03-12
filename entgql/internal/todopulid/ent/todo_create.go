// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/contrib/entgql/internal/todopulid/ent/todo"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoCreate is the builder for creating a Todo entity.
type TodoCreate struct {
	config
	mutation *TodoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TodoCreate) SetCreatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableCreatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TodoCreate) SetStatus(t todo.Status) *TodoCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetPriority sets the "priority" field.
func (tc *TodoCreate) SetPriority(i int) *TodoCreate {
	tc.mutation.SetPriority(i)
	return tc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tc *TodoCreate) SetNillablePriority(i *int) *TodoCreate {
	if i != nil {
		tc.SetPriority(*i)
	}
	return tc
}

// SetText sets the "text" field.
func (tc *TodoCreate) SetText(s string) *TodoCreate {
	tc.mutation.SetText(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TodoCreate) SetID(pu pulid.ID) *TodoCreate {
	tc.mutation.SetID(pu)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TodoCreate) SetNillableID(pu *pulid.ID) *TodoCreate {
	if pu != nil {
		tc.SetID(*pu)
	}
	return tc
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tc *TodoCreate) SetParentID(id pulid.ID) *TodoCreate {
	tc.mutation.SetParentID(id)
	return tc
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tc *TodoCreate) SetNillableParentID(id *pulid.ID) *TodoCreate {
	if id != nil {
		tc = tc.SetParentID(*id)
	}
	return tc
}

// SetParent sets the "parent" edge to the Todo entity.
func (tc *TodoCreate) SetParent(t *Todo) *TodoCreate {
	return tc.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tc *TodoCreate) AddChildIDs(ids ...pulid.ID) *TodoCreate {
	tc.mutation.AddChildIDs(ids...)
	return tc
}

// AddChildren adds the "children" edges to the Todo entity.
func (tc *TodoCreate) AddChildren(t ...*Todo) *TodoCreate {
	ids := make([]pulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddChildIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (tc *TodoCreate) Mutation() *TodoMutation {
	return tc.mutation
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*Todo, error) {
	var (
		err  error
		node *Todo
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodoCreate) SaveX(ctx context.Context) *Todo {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tc *TodoCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := todo.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.Priority(); !ok {
		v := todo.DefaultPriority
		tc.mutation.SetPriority(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := todo.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TodoCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New("ent: missing required field \"priority\"")}
	}
	if _, ok := tc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New("ent: missing required field \"text\"")}
	}
	if v, ok := tc.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf("ent: validator failed for field \"text\": %w", err)}
		}
	}
	return nil
}

func (tc *TodoCreate) sqlSave(ctx context.Context) (*Todo, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (tc *TodoCreate) createSpec() (*Todo, *sqlgraph.CreateSpec) {
	var (
		_node = &Todo{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: todo.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: todo.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tc.mutation.Priority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todo.FieldPriority,
		})
		_node.Priority = value
	}
	if value, ok := tc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldText,
		})
		_node.Text = value
	}
	if nodes := tc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TodoCreateBulk is the builder for creating many Todo entities in bulk.
type TodoCreateBulk struct {
	config
	builders []*TodoCreate
}

// Save creates the Todo entities in the database.
func (tcb *TodoCreateBulk) Save(ctx context.Context) ([]*Todo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Todo, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TodoMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TodoCreateBulk) SaveX(ctx context.Context) []*Todo {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
