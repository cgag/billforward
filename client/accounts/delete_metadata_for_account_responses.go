package accounts

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

type DeleteMetadataForAccountReader struct {
	formats strfmt.Registry
}

func (o *DeleteMetadataForAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMetadataForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMetadataForAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteMetadataForAccountOK creates a DeleteMetadataForAccountOK with default headers values
func NewDeleteMetadataForAccountOK() *DeleteMetadataForAccountOK {
	return &DeleteMetadataForAccountOK{}
}

/*DeleteMetadataForAccountOK

success
*/
type DeleteMetadataForAccountOK struct {
	Payload *models.MetadataKeyValuesPagedMetadata
}

func (o *DeleteMetadataForAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{account-ID}/metadata][%d] deleteMetadataForAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteMetadataForAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataKeyValuesPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetadataForAccountDefault creates a DeleteMetadataForAccountDefault with default headers values
func NewDeleteMetadataForAccountDefault(code int) *DeleteMetadataForAccountDefault {
	return &DeleteMetadataForAccountDefault{
		_statusCode: code,
	}
}

/*DeleteMetadataForAccountDefault

error
*/
type DeleteMetadataForAccountDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the delete metadata for account default response
func (o *DeleteMetadataForAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMetadataForAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{account-ID}/metadata][%d] deleteMetadataForAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMetadataForAccountDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
