package subscriptions

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

type UpsertMetadataForSubscriptionReader struct {
	formats strfmt.Registry
}

func (o *UpsertMetadataForSubscriptionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpsertMetadataForSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpsertMetadataForSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpsertMetadataForSubscriptionOK creates a UpsertMetadataForSubscriptionOK with default headers values
func NewUpsertMetadataForSubscriptionOK() *UpsertMetadataForSubscriptionOK {
	return &UpsertMetadataForSubscriptionOK{}
}

/*UpsertMetadataForSubscriptionOK

success
*/
type UpsertMetadataForSubscriptionOK struct {
	Payload *models.MetadataKeyValuesPagedMetadata
}

func (o *UpsertMetadataForSubscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscription-ID}/metadata][%d] upsertMetadataForSubscriptionOK  %+v", 200, o.Payload)
}

func (o *UpsertMetadataForSubscriptionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataKeyValuesPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertMetadataForSubscriptionDefault creates a UpsertMetadataForSubscriptionDefault with default headers values
func NewUpsertMetadataForSubscriptionDefault(code int) *UpsertMetadataForSubscriptionDefault {
	return &UpsertMetadataForSubscriptionDefault{
		_statusCode: code,
	}
}

/*UpsertMetadataForSubscriptionDefault

error
*/
type UpsertMetadataForSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the upsert metadata for subscription default response
func (o *UpsertMetadataForSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *UpsertMetadataForSubscriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscription-ID}/metadata][%d] upsertMetadataForSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertMetadataForSubscriptionDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
