package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return append([]ent.Field{
		field.Enum("type").Values("TEST", "GROUP").Default("TEST"),
		field.String("type_id"),
		field.String("order_type_id").Optional(),
		field.Bool("is_orderable").Default(true),
		field.String("name").Optional(),
		field.String("display_name").Optional(),
		field.String("unique_emr_code").Optional(),
		field.String("tax_code").Optional(),
		field.Text("weblink").Optional(),
		field.Text("description").Optional(),
		field.Strings("sample_type_list").Optional(),
		field.Strings("subpackages_list").Optional(),
		field.Strings("subtests_list").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	})
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("item_price", ItemPrice.Type),
		//edge.To("test", Item.Type).
		//	Annotations(entsql.Annotation{
		//		OnDelete: entsql.Cascade,
		//	}).From("group").Unique(), //enforces a cascading delete when the related entity is deleted
	}
}

// Indexes of the Item.
func (Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "type_id"),
		index.Fields("type_id"),
		index.Fields("order_type_id"),
	}
}
