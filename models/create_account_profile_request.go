// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateAccountProfileRequest create account profile request
// swagger:model CreateAccountProfileRequest

type CreateAccountProfileRequest struct {

	// { "description" : "", "verbs":["GET"] }
	AccountID string `json:"accountID,omitempty"`

	// { "description" : "Any additional information", "verbs":["POST","PUT","GET"] }
	AdditionalInformation string `json:"additionalInformation,omitempty"`

	// { "description" : "Address associated with the profile", "verbs":["POST","PUT","GET"] }
	Addresses []*CreateProfileAddressRequest `json:"addresses"`

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	CompanyName string `json:"companyName,omitempty"`

	// { "description" : "Date of birth in YYYY-MM-DD format", "verbs":["POST","PUT","GET"] }
	Dob *strfmt.DateTime `json:"dob,omitempty"`

	// { "description" : "E-mail address", "verbs":["POST","PUT","GET"] }
	Email string `json:"email,omitempty"`

	// { "description" : "Fax number", "verbs":["POST","PUT","GET"] }
	Fax string `json:"fax,omitempty"`

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	FirstName string `json:"firstName,omitempty"`

	// { "description" : "Home telephone number", "verbs":["POST","PUT","GET"] }
	Landline string `json:"landline,omitempty"`

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	LastName string `json:"lastName,omitempty"`

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	LogoURL string `json:"logoURL,omitempty"`

	// { "description" : "Mobile telephone number", "verbs":["POST","PUT","GET"] }
	Mobile string `json:"mobile,omitempty"`

	// { "description" : "VAT number", "verbs":["POST","PUT","GET"] }
	VatNumber string `json:"vatNumber,omitempty"`
}

/* polymorph CreateAccountProfileRequest accountID false */

/* polymorph CreateAccountProfileRequest additionalInformation false */

/* polymorph CreateAccountProfileRequest addresses false */

/* polymorph CreateAccountProfileRequest companyName false */

/* polymorph CreateAccountProfileRequest dob false */

/* polymorph CreateAccountProfileRequest email false */

/* polymorph CreateAccountProfileRequest fax false */

/* polymorph CreateAccountProfileRequest firstName false */

/* polymorph CreateAccountProfileRequest landline false */

/* polymorph CreateAccountProfileRequest lastName false */

/* polymorph CreateAccountProfileRequest logoURL false */

/* polymorph CreateAccountProfileRequest mobile false */

/* polymorph CreateAccountProfileRequest vatNumber false */

// Validate validates this create account profile request
func (m *CreateAccountProfileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAccountProfileRequest) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {

		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {

			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAccountProfileRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAccountProfileRequest) UnmarshalBinary(b []byte) error {
	var res CreateAccountProfileRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
