package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProviderRegisterData holds the schema definition for the ProviderRegisterData entity.
type ProviderRegisterData struct {
	ent.Schema
}

/*ID       string    `json:"id"`
Name     string    `json:"name"`
Port     string    `json:"port"`
Address  string    `json:"address"`
check    checkLive `json:"check"`
testCall string    `json:"testCall"`*/
// Fields of the ProviderRegisterData.
func (ProviderRegisterData) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"-"`),
		field.String("name"),
		field.String("port").StructTag(`json:"-"`),
		field.String("address").StructTag(`json:"-"`),
		field.String("description"),
		field.Int("liveInterval").StructTag(`json:"-"`),
		field.Int("liveTimeout").StructTag(`json:"-"`),
	}
}

// Edges of the ProviderRegisterData.
func (ProviderRegisterData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("endpoints", ProviderEndpoint.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

func (ProviderRegisterData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{StructTag: `json:"connections"`},
	}
}
