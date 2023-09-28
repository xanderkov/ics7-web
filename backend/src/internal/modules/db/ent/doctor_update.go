// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hospital/internal/modules/db/ent/account"
	"hospital/internal/modules/db/ent/doctor"
	"hospital/internal/modules/db/ent/patient"
	"hospital/internal/modules/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DoctorUpdate is the builder for updating Doctor entities.
type DoctorUpdate struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// Where appends a list predicates to the DoctorUpdate builder.
func (du *DoctorUpdate) Where(ps ...predicate.Doctor) *DoctorUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetTokenId sets the "tokenId" field.
func (du *DoctorUpdate) SetTokenId(s string) *DoctorUpdate {
	du.mutation.SetTokenId(s)
	return du
}

// SetSurname sets the "surname" field.
func (du *DoctorUpdate) SetSurname(s string) *DoctorUpdate {
	du.mutation.SetSurname(s)
	return du
}

// SetSpeciality sets the "speciality" field.
func (du *DoctorUpdate) SetSpeciality(s string) *DoctorUpdate {
	du.mutation.SetSpeciality(s)
	return du
}

// SetRole sets the "role" field.
func (du *DoctorUpdate) SetRole(s string) *DoctorUpdate {
	du.mutation.SetRole(s)
	return du
}

// AddTreatIDs adds the "treats" edge to the Patient entity by IDs.
func (du *DoctorUpdate) AddTreatIDs(ids ...int) *DoctorUpdate {
	du.mutation.AddTreatIDs(ids...)
	return du
}

// AddTreats adds the "treats" edges to the Patient entity.
func (du *DoctorUpdate) AddTreats(p ...*Patient) *DoctorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddTreatIDs(ids...)
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (du *DoctorUpdate) SetAccountID(id int) *DoctorUpdate {
	du.mutation.SetAccountID(id)
	return du
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (du *DoctorUpdate) SetNillableAccountID(id *int) *DoctorUpdate {
	if id != nil {
		du = du.SetAccountID(*id)
	}
	return du
}

// SetAccount sets the "account" edge to the Account entity.
func (du *DoctorUpdate) SetAccount(a *Account) *DoctorUpdate {
	return du.SetAccountID(a.ID)
}

// Mutation returns the DoctorMutation object of the builder.
func (du *DoctorUpdate) Mutation() *DoctorMutation {
	return du.mutation
}

// ClearTreats clears all "treats" edges to the Patient entity.
func (du *DoctorUpdate) ClearTreats() *DoctorUpdate {
	du.mutation.ClearTreats()
	return du
}

// RemoveTreatIDs removes the "treats" edge to Patient entities by IDs.
func (du *DoctorUpdate) RemoveTreatIDs(ids ...int) *DoctorUpdate {
	du.mutation.RemoveTreatIDs(ids...)
	return du
}

// RemoveTreats removes "treats" edges to Patient entities.
func (du *DoctorUpdate) RemoveTreats(p ...*Patient) *DoctorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemoveTreatIDs(ids...)
}

// ClearAccount clears the "account" edge to the Account entity.
func (du *DoctorUpdate) ClearAccount() *DoctorUpdate {
	du.mutation.ClearAccount()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DoctorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, DoctorMutation](ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(doctor.Table, doctor.Columns, sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.TokenId(); ok {
		_spec.SetField(doctor.FieldTokenId, field.TypeString, value)
	}
	if value, ok := du.mutation.Surname(); ok {
		_spec.SetField(doctor.FieldSurname, field.TypeString, value)
	}
	if value, ok := du.mutation.Speciality(); ok {
		_spec.SetField(doctor.FieldSpeciality, field.TypeString, value)
	}
	if value, ok := du.mutation.Role(); ok {
		_spec.SetField(doctor.FieldRole, field.TypeString, value)
	}
	if du.mutation.TreatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   doctor.TreatsTable,
			Columns: doctor.TreatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedTreatsIDs(); len(nodes) > 0 && !du.mutation.TreatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   doctor.TreatsTable,
			Columns: doctor.TreatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.TreatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   doctor.TreatsTable,
			Columns: doctor.TreatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.AccountTable,
			Columns: []string{doctor.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.AccountTable,
			Columns: []string{doctor.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DoctorUpdateOne is the builder for updating a single Doctor entity.
type DoctorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DoctorMutation
}

// SetTokenId sets the "tokenId" field.
func (duo *DoctorUpdateOne) SetTokenId(s string) *DoctorUpdateOne {
	duo.mutation.SetTokenId(s)
	return duo
}

// SetSurname sets the "surname" field.
func (duo *DoctorUpdateOne) SetSurname(s string) *DoctorUpdateOne {
	duo.mutation.SetSurname(s)
	return duo
}

// SetSpeciality sets the "speciality" field.
func (duo *DoctorUpdateOne) SetSpeciality(s string) *DoctorUpdateOne {
	duo.mutation.SetSpeciality(s)
	return duo
}

// SetRole sets the "role" field.
func (duo *DoctorUpdateOne) SetRole(s string) *DoctorUpdateOne {
	duo.mutation.SetRole(s)
	return duo
}

// AddTreatIDs adds the "treats" edge to the Patient entity by IDs.
func (duo *DoctorUpdateOne) AddTreatIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.AddTreatIDs(ids...)
	return duo
}

// AddTreats adds the "treats" edges to the Patient entity.
func (duo *DoctorUpdateOne) AddTreats(p ...*Patient) *DoctorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddTreatIDs(ids...)
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (duo *DoctorUpdateOne) SetAccountID(id int) *DoctorUpdateOne {
	duo.mutation.SetAccountID(id)
	return duo
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableAccountID(id *int) *DoctorUpdateOne {
	if id != nil {
		duo = duo.SetAccountID(*id)
	}
	return duo
}

// SetAccount sets the "account" edge to the Account entity.
func (duo *DoctorUpdateOne) SetAccount(a *Account) *DoctorUpdateOne {
	return duo.SetAccountID(a.ID)
}

// Mutation returns the DoctorMutation object of the builder.
func (duo *DoctorUpdateOne) Mutation() *DoctorMutation {
	return duo.mutation
}

// ClearTreats clears all "treats" edges to the Patient entity.
func (duo *DoctorUpdateOne) ClearTreats() *DoctorUpdateOne {
	duo.mutation.ClearTreats()
	return duo
}

// RemoveTreatIDs removes the "treats" edge to Patient entities by IDs.
func (duo *DoctorUpdateOne) RemoveTreatIDs(ids ...int) *DoctorUpdateOne {
	duo.mutation.RemoveTreatIDs(ids...)
	return duo
}

// RemoveTreats removes "treats" edges to Patient entities.
func (duo *DoctorUpdateOne) RemoveTreats(p ...*Patient) *DoctorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemoveTreatIDs(ids...)
}

// ClearAccount clears the "account" edge to the Account entity.
func (duo *DoctorUpdateOne) ClearAccount() *DoctorUpdateOne {
	duo.mutation.ClearAccount()
	return duo
}

// Where appends a list predicates to the DoctorUpdate builder.
func (duo *DoctorUpdateOne) Where(ps ...predicate.Doctor) *DoctorUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DoctorUpdateOne) Select(field string, fields ...string) *DoctorUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Doctor entity.
func (duo *DoctorUpdateOne) Save(ctx context.Context) (*Doctor, error) {
	return withHooks[*Doctor, DoctorMutation](ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DoctorUpdateOne) SaveX(ctx context.Context) *Doctor {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DoctorUpdateOne) sqlSave(ctx context.Context) (_node *Doctor, err error) {
	_spec := sqlgraph.NewUpdateSpec(doctor.Table, doctor.Columns, sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Doctor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, doctor.FieldID)
		for _, f := range fields {
			if !doctor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != doctor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.TokenId(); ok {
		_spec.SetField(doctor.FieldTokenId, field.TypeString, value)
	}
	if value, ok := duo.mutation.Surname(); ok {
		_spec.SetField(doctor.FieldSurname, field.TypeString, value)
	}
	if value, ok := duo.mutation.Speciality(); ok {
		_spec.SetField(doctor.FieldSpeciality, field.TypeString, value)
	}
	if value, ok := duo.mutation.Role(); ok {
		_spec.SetField(doctor.FieldRole, field.TypeString, value)
	}
	if duo.mutation.TreatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   doctor.TreatsTable,
			Columns: doctor.TreatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedTreatsIDs(); len(nodes) > 0 && !duo.mutation.TreatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   doctor.TreatsTable,
			Columns: doctor.TreatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.TreatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   doctor.TreatsTable,
			Columns: doctor.TreatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.AccountTable,
			Columns: []string{doctor.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.AccountTable,
			Columns: []string{doctor.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Doctor{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
