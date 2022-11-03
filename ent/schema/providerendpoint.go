package schema

import (
	"entgo.io/ent"
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
		field.String("name"),
		field.String("path"),
		field.String("type"),
	}
}

// Edges of the ProviderEndpoint.
func (ProviderEndpoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requiredData", EndpointData.Type),
		edge.To("providedData", EndpointData.Type),
		edge.From("provider", ProviderRegisterData.Type).Ref("endpoints").Unique(),
	}
}
