package schema

import (
	"entgo.io/ent"
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
		field.String("dataName"),
		field.String("discription"),
		field.String("type"),
	}
}

// Edges of the EndpointData.
func (EndpointData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("endpointRequired", ProviderEndpoint.Type).Ref("requiredData").Unique(),
		edge.From("endpointProvided", ProviderEndpoint.Type).Ref("providedData").Unique(),
	}
}
