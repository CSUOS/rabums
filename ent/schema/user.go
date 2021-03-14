package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique(),
		field.String("user_pw"),

		field.String("user_name"),
		field.Int("user_number"),
		field.String("email"),

		field.Time("created_at").Default(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("user"),
		edge.From("owns", ClientServer.Type).Ref("owner"),
	}
}
