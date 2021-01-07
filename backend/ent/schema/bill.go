package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Pay_number").Positive().Unique(),
		field.Int("Pay_total").Positive(),
		field.Time("Pay_time"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bank", Bank.Type).Ref("bill").Unique(),
		edge.From("confirmation", Confirmation.Type).Ref("bill").Unique(),
		edge.From("owner", User.Type).Ref("bill").Unique(),
	}
}
