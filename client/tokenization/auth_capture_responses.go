package tokenization

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

type AuthCaptureReader struct {
	formats strfmt.Registry
}

func (o *AuthCaptureReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAuthCaptureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAuthCaptureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewAuthCaptureOK creates a AuthCaptureOK with default headers values
func NewAuthCaptureOK() *AuthCaptureOK {
	return &AuthCaptureOK{}
}

/*AuthCaptureOK

success
*/
type AuthCaptureOK struct {
	Payload *models.PaymentMethodPagedMetadata
}

func (o *AuthCaptureOK) Error() string {
	return fmt.Sprintf("[POST /tokenization/auth-capture][%d] authCaptureOK  %+v", 200, o.Payload)
}

func (o *AuthCaptureOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethodPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthCaptureDefault creates a AuthCaptureDefault with default headers values
func NewAuthCaptureDefault(code int) *AuthCaptureDefault {
	return &AuthCaptureDefault{
		_statusCode: code,
	}
}

/*AuthCaptureDefault

error
*/
type AuthCaptureDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the auth capture default response
func (o *AuthCaptureDefault) Code() int {
	return o._statusCode
}

func (o *AuthCaptureDefault) Error() string {
	return fmt.Sprintf("[POST /tokenization/auth-capture][%d] authCapture default  %+v", o._statusCode, o.Payload)
}

func (o *AuthCaptureDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
