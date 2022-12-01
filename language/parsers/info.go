package parsers

import (
	"encoding/json"
	"servicediscoverer/dev"
	"servicediscoverer/ent"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"
	"servicediscoverer/models"
)

type InfoParser struct {
}

func NewInfoParser() *InfoParser {
	return &InfoParser{}
}

func (l *InfoParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	service, err := dev.LocalClient.ProviderRegisterData.
		Query().
		Where(providerregisterdata.Name(tok[0].Data)).
		Only(dev.Ctx)
	if err != nil {
		return err
	}
	var endpoints []*ent.ProviderEndpoint
	switch tok[1].Data {
	case "*":
		endpoints, err = service.QueryEndpoints().All(dev.Ctx)
	default:
		endpoints, err = service.QueryEndpoints().Where(providerendpoint.Name(tok[1].Data)).All(dev.Ctx)
	}
	if err != nil {
		return err
	}
	data, errJs := json.Marshal(endpoints)
	if errJs != nil {
		return err
	}
	information["Info_Data"] = data
	return err
}
