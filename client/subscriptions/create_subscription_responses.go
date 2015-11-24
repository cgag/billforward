package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

type CreateSubscriptionReader struct {
	formats strfmt.Registry
}

func (o *CreateSubscriptionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result CreateSubscriptionOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		var result CreateSubscriptionDefault
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("createSubscription default", &result, response.Code())
	}
}

/*
successful operation
*/
type CreateSubscriptionOK struct {
	Payload *models.SubscriptionQueryResultWrapper
}

func (o *CreateSubscriptionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionQueryResultWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*
an error occurred
*/
type CreateSubscriptionDefault struct {
	Payload *models.GeneralError
}

func (o *CreateSubscriptionDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}