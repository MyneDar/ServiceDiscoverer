package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EndpointData holds the schema definition for the EndpointData entity.
type EndpointData struct {
	ent.Schema
}

// Fields of the EndpointData.
func (EndpointData) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`),
		field.String("dataName"),
		field.String("description"),
		field.String("type"),
	}
}

// Edges of the EndpointData.
func (EndpointData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("endpoint_required", ProviderEndpoint.Type).Ref("required_data").Unique(),
		edge.From("endpoint_provided", ProviderEndpoint.Type).Ref("provided_data").Unique(),
	}
}

func (EndpointData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{StructTag: `json:"-"`},
	}
}
