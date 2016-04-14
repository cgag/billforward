package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetAllProductsParams creates a new GetAllProductsParams object
// with the default values initialized.
func NewGetAllProductsParams() *GetAllProductsParams {
	var (
		includeRetiredDefault bool   = bool(true)
		offsetDefault         int32  = int32(0)
		orderDefault          string = string("DESC")
		orderByDefault        string = string("created")
		recordsDefault        int32  = int32(10)
	)
	return &GetAllProductsParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,
	}
}

/*GetAllProductsParams contains all the parameters to send to the API endpoint
for the get all products operation typically these are written to a http.Request
*/
type GetAllProductsParams struct {

	/*IncludeRetired
	  Whether retired products should be returned.

	*/
	IncludeRetired *bool
	/*Metadata*/
	Metadata *string
	/*Offset
	  The offset from the first product to return.

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
	/*Records
	  The maximum number of products to return.

	*/
	Records *int32
}

// WithIncludeRetired adds the includeRetired to the get all products params
func (o *GetAllProductsParams) WithIncludeRetired(includeRetired *bool) *GetAllProductsParams {
	o.IncludeRetired = includeRetired
	return o
}

// WithMetadata adds the metadata to the get all products params
func (o *GetAllProductsParams) WithMetadata(metadata *string) *GetAllProductsParams {
	o.Metadata = metadata
	return o
}

// WithOffset adds the offset to the get all products params
func (o *GetAllProductsParams) WithOffset(offset *int32) *GetAllProductsParams {
	o.Offset = offset
	return o
}

// WithOrder adds the order to the get all products params
func (o *GetAllProductsParams) WithOrder(order *string) *GetAllProductsParams {
	o.Order = order
	return o
}

// WithOrderBy adds the orderBy to the get all products params
func (o *GetAllProductsParams) WithOrderBy(orderBy *string) *GetAllProductsParams {
	o.OrderBy = orderBy
	return o
}

// WithOrganizations adds the organizations to the get all products params
func (o *GetAllProductsParams) WithOrganizations(organizations []string) *GetAllProductsParams {
	o.Organizations = organizations
	return o
}

// WithRecords adds the records to the get all products params
func (o *GetAllProductsParams) WithRecords(records *int32) *GetAllProductsParams {
	o.Records = records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllProductsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata string
		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := qrMetadata
		if qMetadata != "" {
			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
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
