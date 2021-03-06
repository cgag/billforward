// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReviveSubscriptionRequest ReviveSubscriptionRequest
// swagger:model ReviveSubscriptionRequest
type ReviveSubscriptionRequest struct {

	// subscription ID
	// Required: true
	SubscriptionID *string `json:"subscriptionID"`
}

// Validate validates this revive subscription request
func (m *ReviveSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReviveSubscriptionRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReviveSubscriptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReviveSubscriptionRequest) UnmarshalBinary(b []byte) error {
	var res ReviveSubscriptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
