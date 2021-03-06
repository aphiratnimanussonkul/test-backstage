// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/predicate"
	"github.com/sut63/team17/app/ent/prefix"
	"github.com/sut63/team17/app/ent/professor"
	"github.com/sut63/team17/app/ent/student"
)

// PrefixUpdate is the builder for updating Prefix entities.
type PrefixUpdate struct {
	config
	hooks      []Hook
	mutation   *PrefixMutation
	predicates []predicate.Prefix
}

// Where adds a new predicate for the builder.
func (pu *PrefixUpdate) Where(ps ...predicate.Prefix) *PrefixUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPrefix sets the prefix field.
func (pu *PrefixUpdate) SetPrefix(s string) *PrefixUpdate {
	pu.mutation.SetPrefix(s)
	return pu
}

// AddPreProfIDs adds the pre_prof edge to Professor by ids.
func (pu *PrefixUpdate) AddPreProfIDs(ids ...int) *PrefixUpdate {
	pu.mutation.AddPreProfIDs(ids...)
	return pu
}

// AddPreProf adds the pre_prof edges to Professor.
func (pu *PrefixUpdate) AddPreProf(p ...*Professor) *PrefixUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPreProfIDs(ids...)
}

// AddPrefStudIDs adds the pref_stud edge to Student by ids.
func (pu *PrefixUpdate) AddPrefStudIDs(ids ...int) *PrefixUpdate {
	pu.mutation.AddPrefStudIDs(ids...)
	return pu
}

// AddPrefStud adds the pref_stud edges to Student.
func (pu *PrefixUpdate) AddPrefStud(s ...*Student) *PrefixUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddPrefStudIDs(ids...)
}

// Mutation returns the PrefixMutation object of the builder.
func (pu *PrefixUpdate) Mutation() *PrefixMutation {
	return pu.mutation
}

// RemovePreProfIDs removes the pre_prof edge to Professor by ids.
func (pu *PrefixUpdate) RemovePreProfIDs(ids ...int) *PrefixUpdate {
	pu.mutation.RemovePreProfIDs(ids...)
	return pu
}

// RemovePreProf removes pre_prof edges to Professor.
func (pu *PrefixUpdate) RemovePreProf(p ...*Professor) *PrefixUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePreProfIDs(ids...)
}

// RemovePrefStudIDs removes the pref_stud edge to Student by ids.
func (pu *PrefixUpdate) RemovePrefStudIDs(ids ...int) *PrefixUpdate {
	pu.mutation.RemovePrefStudIDs(ids...)
	return pu
}

// RemovePrefStud removes pref_stud edges to Student.
func (pu *PrefixUpdate) RemovePrefStud(s ...*Student) *PrefixUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemovePrefStudIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PrefixUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Prefix(); ok {
		if err := prefix.PrefixValidator(v); err != nil {
			return 0, &ValidationError{Name: "prefix", err: fmt.Errorf("ent: validator failed for field \"prefix\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrefixMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrefixUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrefixUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrefixUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrefixUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prefix.Table,
			Columns: prefix.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prefix.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Prefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prefix.FieldPrefix,
		})
	}
	if nodes := pu.mutation.RemovedPreProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PreProfTable,
			Columns: []string{prefix.PreProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PreProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PreProfTable,
			Columns: []string{prefix.PreProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedPrefStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PrefStudTable,
			Columns: []string{prefix.PrefStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrefStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PrefStudTable,
			Columns: []string{prefix.PrefStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prefix.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PrefixUpdateOne is the builder for updating a single Prefix entity.
type PrefixUpdateOne struct {
	config
	hooks    []Hook
	mutation *PrefixMutation
}

// SetPrefix sets the prefix field.
func (puo *PrefixUpdateOne) SetPrefix(s string) *PrefixUpdateOne {
	puo.mutation.SetPrefix(s)
	return puo
}

// AddPreProfIDs adds the pre_prof edge to Professor by ids.
func (puo *PrefixUpdateOne) AddPreProfIDs(ids ...int) *PrefixUpdateOne {
	puo.mutation.AddPreProfIDs(ids...)
	return puo
}

// AddPreProf adds the pre_prof edges to Professor.
func (puo *PrefixUpdateOne) AddPreProf(p ...*Professor) *PrefixUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPreProfIDs(ids...)
}

// AddPrefStudIDs adds the pref_stud edge to Student by ids.
func (puo *PrefixUpdateOne) AddPrefStudIDs(ids ...int) *PrefixUpdateOne {
	puo.mutation.AddPrefStudIDs(ids...)
	return puo
}

// AddPrefStud adds the pref_stud edges to Student.
func (puo *PrefixUpdateOne) AddPrefStud(s ...*Student) *PrefixUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddPrefStudIDs(ids...)
}

// Mutation returns the PrefixMutation object of the builder.
func (puo *PrefixUpdateOne) Mutation() *PrefixMutation {
	return puo.mutation
}

// RemovePreProfIDs removes the pre_prof edge to Professor by ids.
func (puo *PrefixUpdateOne) RemovePreProfIDs(ids ...int) *PrefixUpdateOne {
	puo.mutation.RemovePreProfIDs(ids...)
	return puo
}

// RemovePreProf removes pre_prof edges to Professor.
func (puo *PrefixUpdateOne) RemovePreProf(p ...*Professor) *PrefixUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePreProfIDs(ids...)
}

// RemovePrefStudIDs removes the pref_stud edge to Student by ids.
func (puo *PrefixUpdateOne) RemovePrefStudIDs(ids ...int) *PrefixUpdateOne {
	puo.mutation.RemovePrefStudIDs(ids...)
	return puo
}

// RemovePrefStud removes pref_stud edges to Student.
func (puo *PrefixUpdateOne) RemovePrefStud(s ...*Student) *PrefixUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemovePrefStudIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PrefixUpdateOne) Save(ctx context.Context) (*Prefix, error) {
	if v, ok := puo.mutation.Prefix(); ok {
		if err := prefix.PrefixValidator(v); err != nil {
			return nil, &ValidationError{Name: "prefix", err: fmt.Errorf("ent: validator failed for field \"prefix\": %w", err)}
		}
	}

	var (
		err  error
		node *Prefix
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrefixMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrefixUpdateOne) SaveX(ctx context.Context) *Prefix {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PrefixUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrefixUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrefixUpdateOne) sqlSave(ctx context.Context) (pr *Prefix, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prefix.Table,
			Columns: prefix.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prefix.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Prefix.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Prefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prefix.FieldPrefix,
		})
	}
	if nodes := puo.mutation.RemovedPreProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PreProfTable,
			Columns: []string{prefix.PreProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PreProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PreProfTable,
			Columns: []string{prefix.PreProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedPrefStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PrefStudTable,
			Columns: []string{prefix.PrefStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrefStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PrefStudTable,
			Columns: []string{prefix.PrefStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Prefix{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prefix.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
