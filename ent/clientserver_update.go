// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/CSUOS/rabums/ent/clientserver"
	"github.com/CSUOS/rabums/ent/event"
	"github.com/CSUOS/rabums/ent/predicate"
	"github.com/CSUOS/rabums/ent/user"
)

// ClientServerUpdate is the builder for updating ClientServer entities.
type ClientServerUpdate struct {
	config
	hooks    []Hook
	mutation *ClientServerMutation
}

// Where adds a new predicate for the ClientServerUpdate builder.
func (csu *ClientServerUpdate) Where(ps ...predicate.ClientServer) *ClientServerUpdate {
	csu.mutation.predicates = append(csu.mutation.predicates, ps...)
	return csu
}

// SetClientName sets the "client_name" field.
func (csu *ClientServerUpdate) SetClientName(s string) *ClientServerUpdate {
	csu.mutation.SetClientName(s)
	return csu
}

// SetToken sets the "token" field.
func (csu *ClientServerUpdate) SetToken(s string) *ClientServerUpdate {
	csu.mutation.SetToken(s)
	return csu
}

// SetLink sets the "link" field.
func (csu *ClientServerUpdate) SetLink(s string) *ClientServerUpdate {
	csu.mutation.SetLink(s)
	return csu
}

// SetDescription sets the "description" field.
func (csu *ClientServerUpdate) SetDescription(s string) *ClientServerUpdate {
	csu.mutation.SetDescription(s)
	return csu
}

// SetAvailable sets the "available" field.
func (csu *ClientServerUpdate) SetAvailable(b bool) *ClientServerUpdate {
	csu.mutation.SetAvailable(b)
	return csu
}

// SetCreatedAt sets the "created_at" field.
func (csu *ClientServerUpdate) SetCreatedAt(t time.Time) *ClientServerUpdate {
	csu.mutation.SetCreatedAt(t)
	return csu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (csu *ClientServerUpdate) SetNillableCreatedAt(t *time.Time) *ClientServerUpdate {
	if t != nil {
		csu.SetCreatedAt(*t)
	}
	return csu
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (csu *ClientServerUpdate) AddEventIDs(ids ...int) *ClientServerUpdate {
	csu.mutation.AddEventIDs(ids...)
	return csu
}

// AddEvents adds the "events" edges to the Event entity.
func (csu *ClientServerUpdate) AddEvents(e ...*Event) *ClientServerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csu.AddEventIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (csu *ClientServerUpdate) AddOwnerIDs(ids ...int) *ClientServerUpdate {
	csu.mutation.AddOwnerIDs(ids...)
	return csu
}

// AddOwner adds the "owner" edges to the User entity.
func (csu *ClientServerUpdate) AddOwner(u ...*User) *ClientServerUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csu.AddOwnerIDs(ids...)
}

// Mutation returns the ClientServerMutation object of the builder.
func (csu *ClientServerUpdate) Mutation() *ClientServerMutation {
	return csu.mutation
}

// ClearEvents clears all "events" edges to the Event entity.
func (csu *ClientServerUpdate) ClearEvents() *ClientServerUpdate {
	csu.mutation.ClearEvents()
	return csu
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (csu *ClientServerUpdate) RemoveEventIDs(ids ...int) *ClientServerUpdate {
	csu.mutation.RemoveEventIDs(ids...)
	return csu
}

// RemoveEvents removes "events" edges to Event entities.
func (csu *ClientServerUpdate) RemoveEvents(e ...*Event) *ClientServerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csu.RemoveEventIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (csu *ClientServerUpdate) ClearOwner() *ClientServerUpdate {
	csu.mutation.ClearOwner()
	return csu
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (csu *ClientServerUpdate) RemoveOwnerIDs(ids ...int) *ClientServerUpdate {
	csu.mutation.RemoveOwnerIDs(ids...)
	return csu
}

// RemoveOwner removes "owner" edges to User entities.
func (csu *ClientServerUpdate) RemoveOwner(u ...*User) *ClientServerUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csu.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (csu *ClientServerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csu.hooks) == 0 {
		affected, err = csu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClientServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csu.mutation = mutation
			affected, err = csu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csu.hooks) - 1; i >= 0; i-- {
			mut = csu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csu *ClientServerUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *ClientServerUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *ClientServerUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csu *ClientServerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clientserver.Table,
			Columns: clientserver.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clientserver.FieldID,
			},
		},
	}
	if ps := csu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.ClientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldClientName,
		})
	}
	if value, ok := csu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldToken,
		})
	}
	if value, ok := csu.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldLink,
		})
	}
	if value, ok := csu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldDescription,
		})
	}
	if value, ok := csu.mutation.Available(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: clientserver.FieldAvailable,
		})
	}
	if value, ok := csu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clientserver.FieldCreatedAt,
		})
	}
	if csu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   clientserver.EventsTable,
			Columns: []string{clientserver.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.RemovedEventsIDs(); len(nodes) > 0 && !csu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   clientserver.EventsTable,
			Columns: []string{clientserver.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   clientserver.EventsTable,
			Columns: []string{clientserver.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if csu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   clientserver.OwnerTable,
			Columns: clientserver.OwnerPrimaryKey,
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
	if nodes := csu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !csu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   clientserver.OwnerTable,
			Columns: clientserver.OwnerPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   clientserver.OwnerTable,
			Columns: clientserver.OwnerPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clientserver.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ClientServerUpdateOne is the builder for updating a single ClientServer entity.
type ClientServerUpdateOne struct {
	config
	hooks    []Hook
	mutation *ClientServerMutation
}

// SetClientName sets the "client_name" field.
func (csuo *ClientServerUpdateOne) SetClientName(s string) *ClientServerUpdateOne {
	csuo.mutation.SetClientName(s)
	return csuo
}

// SetToken sets the "token" field.
func (csuo *ClientServerUpdateOne) SetToken(s string) *ClientServerUpdateOne {
	csuo.mutation.SetToken(s)
	return csuo
}

// SetLink sets the "link" field.
func (csuo *ClientServerUpdateOne) SetLink(s string) *ClientServerUpdateOne {
	csuo.mutation.SetLink(s)
	return csuo
}

// SetDescription sets the "description" field.
func (csuo *ClientServerUpdateOne) SetDescription(s string) *ClientServerUpdateOne {
	csuo.mutation.SetDescription(s)
	return csuo
}

// SetAvailable sets the "available" field.
func (csuo *ClientServerUpdateOne) SetAvailable(b bool) *ClientServerUpdateOne {
	csuo.mutation.SetAvailable(b)
	return csuo
}

// SetCreatedAt sets the "created_at" field.
func (csuo *ClientServerUpdateOne) SetCreatedAt(t time.Time) *ClientServerUpdateOne {
	csuo.mutation.SetCreatedAt(t)
	return csuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (csuo *ClientServerUpdateOne) SetNillableCreatedAt(t *time.Time) *ClientServerUpdateOne {
	if t != nil {
		csuo.SetCreatedAt(*t)
	}
	return csuo
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (csuo *ClientServerUpdateOne) AddEventIDs(ids ...int) *ClientServerUpdateOne {
	csuo.mutation.AddEventIDs(ids...)
	return csuo
}

// AddEvents adds the "events" edges to the Event entity.
func (csuo *ClientServerUpdateOne) AddEvents(e ...*Event) *ClientServerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csuo.AddEventIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (csuo *ClientServerUpdateOne) AddOwnerIDs(ids ...int) *ClientServerUpdateOne {
	csuo.mutation.AddOwnerIDs(ids...)
	return csuo
}

// AddOwner adds the "owner" edges to the User entity.
func (csuo *ClientServerUpdateOne) AddOwner(u ...*User) *ClientServerUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csuo.AddOwnerIDs(ids...)
}

// Mutation returns the ClientServerMutation object of the builder.
func (csuo *ClientServerUpdateOne) Mutation() *ClientServerMutation {
	return csuo.mutation
}

// ClearEvents clears all "events" edges to the Event entity.
func (csuo *ClientServerUpdateOne) ClearEvents() *ClientServerUpdateOne {
	csuo.mutation.ClearEvents()
	return csuo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (csuo *ClientServerUpdateOne) RemoveEventIDs(ids ...int) *ClientServerUpdateOne {
	csuo.mutation.RemoveEventIDs(ids...)
	return csuo
}

// RemoveEvents removes "events" edges to Event entities.
func (csuo *ClientServerUpdateOne) RemoveEvents(e ...*Event) *ClientServerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csuo.RemoveEventIDs(ids...)
}

// ClearOwner clears all "owner" edges to the User entity.
func (csuo *ClientServerUpdateOne) ClearOwner() *ClientServerUpdateOne {
	csuo.mutation.ClearOwner()
	return csuo
}

// RemoveOwnerIDs removes the "owner" edge to User entities by IDs.
func (csuo *ClientServerUpdateOne) RemoveOwnerIDs(ids ...int) *ClientServerUpdateOne {
	csuo.mutation.RemoveOwnerIDs(ids...)
	return csuo
}

// RemoveOwner removes "owner" edges to User entities.
func (csuo *ClientServerUpdateOne) RemoveOwner(u ...*User) *ClientServerUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csuo.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the updated ClientServer entity.
func (csuo *ClientServerUpdateOne) Save(ctx context.Context) (*ClientServer, error) {
	var (
		err  error
		node *ClientServer
	)
	if len(csuo.hooks) == 0 {
		node, err = csuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClientServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csuo.mutation = mutation
			node, err = csuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csuo.hooks) - 1; i >= 0; i-- {
			mut = csuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *ClientServerUpdateOne) SaveX(ctx context.Context) *ClientServer {
	node, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (csuo *ClientServerUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *ClientServerUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csuo *ClientServerUpdateOne) sqlSave(ctx context.Context) (_node *ClientServer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clientserver.Table,
			Columns: clientserver.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clientserver.FieldID,
			},
		},
	}
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ClientServer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := csuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csuo.mutation.ClientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldClientName,
		})
	}
	if value, ok := csuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldToken,
		})
	}
	if value, ok := csuo.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldLink,
		})
	}
	if value, ok := csuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clientserver.FieldDescription,
		})
	}
	if value, ok := csuo.mutation.Available(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: clientserver.FieldAvailable,
		})
	}
	if value, ok := csuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clientserver.FieldCreatedAt,
		})
	}
	if csuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   clientserver.EventsTable,
			Columns: []string{clientserver.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !csuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   clientserver.EventsTable,
			Columns: []string{clientserver.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   clientserver.EventsTable,
			Columns: []string{clientserver.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if csuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   clientserver.OwnerTable,
			Columns: clientserver.OwnerPrimaryKey,
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
	if nodes := csuo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !csuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   clientserver.OwnerTable,
			Columns: clientserver.OwnerPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   clientserver.OwnerTable,
			Columns: clientserver.OwnerPrimaryKey,
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
	_node = &ClientServer{config: csuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clientserver.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
