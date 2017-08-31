package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// NewCreateAccountParams creates a new CreateAccountParams object
// with the default values initialized.
func NewCreateAccountParams() *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountParamsWithTimeout creates a new CreateAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccountParamsWithTimeout(timeout time.Duration) *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		timeout: timeout,
	}
}

/*CreateAccountParams contains all the parameters to send to the API endpoint
for the create account operation typically these are written to a http.Request
*/
type CreateAccountParams struct {

	/*Request
	  The account object to be created.

	*/
	Request *models.CreateAccountRequest

	timeout time.Duration
}

// WithRequest adds the request to the create account params
func (o *CreateAccountParams) WithRequest(request *models.CreateAccountRequest) *CreateAccountParams {
	o.Request = request
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateAccountRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
