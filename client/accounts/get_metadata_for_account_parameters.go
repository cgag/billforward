package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

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

	timeout time.Duration
}

// WithAccountID adds the accountID to the get metadata for account params
func (o *GetMetadataForAccountParams) WithAccountID(accountID string) *GetMetadataForAccountParams {
	o.AccountID = accountID
	return o
}

// WithOrganizations adds the organizations to the get metadata for account params
func (o *GetMetadataForAccountParams) WithOrganizations(organizations []string) *GetMetadataForAccountParams {
	o.Organizations = organizations
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetadataForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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
