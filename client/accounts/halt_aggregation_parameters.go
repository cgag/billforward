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

	return &HaltAggregationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHaltAggregationParamsWithTimeout creates a new HaltAggregationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHaltAggregationParamsWithTimeout(timeout time.Duration) *HaltAggregationParams {

	return &HaltAggregationParams{

		timeout: timeout,
	}
}

// NewHaltAggregationParamsWithContext creates a new HaltAggregationParams object
// with the default values initialized, and the ability to set a context for a request
func NewHaltAggregationParamsWithContext(ctx context.Context) *HaltAggregationParams {

	return &HaltAggregationParams{

		Context: ctx,
	}
}

// NewHaltAggregationParamsWithHTTPClient creates a new HaltAggregationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHaltAggregationParamsWithHTTPClient(client *http.Client) *HaltAggregationParams {

	return &HaltAggregationParams{
		HTTPClient: client,
	}
}

/*HaltAggregationParams contains all the parameters to send to the API endpoint
for the halt aggregation operation typically these are written to a http.Request
*/
type HaltAggregationParams struct {
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

// WriteToRequest writes these params to a swagger request
func (o *HaltAggregationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
