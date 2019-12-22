// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/template/ent/group"
	"github.com/facebookincubator/ent/entc/integration/template/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	max_users    *int
	addmax_users *int
	predicates   []predicate.Group
}

// Where adds a new predicate for the builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetMaxUsers sets the max_users field.
func (gu *GroupUpdate) SetMaxUsers(i int) *GroupUpdate {
	gu.max_users = &i
	gu.addmax_users = nil
	return gu
}

// AddMaxUsers adds i to max_users.
func (gu *GroupUpdate) AddMaxUsers(i int) *GroupUpdate {
	if gu.addmax_users == nil {
		gu.addmax_users = &i
	} else {
		*gu.addmax_users += i
	}
	return gu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	return gu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
	}
	if ps := gu.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := gu.max_users; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: group.FieldMaxUsers,
		})
	}
	if value := gu.addmax_users; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: group.FieldMaxUsers,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	id           int
	max_users    *int
	addmax_users *int
}

// SetMaxUsers sets the max_users field.
func (guo *GroupUpdateOne) SetMaxUsers(i int) *GroupUpdateOne {
	guo.max_users = &i
	guo.addmax_users = nil
	return guo
}

// AddMaxUsers adds i to max_users.
func (guo *GroupUpdateOne) AddMaxUsers(i int) *GroupUpdateOne {
	if guo.addmax_users == nil {
		guo.addmax_users = &i
	} else {
		*guo.addmax_users += i
	}
	return guo
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return guo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	gr, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (gr *Group, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  guo.id,
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
	}
	if value := guo.max_users; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: group.FieldMaxUsers,
		})
	}
	if value := guo.addmax_users; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: group.FieldMaxUsers,
		})
	}
	gr = &Group{config: guo.config}
	spec.ScanTypes = []interface{}{
		&sql.NullInt64{},
		&sql.NullInt64{},
	}
	spec.Assign = func(values ...interface{}) error {
		if m, n := len(values), len(spec.ScanTypes); m != n {
			return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
		}
		value, ok := values[0].(*sql.NullInt64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field id", value)
		}
		gr.ID = int(value.Int64)
		values = values[1:]
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for field max_users", values[0])
		} else if value.Valid {
			gr.MaxUsers = int(value.Int64)
		}
		return nil
	}
	if err = sqlgraph.UpdateNode(ctx, guo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return gr, nil
}
