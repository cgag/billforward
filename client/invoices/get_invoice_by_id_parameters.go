package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*
GetInvoiceByIDParams contains all the parameters to send to the API endpoint
for the get invoice by ID operation typically these are written to a http.Request
*/
type GetInvoiceByIDParams struct {
	/*
	  Whether retired products should be returned.
	*/
	IncludeRetired bool
	/*
	  The ID of the invoice.
	*/
	InvoiceID string
	/*
	  The offset from the first invoice to return.
	*/
	Offset int32
	/*
	  Ihe direction of any ordering, either ASC or DESC.
	*/
	Order string
	/*
	  Specify a field used to order the result set.
	*/
	OrderBy string
	/*
	  A list of organization-IDs used to restrict the scope of API calls.
	*/
	Organizations []string
	/*
	  The maximum number of invoices to return.
	*/
	Records int32
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param include_retired
	if err := r.SetQueryParam("include_retired", swag.FormatBool(o.IncludeRetired)); err != nil {
		return err
	}

	// path param invoice-ID
	if err := r.SetPathParam("invoice-ID", o.InvoiceID); err != nil {
		return err
	}

	// query param offset
	if err := r.SetQueryParam("offset", swag.FormatInt32(o.Offset)); err != nil {
		return err
	}

	// query param order
	if err := r.SetQueryParam("order", o.Order); err != nil {
		return err
	}

	// query param order_by
	if err := r.SetQueryParam("order_by", o.OrderBy); err != nil {
		return err
	}

	valuesOrganizations := o.Organizations

	// query array param organizations
	if err := r.SetQueryParam("organizations", swag.JoinByFormat(valuesOrganizations, "multi")...); err != nil {
		return err
	}

	// query param records
	if err := r.SetQueryParam("records", swag.FormatInt32(o.Records)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
