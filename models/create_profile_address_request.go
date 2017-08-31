package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CreateProfileAddressRequest create profile address request

swagger:model CreateProfileAddressRequest
*/
type CreateProfileAddressRequest struct {

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	AddressLine1 *string `json:"addressLine1"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine2 string `json:"addressLine2,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine3 string `json:"addressLine3,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	City *string `json:"city"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Country *string `json:"country"`

	/* { "description" : "Phone number", "verbs":["POST","PUT","GET"] }
	 */
	Landline string `json:"landline,omitempty"`

	/* { "description" : "ZIP code or postcode.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Postcode *string `json:"postcode"`

	/* { "description" : "Is this the primary, default, address for the associated profile?", "verbs":["POST","PUT","GET"] }
	 */
	PrimaryAddress *bool `json:"primaryAddress,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Province *string `json:"province"`
}

// Validate validates this create profile address request
func (m *CreateProfileAddressRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePostcode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProvince(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProfileAddressRequest) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	return nil
}

func (m *CreateProfileAddressRequest) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *CreateProfileAddressRequest) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *CreateProfileAddressRequest) validatePostcode(formats strfmt.Registry) error {

	if err := validate.Required("postcode", "body", m.Postcode); err != nil {
		return err
	}

	return nil
}

func (m *CreateProfileAddressRequest) validateProvince(formats strfmt.Registry) error {

	if err := validate.Required("province", "body", m.Province); err != nil {
		return err
	}

	return nil
}
