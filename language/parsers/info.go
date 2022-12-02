package parsers

import (
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
	var err error
	var provider *ent.ProviderRegisterData
	switch tok[1].Data {
	case models.ASTERISK.String():
		provider, err = dev.LocalClient.ProviderRegisterData.
			Query().
			Where(providerregisterdata.Name(tok[0].Data)).
			WithEndpoints(
				func(query *ent.ProviderEndpointQuery) {
					query.WithProvidedData()
					query.WithRequiredData()
					query.All(dev.Ctx)
				},
			).
			Only(dev.Ctx)
	default:
		provider, err = dev.LocalClient.ProviderRegisterData.
			Query().
			Where(providerregisterdata.Name(tok[0].Data)).
			WithEndpoints(
				func(query *ent.ProviderEndpointQuery) {
					query.Where(providerendpoint.Name(tok[1].Data))
					query.WithProvidedData()
					query.WithRequiredData()
					query.All(dev.Ctx)
				},
			).
			Only(dev.Ctx)
	}
	if err != nil {
		return err
	}
	information["Info_Data"] = provider
	return err
}
