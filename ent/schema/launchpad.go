package schema

import (
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MinterContract struct {
	Address string          `json:"address"`
	Abi     json.RawMessage `json:"abi"`
}

// Launchpad holds the schema definition for the Launchpad entity.
type Launchpad struct {
	ent.Schema
}

// Fields of the Launchpad.
func (Launchpad) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_ready").Default(false),
		field.JSON("minter_contract", MinterContract{}),
		field.Bool("whitelisted_sale_open").Default(false),
		field.Bool("public_sale_open").Default(false),
		field.Bool("is_sold_out").Default(false),
		field.Bool("is_canceled").Default(false),
	}
}

// Edges of the Launchpad.
func (Launchpad) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("launchpad").Required().Unique(),
	}
}
