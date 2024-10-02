package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Merchandise struct {
	ent.Schema
}

// Fields of the Merchandise.
func (Merchandise) Fields() []ent.Field {
	return append([]ent.Field{
		field.Enum("type").Values("item", "service").Default("item"),
		field.String("type_id").Optional(),
		field.Int("stock").Default(-1),
		field.Int("pending_period").Default(-1),
		field.String("name").Optional(),
		field.Text("display_name").Optional(),
		field.Text("picture_url").Optional(),
		field.Text("description").Optional(),
		field.Float("price").Optional(),
		field.Float("refundable_ratio").Optional(),
		field.Float("deduction_ratio").Optional(),
		field.Text("deduction_reason").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	})
}

// Edges of the Merchandise.
func (Merchandise) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("provider", Provider.Type).Ref("merchandise").Unique(),
	}
}

// Indexes of the Merchandise.
func (Merchandise) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "type_id"),
		index.Fields("type_id"),
	}
}
