// Code generated by go-swagger; DO NOT EDIT.

package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// CreateAddressReader is a Reader for the CreateAddress structure.
type CreateAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAddressOK creates a CreateAddressOK with default headers values
func NewCreateAddressOK() *CreateAddressOK {
	return &CreateAddressOK{}
}

/*CreateAddressOK handles this case with default header values.

success
*/
type CreateAddressOK struct {
	Payload *models.AddressPagedMetadata
}

func (o *CreateAddressOK) Error() string {
	return fmt.Sprintf("[POST /addresses][%d] createAddressOK  %+v", 200, o.Payload)
}

func (o *CreateAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAddressDefault creates a CreateAddressDefault with default headers values
func NewCreateAddressDefault(code int) *CreateAddressDefault {
	return &CreateAddressDefault{
		_statusCode: code,
	}
}

/*CreateAddressDefault handles this case with default header values.

error
*/
type CreateAddressDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the create address default response
func (o *CreateAddressDefault) Code() int {
	return o._statusCode
}

func (o *CreateAddressDefault) Error() string {
	return fmt.Sprintf("[POST /addresses][%d] createAddress default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
