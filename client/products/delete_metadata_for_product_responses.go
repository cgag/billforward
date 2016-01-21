package products

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

type DeleteMetadataForProductReader struct {
	formats strfmt.Registry
}

func (o *DeleteMetadataForProductReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMetadataForProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMetadataForProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteMetadataForProductOK creates a DeleteMetadataForProductOK with default headers values
func NewDeleteMetadataForProductOK() *DeleteMetadataForProductOK {
	return &DeleteMetadataForProductOK{}
}

/*DeleteMetadataForProductOK

success
*/
type DeleteMetadataForProductOK struct {
	Payload *models.DynamicMetadata
}

func (o *DeleteMetadataForProductOK) Error() string {
	return fmt.Sprintf("[DELETE /products/{product-ID}/metadata][%d] deleteMetadataForProductOK  %+v", 200, o.Payload)
}

func (o *DeleteMetadataForProductOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DynamicMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetadataForProductDefault creates a DeleteMetadataForProductDefault with default headers values
func NewDeleteMetadataForProductDefault(code int) *DeleteMetadataForProductDefault {
	return &DeleteMetadataForProductDefault{
		_statusCode: code,
	}
}

/*DeleteMetadataForProductDefault

error
*/
type DeleteMetadataForProductDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the delete metadata for product default response
func (o *DeleteMetadataForProductDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMetadataForProductDefault) Error() string {
	return fmt.Sprintf("[DELETE /products/{product-ID}/metadata][%d] deleteMetadataForProduct default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMetadataForProductDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
