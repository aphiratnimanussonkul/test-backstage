// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/course"
	"github.com/sut63/team17/app/ent/degree"
	"github.com/sut63/team17/app/ent/faculty"
	"github.com/sut63/team17/app/ent/institution"
)

// CourseCreate is the builder for creating a Course entity.
type CourseCreate struct {
	config
	mutation *CourseMutation
	hooks    []Hook
}

// SetCourse sets the course field.
func (cc *CourseCreate) SetCourse(s string) *CourseCreate {
	cc.mutation.SetCourse(s)
	return cc
}

// SetCourFacuID sets the cour_facu edge to Faculty by id.
func (cc *CourseCreate) SetCourFacuID(id int) *CourseCreate {
	cc.mutation.SetCourFacuID(id)
	return cc
}

// SetNillableCourFacuID sets the cour_facu edge to Faculty by id if the given value is not nil.
func (cc *CourseCreate) SetNillableCourFacuID(id *int) *CourseCreate {
	if id != nil {
		cc = cc.SetCourFacuID(*id)
	}
	return cc
}

// SetCourFacu sets the cour_facu edge to Faculty.
func (cc *CourseCreate) SetCourFacu(f *Faculty) *CourseCreate {
	return cc.SetCourFacuID(f.ID)
}

// SetCourDegrID sets the cour_degr edge to Degree by id.
func (cc *CourseCreate) SetCourDegrID(id int) *CourseCreate {
	cc.mutation.SetCourDegrID(id)
	return cc
}

// SetNillableCourDegrID sets the cour_degr edge to Degree by id if the given value is not nil.
func (cc *CourseCreate) SetNillableCourDegrID(id *int) *CourseCreate {
	if id != nil {
		cc = cc.SetCourDegrID(*id)
	}
	return cc
}

// SetCourDegr sets the cour_degr edge to Degree.
func (cc *CourseCreate) SetCourDegr(d *Degree) *CourseCreate {
	return cc.SetCourDegrID(d.ID)
}

// SetCourInstID sets the cour_inst edge to Institution by id.
func (cc *CourseCreate) SetCourInstID(id int) *CourseCreate {
	cc.mutation.SetCourInstID(id)
	return cc
}

// SetNillableCourInstID sets the cour_inst edge to Institution by id if the given value is not nil.
func (cc *CourseCreate) SetNillableCourInstID(id *int) *CourseCreate {
	if id != nil {
		cc = cc.SetCourInstID(*id)
	}
	return cc
}

// SetCourInst sets the cour_inst edge to Institution.
func (cc *CourseCreate) SetCourInst(i *Institution) *CourseCreate {
	return cc.SetCourInstID(i.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cc *CourseCreate) Mutation() *CourseMutation {
	return cc.mutation
}

// Save creates the Course in the database.
func (cc *CourseCreate) Save(ctx context.Context) (*Course, error) {
	if _, ok := cc.mutation.Course(); !ok {
		return nil, &ValidationError{Name: "course", err: errors.New("ent: missing required field \"course\"")}
	}
	if v, ok := cc.mutation.Course(); ok {
		if err := course.CourseValidator(v); err != nil {
			return nil, &ValidationError{Name: "course", err: fmt.Errorf("ent: validator failed for field \"course\": %w", err)}
		}
	}
	var (
		err  error
		node *Course
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CourseCreate) SaveX(ctx context.Context) *Course {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CourseCreate) sqlSave(ctx context.Context) (*Course, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CourseCreate) createSpec() (*Course, *sqlgraph.CreateSpec) {
	var (
		c     = &Course{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: course.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: course.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Course(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldCourse,
		})
		c.Course = value
	}
	if nodes := cc.mutation.CourFacuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.CourFacuTable,
			Columns: []string{course.CourFacuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CourDegrIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.CourDegrTable,
			Columns: []string{course.CourDegrColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: degree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CourInstIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   course.CourInstTable,
			Columns: []string{course.CourInstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}
