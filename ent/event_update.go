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

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where adds a new predicate for the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.predicates = append(eu.mutation.predicates, ps...)
	return eu
}

// SetEvent sets the "event" field.
func (eu *EventUpdate) SetEvent(s string) *EventUpdate {
	eu.mutation.SetEvent(s)
	return eu
}

// SetMessage sets the "message" field.
func (eu *EventUpdate) SetMessage(s string) *EventUpdate {
	eu.mutation.SetMessage(s)
	return eu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (eu *EventUpdate) SetNillableMessage(s *string) *EventUpdate {
	if s != nil {
		eu.SetMessage(*s)
	}
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EventUpdate) SetCreatedAt(t time.Time) *EventUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreatedAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (eu *EventUpdate) SetUserID(id int) *EventUpdate {
	eu.mutation.SetUserID(id)
	return eu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableUserID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetUserID(*id)
	}
	return eu
}

// SetUser sets the "user" edge to the User entity.
func (eu *EventUpdate) SetUser(u *User) *EventUpdate {
	return eu.SetUserID(u.ID)
}

// SetClientserverID sets the "clientserver" edge to the ClientServer entity by ID.
func (eu *EventUpdate) SetClientserverID(id int) *EventUpdate {
	eu.mutation.SetClientserverID(id)
	return eu
}

// SetNillableClientserverID sets the "clientserver" edge to the ClientServer entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableClientserverID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetClientserverID(*id)
	}
	return eu
}

// SetClientserver sets the "clientserver" edge to the ClientServer entity.
func (eu *EventUpdate) SetClientserver(c *ClientServer) *EventUpdate {
	return eu.SetClientserverID(c.ID)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (eu *EventUpdate) ClearUser() *EventUpdate {
	eu.mutation.ClearUser()
	return eu
}

// ClearClientserver clears the "clientserver" edge to the ClientServer entity.
func (eu *EventUpdate) ClearClientserver() *EventUpdate {
	eu.mutation.ClearClientserver()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEvent,
		})
	}
	if value, ok := eu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldMessage,
		})
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if eu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.UserTable,
			Columns: []string{event.UserColumn},
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
	if nodes := eu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.UserTable,
			Columns: []string{event.UserColumn},
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
	if eu.mutation.ClientserverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.ClientserverTable,
			Columns: []string{event.ClientserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clientserver.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ClientserverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.ClientserverTable,
			Columns: []string{event.ClientserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clientserver.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// SetEvent sets the "event" field.
func (euo *EventUpdateOne) SetEvent(s string) *EventUpdateOne {
	euo.mutation.SetEvent(s)
	return euo
}

// SetMessage sets the "message" field.
func (euo *EventUpdateOne) SetMessage(s string) *EventUpdateOne {
	euo.mutation.SetMessage(s)
	return euo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableMessage(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetMessage(*s)
	}
	return euo
}

// SetCreatedAt sets the "created_at" field.
func (euo *EventUpdateOne) SetCreatedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreatedAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (euo *EventUpdateOne) SetUserID(id int) *EventUpdateOne {
	euo.mutation.SetUserID(id)
	return euo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableUserID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetUserID(*id)
	}
	return euo
}

// SetUser sets the "user" edge to the User entity.
func (euo *EventUpdateOne) SetUser(u *User) *EventUpdateOne {
	return euo.SetUserID(u.ID)
}

// SetClientserverID sets the "clientserver" edge to the ClientServer entity by ID.
func (euo *EventUpdateOne) SetClientserverID(id int) *EventUpdateOne {
	euo.mutation.SetClientserverID(id)
	return euo
}

// SetNillableClientserverID sets the "clientserver" edge to the ClientServer entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableClientserverID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetClientserverID(*id)
	}
	return euo
}

// SetClientserver sets the "clientserver" edge to the ClientServer entity.
func (euo *EventUpdateOne) SetClientserver(c *ClientServer) *EventUpdateOne {
	return euo.SetClientserverID(c.ID)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (euo *EventUpdateOne) ClearUser() *EventUpdateOne {
	euo.mutation.ClearUser()
	return euo
}

// ClearClientserver clears the "clientserver" edge to the ClientServer entity.
func (euo *EventUpdateOne) ClearClientserver() *EventUpdateOne {
	euo.mutation.ClearClientserver()
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Event.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEvent,
		})
	}
	if value, ok := euo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldMessage,
		})
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldCreatedAt,
		})
	}
	if euo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.UserTable,
			Columns: []string{event.UserColumn},
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
	if nodes := euo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.UserTable,
			Columns: []string{event.UserColumn},
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
	if euo.mutation.ClientserverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.ClientserverTable,
			Columns: []string{event.ClientserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clientserver.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ClientserverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.ClientserverTable,
			Columns: []string{event.ClientserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clientserver.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
