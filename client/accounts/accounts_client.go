package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new accounts API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Create an Account.

{"nickname":"Create a new account","response":"createAccountResponse.html","request":"createAccountRequest.html"}
*/
func (a *Client) CreateAccount(params *CreateAccountParams) (*CreateAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "createAccount",
		Method:      "POST",
		PathPattern: "/accounts",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &CreateAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAccountOK), nil
}

/*Delete the account specified by the account-ID parameter.

{"nickname":"Retire","response":"deleteAccount.html"}
*/
func (a *Client) DeleteAccount(params *DeleteAccountParams) (*DeleteAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "deleteAccount",
		Method:      "DELETE",
		PathPattern: "/accounts/{account-ID}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &DeleteAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccountOK), nil
}

/*Remove any associated metadata.

{"nickname":"Clear metadata from account","request" :"deleteAccountMetadataRequest.html","response":"deleteAccountMetadataResponse.html"}
*/
func (a *Client) DeleteMetadataForAccount(params *DeleteMetadataForAccountParams) (*DeleteMetadataForAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMetadataForAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "deleteMetadataForAccount",
		Method:      "DELETE",
		PathPattern: "/accounts/{account-ID}/metadata",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &DeleteMetadataForAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMetadataForAccountOK), nil
}

/*Returns a single account, specified by the account-ID parameter.

{"nickname":"Retrieve an existing account","response":"getAccountByID.html"}
*/
func (a *Client) GetAccountByID(params *GetAccountByIDParams) (*GetAccountByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountByIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getAccountByID",
		Method:      "GET",
		PathPattern: "/accounts/{account-ID}",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetAccountByIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountByIDOK), nil
}

/*Returns a collection of all account objects. By default 10 values are returned. Records are returned in natural order.

{"nickname":"Get all accounts","response":"getAccountAll.html"}
*/
func (a *Client) GetAllAccounts(params *GetAllAccountsParams) (*GetAllAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllAccountsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getAllAccounts",
		Method:      "GET",
		PathPattern: "/accounts",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetAllAccountsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllAccountsOK), nil
}

/*Retrieve any associated metadata.

{"nickname":"Retrieve metadata on account","request":"getAccountMetadataRequest.html","response":"getAccountMetadataResponse.html"}
*/
func (a *Client) GetMetadataForAccount(params *GetMetadataForAccountParams) (*GetMetadataForAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataForAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "getMetadataForAccount",
		Method:      "GET",
		PathPattern: "/accounts/{account-ID}/metadata",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &GetMetadataForAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetadataForAccountOK), nil
}

/*Remove any existing metadata keys and create the provided data.

{"nickname":"Set metadata on account","request":"setAccountMetadataRequest.html","response":"setAccountMetadataResponse.html"}
*/
func (a *Client) SetMetadataForAccount(params *SetMetadataForAccountParams) (*SetMetadataForAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMetadataForAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "setMetadataForAccount",
		Method:      "POST",
		PathPattern: "/accounts/{account-ID}/metadata",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &SetMetadataForAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetMetadataForAccountOK), nil
}

/*Update any existing metadata key-values and insert any new key-values, no keys will be removed.

{"nickname":"Upsert metadata on account","request":"upsertAccountMetadataRequest.html","response":"upsertAccountMetadataResponse.html"}
*/
func (a *Client) UpsertMetadataForAccount(params *UpsertMetadataForAccountParams) (*UpsertMetadataForAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertMetadataForAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "upsertMetadataForAccount",
		Method:      "PUT",
		PathPattern: "/accounts/{account-ID}/metadata",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &UpsertMetadataForAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpsertMetadataForAccountOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
