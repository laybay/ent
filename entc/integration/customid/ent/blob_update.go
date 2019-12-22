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
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// BlobUpdate is the builder for updating Blob entities.
type BlobUpdate struct {
	config
	uuid       *uuid.UUID
	predicates []predicate.Blob
}

// Where adds a new predicate for the builder.
func (bu *BlobUpdate) Where(ps ...predicate.Blob) *BlobUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetUUID sets the uuid field.
func (bu *BlobUpdate) SetUUID(u uuid.UUID) *BlobUpdate {
	bu.uuid = &u
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BlobUpdate) Save(ctx context.Context) (int, error) {
	return bu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlobUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlobUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlobUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BlobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := bu.uuid; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  *value,
			Column: blob.FieldUUID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BlobUpdateOne is the builder for updating a single Blob entity.
type BlobUpdateOne struct {
	config
	id   uuid.UUID
	uuid *uuid.UUID
}

// SetUUID sets the uuid field.
func (buo *BlobUpdateOne) SetUUID(u uuid.UUID) *BlobUpdateOne {
	buo.uuid = &u
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BlobUpdateOne) Save(ctx context.Context) (*Blob, error) {
	return buo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlobUpdateOne) SaveX(ctx context.Context) *Blob {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BlobUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlobUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BlobUpdateOne) sqlSave(ctx context.Context) (b *Blob, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  buo.id,
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	if value := buo.uuid; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  *value,
			Column: blob.FieldUUID,
		})
	}
	b = &Blob{config: buo.config}
	spec.ScanTypes = []interface{}{
		&uuid.UUID{},
		&uuid.UUID{},
	}
	spec.Assign = func(values ...interface{}) error {
		if m, n := len(values), len(spec.ScanTypes); m != n {
			return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
		}
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field id", values[0])
		} else if value != nil {
			b.ID = *value
		}
		values = values[1:]
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field uuid", values[0])
		} else if value != nil {
			b.UUID = *value
		}
		return nil
	}
	if err = sqlgraph.UpdateNode(ctx, buo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
