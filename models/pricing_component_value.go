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

// PricingComponentValue PricingComponentValue
// swagger:model PricingComponentValue

type PricingComponentValue struct {

	// { "description" : "<p>The appliesFrom can be left null. If appliesFrom is set, it indicates when a value came into effect.</p>", "verbs":["POST","PUT","GET"] }
	AppliesFrom strfmt.DateTime `json:"appliesFrom,omitempty"`

	// { "description" : "<p>For <span class=\"label label-default\">setup</span>, <span class=\"label label-default\">subscription</span>, and <span class=\"label label-default\">arrears</span> pricing components if appliesTill is specificed the value will be used whilst the time has not been reached. If appliesTill is null the pricing component value will be used until a new value is added. When a new value is added appliesTill will be set to the time the new value will take effect.</p><p><span class=\"label label-default\">usage</span> pricing applies to the previous billing period as it is charged in-arrears. When adding usage a new pricing component value should be added with appliesTill set to the end of the usages billing period. For example a monthly subscription results in an invoice being generated on the 1<sup>st</sup> of March, the previous months usage period ended on the same date. A usage value should be added to the subscription with the appliesTill set to the invoices periodStart, the 1<sup>st</sup> of March.</p>", "verbs":["POST","PUT","GET"] }
	AppliesTill strfmt.DateTime `json:"appliesTill,omitempty"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "", "verbs":["GET"] }
	// Required: true
	ID *string `json:"id"`

	// { "description" : "", "verbs":["GET"] }
	OrganizationID string `json:"organizationID,omitempty"`

	// pending change
	PendingChange *PendingComponentValueChange `json:"pendingChange,omitempty"`

	// { "description" : "", "verbs":["POST","GET"] }
	// Required: true
	PricingComponentID *string `json:"pricingComponentID"`

	// { "description" : "Name of the pricing-component associated with the pricing-component-value.", "verbs":["GET"] }
	PricingComponentName string `json:"pricingComponentName,omitempty"`

	// { "description" : "Value can be left null if setting the pricing component value on the subscription directly.", "verbs":["GET", "POST"] }
	// Required: true
	SubscriptionID *string `json:"subscriptionID"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// { "description" : "Quantity of a particular pricing component the subscription should have. For example if you have a pricing component for widgets, where $5/widget/month and you set the value to 10 then the customer will be charged $50 ($5 x 10) monthly.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Value *int32 `json:"value"`

	// { "description" : "", "verbs":["GET"] }
	VersionID string `json:"versionID,omitempty"`
}

/* polymorph PricingComponentValue appliesFrom false */

/* polymorph PricingComponentValue appliesTill false */

/* polymorph PricingComponentValue changedBy false */

/* polymorph PricingComponentValue created false */

/* polymorph PricingComponentValue id false */

/* polymorph PricingComponentValue organizationID false */

/* polymorph PricingComponentValue pendingChange false */

/* polymorph PricingComponentValue pricingComponentID false */

/* polymorph PricingComponentValue pricingComponentName false */

/* polymorph PricingComponentValue subscriptionID false */

/* polymorph PricingComponentValue updated false */

/* polymorph PricingComponentValue value false */

/* polymorph PricingComponentValue versionID false */

// Validate validates this pricing component value
func (m *PricingComponentValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePendingChange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingComponentValue) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponentValue) validatePendingChange(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingChange) { // not required
		return nil
	}

	if m.PendingChange != nil {

		if err := m.PendingChange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pendingChange")
			}
			return err
		}
	}

	return nil
}

func (m *PricingComponentValue) validatePricingComponentID(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentID", "body", m.PricingComponentID); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponentValue) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponentValue) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PricingComponentValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PricingComponentValue) UnmarshalBinary(b []byte) error {
	var res PricingComponentValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
