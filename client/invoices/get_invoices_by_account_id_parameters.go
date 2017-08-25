// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewGetInvoicesByAccountIDParams creates a new GetInvoicesByAccountIDParams object
// with the default values initialized.
func NewGetInvoicesByAccountIDParams() *GetInvoicesByAccountIDParams {
	var (
		excludeChildrenDefault = bool(true)
		includeRetiredDefault  = bool(false)
		offsetDefault          = int32(0)
		orderDefault           = string("DESC")
		orderByDefault         = string("created")
		recordsDefault         = int32(10)
	)
	return &GetInvoicesByAccountIDParams{
		ExcludeChildren: &excludeChildrenDefault,
		IncludeRetired:  &includeRetiredDefault,
		Offset:          &offsetDefault,
		Order:           &orderDefault,
		OrderBy:         &orderByDefault,
		Records:         &recordsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoicesByAccountIDParamsWithTimeout creates a new GetInvoicesByAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoicesByAccountIDParamsWithTimeout(timeout time.Duration) *GetInvoicesByAccountIDParams {
	var (
		excludeChildrenDefault = bool(true)
		includeRetiredDefault  = bool(false)
		offsetDefault          = int32(0)
		orderDefault           = string("DESC")
		orderByDefault         = string("created")
		recordsDefault         = int32(10)
	)
	return &GetInvoicesByAccountIDParams{
		ExcludeChildren: &excludeChildrenDefault,
		IncludeRetired:  &includeRetiredDefault,
		Offset:          &offsetDefault,
		Order:           &orderDefault,
		OrderBy:         &orderByDefault,
		Records:         &recordsDefault,

		timeout: timeout,
	}
}

// NewGetInvoicesByAccountIDParamsWithContext creates a new GetInvoicesByAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoicesByAccountIDParamsWithContext(ctx context.Context) *GetInvoicesByAccountIDParams {
	var (
		excludeChildrenDefault = bool(true)
		includeRetiredDefault  = bool(false)
		offsetDefault          = int32(0)
		orderDefault           = string("DESC")
		orderByDefault         = string("created")
		recordsDefault         = int32(10)
	)
	return &GetInvoicesByAccountIDParams{
		ExcludeChildren: &excludeChildrenDefault,
		IncludeRetired:  &includeRetiredDefault,
		Offset:          &offsetDefault,
		Order:           &orderDefault,
		OrderBy:         &orderByDefault,
		Records:         &recordsDefault,

		Context: ctx,
	}
}

// NewGetInvoicesByAccountIDParamsWithHTTPClient creates a new GetInvoicesByAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoicesByAccountIDParamsWithHTTPClient(client *http.Client) *GetInvoicesByAccountIDParams {
	var (
		excludeChildrenDefault = bool(true)
		includeRetiredDefault  = bool(false)
		offsetDefault          = int32(0)
		orderDefault           = string("DESC")
		orderByDefault         = string("created")
		recordsDefault         = int32(10)
	)
	return &GetInvoicesByAccountIDParams{
		ExcludeChildren: &excludeChildrenDefault,
		IncludeRetired:  &includeRetiredDefault,
		Offset:          &offsetDefault,
		Order:           &orderDefault,
		OrderBy:         &orderByDefault,
		Records:         &recordsDefault,
		HTTPClient:      client,
	}
}

/*GetInvoicesByAccountIDParams contains all the parameters to send to the API endpoint
for the get invoices by account ID operation typically these are written to a http.Request
*/
type GetInvoicesByAccountIDParams struct {

	/*AccountID
	  ID of the account.

	*/
	AccountID string
	/*ExcludeChildren
	  Should child invoices be excluded.

	*/
	ExcludeChildren *bool
	/*IncludeRetired
	  Whether retired products should be returned.

	*/
	IncludeRetired *bool
	/*Offset
	  The offset from the first invoice to return.

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
	  The maximum number of invoices to return.

	*/
	Records *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithTimeout(timeout time.Duration) *GetInvoicesByAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithContext(ctx context.Context) *GetInvoicesByAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithHTTPClient(client *http.Client) *GetInvoicesByAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithAccountID(accountID string) *GetInvoicesByAccountIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExcludeChildren adds the excludeChildren to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithExcludeChildren(excludeChildren *bool) *GetInvoicesByAccountIDParams {
	o.SetExcludeChildren(excludeChildren)
	return o
}

// SetExcludeChildren adds the excludeChildren to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetExcludeChildren(excludeChildren *bool) {
	o.ExcludeChildren = excludeChildren
}

// WithIncludeRetired adds the includeRetired to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithIncludeRetired(includeRetired *bool) *GetInvoicesByAccountIDParams {
	o.SetIncludeRetired(includeRetired)
	return o
}

// SetIncludeRetired adds the includeRetired to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetIncludeRetired(includeRetired *bool) {
	o.IncludeRetired = includeRetired
}

// WithOffset adds the offset to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithOffset(offset *int32) *GetInvoicesByAccountIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrder adds the order to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithOrder(order *string) *GetInvoicesByAccountIDParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetOrder(order *string) {
	o.Order = order
}

// WithOrderBy adds the orderBy to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithOrderBy(orderBy *string) *GetInvoicesByAccountIDParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizations adds the organizations to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithOrganizations(organizations []string) *GetInvoicesByAccountIDParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WithRecords adds the records to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) WithRecords(records *int32) *GetInvoicesByAccountIDParams {
	o.SetRecords(records)
	return o
}

// SetRecords adds the records to the get invoices by account ID params
func (o *GetInvoicesByAccountIDParams) SetRecords(records *int32) {
	o.Records = records
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoicesByAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account-ID
	if err := r.SetPathParam("account-ID", o.AccountID); err != nil {
		return err
	}

	if o.ExcludeChildren != nil {

		// query param exclude_children
		var qrExcludeChildren bool
		if o.ExcludeChildren != nil {
			qrExcludeChildren = *o.ExcludeChildren
		}
		qExcludeChildren := swag.FormatBool(qrExcludeChildren)
		if qExcludeChildren != "" {
			if err := r.SetQueryParam("exclude_children", qExcludeChildren); err != nil {
				return err
			}
		}

	}

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
