// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// CreateWebhookReader is a Reader for the CreateWebhook structure.
type CreateWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateWebhookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateWebhookOK creates a CreateWebhookOK with default headers values
func NewCreateWebhookOK() *CreateWebhookOK {
	return &CreateWebhookOK{}
}

/*CreateWebhookOK handles this case with default header values.

success
*/
type CreateWebhookOK struct {
	Payload *models.WebhookPagedMetadata
}

func (o *CreateWebhookOK) Error() string {
	return fmt.Sprintf("[POST /webhooks][%d] createWebhookOK  %+v", 200, o.Payload)
}

func (o *CreateWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWebhookDefault creates a CreateWebhookDefault with default headers values
func NewCreateWebhookDefault(code int) *CreateWebhookDefault {
	return &CreateWebhookDefault{
		_statusCode: code,
	}
}

/*CreateWebhookDefault handles this case with default header values.

error
*/
type CreateWebhookDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the create webhook default response
func (o *CreateWebhookDefault) Code() int {
	return o._statusCode
}

func (o *CreateWebhookDefault) Error() string {
	return fmt.Sprintf("[POST /webhooks][%d] createWebhook default  %+v", o._statusCode, o.Payload)
}

func (o *CreateWebhookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
