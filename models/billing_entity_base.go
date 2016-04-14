package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*BillingEntityBase Billing entities are models in the BillForward system, of objects in the real-world or otherwise.

swagger:model BillingEntityBase
*/
type BillingEntityBase struct {

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`
}

// Validate validates this billing entity base
func (m *BillingEntityBase) Validate(formats strfmt.Registry) error {
	return nil
}
