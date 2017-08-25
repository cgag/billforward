package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// NewCreateWebhookV2Params creates a new CreateWebhookV2Params object
// with the default values initialized.
func NewCreateWebhookV2Params() *CreateWebhookV2Params {
	var ()
	return &CreateWebhookV2Params{}
}

/*CreateWebhookV2Params contains all the parameters to send to the API endpoint
for the create webhook v2 operation typically these are written to a http.Request
*/
type CreateWebhookV2Params struct {

	/*Request*/
	Request *models.CreateWebhookRequest
}

// WithRequest adds the request to the create webhook v2 params
func (o *CreateWebhookV2Params) WithRequest(Request *models.CreateWebhookRequest) *CreateWebhookV2Params {
	o.Request = Request
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWebhookV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateWebhookRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}