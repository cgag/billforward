// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// GetMetadataForProductReader is a Reader for the GetMetadataForProduct structure.
type GetMetadataForProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetadataForProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetadataForProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMetadataForProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetadataForProductOK creates a GetMetadataForProductOK with default headers values
func NewGetMetadataForProductOK() *GetMetadataForProductOK {
	return &GetMetadataForProductOK{}
}

/*GetMetadataForProductOK handles this case with default header values.

success
*/
type GetMetadataForProductOK struct {
	Payload models.DynamicMetadata
}

func (o *GetMetadataForProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{product-ID}/metadata][%d] getMetadataForProductOK  %+v", 200, o.Payload)
}

func (o *GetMetadataForProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetadataForProductDefault creates a GetMetadataForProductDefault with default headers values
func NewGetMetadataForProductDefault(code int) *GetMetadataForProductDefault {
	return &GetMetadataForProductDefault{
		_statusCode: code,
	}
}

/*GetMetadataForProductDefault handles this case with default header values.

error
*/
type GetMetadataForProductDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get metadata for product default response
func (o *GetMetadataForProductDefault) Code() int {
	return o._statusCode
}

func (o *GetMetadataForProductDefault) Error() string {
	return fmt.Sprintf("[GET /products/{product-ID}/metadata][%d] getMetadataForProduct default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetadataForProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
