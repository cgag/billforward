// Code generated by go-swagger; DO NOT EDIT.

package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// GetProfileByEmailAddressReader is a Reader for the GetProfileByEmailAddress structure.
type GetProfileByEmailAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfileByEmailAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfileByEmailAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProfileByEmailAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProfileByEmailAddressOK creates a GetProfileByEmailAddressOK with default headers values
func NewGetProfileByEmailAddressOK() *GetProfileByEmailAddressOK {
	return &GetProfileByEmailAddressOK{}
}

/*GetProfileByEmailAddressOK handles this case with default header values.

success
*/
type GetProfileByEmailAddressOK struct {
	Payload *models.ProfilePagedMetadata
}

func (o *GetProfileByEmailAddressOK) Error() string {
	return fmt.Sprintf("[GET /profiles/email/{email}][%d] getProfileByEmailAddressOK  %+v", 200, o.Payload)
}

func (o *GetProfileByEmailAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfilePagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileByEmailAddressDefault creates a GetProfileByEmailAddressDefault with default headers values
func NewGetProfileByEmailAddressDefault(code int) *GetProfileByEmailAddressDefault {
	return &GetProfileByEmailAddressDefault{
		_statusCode: code,
	}
}

/*GetProfileByEmailAddressDefault handles this case with default header values.

error
*/
type GetProfileByEmailAddressDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get profile by email address default response
func (o *GetProfileByEmailAddressDefault) Code() int {
	return o._statusCode
}

func (o *GetProfileByEmailAddressDefault) Error() string {
	return fmt.Sprintf("[GET /profiles/email/{email}][%d] getProfileByEmailAddress default  %+v", o._statusCode, o.Payload)
}

func (o *GetProfileByEmailAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
