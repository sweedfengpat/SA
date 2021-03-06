// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/patza/app/ent/bill"
	"github.com/patza/app/ent/confirmation"
	"github.com/patza/app/ent/predicate"
	"github.com/patza/app/ent/user"
)

// ConfirmationUpdate is the builder for updating Confirmation entities.
type ConfirmationUpdate struct {
	config
	hooks      []Hook
	mutation   *ConfirmationMutation
	predicates []predicate.Confirmation
}

// Where adds a new predicate for the builder.
func (cu *ConfirmationUpdate) Where(ps ...predicate.Confirmation) *ConfirmationUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetAdddate sets the adddate field.
func (cu *ConfirmationUpdate) SetAdddate(t time.Time) *ConfirmationUpdate {
	cu.mutation.SetAdddate(t)
	return cu
}

// SetBookingstart sets the bookingstart field.
func (cu *ConfirmationUpdate) SetBookingstart(t time.Time) *ConfirmationUpdate {
	cu.mutation.SetBookingstart(t)
	return cu
}

// SetBookingend sets the bookingend field.
func (cu *ConfirmationUpdate) SetBookingend(t time.Time) *ConfirmationUpdate {
	cu.mutation.SetBookingend(t)
	return cu
}

// SetHourstime sets the hourstime field.
func (cu *ConfirmationUpdate) SetHourstime(i int) *ConfirmationUpdate {
	cu.mutation.ResetHourstime()
	cu.mutation.SetHourstime(i)
	return cu
}

// AddHourstime adds i to hourstime.
func (cu *ConfirmationUpdate) AddHourstime(i int) *ConfirmationUpdate {
	cu.mutation.AddHourstime(i)
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *ConfirmationUpdate) SetOwnerID(id int) *ConfirmationUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *ConfirmationUpdate) SetNillableOwnerID(id *int) *ConfirmationUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *ConfirmationUpdate) SetOwner(u *User) *ConfirmationUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddBillIDs adds the bill edge to Bill by ids.
func (cu *ConfirmationUpdate) AddBillIDs(ids ...int) *ConfirmationUpdate {
	cu.mutation.AddBillIDs(ids...)
	return cu
}

// AddBill adds the bill edges to Bill.
func (cu *ConfirmationUpdate) AddBill(b ...*Bill) *ConfirmationUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBillIDs(ids...)
}

// Mutation returns the ConfirmationMutation object of the builder.
func (cu *ConfirmationUpdate) Mutation() *ConfirmationMutation {
	return cu.mutation
}

// ClearOwner clears the owner edge to User.
func (cu *ConfirmationUpdate) ClearOwner() *ConfirmationUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// RemoveBillIDs removes the bill edge to Bill by ids.
func (cu *ConfirmationUpdate) RemoveBillIDs(ids ...int) *ConfirmationUpdate {
	cu.mutation.RemoveBillIDs(ids...)
	return cu
}

// RemoveBill removes bill edges to Bill.
func (cu *ConfirmationUpdate) RemoveBill(b ...*Bill) *ConfirmationUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *ConfirmationUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfirmationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfirmationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfirmationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfirmationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConfirmationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   confirmation.Table,
			Columns: confirmation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: confirmation.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Adddate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: confirmation.FieldAdddate,
		})
	}
	if value, ok := cu.mutation.Bookingstart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: confirmation.FieldBookingstart,
		})
	}
	if value, ok := cu.mutation.Bookingend(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: confirmation.FieldBookingend,
		})
	}
	if value, ok := cu.mutation.Hourstime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: confirmation.FieldHourstime,
		})
	}
	if value, ok := cu.mutation.AddedHourstime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: confirmation.FieldHourstime,
		})
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   confirmation.OwnerTable,
			Columns: []string{confirmation.OwnerColumn},
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
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   confirmation.OwnerTable,
			Columns: []string{confirmation.OwnerColumn},
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
	if nodes := cu.mutation.RemovedBillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   confirmation.BillTable,
			Columns: []string{confirmation.BillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   confirmation.BillTable,
			Columns: []string{confirmation.BillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{confirmation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ConfirmationUpdateOne is the builder for updating a single Confirmation entity.
type ConfirmationUpdateOne struct {
	config
	hooks    []Hook
	mutation *ConfirmationMutation
}

// SetAdddate sets the adddate field.
func (cuo *ConfirmationUpdateOne) SetAdddate(t time.Time) *ConfirmationUpdateOne {
	cuo.mutation.SetAdddate(t)
	return cuo
}

// SetBookingstart sets the bookingstart field.
func (cuo *ConfirmationUpdateOne) SetBookingstart(t time.Time) *ConfirmationUpdateOne {
	cuo.mutation.SetBookingstart(t)
	return cuo
}

// SetBookingend sets the bookingend field.
func (cuo *ConfirmationUpdateOne) SetBookingend(t time.Time) *ConfirmationUpdateOne {
	cuo.mutation.SetBookingend(t)
	return cuo
}

// SetHourstime sets the hourstime field.
func (cuo *ConfirmationUpdateOne) SetHourstime(i int) *ConfirmationUpdateOne {
	cuo.mutation.ResetHourstime()
	cuo.mutation.SetHourstime(i)
	return cuo
}

// AddHourstime adds i to hourstime.
func (cuo *ConfirmationUpdateOne) AddHourstime(i int) *ConfirmationUpdateOne {
	cuo.mutation.AddHourstime(i)
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *ConfirmationUpdateOne) SetOwnerID(id int) *ConfirmationUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *ConfirmationUpdateOne) SetNillableOwnerID(id *int) *ConfirmationUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *ConfirmationUpdateOne) SetOwner(u *User) *ConfirmationUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddBillIDs adds the bill edge to Bill by ids.
func (cuo *ConfirmationUpdateOne) AddBillIDs(ids ...int) *ConfirmationUpdateOne {
	cuo.mutation.AddBillIDs(ids...)
	return cuo
}

// AddBill adds the bill edges to Bill.
func (cuo *ConfirmationUpdateOne) AddBill(b ...*Bill) *ConfirmationUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBillIDs(ids...)
}

// Mutation returns the ConfirmationMutation object of the builder.
func (cuo *ConfirmationUpdateOne) Mutation() *ConfirmationMutation {
	return cuo.mutation
}

// ClearOwner clears the owner edge to User.
func (cuo *ConfirmationUpdateOne) ClearOwner() *ConfirmationUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// RemoveBillIDs removes the bill edge to Bill by ids.
func (cuo *ConfirmationUpdateOne) RemoveBillIDs(ids ...int) *ConfirmationUpdateOne {
	cuo.mutation.RemoveBillIDs(ids...)
	return cuo
}

// RemoveBill removes bill edges to Bill.
func (cuo *ConfirmationUpdateOne) RemoveBill(b ...*Bill) *ConfirmationUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBillIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *ConfirmationUpdateOne) Save(ctx context.Context) (*Confirmation, error) {

	var (
		err  error
		node *Confirmation
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfirmationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfirmationUpdateOne) SaveX(ctx context.Context) *Confirmation {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *ConfirmationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfirmationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConfirmationUpdateOne) sqlSave(ctx context.Context) (c *Confirmation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   confirmation.Table,
			Columns: confirmation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: confirmation.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Confirmation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Adddate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: confirmation.FieldAdddate,
		})
	}
	if value, ok := cuo.mutation.Bookingstart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: confirmation.FieldBookingstart,
		})
	}
	if value, ok := cuo.mutation.Bookingend(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: confirmation.FieldBookingend,
		})
	}
	if value, ok := cuo.mutation.Hourstime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: confirmation.FieldHourstime,
		})
	}
	if value, ok := cuo.mutation.AddedHourstime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: confirmation.FieldHourstime,
		})
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   confirmation.OwnerTable,
			Columns: []string{confirmation.OwnerColumn},
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
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   confirmation.OwnerTable,
			Columns: []string{confirmation.OwnerColumn},
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
	if nodes := cuo.mutation.RemovedBillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   confirmation.BillTable,
			Columns: []string{confirmation.BillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   confirmation.BillTable,
			Columns: []string{confirmation.BillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Confirmation{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{confirmation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
