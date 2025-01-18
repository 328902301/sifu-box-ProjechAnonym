// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sifu-box/ent/provider"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProviderCreate is the builder for creating a Provider entity.
type ProviderCreate struct {
	config
	mutation *ProviderMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProviderCreate) SetName(s string) *ProviderCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetPath sets the "path" field.
func (pc *ProviderCreate) SetPath(s string) *ProviderCreate {
	pc.mutation.SetPath(s)
	return pc
}

// SetDetour sets the "detour" field.
func (pc *ProviderCreate) SetDetour(s string) *ProviderCreate {
	pc.mutation.SetDetour(s)
	return pc
}

// SetNillableDetour sets the "detour" field if the given value is not nil.
func (pc *ProviderCreate) SetNillableDetour(s *string) *ProviderCreate {
	if s != nil {
		pc.SetDetour(*s)
	}
	return pc
}

// SetRemote sets the "remote" field.
func (pc *ProviderCreate) SetRemote(b bool) *ProviderCreate {
	pc.mutation.SetRemote(b)
	return pc
}

// Mutation returns the ProviderMutation object of the builder.
func (pc *ProviderCreate) Mutation() *ProviderMutation {
	return pc.mutation
}

// Save creates the Provider in the database.
func (pc *ProviderCreate) Save(ctx context.Context) (*Provider, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProviderCreate) SaveX(ctx context.Context) *Provider {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProviderCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProviderCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProviderCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Provider.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := provider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Provider.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Provider.path"`)}
	}
	if v, ok := pc.mutation.Path(); ok {
		if err := provider.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Provider.path": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Detour(); ok {
		if err := provider.DetourValidator(v); err != nil {
			return &ValidationError{Name: "detour", err: fmt.Errorf(`ent: validator failed for field "Provider.detour": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Remote(); !ok {
		return &ValidationError{Name: "remote", err: errors.New(`ent: missing required field "Provider.remote"`)}
	}
	return nil
}

func (pc *ProviderCreate) sqlSave(ctx context.Context) (*Provider, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProviderCreate) createSpec() (*Provider, *sqlgraph.CreateSpec) {
	var (
		_node = &Provider{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(provider.Table, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(provider.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Path(); ok {
		_spec.SetField(provider.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := pc.mutation.Detour(); ok {
		_spec.SetField(provider.FieldDetour, field.TypeString, value)
		_node.Detour = value
	}
	if value, ok := pc.mutation.Remote(); ok {
		_spec.SetField(provider.FieldRemote, field.TypeBool, value)
		_node.Remote = value
	}
	return _node, _spec
}

// ProviderCreateBulk is the builder for creating many Provider entities in bulk.
type ProviderCreateBulk struct {
	config
	err      error
	builders []*ProviderCreate
}

// Save creates the Provider entities in the database.
func (pcb *ProviderCreateBulk) Save(ctx context.Context) ([]*Provider, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Provider, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProviderCreateBulk) SaveX(ctx context.Context) []*Provider {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProviderCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
