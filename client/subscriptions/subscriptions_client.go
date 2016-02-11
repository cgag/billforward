package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new subscriptions API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscriptions API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
AddPaymentMethodToSubscription enables the payment method to pay invoices of this subscription

{"nickname":"Add payment-method to subscription","response":"addPaymentMethod.html","request":"addPaymentMethod.request.html"}
*/
func (a *Client) AddPaymentMethodToSubscription(params *AddPaymentMethodToSubscriptionParams) (*AddPaymentMethodToSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPaymentMethodToSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "addPaymentMethodToSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions/{subscription-ID}/payment-methods",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddPaymentMethodToSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddPaymentMethodToSubscriptionOK), nil
}

/*
AvailablePaymentMethodsForSubscription returns all available payment methods for the specified subscription by default 10 values are returned records are returned in natural order

{ "nickname" : "List on subscription","response" : "getAvailablePaymentMethods.html"}
*/
func (a *Client) AvailablePaymentMethodsForSubscription(params *AvailablePaymentMethodsForSubscriptionParams) (*AvailablePaymentMethodsForSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailablePaymentMethodsForSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "availablePaymentMethodsForSubscription",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscription-ID}/payment-methods",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvailablePaymentMethodsForSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailablePaymentMethodsForSubscriptionOK), nil
}

/*
BatchCreateSubscriptions creates multiple subscriptions

{"nickname":"Create multiple subscriptions","response":"createMultipleSubscriptionViaHelper.html","request":"createMultipleSubscriptionViaHelper.request.html"}
*/
func (a *Client) BatchCreateSubscriptions(params *BatchCreateSubscriptionsParams) (*BatchCreateSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchCreateSubscriptionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "batchCreateSubscriptions",
		Method:             "POST",
		PathPattern:        "/subscriptions/batch",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchCreateSubscriptionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*BatchCreateSubscriptionsOK), nil
}

/*
CancelSubscription retires the subscription specified by the subscription ID parameter retiring a subscription causes it to cancel based on the specified retirement settings for the product

{"nickname":"Cancel subscription","response":"deleteSubscription.html","request":"deleteSubscriptionRequest.html"}
*/
func (a *Client) CancelSubscription(params *CancelSubscriptionParams) (*CancelSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "cancelSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions/{subscription-ID}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelSubscriptionOK), nil
}

/*
CreateAggregatingSubscription creates an aggregating subscription

{"nickname":"Create aggregating subscription","response":"createAggregatingSubscription.html","request":"createAggregatingSubscription.request.html"}
*/
func (a *Client) CreateAggregatingSubscription(params *CreateAggregatingSubscriptionParams) (*CreateAggregatingSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAggregatingSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createAggregatingSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions/aggregating",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAggregatingSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAggregatingSubscriptionOK), nil
}

/*
CreateSubscription creates a new subscription

{"nickname":"Create a new subscription","request":"createSubscriptionRequest.html","response":"createSubscriptionResponse.html"}
*/
func (a *Client) CreateSubscription(params *CreateSubscriptionParams) (*CreateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSubscriptionOK), nil
}

/*
CreateSubscriptionV2 creates a subscription v2

{"nickname":"Create a subscription (V2)","response":"createSubscriptionViaHelper.html","request":"createSubscriptionViaHelper.request.html"}
*/
func (a *Client) CreateSubscriptionV2(params *CreateSubscriptionV2Params) (*CreateSubscriptionV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionV2Params()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createSubscriptionV2",
		Method:             "POST",
		PathPattern:        "/subscriptions/create",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSubscriptionV2Reader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSubscriptionV2OK), nil
}

/*
DeleteMetadataForSubscription removes any associated metadata

{"nickname":"Clear from subscription","request" :"deleteSubscriptionMetadataRequest.html","response":"deleteSubscriptionMetadataResponse.html"}
*/
func (a *Client) DeleteMetadataForSubscription(params *DeleteMetadataForSubscriptionParams) (*DeleteMetadataForSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMetadataForSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteMetadataForSubscription",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscription-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMetadataForSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMetadataForSubscriptionOK), nil
}

/*
GetMetadataForSubscription retrieves any associated metadata

{"nickname":"Retrieve on subscription","request":"getSubscriptionMetadataRequest.html","response":"getSubscriptionMetadataResponse.html"}
*/
func (a *Client) GetMetadataForSubscription(params *GetMetadataForSubscriptionParams) (*GetMetadataForSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataForSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getMetadataForSubscription",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscription-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataForSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetadataForSubscriptionOK), nil
}

/*
GetSubscriptionByAccountID retrieves a collection of subscriptions specified by the account ID parameter by default 10 values are returned records are returned in natural order

{"nickname":"Retrieve by account","response":"getSubscriptionByAccount.html"}
*/
func (a *Client) GetSubscriptionByAccountID(params *GetSubscriptionByAccountIDParams) (*GetSubscriptionByAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionByAccountIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getSubscriptionByAccountID",
		Method:             "GET",
		PathPattern:        "/subscriptions/account/{account-ID}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSubscriptionByAccountIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionByAccountIDOK), nil
}

/*
GetSubscriptionByID retrieves a single subscription specified by the ID parameter

{"nickname":"Retrieve an existing subscription","response":"getSubscriptionByID.html"}
*/
func (a *Client) GetSubscriptionByID(params *GetSubscriptionByIDParams) (*GetSubscriptionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionByIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getSubscriptionByID",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscription-ID}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSubscriptionByIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionByIDOK), nil
}

/*
RemovePaymentMethodFromSubscription removes the specified payment method for the given subscription

{"nickname":"Remove payment-method","response":"removePaymentMethod.html","request":"removePaymentMethod.request.html"}
*/
func (a *Client) RemovePaymentMethodFromSubscription(params *RemovePaymentMethodFromSubscriptionParams) (*RemovePaymentMethodFromSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemovePaymentMethodFromSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "removePaymentMethodFromSubscription",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscription-ID}/payment-methods/{payment-method-ID}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemovePaymentMethodFromSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemovePaymentMethodFromSubscriptionOK), nil
}

/*
ReviveSubscription revives a cancelled subscription and returns it to its previous state

{"nickname":"Revive subscription","request":"reviveSubscriptionRequest.html", "response":"reviveSubscription.html"}
*/
func (a *Client) ReviveSubscription(params *ReviveSubscriptionParams) (*ReviveSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReviveSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "reviveSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions/{subscription-ID}/revive",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReviveSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReviveSubscriptionOK), nil
}

/*
SetMetadataForSubscription removes any existing metadata keys and create the provided data

{"nickname":"Set on subscription","request":"setSubscriptionMetadataRequest.html","response":"setSubscriptionMetadataResponse.html"}
*/
func (a *Client) SetMetadataForSubscription(params *SetMetadataForSubscriptionParams) (*SetMetadataForSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMetadataForSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "setMetadataForSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions/{subscription-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetMetadataForSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetMetadataForSubscriptionOK), nil
}

/*
UpdateSubscriptionV2 updates a subscription v2

{"nickname":"Update subscription (V2)","response":"updateSubscriptionViaHelper.html","request":"updateSubscriptionViaHelper.request.html"}
*/
func (a *Client) UpdateSubscriptionV2(params *UpdateSubscriptionV2Params) (*UpdateSubscriptionV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSubscriptionV2Params()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateSubscriptionV2",
		Method:             "PUT",
		PathPattern:        "/subscriptions/update",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSubscriptionV2Reader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSubscriptionV2OK), nil
}

/*
UpsertMetadataForSubscription updates any existing metadata key values and insert any new key values no keys will be removed

{"nickname":"Upsert on subscription","request":"upsertSubscriptionMetadataRequest.html","response":"upsertSubscriptionMetadataResponse.html"}
*/
func (a *Client) UpsertMetadataForSubscription(params *UpsertMetadataForSubscriptionParams) (*UpsertMetadataForSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertMetadataForSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "upsertMetadataForSubscription",
		Method:             "PUT",
		PathPattern:        "/subscriptions/{subscription-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertMetadataForSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpsertMetadataForSubscriptionOK), nil
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
