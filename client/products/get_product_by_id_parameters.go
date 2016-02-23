package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetProductByIDParams creates a new GetProductByIDParams object
// with the default values initialized.
func NewGetProductByIDParams() *GetProductByIDParams {
	var (
		includeRetiredDefault bool   = bool(true)
		orderDefault          string = string("DESC")
		orderByDefault        string = string("created")
		recordsDefault        int32  = int32(10)
	)
	return &GetProductByIDParams{
		IncludeRetired: &includeRetiredDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,
	}
}

/*GetProductByIDParams contains all the parameters to send to the API endpoint
for the get product by ID operation typically these are written to a http.Request
*/
type GetProductByIDParams struct {

	/*IncludeRetired
	  Whether retired products should be returned.

	*/
	IncludeRetired *bool
	/*Offset
	  The offset from the first product-rate-plan to return.

	*/
	Offset *int32
	/*Order
	  Ihe direction of any ordering, either ASC or DESC.

	*/
	Order *string
	/*OrderBy
	  Specify a field used to order the result set.

	*/
	OrderBy *string
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*ProductID
	  ID or name of the product.

	*/
	ProductID string
	/*Records
	  The maximum number of product-rate-plans to return.

	*/
	Records *int32
}

// WithIncludeRetired adds the includeRetired to the get product by ID params
func (o *GetProductByIDParams) WithIncludeRetired(includeRetired *bool) *GetProductByIDParams {
	o.IncludeRetired = includeRetired
	return o
}

// WithOffset adds the offset to the get product by ID params
func (o *GetProductByIDParams) WithOffset(offset *int32) *GetProductByIDParams {
	o.Offset = offset
	return o
}

// WithOrder adds the order to the get product by ID params
func (o *GetProductByIDParams) WithOrder(order *string) *GetProductByIDParams {
	o.Order = order
	return o
}

// WithOrderBy adds the orderBy to the get product by ID params
func (o *GetProductByIDParams) WithOrderBy(orderBy *string) *GetProductByIDParams {
	o.OrderBy = orderBy
	return o
}

// WithOrganizations adds the organizations to the get product by ID params
func (o *GetProductByIDParams) WithOrganizations(organizations []string) *GetProductByIDParams {
	o.Organizations = organizations
	return o
}

// WithProductID adds the productId to the get product by ID params
func (o *GetProductByIDParams) WithProductID(productId string) *GetProductByIDParams {
	o.ProductID = productId
	return o
}

// WithRecords adds the records to the get product by ID params
func (o *GetProductByIDParams) WithRecords(records *int32) *GetProductByIDParams {
	o.Records = records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductByIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.IncludeRetired != nil {

		// query param include_retired
		var qrIncludeRetired bool
		if o.IncludeRetired != nil {
			qrIncludeRetired = *o.IncludeRetired
		}
		qIncludeRetired := swag.FormatBool(qrIncludeRetired)
		if qIncludeRetired != "" {
			if err := r.SetQueryParam("include_retired", qIncludeRetired); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	// path param product-ID
	if err := r.SetPathParam("product-ID", o.ProductID); err != nil {
		return err
	}

	if o.Records != nil {

		// query param records
		var qrRecords int32
		if o.Records != nil {
			qrRecords = *o.Records
		}
		qRecords := swag.FormatInt32(qrRecords)
		if qRecords != "" {
			if err := r.SetQueryParam("records", qRecords); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
