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

// NewDeleteMetadataForAccountParams creates a new DeleteMetadataForAccountParams object
// with the default values initialized.
func NewDeleteMetadataForAccountParams() *DeleteMetadataForAccountParams {
	var ()
	return &DeleteMetadataForAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMetadataForAccountParamsWithTimeout creates a new DeleteMetadataForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMetadataForAccountParamsWithTimeout(timeout time.Duration) *DeleteMetadataForAccountParams {
	var ()
	return &DeleteMetadataForAccountParams{

		timeout: timeout,
	}
}

/*DeleteMetadataForAccountParams contains all the parameters to send to the API endpoint
for the delete metadata for account operation typically these are written to a http.Request
*/
type DeleteMetadataForAccountParams struct {

	/*AccountID*/
	AccountID string
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string

	timeout time.Duration
}

// WithAccountID adds the accountID to the delete metadata for account params
func (o *DeleteMetadataForAccountParams) WithAccountID(accountID string) *DeleteMetadataForAccountParams {
	o.AccountID = accountID
	return o
}

// WithOrganizations adds the organizations to the delete metadata for account params
func (o *DeleteMetadataForAccountParams) WithOrganizations(organizations []string) *DeleteMetadataForAccountParams {
	o.Organizations = organizations
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMetadataForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
