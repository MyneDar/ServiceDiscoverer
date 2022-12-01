package models

import "servicediscoverer/ent"

type ServiceInfoDTO struct {
	Name        string
	Description string
	Endpoints   []*ent.ProviderEndpoint
}
