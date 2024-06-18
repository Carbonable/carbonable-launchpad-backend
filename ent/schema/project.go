package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type (
	PaymentToken struct {
		Symbol  string `json:"symbol"`
		Address string `json:"address"`
	}
	Milestone struct {
		Boost    string `json:"boost"`
		Area     string `json:"ha"`
		TonPrice string `json:"ton"`
		Ceil     int    `json:"ceil"`
		Id       int    `json:"id"`
	}
	Metadata struct {
		Rating     string      `json:"rating"`
		TonPrice   string      `json:"ton_price"`
		Milestones []Milestone `json:"milestones"`
	}
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("address"),
		field.Int("slot"),
		field.String("name"),
		field.String("slug"),
		field.Int("value_decimal"),
		field.String("forecasted_apr").Optional(),
		field.String("total_value"),
		field.JSON("payment_token", PaymentToken{}),
		field.JSON("metadata", Metadata{}),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mint", Mint.Type).Unique(),
		edge.To("launchpad", Launchpad.Type).Unique(),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address", "slot").Unique(),
	}
}
