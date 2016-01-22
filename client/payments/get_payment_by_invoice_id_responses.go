package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

type GetPaymentByInvoiceIDReader struct {
	formats strfmt.Registry
}

func (o *GetPaymentByInvoiceIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentByInvoiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPaymentByInvoiceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPaymentByInvoiceIDOK creates a GetPaymentByInvoiceIDOK with default headers values
func NewGetPaymentByInvoiceIDOK() *GetPaymentByInvoiceIDOK {
	return &GetPaymentByInvoiceIDOK{}
}

/*GetPaymentByInvoiceIDOK

success
*/
type GetPaymentByInvoiceIDOK struct {
	Payload *models.PaymentPagedMetadata
}

func (o *GetPaymentByInvoiceIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/invoice/{invoice-ID}][%d] getPaymentByInvoiceIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentByInvoiceIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentByInvoiceIDDefault creates a GetPaymentByInvoiceIDDefault with default headers values
func NewGetPaymentByInvoiceIDDefault(code int) *GetPaymentByInvoiceIDDefault {
	return &GetPaymentByInvoiceIDDefault{
		_statusCode: code,
	}
}

/*GetPaymentByInvoiceIDDefault

error
*/
type GetPaymentByInvoiceIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get payment by invoice ID default response
func (o *GetPaymentByInvoiceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPaymentByInvoiceIDDefault) Error() string {
	return fmt.Sprintf("[GET /payments/invoice/{invoice-ID}][%d] getPaymentByInvoiceID default  %+v", o._statusCode, o.Payload)
}

func (o *GetPaymentByInvoiceIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}