// Code generated by go-swagger; DO NOT EDIT.

package addresses

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

	"github.com/authclub/billforward/models"
)

// NewUpdateAddressParams creates a new UpdateAddressParams object
// with the default values initialized.
func NewUpdateAddressParams() *UpdateAddressParams {
	var ()
	return &UpdateAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAddressParamsWithTimeout creates a new UpdateAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAddressParamsWithTimeout(timeout time.Duration) *UpdateAddressParams {
	var ()
	return &UpdateAddressParams{

		timeout: timeout,
	}
}

// NewUpdateAddressParamsWithContext creates a new UpdateAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAddressParamsWithContext(ctx context.Context) *UpdateAddressParams {
	var ()
	return &UpdateAddressParams{

		Context: ctx,
	}
}

// NewUpdateAddressParamsWithHTTPClient creates a new UpdateAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAddressParamsWithHTTPClient(client *http.Client) *UpdateAddressParams {
	var ()
	return &UpdateAddressParams{
		HTTPClient: client,
	}
}

/*UpdateAddressParams contains all the parameters to send to the API endpoint
for the update address operation typically these are written to a http.Request
*/
type UpdateAddressParams struct {

	/*Request
	  The address object to be created.

	*/
	Request *models.UpdateAddressRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update address params
func (o *UpdateAddressParams) WithTimeout(timeout time.Duration) *UpdateAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update address params
func (o *UpdateAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update address params
func (o *UpdateAddressParams) WithContext(ctx context.Context) *UpdateAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update address params
func (o *UpdateAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update address params
func (o *UpdateAddressParams) WithHTTPClient(client *http.Client) *UpdateAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update address params
func (o *UpdateAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the update address params
func (o *UpdateAddressParams) WithRequest(request *models.UpdateAddressRequest) *UpdateAddressParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update address params
func (o *UpdateAddressParams) SetRequest(request *models.UpdateAddressRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request == nil {
		o.Request = new(models.UpdateAddressRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
