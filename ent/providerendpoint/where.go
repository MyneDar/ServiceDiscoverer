// Code generated by ent, DO NOT EDIT.

package providerendpoint

import (
	"servicediscoverer/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldName, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldPath, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldType, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldDescription, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContainsFold(FieldName, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContainsFold(FieldPath, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContainsFold(FieldType, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(sql.FieldContainsFold(FieldDescription, v))
}

// HasRequiredData applies the HasEdge predicate on the "required_data" edge.
func HasRequiredData() predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RequiredDataTable, RequiredDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRequiredDataWith applies the HasEdge predicate on the "required_data" edge with a given conditions (other predicates).
func HasRequiredDataWith(preds ...predicate.EndpointData) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RequiredDataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RequiredDataTable, RequiredDataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProvidedData applies the HasEdge predicate on the "provided_data" edge.
func HasProvidedData() predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvidedDataTable, ProvidedDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvidedDataWith applies the HasEdge predicate on the "provided_data" edge with a given conditions (other predicates).
func HasProvidedDataWith(preds ...predicate.EndpointData) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvidedDataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvidedDataTable, ProvidedDataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.ProviderRegisterData) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProviderEndpoint) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProviderEndpoint) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProviderEndpoint) predicate.ProviderEndpoint {
	return predicate.ProviderEndpoint(func(s *sql.Selector) {
		p(s.Not())
	})
}
