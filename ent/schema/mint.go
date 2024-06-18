package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mint holds the schema definition for the Mint entity.
type Mint struct {
	ent.Schema
}

// Fields of the Mint.
func (Mint) Fields() []ent.Field {
	return []ent.Field{
		field.String("min_value_per_tx"),
		field.String("max_value_per_tx"),
		field.String("minter_address"),
	}
}

// Edges of the Mint.
func (Mint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("mint").Required().Unique(),
	}
}
