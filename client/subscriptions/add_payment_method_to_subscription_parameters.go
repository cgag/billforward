// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// NewAddPaymentMethodToSubscriptionParams creates a new AddPaymentMethodToSubscriptionParams object
// with the default values initialized.
func NewAddPaymentMethodToSubscriptionParams() *AddPaymentMethodToSubscriptionParams {
	var ()
	return &AddPaymentMethodToSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPaymentMethodToSubscriptionParamsWithTimeout creates a new AddPaymentMethodToSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPaymentMethodToSubscriptionParamsWithTimeout(timeout time.Duration) *AddPaymentMethodToSubscriptionParams {
	var ()
	return &AddPaymentMethodToSubscriptionParams{

		timeout: timeout,
	}
}

// NewAddPaymentMethodToSubscriptionParamsWithContext creates a new AddPaymentMethodToSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPaymentMethodToSubscriptionParamsWithContext(ctx context.Context) *AddPaymentMethodToSubscriptionParams {
	var ()
	return &AddPaymentMethodToSubscriptionParams{

		Context: ctx,
	}
}

// NewAddPaymentMethodToSubscriptionParamsWithHTTPClient creates a new AddPaymentMethodToSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPaymentMethodToSubscriptionParamsWithHTTPClient(client *http.Client) *AddPaymentMethodToSubscriptionParams {
	var ()
	return &AddPaymentMethodToSubscriptionParams{
		HTTPClient: client,
	}
}

/*AddPaymentMethodToSubscriptionParams contains all the parameters to send to the API endpoint
for the add payment method to subscription operation typically these are written to a http.Request
*/
type AddPaymentMethodToSubscriptionParams struct {

	/*PaymentMethod*/
	PaymentMethod *models.AddPaymentMethodRequest
	/*SubscriptionID*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithTimeout(timeout time.Duration) *AddPaymentMethodToSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithContext(ctx context.Context) *AddPaymentMethodToSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithHTTPClient(client *http.Client) *AddPaymentMethodToSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentMethod adds the paymentMethod to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithPaymentMethod(paymentMethod *models.AddPaymentMethodRequest) *AddPaymentMethodToSubscriptionParams {
	o.SetPaymentMethod(paymentMethod)
	return o
}

// SetPaymentMethod adds the paymentMethod to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) SetPaymentMethod(paymentMethod *models.AddPaymentMethodRequest) {
	o.PaymentMethod = paymentMethod
}

// WithSubscriptionID adds the subscriptionID to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithSubscriptionID(subscriptionID string) *AddPaymentMethodToSubscriptionParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddPaymentMethodToSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentMethod == nil {
		o.PaymentMethod = new(models.AddPaymentMethodRequest)
	}

	if err := r.SetBodyParam(o.PaymentMethod); err != nil {
		return err
	}

	// path param subscription-ID
	if err := r.SetPathParam("subscription-ID", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
