// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewGetMetadataForAccountParams creates a new GetMetadataForAccountParams object
// with the default values initialized.
func NewGetMetadataForAccountParams() *GetMetadataForAccountParams {
	var ()
	return &GetMetadataForAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetadataForAccountParamsWithTimeout creates a new GetMetadataForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMetadataForAccountParamsWithTimeout(timeout time.Duration) *GetMetadataForAccountParams {
	var ()
	return &GetMetadataForAccountParams{

		timeout: timeout,
	}
}

// NewGetMetadataForAccountParamsWithContext creates a new GetMetadataForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMetadataForAccountParamsWithContext(ctx context.Context) *GetMetadataForAccountParams {
	var ()
	return &GetMetadataForAccountParams{

		Context: ctx,
	}
}

// NewGetMetadataForAccountParamsWithHTTPClient creates a new GetMetadataForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMetadataForAccountParamsWithHTTPClient(client *http.Client) *GetMetadataForAccountParams {
	var ()
	return &GetMetadataForAccountParams{
		HTTPClient: client,
	}
}

/*GetMetadataForAccountParams contains all the parameters to send to the API endpoint
for the get metadata for account operation typically these are written to a http.Request
*/
type GetMetadataForAccountParams struct {

	/*AccountID*/
	AccountID string
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get metadata for account params
func (o *GetMetadataForAccountParams) WithTimeout(timeout time.Duration) *GetMetadataForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metadata for account params
func (o *GetMetadataForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metadata for account params
func (o *GetMetadataForAccountParams) WithContext(ctx context.Context) *GetMetadataForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metadata for account params
func (o *GetMetadataForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metadata for account params
func (o *GetMetadataForAccountParams) WithHTTPClient(client *http.Client) *GetMetadataForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metadata for account params
func (o *GetMetadataForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get metadata for account params
func (o *GetMetadataForAccountParams) WithAccountID(accountID string) *GetMetadataForAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get metadata for account params
func (o *GetMetadataForAccountParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithOrganizations adds the organizations to the get metadata for account params
func (o *GetMetadataForAccountParams) WithOrganizations(organizations []string) *GetMetadataForAccountParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the get metadata for account params
func (o *GetMetadataForAccountParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetadataForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account-ID
	if err := r.SetPathParam("account-ID", o.AccountID); err != nil {
		return err
	}

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
