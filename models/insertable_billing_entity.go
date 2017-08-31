package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*InsertableBillingEntity Insertable entities are those entities that can be created.

swagger:model InsertableBillingEntity
*/
type InsertableBillingEntity struct {

	/* { "description" : "The UTC DateTime when the pricing-component-value-change was processed.", "verbs":["POST","PUT","GET"] }
	 */
	Applied strfmt.DateTime `json:"applied,omitempty"`

	/* { "description" : "The UTC DateTime when the pricing-component-value-change was calculated.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	AsOf *strfmt.DateTime `json:"asOf"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "ID of the pricing-component-value-change.", "verbs":["POST","PUT","GET"] }
	 */
	ID string `json:"id,omitempty"`

	/* { "description" : "ID of the invoice associated with the pricing-component-value-change.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	InvoiceID *string `json:"invoiceID"`

	/* { "description" : "The value change mode.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Mode *string `json:"mode"`

	/* { "description" : "The new value of the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	NewValue *int32 `json:"newValue"`

	/* { "description" : "The new value of the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OldValue *int32 `json:"oldValue"`

	/* { "description" : "The organizationID.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OrganizationID *string `json:"organizationID"`

	/* { "description" : "ID of the pricing-component associated with the pricing-component-value-change.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	PricingComponentID *string `json:"pricingComponentID"`

	/* { "description" : "The value change state.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	State *string `json:"state"`

	/* { "description" : "ID of the subscription associated with the pricing-component-value-change.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	SubscriptionID *string `json:"subscriptionID"`

	/* { "description" : "ID of the unit-of-measure associated with the pricing-component-value-change.", "verbs":["POST","PUT","GET"] }
	 */
	UnitOfMeasureID string `json:"unitOfMeasureID,omitempty"`
}

// Validate validates this insertable billing entity
func (m *InsertableBillingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsOf(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOldValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InsertableBillingEntity) validateAsOf(formats strfmt.Registry) error {

	if err := validate.Required("asOf", "body", m.AsOf); err != nil {
		return err
	}

	return nil
}

func (m *InsertableBillingEntity) validateInvoiceID(formats strfmt.Registry) error {

	if err := validate.Required("invoiceID", "body", m.InvoiceID); err != nil {
		return err
	}

	return nil
}

var insertableBillingEntityTypeModePropEnum []interface{}

// prop value enum
func (m *InsertableBillingEntity) validateModeEnum(path, location string, value string) error {
	if insertableBillingEntityTypeModePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			insertableBillingEntityTypeModePropEnum = append(insertableBillingEntityTypeModePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, insertableBillingEntityTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InsertableBillingEntity) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *InsertableBillingEntity) validateNewValue(formats strfmt.Registry) error {

	if err := validate.Required("newValue", "body", m.NewValue); err != nil {
		return err
	}

	return nil
}

func (m *InsertableBillingEntity) validateOldValue(formats strfmt.Registry) error {

	if err := validate.Required("oldValue", "body", m.OldValue); err != nil {
		return err
	}

	return nil
}

func (m *InsertableBillingEntity) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *InsertableBillingEntity) validatePricingComponentID(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentID", "body", m.PricingComponentID); err != nil {
		return err
	}

	return nil
}

var insertableBillingEntityTypeStatePropEnum []interface{}

// prop value enum
func (m *InsertableBillingEntity) validateStateEnum(path, location string, value string) error {
	if insertableBillingEntityTypeStatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["New","Accepted","Rejected","ChargeCreated"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			insertableBillingEntityTypeStatePropEnum = append(insertableBillingEntityTypeStatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, insertableBillingEntityTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InsertableBillingEntity) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *InsertableBillingEntity) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}
