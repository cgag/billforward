package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*PricingComponentTier PricingComponentTier

swagger:model PricingComponentTier
*/
type PricingComponentTier struct {

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created *strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	CrmID *string `json:"crmID,omitempty"`

	/* { "description" : "", "verbs":["GET"] }
	 */
	ID *string `json:"id,omitempty"`

	/* { "description" : "The lower threshold of the tier.", "verbs":["POST","PUT","GET"] }
	 */
	LowerThreshold *int32 `json:"lowerThreshold,omitempty"`

	/* { "description" : "Organization associated with the pricing-component-tier.", "verbs":[] }
	 */
	OrganizationID *string `json:"organizationID,omitempty"`

	/* { "description" : "Cost associated with tier. When the pricingType is fixed this is the total value. When pricingType is unit, this is the cost of each unit. ", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Price float64 `json:"price,omitempty"`

	/* { "description" : "ID of the pricing-component associated with the pricing-component-tier.", "verbs":["POST","PUT","GET"] }
	 */
	PricingComponentID *string `json:"pricingComponentID,omitempty"`

	/* { "description" : "Version ID of the associated pricing-component", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	PricingComponentVersionID string `json:"pricingComponentVersionID,omitempty"`

	/* { "description" : "Pricing calculation used to price items in this pricing tier. Unit pricing means every distinct value is used in the calculation. Fixed means that the total price of the tier is fixed regardless of the purchased amount.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	PricingType string `json:"pricingType,omitempty"`

	/* { "description" : "ID of the product-rate-plan associated with the pricing-component-tier.", "verbs":["POST","PUT","GET"] }
	 */
	ProductRatePlanID *string `json:"productRatePlanID,omitempty"`

	/* {  "default" : "&infin;",  "description" : "The upper threshold of the tier. If this is left null the tier will be infinite", "verbs":["POST","PUT","GET"] }
	 */
	UpperThreshold *int32 `json:"upperThreshold,omitempty"`
}

// Validate validates this pricing component tier
func (m *PricingComponentTier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentVersionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingComponentTier) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", float64(m.Price)); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponentTier) validatePricingComponentVersionID(formats strfmt.Registry) error {

	if err := validate.RequiredString("pricingComponentVersionID", "body", string(m.PricingComponentVersionID)); err != nil {
		return err
	}

	return nil
}

var pricingComponentTierPricingTypeEnum []interface{}

func (m *PricingComponentTier) validatePricingTypeEnum(path, location string, value string) error {
	if pricingComponentTierPricingTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["unit","fixed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentTierPricingTypeEnum = append(pricingComponentTierPricingTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, pricingComponentTierPricingTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *PricingComponentTier) validatePricingType(formats strfmt.Registry) error {

	if err := validate.RequiredString("pricingType", "body", string(m.PricingType)); err != nil {
		return err
	}

	if err := m.validatePricingTypeEnum("pricingType", "body", m.PricingType); err != nil {
		return err
	}

	return nil
}
