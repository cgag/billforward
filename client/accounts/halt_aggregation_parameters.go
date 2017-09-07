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

	strfmt "github.com/go-openapi/strfmt"
)

// NewHaltAggregationParams creates a new HaltAggregationParams object
// with the default values initialized.
func NewHaltAggregationParams() *HaltAggregationParams {
	var ()
	return &HaltAggregationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHaltAggregationParamsWithTimeout creates a new HaltAggregationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHaltAggregationParamsWithTimeout(timeout time.Duration) *HaltAggregationParams {
	var ()
	return &HaltAggregationParams{

		timeout: timeout,
	}
}

// NewHaltAggregationParamsWithContext creates a new HaltAggregationParams object
// with the default values initialized, and the ability to set a context for a request
func NewHaltAggregationParamsWithContext(ctx context.Context) *HaltAggregationParams {
	var ()
	return &HaltAggregationParams{

		Context: ctx,
	}
}

// NewHaltAggregationParamsWithHTTPClient creates a new HaltAggregationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHaltAggregationParamsWithHTTPClient(client *http.Client) *HaltAggregationParams {
	var ()
	return &HaltAggregationParams{
		HTTPClient: client,
	}
}

/*HaltAggregationParams contains all the parameters to send to the API endpoint
for the halt aggregation operation typically these are written to a http.Request
*/
type HaltAggregationParams struct {

	/*AccountID*/
	AccountID string
	/*Dummy
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Dummy HaltAggregationBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the halt aggregation params
func (o *HaltAggregationParams) WithTimeout(timeout time.Duration) *HaltAggregationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the halt aggregation params
func (o *HaltAggregationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the halt aggregation params
func (o *HaltAggregationParams) WithContext(ctx context.Context) *HaltAggregationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the halt aggregation params
func (o *HaltAggregationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the halt aggregation params
func (o *HaltAggregationParams) WithHTTPClient(client *http.Client) *HaltAggregationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the halt aggregation params
func (o *HaltAggregationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the halt aggregation params
func (o *HaltAggregationParams) WithAccountID(accountID string) *HaltAggregationParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the halt aggregation params
func (o *HaltAggregationParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithDummy adds the dummy to the halt aggregation params
func (o *HaltAggregationParams) WithDummy(dummy HaltAggregationBody) *HaltAggregationParams {
	o.SetDummy(dummy)
	return o
}

// SetDummy adds the dummy to the halt aggregation params
func (o *HaltAggregationParams) SetDummy(dummy HaltAggregationBody) {
	o.Dummy = dummy
}

// WriteToRequest writes these params to a swagger request
func (o *HaltAggregationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account-ID
	if err := r.SetPathParam("account-ID", o.AccountID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Dummy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
