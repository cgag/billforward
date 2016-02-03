package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*Account Account

swagger:model Account
*/
type Account struct {

	/* { "description" : "If present, this will be the product rate plan to use when creating an aggregating subscription.  An account level aggregating subscription will be created when the first subscription is created against the account.", "verbs":[] }
	 */
	AggregatingProductRatePlanID *string `json:"aggregatingProductRatePlanID,omitempty"`

	/* { "description" : "The consistent ID of the account level aggregating subscription, if one exists.", "verbs":[] }
	 */
	AggregatingSubscriptionID *string `json:"aggregatingSubscriptionID,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "CRM ID of the account.", "verbs":["GET"] }
	 */
	CrmID *string `json:"crmID,omitempty"`

	/* {  "default" : "false",  "description" : "Indicates if an account has been retired. If an account has been retired it can still be retrieved using the appropriate flag on API requests.", "verbs":["GET"] }
	 */
	Deleted *bool `json:"deleted,omitempty"`

	/* { "description" : "ID of the account.", "verbs":["GET"] }

	Required: true
	*/
	ID string `json:"id,omitempty"`

	/* { "description" : "Add metadata.", "verbs":["POST"] }
	 */
	Metadata DynamicMetadata `json:"metadata,omitempty"`

	/* { "description" : "Organization associated with the account.", "verbs":[] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "The payment-methods associated with the account.", "verbs":["GET"] }
	 */
	PaymentMethods []*PaymentMethod `json:"paymentMethods,omitempty"`

	/* Profile profile
	 */
	Profile *Profile `json:"profile,omitempty"`

	/* { "description" : "Number of distinct, paid subscriptions associated with this account. Initially the value will be 0 when no successful subscriptions have been taken. A subscription cancelled whilst in trial is counted as successful.", "verbs":["GET"] }
	 */
	SuccessfulSubscriptions int32 `json:"successfulSubscriptions,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "User associated with the account. If this is null, no user is currently assocaited with the account. A user is only set when an account is associated with a user account.", "verbs":[] }
	 */
	UserID *string `json:"userID,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethods(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.RequiredString("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *Account) validatePaymentMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentMethods); i++ {

		if m.PaymentMethods[i] != nil {

			if err := m.PaymentMethods[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
