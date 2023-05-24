package parsers

import (
	"servicediscoverer/dev"
	"servicediscoverer/ent"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"
	"servicediscoverer/models"
)

// InfoParser is a parser for info keyword and for the arguments that shows us what information needed to be returned
type InfoParser struct{}

// NewInfoParser is a function that returns a pointer to a new InfoParser struct
func NewInfoParser() *InfoParser {
	return &InfoParser{}
}

// Process is a function that checks what information should be returned
func (l *InfoParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	// initialize the variables
	var err error
	var provider interface{}
	var temp *ent.ProviderRegisterData

	// Check if the token/s length is only 1 or not
	if len(tok) == 1 {
		// If it is
		provider, err = dev.LocalClient.ProviderRegisterData.
			Query().
			WithEndpoints(
				func(query *ent.ProviderEndpointQuery) {
					query.WithProvidedData()
					query.WithRequiredData()
					_, err2 := query.All(dev.Ctx)
					if err2 != nil {
						err = err2
					}
				},
			).All(dev.Ctx)
	} else {
		switch tok[1].Data {
		case models.ASTERISK.String():
			provider, err = dev.LocalClient.ProviderRegisterData.
				Query().
				Where(providerregisterdata.Name(tok[0].Data)).
				WithEndpoints(
					func(query *ent.ProviderEndpointQuery) {
						query.WithProvidedData()
						query.WithRequiredData()
						_, err2 := query.All(dev.Ctx)
						if err2 != nil {
							err = err2
						}
					},
				).
				Only(dev.Ctx)
		default:
			temp, err = dev.LocalClient.ProviderRegisterData.
				Query().
				Where(providerregisterdata.Name(tok[0].Data)).
				WithEndpoints(
					func(query *ent.ProviderEndpointQuery) {
						query.Where(providerendpoint.Name(tok[1].Data))
						query.WithProvidedData()
						query.WithRequiredData()
						_, err2 := query.All(dev.Ctx)
						if err2 != nil {
							err = err2
						}
					},
				).
				Only(dev.Ctx)
			if err != nil {
				return err
			}
			provider = temp.Edges.Endpoints[0]
		}
	}

	// If the error is not nil, return it
	if err != nil {
		return err
	}

	// Put the provider in the information map
	information["Info_Data"] = provider
	return err
}
