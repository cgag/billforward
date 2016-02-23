package invoices

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

// GetInvoiceAsPDFReader is a Reader for the GetInvoiceAsPDF structure.
type GetInvoiceAsPDFReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetInvoiceAsPDFReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceAsPDFOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetInvoiceAsPDFDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetInvoiceAsPDFOK creates a GetInvoiceAsPDFOK with default headers values
func NewGetInvoiceAsPDFOK() *GetInvoiceAsPDFOK {
	return &GetInvoiceAsPDFOK{}
}

/*GetInvoiceAsPDFOK handles this case with default header values.

success
*/
type GetInvoiceAsPDFOK struct {
	Payload httpkit.File
}

func (o *GetInvoiceAsPDFOK) Error() string {
	return fmt.Sprintf("[GET /invoices/{ID}.pdf][%d] getInvoiceAsPDFOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceAsPDFOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceAsPDFDefault creates a GetInvoiceAsPDFDefault with default headers values
func NewGetInvoiceAsPDFDefault(code int) *GetInvoiceAsPDFDefault {
	return &GetInvoiceAsPDFDefault{
		_statusCode: code,
	}
}

/*GetInvoiceAsPDFDefault handles this case with default header values.

error
*/
type GetInvoiceAsPDFDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get invoice as p d f default response
func (o *GetInvoiceAsPDFDefault) Code() int {
	return o._statusCode
}

func (o *GetInvoiceAsPDFDefault) Error() string {
	return fmt.Sprintf("[GET /invoices/{ID}.pdf][%d] getInvoiceAsPDF default  %+v", o._statusCode, o.Payload)
}

func (o *GetInvoiceAsPDFDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
