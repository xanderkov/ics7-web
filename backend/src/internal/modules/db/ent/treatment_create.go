// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hospital/internal/modules/db/ent/patient"
	"hospital/internal/modules/db/ent/treatment"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TreatmentCreate is the builder for creating a Treatment entity.
type TreatmentCreate struct {
	config
	mutation *TreatmentMutation
	hooks    []Hook
}

// SetTablets sets the "tablets" field.
func (tc *TreatmentCreate) SetTablets(s string) *TreatmentCreate {
	tc.mutation.SetTablets(s)
	return tc
}

// SetPsychologicalTreatment sets the "psychologicalTreatment" field.
func (tc *TreatmentCreate) SetPsychologicalTreatment(s string) *TreatmentCreate {
	tc.mutation.SetPsychologicalTreatment(s)
	return tc
}

// SetSurvey sets the "survey" field.
func (tc *TreatmentCreate) SetSurvey(s string) *TreatmentCreate {
	tc.mutation.SetSurvey(s)
	return tc
}

// AddCuredIDs adds the "cured" edge to the Patient entity by IDs.
func (tc *TreatmentCreate) AddCuredIDs(ids ...int) *TreatmentCreate {
	tc.mutation.AddCuredIDs(ids...)
	return tc
}

// AddCured adds the "cured" edges to the Patient entity.
func (tc *TreatmentCreate) AddCured(p ...*Patient) *TreatmentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddCuredIDs(ids...)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tc *TreatmentCreate) Mutation() *TreatmentMutation {
	return tc.mutation
}

// Save creates the Treatment in the database.
func (tc *TreatmentCreate) Save(ctx context.Context) (*Treatment, error) {
	return withHooks[*Treatment, TreatmentMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TreatmentCreate) SaveX(ctx context.Context) *Treatment {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TreatmentCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TreatmentCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TreatmentCreate) check() error {
	if _, ok := tc.mutation.Tablets(); !ok {
		return &ValidationError{Name: "tablets", err: errors.New(`ent: missing required field "Treatment.tablets"`)}
	}
	if _, ok := tc.mutation.PsychologicalTreatment(); !ok {
		return &ValidationError{Name: "psychologicalTreatment", err: errors.New(`ent: missing required field "Treatment.psychologicalTreatment"`)}
	}
	if _, ok := tc.mutation.Survey(); !ok {
		return &ValidationError{Name: "survey", err: errors.New(`ent: missing required field "Treatment.survey"`)}
	}
	return nil
}

func (tc *TreatmentCreate) sqlSave(ctx context.Context) (*Treatment, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TreatmentCreate) createSpec() (*Treatment, *sqlgraph.CreateSpec) {
	var (
		_node = &Treatment{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(treatment.Table, sqlgraph.NewFieldSpec(treatment.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Tablets(); ok {
		_spec.SetField(treatment.FieldTablets, field.TypeString, value)
		_node.Tablets = value
	}
	if value, ok := tc.mutation.PsychologicalTreatment(); ok {
		_spec.SetField(treatment.FieldPsychologicalTreatment, field.TypeString, value)
		_node.PsychologicalTreatment = value
	}
	if value, ok := tc.mutation.Survey(); ok {
		_spec.SetField(treatment.FieldSurvey, field.TypeString, value)
		_node.Survey = value
	}
	if nodes := tc.mutation.CuredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   treatment.CuredTable,
			Columns: []string{treatment.CuredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TreatmentCreateBulk is the builder for creating many Treatment entities in bulk.
type TreatmentCreateBulk struct {
	config
	builders []*TreatmentCreate
}

// Save creates the Treatment entities in the database.
func (tcb *TreatmentCreateBulk) Save(ctx context.Context) ([]*Treatment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Treatment, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TreatmentMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TreatmentCreateBulk) SaveX(ctx context.Context) []*Treatment {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TreatmentCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TreatmentCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}