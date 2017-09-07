// Code generated by go-swagger; DO NOT EDIT.

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

// GetSubscriptionByIDReader is a Reader for the GetSubscriptionByID structure.
type GetSubscriptionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSubscriptionByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionByIDOK creates a GetSubscriptionByIDOK with default headers values
func NewGetSubscriptionByIDOK() *GetSubscriptionByIDOK {
	return &GetSubscriptionByIDOK{}
}

/*GetSubscriptionByIDOK handles this case with default header values.

success
*/
type GetSubscriptionByIDOK struct {
	Payload *models.SubscriptionPagedMetadata
}

func (o *GetSubscriptionByIDOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscription-ID}][%d] getSubscriptionByIdOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionByIDDefault creates a GetSubscriptionByIDDefault with default headers values
func NewGetSubscriptionByIDDefault(code int) *GetSubscriptionByIDDefault {
	return &GetSubscriptionByIDDefault{
		_statusCode: code,
	}
}

/*GetSubscriptionByIDDefault handles this case with default header values.

error
*/
type GetSubscriptionByIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get subscription by ID default response
func (o *GetSubscriptionByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionByIDDefault) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscription-ID}][%d] getSubscriptionByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubscriptionByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
