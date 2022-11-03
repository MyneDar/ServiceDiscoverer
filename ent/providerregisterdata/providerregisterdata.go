// Code generated by ent, DO NOT EDIT.

package providerregisterdata

const (
	// Label holds the string label denoting the providerregisterdata type in the database.
	Label = "provider_register_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLiveInterval holds the string denoting the liveinterval field in the database.
	FieldLiveInterval = "live_interval"
	// FieldLiveTimeout holds the string denoting the livetimeout field in the database.
	FieldLiveTimeout = "live_timeout"
	// EdgeEndpoints holds the string denoting the endpoints edge name in mutations.
	EdgeEndpoints = "endpoints"
	// Table holds the table name of the providerregisterdata in the database.
	Table = "provider_register_data"
	// EndpointsTable is the table that holds the endpoints relation/edge.
	EndpointsTable = "provider_endpoints"
	// EndpointsInverseTable is the table name for the ProviderEndpoint entity.
	// It exists in this package in order to avoid circular dependency with the "providerendpoint" package.
	EndpointsInverseTable = "provider_endpoints"
	// EndpointsColumn is the table column denoting the endpoints relation/edge.
	EndpointsColumn = "provider_register_data_endpoints"
)

// Columns holds all SQL columns for providerregisterdata fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPort,
	FieldAddress,
	FieldDescription,
	FieldLiveInterval,
	FieldLiveTimeout,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
