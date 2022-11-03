package interfaces

import (
	"servicediscoverer/ent"
)

type provider interface {
	CreateRegisterData() ent.ProviderRegisterData
}
