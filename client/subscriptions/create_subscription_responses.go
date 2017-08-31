package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// CreateSubscriptionReader is a Reader for the CreateSubscription structure.
type CreateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewCreateSubscriptionOK creates a CreateSubscriptionOK with default headers values
func NewCreateSubscriptionOK() *CreateSubscriptionOK {
	return &CreateSubscriptionOK{}
}

/*CreateSubscriptionOK handles this case with default header values.

success
*/
type CreateSubscriptionOK struct {
	Payload *models.SubscriptionPagedMetadata
}

func (o *CreateSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] createSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CreateSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubscriptionDefault creates a CreateSubscriptionDefault with default headers values
func NewCreateSubscriptionDefault(code int) *CreateSubscriptionDefault {
	return &CreateSubscriptionDefault{
		_statusCode: code,
	}
}

/*CreateSubscriptionDefault handles this case with default header values.

error
*/
type CreateSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the create subscription default response
func (o *CreateSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] createSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
