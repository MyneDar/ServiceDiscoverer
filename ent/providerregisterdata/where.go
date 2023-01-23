// Code generated by ent, DO NOT EDIT.

package providerregisterdata

import (
	"servicediscoverer/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldName, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldPort, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldAddress, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldDescription, v))
}

// LiveInterval applies equality check predicate on the "liveInterval" field. It's identical to LiveIntervalEQ.
func LiveInterval(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldLiveInterval, v))
}

// LiveTimeout applies equality check predicate on the "liveTimeout" field. It's identical to LiveTimeoutEQ.
func LiveTimeout(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldLiveTimeout, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContainsFold(FieldName, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldPort, v))
}

// PortContains applies the Contains predicate on the "port" field.
func PortContains(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContains(FieldPort, v))
}

// PortHasPrefix applies the HasPrefix predicate on the "port" field.
func PortHasPrefix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasPrefix(FieldPort, v))
}

// PortHasSuffix applies the HasSuffix predicate on the "port" field.
func PortHasSuffix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasSuffix(FieldPort, v))
}

// PortEqualFold applies the EqualFold predicate on the "port" field.
func PortEqualFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEqualFold(FieldPort, v))
}

// PortContainsFold applies the ContainsFold predicate on the "port" field.
func PortContainsFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContainsFold(FieldPort, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContainsFold(FieldAddress, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldContainsFold(FieldDescription, v))
}

// LiveIntervalEQ applies the EQ predicate on the "liveInterval" field.
func LiveIntervalEQ(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldLiveInterval, v))
}

// LiveIntervalNEQ applies the NEQ predicate on the "liveInterval" field.
func LiveIntervalNEQ(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldLiveInterval, v))
}

// LiveIntervalIn applies the In predicate on the "liveInterval" field.
func LiveIntervalIn(vs ...int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldLiveInterval, vs...))
}

// LiveIntervalNotIn applies the NotIn predicate on the "liveInterval" field.
func LiveIntervalNotIn(vs ...int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldLiveInterval, vs...))
}

// LiveIntervalGT applies the GT predicate on the "liveInterval" field.
func LiveIntervalGT(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldLiveInterval, v))
}

// LiveIntervalGTE applies the GTE predicate on the "liveInterval" field.
func LiveIntervalGTE(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldLiveInterval, v))
}

// LiveIntervalLT applies the LT predicate on the "liveInterval" field.
func LiveIntervalLT(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldLiveInterval, v))
}

// LiveIntervalLTE applies the LTE predicate on the "liveInterval" field.
func LiveIntervalLTE(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldLiveInterval, v))
}

// LiveTimeoutEQ applies the EQ predicate on the "liveTimeout" field.
func LiveTimeoutEQ(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldEQ(FieldLiveTimeout, v))
}

// LiveTimeoutNEQ applies the NEQ predicate on the "liveTimeout" field.
func LiveTimeoutNEQ(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNEQ(FieldLiveTimeout, v))
}

// LiveTimeoutIn applies the In predicate on the "liveTimeout" field.
func LiveTimeoutIn(vs ...int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldIn(FieldLiveTimeout, vs...))
}

// LiveTimeoutNotIn applies the NotIn predicate on the "liveTimeout" field.
func LiveTimeoutNotIn(vs ...int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldNotIn(FieldLiveTimeout, vs...))
}

// LiveTimeoutGT applies the GT predicate on the "liveTimeout" field.
func LiveTimeoutGT(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGT(FieldLiveTimeout, v))
}

// LiveTimeoutGTE applies the GTE predicate on the "liveTimeout" field.
func LiveTimeoutGTE(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldGTE(FieldLiveTimeout, v))
}

// LiveTimeoutLT applies the LT predicate on the "liveTimeout" field.
func LiveTimeoutLT(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLT(FieldLiveTimeout, v))
}

// LiveTimeoutLTE applies the LTE predicate on the "liveTimeout" field.
func LiveTimeoutLTE(v int) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(sql.FieldLTE(FieldLiveTimeout, v))
}

// HasEndpoints applies the HasEdge predicate on the "endpoints" edge.
func HasEndpoints() predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EndpointsTable, EndpointsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEndpointsWith applies the HasEdge predicate on the "endpoints" edge with a given conditions (other predicates).
func HasEndpointsWith(preds ...predicate.ProviderEndpoint) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EndpointsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EndpointsTable, EndpointsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProviderRegisterData) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProviderRegisterData) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(func(s *sql.Selector) {
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
func Not(p predicate.ProviderRegisterData) predicate.ProviderRegisterData {
	return predicate.ProviderRegisterData(func(s *sql.Selector) {
		p(s.Not())
	})
}
