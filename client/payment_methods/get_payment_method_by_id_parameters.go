package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPaymentMethodByIDParams creates a new GetPaymentMethodByIDParams object
// with the default values initialized.
func NewGetPaymentMethodByIDParams() *GetPaymentMethodByIDParams {
	var ()
	return &GetPaymentMethodByIDParams{}
}

/*GetPaymentMethodByIDParams contains all the parameters to send to the API endpoint
for the get payment method by ID operation typically these are written to a http.Request
*/
type GetPaymentMethodByIDParams struct {

	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*PaymentMethodID*/
	PaymentMethodID string
}

// WithOrganizations adds the organizations to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithOrganizations(organizations []string) *GetPaymentMethodByIDParams {
	o.Organizations = organizations
	return o
}

// WithPaymentMethodID adds the paymentMethodId to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithPaymentMethodID(paymentMethodId string) *GetPaymentMethodByIDParams {
	o.PaymentMethodID = paymentMethodId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentMethodByIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	// path param payment-method-ID
	if err := r.SetPathParam("payment-method-ID", o.PaymentMethodID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
