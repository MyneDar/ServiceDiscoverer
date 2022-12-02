package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProviderEndpoint holds the schema definition for the ProviderEndpoint entity.
type ProviderEndpoint struct {
	ent.Schema
}

// Fields of the ProviderEndpoint.
func (ProviderEndpoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`),
		field.String("name"),
		field.String("path"),
		field.String("type"),
	}
}

// Edges of the ProviderEndpoint.
func (ProviderEndpoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("required_data", EndpointData.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("provided_data", EndpointData.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.From("provider", ProviderRegisterData.Type).Ref("endpoints").Unique(),
	}
}
func (ProviderEndpoint) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{StructTag: `json:"data"`},
	}
}
