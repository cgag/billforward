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

// UpsertMetadataForProductReader is a Reader for the UpsertMetadataForProduct structure.
type UpsertMetadataForProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertMetadataForProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpsertMetadataForProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpsertMetadataForProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpsertMetadataForProductOK creates a UpsertMetadataForProductOK with default headers values
func NewUpsertMetadataForProductOK() *UpsertMetadataForProductOK {
	return &UpsertMetadataForProductOK{}
}

/*UpsertMetadataForProductOK handles this case with default header values.

success
*/
type UpsertMetadataForProductOK struct {
	Payload models.DynamicMetadata
}

func (o *UpsertMetadataForProductOK) Error() string {
	return fmt.Sprintf("[PUT /products/{product-ID}/metadata][%d] upsertMetadataForProductOK  %+v", 200, o.Payload)
}

func (o *UpsertMetadataForProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertMetadataForProductDefault creates a UpsertMetadataForProductDefault with default headers values
func NewUpsertMetadataForProductDefault(code int) *UpsertMetadataForProductDefault {
	return &UpsertMetadataForProductDefault{
		_statusCode: code,
	}
}

/*UpsertMetadataForProductDefault handles this case with default header values.

error
*/
type UpsertMetadataForProductDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the upsert metadata for product default response
func (o *UpsertMetadataForProductDefault) Code() int {
	return o._statusCode
}

func (o *UpsertMetadataForProductDefault) Error() string {
	return fmt.Sprintf("[PUT /products/{product-ID}/metadata][%d] upsertMetadataForProduct default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertMetadataForProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
