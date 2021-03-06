// Code generated by go-swagger; DO NOT EDIT.

package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payment methods API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment methods API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePaymentMethod deletes the payment method specified by the payment method ID parameter

{"nickname":"Delete payment method","response":"deletePaymentMethod.html"}
*/
func (a *Client) DeletePaymentMethod(params *DeletePaymentMethodParams) (*DeletePaymentMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentMethodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePaymentMethod",
		Method:             "DELETE",
		PathPattern:        "/payment-methods/{payment-method-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePaymentMethodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePaymentMethodOK), nil

}

/*
GetAllPaymentMethods returns a collection of all payment methods by default 10 values are returned records are returned in natural order

{"nickname":"Get all payment methods","response":"getPaymentMethodAll.html"}
*/
func (a *Client) GetAllPaymentMethods(params *GetAllPaymentMethodsParams) (*GetAllPaymentMethodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPaymentMethodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllPaymentMethods",
		Method:             "GET",
		PathPattern:        "/payment-methods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllPaymentMethodsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllPaymentMethodsOK), nil

}

/*
GetPaymentMethodByAccountID returns a collection of payment methods specified by the account ID parameter by default 10 values are returned records are returned in natural order

{"nickname":"Retrieve by account","response":"getPaymentMethodByAccount.html"}
*/
func (a *Client) GetPaymentMethodByAccountID(params *GetPaymentMethodByAccountIDParams) (*GetPaymentMethodByAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodByAccountIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentMethodByAccountID",
		Method:             "GET",
		PathPattern:        "/payment-methods/account/{account-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentMethodByAccountIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentMethodByAccountIDOK), nil

}

/*
GetPaymentMethodByID returns a single payment method specified by the payment method ID parameter

{"nickname":"Get existing payment method","response":"getPaymentMethodByID.html"}
*/
func (a *Client) GetPaymentMethodByID(params *GetPaymentMethodByIDParams) (*GetPaymentMethodByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentMethodByID",
		Method:             "GET",
		PathPattern:        "/payment-methods/{payment-method-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentMethodByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentMethodByIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
