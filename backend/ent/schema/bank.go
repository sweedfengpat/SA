package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	ent.Schema
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.String("Bank").Unique(),
	}
}

// Edges of the Bank.
func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bill", Bill.Type).StorageKey(edge.Column("bank_id")),
	}
}
