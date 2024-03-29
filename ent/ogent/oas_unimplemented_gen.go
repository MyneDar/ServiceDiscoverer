// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateEndpointData implements createEndpointData operation.
//
// Creates a new EndpointData and persists it to storage.
//
// POST /admin/endpoint-data
func (UnimplementedHandler) CreateEndpointData(ctx context.Context, req *CreateEndpointDataReq) (r CreateEndpointDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateProviderEndpoint implements createProviderEndpoint operation.
//
// Creates a new ProviderEndpoint and persists it to storage.
//
// POST /admin/provider-endpoints
func (UnimplementedHandler) CreateProviderEndpoint(ctx context.Context, req *CreateProviderEndpointReq) (r CreateProviderEndpointRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateProviderRegisterData implements createProviderRegisterData operation.
//
// Creates a new ProviderRegisterData and persists it to storage.
//
// POST /admin/provider-register-data
func (UnimplementedHandler) CreateProviderRegisterData(ctx context.Context, req *CreateProviderRegisterDataReq) (r CreateProviderRegisterDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteEndpointData implements deleteEndpointData operation.
//
// Deletes the EndpointData with the requested ID.
//
// DELETE /admin/endpoint-data/{id}
func (UnimplementedHandler) DeleteEndpointData(ctx context.Context, params DeleteEndpointDataParams) (r DeleteEndpointDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteProviderEndpoint implements deleteProviderEndpoint operation.
//
// Deletes the ProviderEndpoint with the requested ID.
//
// DELETE /admin/provider-endpoints/{id}
func (UnimplementedHandler) DeleteProviderEndpoint(ctx context.Context, params DeleteProviderEndpointParams) (r DeleteProviderEndpointRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteProviderRegisterData implements deleteProviderRegisterData operation.
//
// Deletes the ProviderRegisterData with the requested ID.
//
// DELETE /admin/provider-register-data/{id}
func (UnimplementedHandler) DeleteProviderRegisterData(ctx context.Context, params DeleteProviderRegisterDataParams) (r DeleteProviderRegisterDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListEndpointData implements listEndpointData operation.
//
// List EndpointData.
//
// GET /admin/endpoint-data
func (UnimplementedHandler) ListEndpointData(ctx context.Context, params ListEndpointDataParams) (r ListEndpointDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProviderEndpoint implements listProviderEndpoint operation.
//
// List ProviderEndpoints.
//
// GET /admin/provider-endpoints
func (UnimplementedHandler) ListProviderEndpoint(ctx context.Context, params ListProviderEndpointParams) (r ListProviderEndpointRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProviderEndpointProvidedData implements listProviderEndpointProvidedData operation.
//
// List attached ProvidedData.
//
// GET /admin/provider-endpoints/{id}/provided-data
func (UnimplementedHandler) ListProviderEndpointProvidedData(ctx context.Context, params ListProviderEndpointProvidedDataParams) (r ListProviderEndpointProvidedDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProviderEndpointRequiredData implements listProviderEndpointRequiredData operation.
//
// List attached RequiredData.
//
// GET /admin/provider-endpoints/{id}/required-data
func (UnimplementedHandler) ListProviderEndpointRequiredData(ctx context.Context, params ListProviderEndpointRequiredDataParams) (r ListProviderEndpointRequiredDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProviderRegisterData implements listProviderRegisterData operation.
//
// List ProviderRegisterData.
//
// GET /admin/provider-register-data
func (UnimplementedHandler) ListProviderRegisterData(ctx context.Context, params ListProviderRegisterDataParams) (r ListProviderRegisterDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProviderRegisterDataEndpoints implements listProviderRegisterDataEndpoints operation.
//
// List attached Endpoints.
//
// GET /admin/provider-register-data/{id}/endpoints
func (UnimplementedHandler) ListProviderRegisterDataEndpoints(ctx context.Context, params ListProviderRegisterDataEndpointsParams) (r ListProviderRegisterDataEndpointsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEndpointData implements readEndpointData operation.
//
// Finds the EndpointData with the requested ID and returns it.
//
// GET /admin/endpoint-data/{id}
func (UnimplementedHandler) ReadEndpointData(ctx context.Context, params ReadEndpointDataParams) (r ReadEndpointDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEndpointDataEndpointProvided implements readEndpointDataEndpointProvided operation.
//
// Find the attached ProviderEndpoint of the EndpointData with the given ID.
//
// GET /admin/endpoint-data/{id}/endpoint-provided
func (UnimplementedHandler) ReadEndpointDataEndpointProvided(ctx context.Context, params ReadEndpointDataEndpointProvidedParams) (r ReadEndpointDataEndpointProvidedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEndpointDataEndpointRequired implements readEndpointDataEndpointRequired operation.
//
// Find the attached ProviderEndpoint of the EndpointData with the given ID.
//
// GET /admin/endpoint-data/{id}/endpoint-required
func (UnimplementedHandler) ReadEndpointDataEndpointRequired(ctx context.Context, params ReadEndpointDataEndpointRequiredParams) (r ReadEndpointDataEndpointRequiredRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadProviderEndpoint implements readProviderEndpoint operation.
//
// Finds the ProviderEndpoint with the requested ID and returns it.
//
// GET /admin/provider-endpoints/{id}
func (UnimplementedHandler) ReadProviderEndpoint(ctx context.Context, params ReadProviderEndpointParams) (r ReadProviderEndpointRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadProviderEndpointProvider implements readProviderEndpointProvider operation.
//
// Find the attached ProviderRegisterData of the ProviderEndpoint with the given ID.
//
// GET /admin/provider-endpoints/{id}/provider
func (UnimplementedHandler) ReadProviderEndpointProvider(ctx context.Context, params ReadProviderEndpointProviderParams) (r ReadProviderEndpointProviderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadProviderRegisterData implements readProviderRegisterData operation.
//
// Finds the ProviderRegisterData with the requested ID and returns it.
//
// GET /admin/provider-register-data/{id}
func (UnimplementedHandler) ReadProviderRegisterData(ctx context.Context, params ReadProviderRegisterDataParams) (r ReadProviderRegisterDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateEndpointData implements updateEndpointData operation.
//
// Updates a EndpointData and persists changes to storage.
//
// PATCH /admin/endpoint-data/{id}
func (UnimplementedHandler) UpdateEndpointData(ctx context.Context, req *UpdateEndpointDataReq, params UpdateEndpointDataParams) (r UpdateEndpointDataRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateProviderEndpoint implements updateProviderEndpoint operation.
//
// Updates a ProviderEndpoint and persists changes to storage.
//
// PATCH /admin/provider-endpoints/{id}
func (UnimplementedHandler) UpdateProviderEndpoint(ctx context.Context, req *UpdateProviderEndpointReq, params UpdateProviderEndpointParams) (r UpdateProviderEndpointRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateProviderRegisterData implements updateProviderRegisterData operation.
//
// Updates a ProviderRegisterData and persists changes to storage.
//
// PATCH /admin/provider-register-data/{id}
func (UnimplementedHandler) UpdateProviderRegisterData(ctx context.Context, req *UpdateProviderRegisterDataReq, params UpdateProviderRegisterDataParams) (r UpdateProviderRegisterDataRes, _ error) {
	return r, ht.ErrNotImplemented
}
