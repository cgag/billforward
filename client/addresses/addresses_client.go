package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new addresses API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for addresses API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateAddress creates

{"nickname":"Create a new address","response":"createAddressResponse.html","request":"createAddressRequest.html"}
*/
func (a *Client) CreateAddress(params *CreateAddressParams) (*CreateAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAddress",
		Method:             "POST",
		PathPattern:        "/addresses",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAddressReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAddressOK), nil
}

/*
UpdateAddress updates

{"nickname":"Update an address","response":"updateAddressResponse.html","request":"updateAddressRequest.html"}
*/
func (a *Client) UpdateAddress(params *UpdateAddressParams) (*UpdateAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAddress",
		Method:             "PUT",
		PathPattern:        "/addresses",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAddressReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAddressOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
