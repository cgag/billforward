// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// NewGetPaymentByIDParams creates a new GetPaymentByIDParams object
// with the default values initialized.
func NewGetPaymentByIDParams() *GetPaymentByIDParams {
	var ()
	return &GetPaymentByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentByIDParamsWithTimeout creates a new GetPaymentByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentByIDParamsWithTimeout(timeout time.Duration) *GetPaymentByIDParams {
	var ()
	return &GetPaymentByIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentByIDParamsWithContext creates a new GetPaymentByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentByIDParamsWithContext(ctx context.Context) *GetPaymentByIDParams {
	var ()
	return &GetPaymentByIDParams{

		Context: ctx,
	}
}

// NewGetPaymentByIDParamsWithHTTPClient creates a new GetPaymentByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentByIDParamsWithHTTPClient(client *http.Client) *GetPaymentByIDParams {
	var ()
	return &GetPaymentByIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentByIDParams contains all the parameters to send to the API endpoint
for the get payment by ID operation typically these are written to a http.Request
*/
type GetPaymentByIDParams struct {

	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*PaymentID
	  The unique string-ID of the payment.

	*/
	PaymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payment by ID params
func (o *GetPaymentByIDParams) WithTimeout(timeout time.Duration) *GetPaymentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment by ID params
func (o *GetPaymentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment by ID params
func (o *GetPaymentByIDParams) WithContext(ctx context.Context) *GetPaymentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment by ID params
func (o *GetPaymentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment by ID params
func (o *GetPaymentByIDParams) WithHTTPClient(client *http.Client) *GetPaymentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment by ID params
func (o *GetPaymentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizations adds the organizations to the get payment by ID params
func (o *GetPaymentByIDParams) WithOrganizations(organizations []string) *GetPaymentByIDParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the get payment by ID params
func (o *GetPaymentByIDParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WithPaymentID adds the paymentID to the get payment by ID params
func (o *GetPaymentByIDParams) WithPaymentID(paymentID string) *GetPaymentByIDParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the get payment by ID params
func (o *GetPaymentByIDParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param payment-ID
	if err := r.SetPathParam("payment-ID", o.PaymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
