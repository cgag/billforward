// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceLine An invoice-line represents the portion of an invoice specific to one particular pricing-component and its associated pricing-component-value.
// swagger:model InvoiceLine

type InvoiceLine struct {

	// { "description" : "A human readable explanation of how the value of the invoice-line was calculated.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Calculation *string `json:"calculation"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "charge-type.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ChargeType *string `json:"chargeType"`

	// The ID of the invoice that is associated with the invoice-line.
	// Required: true
	ChildInvoiceID *string `json:"childInvoiceID"`

	// { "description" : "The component value for the unit-of-measure that is associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ComponentValue *int32 `json:"componentValue"`

	// { "description" : "The cost of the invoice-line including tax.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Cost *float64 `json:"cost"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "The human readable description of the invoice-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Description *string `json:"description"`

	// { "description" : "ID of the invoice-line.", "verbs":["POST","PUT","GET"] }
	ID string `json:"id,omitempty"`

	// { "description" : "invoice associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	InvoiceID *string `json:"invoiceID"`

	// { "description" : "The human readable name of the invoice-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Name *string `json:"name"`

	// { "description" : "ID of the organization associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	OrganizationID *string `json:"organizationID"`

	// The period end of the charge.
	// Required: true
	PeriodEnd *strfmt.DateTime `json:"periodEnd"`

	// The period start of the charge.
	// Required: true
	PeriodStart *strfmt.DateTime `json:"periodStart"`

	// The ID of the pricing-component that is associated with the invoice-line.
	// Required: true
	PricingComponentID *string `json:"pricingComponentID"`

	// The name of the pricing-component that is associated with the invoice-line.
	// Required: true
	PricingComponentName *string `json:"pricingComponentName"`

	// { "description" : "The type of the pricing component associated with the invoice line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	PricingComponentType *string `json:"pricingComponentType"`

	// { "description" : "the product ID associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	ProductID string `json:"productID,omitempty"`

	// { "description" : "the product name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	ProductName string `json:"productName,omitempty"`

	// { "description" : "the product rate plan ID associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	ProductRatePlanID string `json:"productRatePlanID,omitempty"`

	// { "description" : "the product rate plan name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	ProductRatePlanName string `json:"productRatePlanName,omitempty"`

	// The public name of the pricing-component that is associated with the invoice-line.
	// Required: true
	PublicPricingComponentName *string `json:"publicPricingComponentName"`

	// { "description" : "the public product name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	PublicProductName string `json:"publicProductName,omitempty"`

	// { "description" : "the public product rate plan name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	PublicProductRatePlanName string `json:"publicProductRatePlanName,omitempty"`

	// The ID of the subscription-charge that is associated with the invoice-line.
	// Required: true
	SubscriptionChargeID *string `json:"subscriptionChargeID"`

	// { "description" : "the subscription ID associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// { "description" : "The cumulative tax of the invoice-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Tax *float64 `json:"tax"`

	// The type of the invoice-line.
	// Required: true
	Type *string `json:"type"`

	// { "description" : "The unit-of-measure associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

	// { "description" : "unit-of-measure associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	UnitOfMeasureID string `json:"unitOfMeasureID,omitempty"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

/* polymorph InvoiceLine calculation false */

/* polymorph InvoiceLine changedBy false */

/* polymorph InvoiceLine chargeType false */

/* polymorph InvoiceLine childInvoiceID false */

/* polymorph InvoiceLine componentValue false */

/* polymorph InvoiceLine cost false */

/* polymorph InvoiceLine created false */

/* polymorph InvoiceLine description false */

/* polymorph InvoiceLine id false */

/* polymorph InvoiceLine invoiceID false */

/* polymorph InvoiceLine name false */

/* polymorph InvoiceLine organizationID false */

/* polymorph InvoiceLine periodEnd false */

/* polymorph InvoiceLine periodStart false */

/* polymorph InvoiceLine pricingComponentID false */

/* polymorph InvoiceLine pricingComponentName false */

/* polymorph InvoiceLine pricingComponentType false */

/* polymorph InvoiceLine productID false */

/* polymorph InvoiceLine productName false */

/* polymorph InvoiceLine productRatePlanID false */

/* polymorph InvoiceLine productRatePlanName false */

/* polymorph InvoiceLine publicPricingComponentName false */

/* polymorph InvoiceLine publicProductName false */

/* polymorph InvoiceLine publicProductRatePlanName false */

/* polymorph InvoiceLine subscriptionChargeID false */

/* polymorph InvoiceLine subscriptionID false */

/* polymorph InvoiceLine tax false */

/* polymorph InvoiceLine type false */

/* polymorph InvoiceLine unitOfMeasure false */

/* polymorph InvoiceLine unitOfMeasureID false */

/* polymorph InvoiceLine updated false */

// Validate validates this invoice line
func (m *InvoiceLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChargeType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChildInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateComponentValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeriodEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeriodStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublicPricingComponentName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionChargeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTax(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasure(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceLine) validateCalculation(formats strfmt.Registry) error {

	if err := validate.Required("calculation", "body", m.Calculation); err != nil {
		return err
	}

	return nil
}

var invoiceLineTypeChargeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Credit","Debit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceLineTypeChargeTypePropEnum = append(invoiceLineTypeChargeTypePropEnum, v)
	}
}

const (
	// InvoiceLineChargeTypeCredit captures enum value "Credit"
	InvoiceLineChargeTypeCredit string = "Credit"
	// InvoiceLineChargeTypeDebit captures enum value "Debit"
	InvoiceLineChargeTypeDebit string = "Debit"
)

// prop value enum
func (m *InvoiceLine) validateChargeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoiceLineTypeChargeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceLine) validateChargeType(formats strfmt.Registry) error {

	if err := validate.Required("chargeType", "body", m.ChargeType); err != nil {
		return err
	}

	// value enum
	if err := m.validateChargeTypeEnum("chargeType", "body", *m.ChargeType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateChildInvoiceID(formats strfmt.Registry) error {

	if err := validate.Required("childInvoiceID", "body", m.ChildInvoiceID); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateComponentValue(formats strfmt.Registry) error {

	if err := validate.Required("componentValue", "body", m.ComponentValue); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateInvoiceID(formats strfmt.Registry) error {

	if err := validate.Required("invoiceID", "body", m.InvoiceID); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePeriodEnd(formats strfmt.Registry) error {

	if err := validate.Required("periodEnd", "body", m.PeriodEnd); err != nil {
		return err
	}

	if err := validate.FormatOf("periodEnd", "body", "date-time", m.PeriodEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePeriodStart(formats strfmt.Registry) error {

	if err := validate.Required("periodStart", "body", m.PeriodStart); err != nil {
		return err
	}

	if err := validate.FormatOf("periodStart", "body", "date-time", m.PeriodStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePricingComponentID(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentID", "body", m.PricingComponentID); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePricingComponentName(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentName", "body", m.PricingComponentName); err != nil {
		return err
	}

	return nil
}

var invoiceLineTypePricingComponentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["setup","subscription","arrears","usage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceLineTypePricingComponentTypePropEnum = append(invoiceLineTypePricingComponentTypePropEnum, v)
	}
}

const (
	// InvoiceLinePricingComponentTypeSetup captures enum value "setup"
	InvoiceLinePricingComponentTypeSetup string = "setup"
	// InvoiceLinePricingComponentTypeSubscription captures enum value "subscription"
	InvoiceLinePricingComponentTypeSubscription string = "subscription"
	// InvoiceLinePricingComponentTypeArrears captures enum value "arrears"
	InvoiceLinePricingComponentTypeArrears string = "arrears"
	// InvoiceLinePricingComponentTypeUsage captures enum value "usage"
	InvoiceLinePricingComponentTypeUsage string = "usage"
)

// prop value enum
func (m *InvoiceLine) validatePricingComponentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoiceLineTypePricingComponentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceLine) validatePricingComponentType(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentType", "body", m.PricingComponentType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePricingComponentTypeEnum("pricingComponentType", "body", *m.PricingComponentType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePublicPricingComponentName(formats strfmt.Registry) error {

	if err := validate.Required("publicPricingComponentName", "body", m.PublicPricingComponentName); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateSubscriptionChargeID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionChargeID", "body", m.SubscriptionChargeID); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateTax(formats strfmt.Registry) error {

	if err := validate.Required("tax", "body", m.Tax); err != nil {
		return err
	}

	return nil
}

var invoiceLineTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PricingComponent","Coupon","Migration","AggregatedInvoice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceLineTypeTypePropEnum = append(invoiceLineTypeTypePropEnum, v)
	}
}

const (
	// InvoiceLineTypePricingComponent captures enum value "PricingComponent"
	InvoiceLineTypePricingComponent string = "PricingComponent"
	// InvoiceLineTypeCoupon captures enum value "Coupon"
	InvoiceLineTypeCoupon string = "Coupon"
	// InvoiceLineTypeMigration captures enum value "Migration"
	InvoiceLineTypeMigration string = "Migration"
	// InvoiceLineTypeAggregatedInvoice captures enum value "AggregatedInvoice"
	InvoiceLineTypeAggregatedInvoice string = "AggregatedInvoice"
)

// prop value enum
func (m *InvoiceLine) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoiceLineTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceLine) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateUnitOfMeasure(formats strfmt.Registry) error {

	if swag.IsZero(m.UnitOfMeasure) { // not required
		return nil
	}

	if m.UnitOfMeasure != nil {

		if err := m.UnitOfMeasure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unitOfMeasure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceLine) UnmarshalBinary(b []byte) error {
	var res InvoiceLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
