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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRetireWebhookParams creates a new RetireWebhookParams object
// with the default values initialized.
func NewRetireWebhookParams() *RetireWebhookParams {
	var ()
	return &RetireWebhookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetireWebhookParamsWithTimeout creates a new RetireWebhookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetireWebhookParamsWithTimeout(timeout time.Duration) *RetireWebhookParams {
	var ()
	return &RetireWebhookParams{

		timeout: timeout,
	}
}

// NewRetireWebhookParamsWithContext creates a new RetireWebhookParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetireWebhookParamsWithContext(ctx context.Context) *RetireWebhookParams {
	var ()
	return &RetireWebhookParams{

		Context: ctx,
	}
}

// NewRetireWebhookParamsWithHTTPClient creates a new RetireWebhookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetireWebhookParamsWithHTTPClient(client *http.Client) *RetireWebhookParams {
	var ()
	return &RetireWebhookParams{
		HTTPClient: client,
	}
}

/*RetireWebhookParams contains all the parameters to send to the API endpoint
for the retire webhook operation typically these are written to a http.Request
*/
type RetireWebhookParams struct {

	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*WebhookID
	  ID of the webhook.

	*/
	WebhookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retire webhook params
func (o *RetireWebhookParams) WithTimeout(timeout time.Duration) *RetireWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retire webhook params
func (o *RetireWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retire webhook params
func (o *RetireWebhookParams) WithContext(ctx context.Context) *RetireWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retire webhook params
func (o *RetireWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retire webhook params
func (o *RetireWebhookParams) WithHTTPClient(client *http.Client) *RetireWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retire webhook params
func (o *RetireWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizations adds the organizations to the retire webhook params
func (o *RetireWebhookParams) WithOrganizations(organizations []string) *RetireWebhookParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the retire webhook params
func (o *RetireWebhookParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WithWebhookID adds the webhookID to the retire webhook params
func (o *RetireWebhookParams) WithWebhookID(webhookID string) *RetireWebhookParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the retire webhook params
func (o *RetireWebhookParams) SetWebhookID(webhookID string) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *RetireWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	// path param webhook-ID
	if err := r.SetPathParam("webhook-ID", o.WebhookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
