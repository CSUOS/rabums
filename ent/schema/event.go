package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Event struct {
	ent.Schema
}

// Fields of the User.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("event"),
		field.String("message").Default(""),

		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
		edge.To("clientserver", ClientServer.Type).Unique(),
	}
}
