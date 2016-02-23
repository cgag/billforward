package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// GetPaymentMethodByIDReader is a Reader for the GetPaymentMethodByID structure.
type GetPaymentMethodByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPaymentMethodByIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPaymentMethodByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPaymentMethodByIDOK creates a GetPaymentMethodByIDOK with default headers values
func NewGetPaymentMethodByIDOK() *GetPaymentMethodByIDOK {
	return &GetPaymentMethodByIDOK{}
}

/*GetPaymentMethodByIDOK handles this case with default header values.

success
*/
type GetPaymentMethodByIDOK struct {
	Payload *models.PaymentMethodPagedMetadata
}

func (o *GetPaymentMethodByIDOK) Error() string {
	return fmt.Sprintf("[GET /payment-methods/{payment-method-ID}][%d] getPaymentMethodByIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodByIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethodPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodByIDDefault creates a GetPaymentMethodByIDDefault with default headers values
func NewGetPaymentMethodByIDDefault(code int) *GetPaymentMethodByIDDefault {
	return &GetPaymentMethodByIDDefault{
		_statusCode: code,
	}
}

/*GetPaymentMethodByIDDefault handles this case with default header values.

error
*/
type GetPaymentMethodByIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get payment method by ID default response
func (o *GetPaymentMethodByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPaymentMethodByIDDefault) Error() string {
	return fmt.Sprintf("[GET /payment-methods/{payment-method-ID}][%d] getPaymentMethodByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetPaymentMethodByIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
