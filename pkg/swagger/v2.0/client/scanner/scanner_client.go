// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scanner API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scanner API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateScanner(params *CreateScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateScannerCreated, error)

	DeleteScanner(params *DeleteScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteScannerOK, error)

	GetScanner(params *GetScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScannerOK, error)

	GetScannerMetadata(params *GetScannerMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScannerMetadataOK, error)

	ListScanners(params *ListScannersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListScannersOK, error)

	PingScanner(params *PingScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PingScannerOK, error)

	SetScannerAsDefault(params *SetScannerAsDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetScannerAsDefaultOK, error)

	UpdateScanner(params *UpdateScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateScannerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateScanner creates a scanner registration

  Creats a new scanner registration with the given data.

*/
func (a *Client) CreateScanner(params *CreateScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateScannerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScannerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createScanner",
		Method:             "POST",
		PathPattern:        "/scanners",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateScannerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateScannerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createScanner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteScanner deletes a scanner registration

  Deletes the specified scanner registration.

*/
func (a *Client) DeleteScanner(params *DeleteScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteScannerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteScannerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteScanner",
		Method:             "DELETE",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteScannerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteScannerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteScanner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetScanner gets a scanner registration details

  Retruns the details of the specified scanner registration.

*/
func (a *Client) GetScanner(params *GetScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScannerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScannerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScanner",
		Method:             "GET",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScannerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScannerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetScannerMetadata gets the metadata of the specified scanner registration

  Get the metadata of the specified scanner registration, including the capabilities and customized properties.

*/
func (a *Client) GetScannerMetadata(params *GetScannerMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScannerMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScannerMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScannerMetadata",
		Method:             "GET",
		PathPattern:        "/scanners/{registration_id}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScannerMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScannerMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScannerMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListScanners lists scanner registrations

  Returns a list of currently configured scanner registrations.

*/
func (a *Client) ListScanners(params *ListScannersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListScannersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScannersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listScanners",
		Method:             "GET",
		PathPattern:        "/scanners",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListScannersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListScannersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listScanners: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PingScanner tests scanner registration settings

  Pings scanner adapter to test endpoint URL and authorization settings.

*/
func (a *Client) PingScanner(params *PingScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PingScannerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingScannerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pingScanner",
		Method:             "POST",
		PathPattern:        "/scanners/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PingScannerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PingScannerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pingScanner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetScannerAsDefault sets system default scanner registration

  Set the specified scanner registration as the system default one.

*/
func (a *Client) SetScannerAsDefault(params *SetScannerAsDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetScannerAsDefaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetScannerAsDefaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setScannerAsDefault",
		Method:             "PATCH",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetScannerAsDefaultReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetScannerAsDefaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setScannerAsDefault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateScanner updates a scanner registration

  Updates the specified scanner registration.

*/
func (a *Client) UpdateScanner(params *UpdateScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateScannerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateScannerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateScanner",
		Method:             "PUT",
		PathPattern:        "/scanners/{registration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateScannerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateScannerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateScanner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
