package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// GetInvoicesByAccountIDReader is a Reader for the GetInvoicesByAccountID structure.
type GetInvoicesByAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetInvoicesByAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicesByAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetInvoicesByAccountIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetInvoicesByAccountIDOK creates a GetInvoicesByAccountIDOK with default headers values
func NewGetInvoicesByAccountIDOK() *GetInvoicesByAccountIDOK {
	return &GetInvoicesByAccountIDOK{}
}

/*GetInvoicesByAccountIDOK handles this case with default header values.

success
*/
type GetInvoicesByAccountIDOK struct {
	Payload *models.InvoicePagedMetadata
}

func (o *GetInvoicesByAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /invoices/account/{account-ID}][%d] getInvoicesByAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesByAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoicePagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesByAccountIDDefault creates a GetInvoicesByAccountIDDefault with default headers values
func NewGetInvoicesByAccountIDDefault(code int) *GetInvoicesByAccountIDDefault {
	return &GetInvoicesByAccountIDDefault{
		_statusCode: code,
	}
}

/*GetInvoicesByAccountIDDefault handles this case with default header values.

error
*/
type GetInvoicesByAccountIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get invoices by account ID default response
func (o *GetInvoicesByAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *GetInvoicesByAccountIDDefault) Error() string {
	return fmt.Sprintf("[GET /invoices/account/{account-ID}][%d] getInvoicesByAccountID default  %+v", o._statusCode, o.Payload)
}

func (o *GetInvoicesByAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
