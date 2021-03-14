package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ClientServer struct {
	ent.Schema
}

// Fields of the User.
func (ClientServer) Fields() []ent.Field {
	return []ent.Field{
		field.String("client_name"),
		field.String("token"),
		field.String("link"),
		field.Text("description"),
		field.Bool("available"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (ClientServer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("clientserver"),
		edge.To("owner", User.Type),
	}
}
