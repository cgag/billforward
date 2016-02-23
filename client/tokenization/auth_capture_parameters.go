package tokenization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// NewAuthCaptureParams creates a new AuthCaptureParams object
// with the default values initialized.
func NewAuthCaptureParams() *AuthCaptureParams {
	var ()
	return &AuthCaptureParams{}
}

/*AuthCaptureParams contains all the parameters to send to the API endpoint
for the auth capture operation typically these are written to a http.Request
*/
type AuthCaptureParams struct {

	/*AuthCaptureRequest
	  The auth capture request.

	*/
	AuthCaptureRequest models.AuthCaptureRequest
}

// WithAuthCaptureRequest adds the authCaptureRequest to the auth capture params
func (o *AuthCaptureParams) WithAuthCaptureRequest(authCaptureRequest models.AuthCaptureRequest) *AuthCaptureParams {
	o.AuthCaptureRequest = authCaptureRequest
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *AuthCaptureParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.AuthCaptureRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
