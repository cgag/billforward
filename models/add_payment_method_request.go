// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AddPaymentMethodRequest AddPaymentMethodRequest
// swagger:model AddPaymentMethodRequest
type AddPaymentMethodRequest struct {

	// { "description" : "ID of the payment method to add", "verbs":["POST","GET"] }
	ID string `json:"id,omitempty"`
}

// Validate validates this add payment method request
func (m *AddPaymentMethodRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AddPaymentMethodRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddPaymentMethodRequest) UnmarshalBinary(b []byte) error {
	var res AddPaymentMethodRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
