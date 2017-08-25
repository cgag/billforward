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

// NewAdvanceSubscriptionParams creates a new AdvanceSubscriptionParams object
// with the default values initialized.
func NewAdvanceSubscriptionParams() *AdvanceSubscriptionParams {
	var ()
	return &AdvanceSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdvanceSubscriptionParamsWithTimeout creates a new AdvanceSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdvanceSubscriptionParamsWithTimeout(timeout time.Duration) *AdvanceSubscriptionParams {
	var ()
	return &AdvanceSubscriptionParams{

		timeout: timeout,
	}
}

// NewAdvanceSubscriptionParamsWithContext creates a new AdvanceSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdvanceSubscriptionParamsWithContext(ctx context.Context) *AdvanceSubscriptionParams {
	var ()
	return &AdvanceSubscriptionParams{

		Context: ctx,
	}
}

// NewAdvanceSubscriptionParamsWithHTTPClient creates a new AdvanceSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdvanceSubscriptionParamsWithHTTPClient(client *http.Client) *AdvanceSubscriptionParams {
	var ()
	return &AdvanceSubscriptionParams{
		HTTPClient: client,
	}
}

/*AdvanceSubscriptionParams contains all the parameters to send to the API endpoint
for the advance subscription operation typically these are written to a http.Request
*/
type AdvanceSubscriptionParams struct {

	/*Request
	  The request

	*/
	Request *models.TimeRequest
	/*SubscriptionID
	  ID of the subscription.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the advance subscription params
func (o *AdvanceSubscriptionParams) WithTimeout(timeout time.Duration) *AdvanceSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the advance subscription params
func (o *AdvanceSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the advance subscription params
func (o *AdvanceSubscriptionParams) WithContext(ctx context.Context) *AdvanceSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the advance subscription params
func (o *AdvanceSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the advance subscription params
func (o *AdvanceSubscriptionParams) WithHTTPClient(client *http.Client) *AdvanceSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the advance subscription params
func (o *AdvanceSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the advance subscription params
func (o *AdvanceSubscriptionParams) WithRequest(request *models.TimeRequest) *AdvanceSubscriptionParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the advance subscription params
func (o *AdvanceSubscriptionParams) SetRequest(request *models.TimeRequest) {
	o.Request = request
}

// WithSubscriptionID adds the subscriptionID to the advance subscription params
func (o *AdvanceSubscriptionParams) WithSubscriptionID(subscriptionID string) *AdvanceSubscriptionParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the advance subscription params
func (o *AdvanceSubscriptionParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *AdvanceSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request == nil {
		o.Request = new(models.TimeRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
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
