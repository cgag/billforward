package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// UpsertMetadataForRatePlanReader is a Reader for the UpsertMetadataForRatePlan structure.
type UpsertMetadataForRatePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpsertMetadataForRatePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpsertMetadataForRatePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpsertMetadataForRatePlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpsertMetadataForRatePlanOK creates a UpsertMetadataForRatePlanOK with default headers values
func NewUpsertMetadataForRatePlanOK() *UpsertMetadataForRatePlanOK {
	return &UpsertMetadataForRatePlanOK{}
}

/*UpsertMetadataForRatePlanOK handles this case with default header values.

success
*/
type UpsertMetadataForRatePlanOK struct {
	Payload models.DynamicMetadata
}

func (o *UpsertMetadataForRatePlanOK) Error() string {
	return fmt.Sprintf("[PUT /product-rate-plans/{product-rate-plan-ID}/metadata][%d] upsertMetadataForRatePlanOK  %+v", 200, o.Payload)
}

func (o *UpsertMetadataForRatePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertMetadataForRatePlanDefault creates a UpsertMetadataForRatePlanDefault with default headers values
func NewUpsertMetadataForRatePlanDefault(code int) *UpsertMetadataForRatePlanDefault {
	return &UpsertMetadataForRatePlanDefault{
		_statusCode: code,
	}
}

/*UpsertMetadataForRatePlanDefault handles this case with default header values.

error
*/
type UpsertMetadataForRatePlanDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the upsert metadata for rate plan default response
func (o *UpsertMetadataForRatePlanDefault) Code() int {
	return o._statusCode
}

func (o *UpsertMetadataForRatePlanDefault) Error() string {
	return fmt.Sprintf("[PUT /product-rate-plans/{product-rate-plan-ID}/metadata][%d] upsertMetadataForRatePlan default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertMetadataForRatePlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
