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

	"github.com/authclub/billforward/models"
)

// NewSetMetadataForAccountParams creates a new SetMetadataForAccountParams object
// with the default values initialized.
func NewSetMetadataForAccountParams() *SetMetadataForAccountParams {
	var ()
	return &SetMetadataForAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetMetadataForAccountParamsWithTimeout creates a new SetMetadataForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetMetadataForAccountParamsWithTimeout(timeout time.Duration) *SetMetadataForAccountParams {
	var ()
	return &SetMetadataForAccountParams{

		timeout: timeout,
	}
}

// NewSetMetadataForAccountParamsWithContext creates a new SetMetadataForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetMetadataForAccountParamsWithContext(ctx context.Context) *SetMetadataForAccountParams {
	var ()
	return &SetMetadataForAccountParams{

		Context: ctx,
	}
}

// NewSetMetadataForAccountParamsWithHTTPClient creates a new SetMetadataForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetMetadataForAccountParamsWithHTTPClient(client *http.Client) *SetMetadataForAccountParams {
	var ()
	return &SetMetadataForAccountParams{
		HTTPClient: client,
	}
}

/*SetMetadataForAccountParams contains all the parameters to send to the API endpoint
for the set metadata for account operation typically these are written to a http.Request
*/
type SetMetadataForAccountParams struct {

	/*AccountID*/
	AccountID string
	/*Metadata*/
	Metadata models.DynamicMetadata
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set metadata for account params
func (o *SetMetadataForAccountParams) WithTimeout(timeout time.Duration) *SetMetadataForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set metadata for account params
func (o *SetMetadataForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set metadata for account params
func (o *SetMetadataForAccountParams) WithContext(ctx context.Context) *SetMetadataForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set metadata for account params
func (o *SetMetadataForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set metadata for account params
func (o *SetMetadataForAccountParams) WithHTTPClient(client *http.Client) *SetMetadataForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set metadata for account params
func (o *SetMetadataForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the set metadata for account params
func (o *SetMetadataForAccountParams) WithAccountID(accountID string) *SetMetadataForAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the set metadata for account params
func (o *SetMetadataForAccountParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithMetadata adds the metadata to the set metadata for account params
func (o *SetMetadataForAccountParams) WithMetadata(metadata models.DynamicMetadata) *SetMetadataForAccountParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the set metadata for account params
func (o *SetMetadataForAccountParams) SetMetadata(metadata models.DynamicMetadata) {
	o.Metadata = metadata
}

// WithOrganizations adds the organizations to the set metadata for account params
func (o *SetMetadataForAccountParams) WithOrganizations(organizations []string) *SetMetadataForAccountParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the set metadata for account params
func (o *SetMetadataForAccountParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WriteToRequest writes these params to a swagger request
func (o *SetMetadataForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account-ID
	if err := r.SetPathParam("account-ID", o.AccountID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Metadata); err != nil {
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
