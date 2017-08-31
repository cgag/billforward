package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// GetPaymentMethodByAccountIDReader is a Reader for the GetPaymentMethodByAccountID structure.
type GetPaymentMethodByAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPaymentMethodByAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodByAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPaymentMethodByAccountIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPaymentMethodByAccountIDOK creates a GetPaymentMethodByAccountIDOK with default headers values
func NewGetPaymentMethodByAccountIDOK() *GetPaymentMethodByAccountIDOK {
	return &GetPaymentMethodByAccountIDOK{}
}

/*GetPaymentMethodByAccountIDOK handles this case with default header values.

success
*/
type GetPaymentMethodByAccountIDOK struct {
	Payload *models.PaymentMethodPagedMetadata
}

func (o *GetPaymentMethodByAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /payment-methods/account/{account-ID}][%d] getPaymentMethodByAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodByAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethodPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodByAccountIDDefault creates a GetPaymentMethodByAccountIDDefault with default headers values
func NewGetPaymentMethodByAccountIDDefault(code int) *GetPaymentMethodByAccountIDDefault {
	return &GetPaymentMethodByAccountIDDefault{
		_statusCode: code,
	}
}

/*GetPaymentMethodByAccountIDDefault handles this case with default header values.

error
*/
type GetPaymentMethodByAccountIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get payment method by account ID default response
func (o *GetPaymentMethodByAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPaymentMethodByAccountIDDefault) Error() string {
	return fmt.Sprintf("[GET /payment-methods/account/{account-ID}][%d] getPaymentMethodByAccountID default  %+v", o._statusCode, o.Payload)
}

func (o *GetPaymentMethodByAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
