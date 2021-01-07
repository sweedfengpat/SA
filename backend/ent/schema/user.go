package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.Time("birth"),
		field.Int("tel"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("confirmation", Confirmation.Type).StorageKey(edge.Column("owner_id")),
		edge.To("bill", Bill.Type).StorageKey(edge.Column("owner_id")),
	}
}
