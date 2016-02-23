package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// NewCreateSubscriptionV2Params creates a new CreateSubscriptionV2Params object
// with the default values initialized.
func NewCreateSubscriptionV2Params() *CreateSubscriptionV2Params {
	var ()
	return &CreateSubscriptionV2Params{}
}

/*CreateSubscriptionV2Params contains all the parameters to send to the API endpoint
for the create subscription v2 operation typically these are written to a http.Request
*/
type CreateSubscriptionV2Params struct {

	/*Request*/
	Request *models.CreateSubscriptionRequest
}

// WithRequest adds the request to the create subscription v2 params
func (o *CreateSubscriptionV2Params) WithRequest(request *models.CreateSubscriptionRequest) *CreateSubscriptionV2Params {
	o.Request = request
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionV2Params) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateSubscriptionRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
