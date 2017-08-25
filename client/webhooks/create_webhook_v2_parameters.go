// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// NewCreateWebhookV2Params creates a new CreateWebhookV2Params object
// with the default values initialized.
func NewCreateWebhookV2Params() *CreateWebhookV2Params {
	var ()
	return &CreateWebhookV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWebhookV2ParamsWithTimeout creates a new CreateWebhookV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWebhookV2ParamsWithTimeout(timeout time.Duration) *CreateWebhookV2Params {
	var ()
	return &CreateWebhookV2Params{

		timeout: timeout,
	}
}

// NewCreateWebhookV2ParamsWithContext creates a new CreateWebhookV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWebhookV2ParamsWithContext(ctx context.Context) *CreateWebhookV2Params {
	var ()
	return &CreateWebhookV2Params{

		Context: ctx,
	}
}

// NewCreateWebhookV2ParamsWithHTTPClient creates a new CreateWebhookV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWebhookV2ParamsWithHTTPClient(client *http.Client) *CreateWebhookV2Params {
	var ()
	return &CreateWebhookV2Params{
		HTTPClient: client,
	}
}

/*CreateWebhookV2Params contains all the parameters to send to the API endpoint
for the create webhook v2 operation typically these are written to a http.Request
*/
type CreateWebhookV2Params struct {

	/*Request*/
	Request *models.CreateWebhookRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create webhook v2 params
func (o *CreateWebhookV2Params) WithTimeout(timeout time.Duration) *CreateWebhookV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create webhook v2 params
func (o *CreateWebhookV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create webhook v2 params
func (o *CreateWebhookV2Params) WithContext(ctx context.Context) *CreateWebhookV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create webhook v2 params
func (o *CreateWebhookV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create webhook v2 params
func (o *CreateWebhookV2Params) WithHTTPClient(client *http.Client) *CreateWebhookV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create webhook v2 params
func (o *CreateWebhookV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create webhook v2 params
func (o *CreateWebhookV2Params) WithRequest(request *models.CreateWebhookRequest) *CreateWebhookV2Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create webhook v2 params
func (o *CreateWebhookV2Params) SetRequest(request *models.CreateWebhookRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWebhookV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateWebhookRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
