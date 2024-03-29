// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"servicediscoverer/ent"
	"servicediscoverer/ent/endpointdata"
	"servicediscoverer/ent/providerendpoint"
	"servicediscoverer/ent/providerregisterdata"

	"github.com/go-faster/jx"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateEndpointData handles POST /endpoint-data-slice requests.
func (h *OgentHandler) CreateEndpointData(ctx context.Context, req *CreateEndpointDataReq) (CreateEndpointDataRes, error) {
	b := h.client.EndpointData.Create()
	// Add all fields.
	b.SetDataName(req.DataName)
	b.SetDescription(req.Description)
	b.SetType(req.Type)
	// Add all edges.
	if v, ok := req.EndpointRequired.Get(); ok {
		b.SetEndpointRequiredID(v)
	}
	if v, ok := req.EndpointProvided.Get(); ok {
		b.SetEndpointProvidedID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.EndpointData.Query().Where(endpointdata.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewEndpointDataCreate(e), nil
}

// ReadEndpointData handles GET /endpoint-data-slice/{id} requests.
func (h *OgentHandler) ReadEndpointData(ctx context.Context, params ReadEndpointDataParams) (ReadEndpointDataRes, error) {
	q := h.client.EndpointData.Query().Where(endpointdata.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewEndpointDataRead(e), nil
}

// UpdateEndpointData handles PATCH /endpoint-data-slice/{id} requests.
func (h *OgentHandler) UpdateEndpointData(ctx context.Context, req *UpdateEndpointDataReq, params UpdateEndpointDataParams) (UpdateEndpointDataRes, error) {
	b := h.client.EndpointData.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.DataName.Get(); ok {
		b.SetDataName(v)
	}
	if v, ok := req.Description.Get(); ok {
		b.SetDescription(v)
	}
	if v, ok := req.Type.Get(); ok {
		b.SetType(v)
	}
	// Add all edges.
	if v, ok := req.EndpointRequired.Get(); ok {
		b.SetEndpointRequiredID(v)
	}
	if v, ok := req.EndpointProvided.Get(); ok {
		b.SetEndpointProvidedID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.EndpointData.Query().Where(endpointdata.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewEndpointDataUpdate(e), nil
}

// DeleteEndpointData handles DELETE /endpoint-data-slice/{id} requests.
func (h *OgentHandler) DeleteEndpointData(ctx context.Context, params DeleteEndpointDataParams) (DeleteEndpointDataRes, error) {
	err := h.client.EndpointData.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteEndpointDataNoContent), nil

}

// ListEndpointData handles GET /endpoint-data-slice requests.
func (h *OgentHandler) ListEndpointData(ctx context.Context, params ListEndpointDataParams) (ListEndpointDataRes, error) {
	q := h.client.EndpointData.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewEndpointDataLists(es)
	return (*ListEndpointDataOKApplicationJSON)(&r), nil
}

// ReadEndpointDataEndpointRequired handles GET /endpoint-data-slice/{id}/endpoint-required requests.
func (h *OgentHandler) ReadEndpointDataEndpointRequired(ctx context.Context, params ReadEndpointDataEndpointRequiredParams) (ReadEndpointDataEndpointRequiredRes, error) {
	q := h.client.EndpointData.Query().Where(endpointdata.IDEQ(params.ID)).QueryEndpointRequired()
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewEndpointDataEndpointRequiredRead(e), nil
}

// ReadEndpointDataEndpointProvided handles GET /endpoint-data-slice/{id}/endpoint-provided requests.
func (h *OgentHandler) ReadEndpointDataEndpointProvided(ctx context.Context, params ReadEndpointDataEndpointProvidedParams) (ReadEndpointDataEndpointProvidedRes, error) {
	q := h.client.EndpointData.Query().Where(endpointdata.IDEQ(params.ID)).QueryEndpointProvided()
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewEndpointDataEndpointProvidedRead(e), nil
}

// CreateProviderEndpoint handles POST /provider-endpoints requests.
func (h *OgentHandler) CreateProviderEndpoint(ctx context.Context, req *CreateProviderEndpointReq) (CreateProviderEndpointRes, error) {
	b := h.client.ProviderEndpoint.Create()
	// Add all fields.
	b.SetName(req.Name)
	b.SetPath(req.Path)
	b.SetType(req.Type)
	b.SetDescription(req.Description)
	// Add all edges.
	b.AddRequiredDatumIDs(req.RequiredData...)
	b.AddProvidedDatumIDs(req.ProvidedData...)
	if v, ok := req.Provider.Get(); ok {
		b.SetProviderID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.ProviderEndpoint.Query().Where(providerendpoint.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewProviderEndpointCreate(e), nil
}

// ReadProviderEndpoint handles GET /provider-endpoints/{id} requests.
func (h *OgentHandler) ReadProviderEndpoint(ctx context.Context, params ReadProviderEndpointParams) (ReadProviderEndpointRes, error) {
	q := h.client.ProviderEndpoint.Query().Where(providerendpoint.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewProviderEndpointRead(e), nil
}

// UpdateProviderEndpoint handles PATCH /provider-endpoints/{id} requests.
func (h *OgentHandler) UpdateProviderEndpoint(ctx context.Context, req *UpdateProviderEndpointReq, params UpdateProviderEndpointParams) (UpdateProviderEndpointRes, error) {
	b := h.client.ProviderEndpoint.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	if v, ok := req.Path.Get(); ok {
		b.SetPath(v)
	}
	if v, ok := req.Type.Get(); ok {
		b.SetType(v)
	}
	if v, ok := req.Description.Get(); ok {
		b.SetDescription(v)
	}
	// Add all edges.
	if req.RequiredData != nil {
		b.ClearRequiredData().AddRequiredDatumIDs(req.RequiredData...)
	}
	if req.ProvidedData != nil {
		b.ClearProvidedData().AddProvidedDatumIDs(req.ProvidedData...)
	}
	if v, ok := req.Provider.Get(); ok {
		b.SetProviderID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.ProviderEndpoint.Query().Where(providerendpoint.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewProviderEndpointUpdate(e), nil
}

// DeleteProviderEndpoint handles DELETE /provider-endpoints/{id} requests.
func (h *OgentHandler) DeleteProviderEndpoint(ctx context.Context, params DeleteProviderEndpointParams) (DeleteProviderEndpointRes, error) {
	err := h.client.ProviderEndpoint.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteProviderEndpointNoContent), nil

}

// ListProviderEndpoint handles GET /provider-endpoints requests.
func (h *OgentHandler) ListProviderEndpoint(ctx context.Context, params ListProviderEndpointParams) (ListProviderEndpointRes, error) {
	q := h.client.ProviderEndpoint.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewProviderEndpointLists(es)
	return (*ListProviderEndpointOKApplicationJSON)(&r), nil
}

// ListProviderEndpointRequiredData handles GET /provider-endpoints/{id}/required-data requests.
func (h *OgentHandler) ListProviderEndpointRequiredData(ctx context.Context, params ListProviderEndpointRequiredDataParams) (ListProviderEndpointRequiredDataRes, error) {
	q := h.client.ProviderEndpoint.Query().Where(providerendpoint.IDEQ(params.ID)).QueryRequiredData()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewProviderEndpointRequiredDataLists(es)
	return (*ListProviderEndpointRequiredDataOKApplicationJSON)(&r), nil
}

// ListProviderEndpointProvidedData handles GET /provider-endpoints/{id}/provided-data requests.
func (h *OgentHandler) ListProviderEndpointProvidedData(ctx context.Context, params ListProviderEndpointProvidedDataParams) (ListProviderEndpointProvidedDataRes, error) {
	q := h.client.ProviderEndpoint.Query().Where(providerendpoint.IDEQ(params.ID)).QueryProvidedData()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewProviderEndpointProvidedDataLists(es)
	return (*ListProviderEndpointProvidedDataOKApplicationJSON)(&r), nil
}

// ReadProviderEndpointProvider handles GET /provider-endpoints/{id}/provider requests.
func (h *OgentHandler) ReadProviderEndpointProvider(ctx context.Context, params ReadProviderEndpointProviderParams) (ReadProviderEndpointProviderRes, error) {
	q := h.client.ProviderEndpoint.Query().Where(providerendpoint.IDEQ(params.ID)).QueryProvider()
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewProviderEndpointProviderRead(e), nil
}

// CreateProviderRegisterData handles POST /provider-register-data-slice requests.
func (h *OgentHandler) CreateProviderRegisterData(ctx context.Context, req *CreateProviderRegisterDataReq) (CreateProviderRegisterDataRes, error) {
	b := h.client.ProviderRegisterData.Create()
	// Add all fields.
	b.SetName(req.Name)
	b.SetPort(req.Port)
	b.SetAddress(req.Address)
	b.SetDescription(req.Description)
	b.SetLiveInterval(req.LiveInterval)
	b.SetLiveTimeout(req.LiveTimeout)
	// Add all edges.
	b.AddEndpointIDs(req.Endpoints...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.ProviderRegisterData.Query().Where(providerregisterdata.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewProviderRegisterDataCreate(e), nil
}

// ReadProviderRegisterData handles GET /provider-register-data-slice/{id} requests.
func (h *OgentHandler) ReadProviderRegisterData(ctx context.Context, params ReadProviderRegisterDataParams) (ReadProviderRegisterDataRes, error) {
	q := h.client.ProviderRegisterData.Query().Where(providerregisterdata.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewProviderRegisterDataRead(e), nil
}

// UpdateProviderRegisterData handles PATCH /provider-register-data-slice/{id} requests.
func (h *OgentHandler) UpdateProviderRegisterData(ctx context.Context, req *UpdateProviderRegisterDataReq, params UpdateProviderRegisterDataParams) (UpdateProviderRegisterDataRes, error) {
	b := h.client.ProviderRegisterData.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	if v, ok := req.Port.Get(); ok {
		b.SetPort(v)
	}
	if v, ok := req.Address.Get(); ok {
		b.SetAddress(v)
	}
	if v, ok := req.Description.Get(); ok {
		b.SetDescription(v)
	}
	if v, ok := req.LiveInterval.Get(); ok {
		b.SetLiveInterval(v)
	}
	if v, ok := req.LiveTimeout.Get(); ok {
		b.SetLiveTimeout(v)
	}
	// Add all edges.
	if req.Endpoints != nil {
		b.ClearEndpoints().AddEndpointIDs(req.Endpoints...)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.ProviderRegisterData.Query().Where(providerregisterdata.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewProviderRegisterDataUpdate(e), nil
}

// DeleteProviderRegisterData handles DELETE /provider-register-data-slice/{id} requests.
func (h *OgentHandler) DeleteProviderRegisterData(ctx context.Context, params DeleteProviderRegisterDataParams) (DeleteProviderRegisterDataRes, error) {
	err := h.client.ProviderRegisterData.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteProviderRegisterDataNoContent), nil

}

// ListProviderRegisterData handles GET /provider-register-data-slice requests.
func (h *OgentHandler) ListProviderRegisterData(ctx context.Context, params ListProviderRegisterDataParams) (ListProviderRegisterDataRes, error) {
	q := h.client.ProviderRegisterData.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewProviderRegisterDataLists(es)
	return (*ListProviderRegisterDataOKApplicationJSON)(&r), nil
}

// ListProviderRegisterDataEndpoints handles GET /provider-register-data-slice/{id}/endpoints requests.
func (h *OgentHandler) ListProviderRegisterDataEndpoints(ctx context.Context, params ListProviderRegisterDataEndpointsParams) (ListProviderRegisterDataEndpointsRes, error) {
	q := h.client.ProviderRegisterData.Query().Where(providerregisterdata.IDEQ(params.ID)).QueryEndpoints()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewProviderRegisterDataEndpointsLists(es)
	return (*ListProviderRegisterDataEndpointsOKApplicationJSON)(&r), nil
}
