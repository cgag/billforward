package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*AggregatingComponent The aggregating component defines a component which should be re-priced upon invoice aggregation

swagger:model AggregatingComponent
*/
type AggregatingComponent struct {

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* {"default":false,"description":"Whether the AggregatingComponent has been retired.","verbs":["GET"]}
	 */
	Deleted *bool `json:"deleted,omitempty"`

	/* {"description":"Unique ID by which the AggregatingComponent can be looked up.","verbs":["GET"]}
	 */
	ID string `json:"id,omitempty"`

	/* {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested pricing component should be created. If omitted: this will be auto-populated using your authentication credentials.","verbs":["POST","GET"]}
	 */
	OrganizationID string `json:"organizationID,omitempty"`

	/* {"description":"ID of the pricing component to which this AggregatingComponent's aggregation applies. The subscriber to the aggregating rate plan (which contains this AggregatingComponent), will consult its children at the end of each billing period, and collect from those children all charges whose pricing component matches this ID. Those charges' quantities will be counted, and used when calculating the price of consuming this AggregatingComponent. The aggregating subscription then raises a discount charge &mdash; to account for the more favourable price tiering that emerges when aggregating.","verbs":["POST"]}

	Required: true
	*/
	PricingComponentID *string `json:"pricingComponentID"`

	/* {"description":"Name of the pricing component to which this AggregatingComponent's aggregation applies. The subscriber to the aggregating rate plan (which contains this AggregatingComponent), will consult its children at the end of each billing period, and collect from those children all charges whose pricing component matches this ID. Those charges' quantities will be counted, and used when calculating the price of consuming this AggregatingComponent. The aggregating subscription then raises a discount charge &mdash; to account for the more favourable price tiering that emerges when aggregating.","verbs":["POST"]}

	Required: true
	*/
	PricingComponentName *string `json:"pricingComponentName"`

	/* {"description":"ID of the rate plan upon which this AggregatingComponent is defined.","verbs":["POST","GET"]}
	 */
	ProductRatePlanID string `json:"productRatePlanID,omitempty"`

	/* {"description":"Name of the rate plan upon which this AggregatingComponent is defined.","verbs":["POST","GET"]}
	 */
	ProductRatePlanName string `json:"productRatePlanName,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this aggregating component
func (m *AggregatingComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePricingComponentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregatingComponent) validatePricingComponentID(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentID", "body", m.PricingComponentID); err != nil {
		return err
	}

	return nil
}

func (m *AggregatingComponent) validatePricingComponentName(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponentName", "body", m.PricingComponentName); err != nil {
		return err
	}

	return nil
}
