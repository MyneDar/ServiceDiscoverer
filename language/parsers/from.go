package parsers

import (
	_ "github.com/lib/pq"
	"servicediscoverer/dev"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"
	"servicediscoverer/models"
)

// FromParser is a parser for FROM keyword and the following serviceName.endpoint arguments after it.
type FromParser struct{}

// NewFromParser is a function that returns a pointer to a new FromParser struct
func NewFromParser() *FromParser {
	return &FromParser{}
}

// Process is a function that checks process the first
func (l *FromParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {

	// Get the service from the db by its name
	service, err := dev.LocalClient.ProviderRegisterData.
		Query().
		Where(providerregisterdata.Name(tok[0].Data)).
		Only(dev.Ctx)
	if err != nil {
		return err
	}

	// Get the endpoint from the db by its name
	endpoints, err := service.QueryEndpoints().Where(providerendpoint.Name(tok[1].Data)).All(dev.Ctx)
	if err != nil {
		return err
	}

	var dataMap map[string]string
	// Create an url for the call
	for _, value := range endpoints {
		dataMap[value.Type] = service.Address + ":" + service.Port + "/" + value.Path
	}

	//put the url in the information map
	information["From_Data"] = dataMap

	return err
}
