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

// NewCreateWebhookParams creates a new CreateWebhookParams object
// with the default values initialized.
func NewCreateWebhookParams() *CreateWebhookParams {
	var ()
	return &CreateWebhookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWebhookParamsWithTimeout creates a new CreateWebhookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWebhookParamsWithTimeout(timeout time.Duration) *CreateWebhookParams {
	var ()
	return &CreateWebhookParams{

		timeout: timeout,
	}
}

// NewCreateWebhookParamsWithContext creates a new CreateWebhookParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWebhookParamsWithContext(ctx context.Context) *CreateWebhookParams {
	var ()
	return &CreateWebhookParams{

		Context: ctx,
	}
}

// NewCreateWebhookParamsWithHTTPClient creates a new CreateWebhookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWebhookParamsWithHTTPClient(client *http.Client) *CreateWebhookParams {
	var ()
	return &CreateWebhookParams{
		HTTPClient: client,
	}
}

/*CreateWebhookParams contains all the parameters to send to the API endpoint
for the create webhook operation typically these are written to a http.Request
*/
type CreateWebhookParams struct {

	/*Webhook
	  The webhook object to be created.

	*/
	Webhook *models.CreateWebhookRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create webhook params
func (o *CreateWebhookParams) WithTimeout(timeout time.Duration) *CreateWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create webhook params
func (o *CreateWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create webhook params
func (o *CreateWebhookParams) WithContext(ctx context.Context) *CreateWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create webhook params
func (o *CreateWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create webhook params
func (o *CreateWebhookParams) WithHTTPClient(client *http.Client) *CreateWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create webhook params
func (o *CreateWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebhook adds the webhook to the create webhook params
func (o *CreateWebhookParams) WithWebhook(webhook *models.CreateWebhookRequest) *CreateWebhookParams {
	o.SetWebhook(webhook)
	return o
}

// SetWebhook adds the webhook to the create webhook params
func (o *CreateWebhookParams) SetWebhook(webhook *models.CreateWebhookRequest) {
	o.Webhook = webhook
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Webhook == nil {
		o.Webhook = new(models.CreateWebhookRequest)
	}

	if err := r.SetBodyParam(o.Webhook); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
