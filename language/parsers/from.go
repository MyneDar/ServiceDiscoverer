package parsers

import (
	"context"
	_ "github.com/lib/pq"
	"servicediscoverer/dev"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"
	"servicediscoverer/models"
)

type FromParser struct {
}

func NewFromParser() *FromParser {
	return &FromParser{}
}

func (l *FromParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	service, err := dev.LocalClient.ProviderRegisterData.
		Query().
		Where(providerregisterdata.Name(tok[0].Data)).
		Only(context.Background())
	if err != nil {
		return err
	}
	endpoints, err := service.QueryEndpoints().Where(providerendpoint.Name(tok[1].Data)).All(context.Background())
	if err != nil {
		return err
	}

	var dataMap map[string]string
	for _, value := range endpoints {
		dataMap[value.Type] = service.Address + ":" + service.Port + "/" + value.Path
	}

	information["From_Data"] = dataMap

	return err
}
