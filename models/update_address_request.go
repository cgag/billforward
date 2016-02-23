package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*UpdateAddressRequest CreateAddressRequest

swagger:model UpdateAddressRequest
*/
type UpdateAddressRequest struct {

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine1 *string `json:"addressLine1,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine2 *string `json:"addressLine2,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	AddressLine3 *string `json:"addressLine3,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	City *string `json:"city,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	Country *string `json:"country,omitempty"`

	/* ID id

	Required: true
	*/
	ID string `json:"id,omitempty"`

	/* { "description" : "Phone number", "verbs":["POST","PUT","GET"] }
	 */
	Landline *string `json:"landline,omitempty"`

	/* { "description" : "ZIP code or postcode.", "verbs":["POST","PUT","GET"] }
	 */
	Postcode *string `json:"postcode,omitempty"`

	/* { "description" : "Is this the primary, default, address for the associated profile?", "verbs":["POST","PUT","GET"] }
	 */
	PrimaryAddress *bool `json:"primaryAddress,omitempty"`

	/* { "description" : "", "verbs":["GET"] }

	Required: true
	*/
	ProfileID string `json:"profileID,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	Province *string `json:"province,omitempty"`
}

// Validate validates this update address request
func (m *UpdateAddressRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateAddressRequest) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *UpdateAddressRequest) validateProfileID(formats strfmt.Registry) error {

	if err := validate.RequiredString("profileID", "body", string(m.ProfileID)); err != nil {
		return err
	}

	return nil
}
