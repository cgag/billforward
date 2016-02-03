package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/authclub/billforward/models"
)

// NewSetMetadataForAccountParams creates a new SetMetadataForAccountParams object
// with the default values initialized.
func NewSetMetadataForAccountParams() *SetMetadataForAccountParams {
	var ()
	return &SetMetadataForAccountParams{}
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
}

// WithAccountID adds the accountId to the set metadata for account params
func (o *SetMetadataForAccountParams) WithAccountID(accountId string) *SetMetadataForAccountParams {
	o.AccountID = accountId
	return o
}

// WithMetadata adds the metadata to the set metadata for account params
func (o *SetMetadataForAccountParams) WithMetadata(metadata models.DynamicMetadata) *SetMetadataForAccountParams {
	o.Metadata = metadata
	return o
}

// WithOrganizations adds the organizations to the set metadata for account params
func (o *SetMetadataForAccountParams) WithOrganizations(organizations []string) *SetMetadataForAccountParams {
	o.Organizations = organizations
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SetMetadataForAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
