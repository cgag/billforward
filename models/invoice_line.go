package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*InvoiceLine An invoice-line represents the portion of an invoice specific to one particular pricing-component and its associated pricing-component-value.

swagger:model InvoiceLine
*/
type InvoiceLine struct {

	/* { "description" : "A human readable explanation of how the value of the invoice-line was calculated.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Calculation string `json:"calculation,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "charge-type.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ChargeType string `json:"chargeType,omitempty"`

	/* The ID of the invoice that is associated with the invoice-line.

	Required: true
	*/
	ChildInvoiceID string `json:"childInvoiceID,omitempty"`

	/* { "description" : "The component value for the unit-of-measure that is associated with the invoice-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ComponentValue int32 `json:"componentValue,omitempty"`

	/* { "description" : "The cost of the invoice-line including tax.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Cost float64 `json:"cost,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created *strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "The human readable description of the invoice-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Description string `json:"description,omitempty"`

	/* { "description" : "ID of the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	ID *string `json:"id,omitempty"`

	/* { "description" : "invoice associated with the invoice-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	InvoiceID string `json:"invoiceID,omitempty"`

	/* { "description" : "The human readable name of the invoice-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* { "description" : "ID of the organization associated with the invoice-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* The period end of the charge.

	Required: true
	*/
	PeriodEnd strfmt.DateTime `json:"periodEnd,omitempty"`

	/* The period start of the charge.

	Required: true
	*/
	PeriodStart strfmt.DateTime `json:"periodStart,omitempty"`

	/* The ID of the pricing-component that is associated with the invoice-line.

	Required: true
	*/
	PricingComponentID string `json:"pricingComponentID,omitempty"`

	/* The name of the pricing-component that is associated with the invoice-line.

	Required: true
	*/
	PricingComponentName string `json:"pricingComponentName,omitempty"`

	/* { "description" : "The type of the pricing component associated with the invoice line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	PricingComponentType string `json:"pricingComponentType,omitempty"`

	/* { "description" : "the product ID associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	ProductID *string `json:"productID,omitempty"`

	/* { "description" : "the product name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	ProductName *string `json:"productName,omitempty"`

	/* { "description" : "the product rate plan ID associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	ProductRatePlanID *string `json:"productRatePlanID,omitempty"`

	/* { "description" : "the product rate plan name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	ProductRatePlanName *string `json:"productRatePlanName,omitempty"`

	/* The public name of the pricing-component that is associated with the invoice-line.

	Required: true
	*/
	PublicPricingComponentName string `json:"publicPricingComponentName,omitempty"`

	/* { "description" : "the public product name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	PublicProductName *string `json:"publicProductName,omitempty"`

	/* { "description" : "the public product rate plan name associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	PublicProductRatePlanName *string `json:"publicProductRatePlanName,omitempty"`

	/* The ID of the subscription-charge that is associated with the invoice-line.

	Required: true
	*/
	SubscriptionChargeID string `json:"subscriptionChargeID,omitempty"`

	/* { "description" : "the subscription ID associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	SubscriptionID *string `json:"subscriptionID,omitempty"`

	/* { "description" : "The cumulative tax of the invoice-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Tax float64 `json:"tax,omitempty"`

	/* The type of the invoice-line.

	Required: true
	*/
	Type string `json:"type,omitempty"`

	/* { "description" : "The unit-of-measure associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

	/* { "description" : "unit-of-measure associated with the invoice-line.", "verbs":["POST","PUT","GET"] }
	 */
	UnitOfMeasureID *string `json:"unitOfMeasureID,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceLine) validateCalculation(formats strfmt.Registry) error {

	if err := validate.RequiredString("calculation", "body", string(m.Calculation)); err != nil {
		return err
	}

	return nil
}

var invoiceLineChargeTypeEnum []interface{}

func (m *InvoiceLine) validateChargeTypeEnum(path, location string, value string) error {
	if invoiceLineChargeTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Credit","Debit"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceLineChargeTypeEnum = append(invoiceLineChargeTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, invoiceLineChargeTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceLine) validateChargeType(formats strfmt.Registry) error {

	if err := validate.RequiredString("chargeType", "body", string(m.ChargeType)); err != nil {
		return err
	}

	if err := m.validateChargeTypeEnum("chargeType", "body", m.ChargeType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateChildInvoiceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("childInvoiceID", "body", string(m.ChildInvoiceID)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateComponentValue(formats strfmt.Registry) error {

	if err := validate.Required("componentValue", "body", int32(m.ComponentValue)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", float64(m.Cost)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateInvoiceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("invoiceID", "body", string(m.InvoiceID)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.RequiredString("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePeriodEnd(formats strfmt.Registry) error {

	if err := validate.Required("periodEnd", "body", strfmt.DateTime(m.PeriodEnd)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePeriodStart(formats strfmt.Registry) error {

	if err := validate.Required("periodStart", "body", strfmt.DateTime(m.PeriodStart)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePricingComponentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("pricingComponentID", "body", string(m.PricingComponentID)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePricingComponentName(formats strfmt.Registry) error {

	if err := validate.RequiredString("pricingComponentName", "body", string(m.PricingComponentName)); err != nil {
		return err
	}

	return nil
}

var invoiceLinePricingComponentTypeEnum []interface{}

func (m *InvoiceLine) validatePricingComponentTypeEnum(path, location string, value string) error {
	if invoiceLinePricingComponentTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["setup","subscription","arrears","usage"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceLinePricingComponentTypeEnum = append(invoiceLinePricingComponentTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, invoiceLinePricingComponentTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceLine) validatePricingComponentType(formats strfmt.Registry) error {

	if err := validate.RequiredString("pricingComponentType", "body", string(m.PricingComponentType)); err != nil {
		return err
	}

	if err := m.validatePricingComponentTypeEnum("pricingComponentType", "body", m.PricingComponentType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validatePublicPricingComponentName(formats strfmt.Registry) error {

	if err := validate.RequiredString("publicPricingComponentName", "body", string(m.PublicPricingComponentName)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateSubscriptionChargeID(formats strfmt.Registry) error {

	if err := validate.RequiredString("subscriptionChargeID", "body", string(m.SubscriptionChargeID)); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceLine) validateTax(formats strfmt.Registry) error {

	if err := validate.Required("tax", "body", float64(m.Tax)); err != nil {
		return err
	}

	return nil
}

var invoiceLineTypeEnum []interface{}

func (m *InvoiceLine) validateTypeEnum(path, location string, value string) error {
	if invoiceLineTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["PricingComponent","Coupon","Migration","AggregatedInvoice"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceLineTypeEnum = append(invoiceLineTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, invoiceLineTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceLine) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
