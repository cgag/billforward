package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*UpdateSubscriptionRequest Entity for requesting that a subscription be updated.

swagger:model UpdateSubscriptionRequest
*/
type UpdateSubscriptionRequest struct {

	/* {"description":"Whether this subscription should become an 'aggregating subscription', collecting charges (starting now) from all other subscriptions (current and future) belonging to this BillForward Account.","verbs":["POST"]}
	 */
	AggregateAllSubscriptionsOnAccount *bool `json:"aggregateAllSubscriptionsOnAccount,omitempty"`

	/* {"description":"New description to assign to the updated subscription. This is primarily for your benefit &mdash; for example, you could write here the mechanism through which you obtained this customer. (e.g. 'Customer obtained through Lazy Wednesdays promotion').","verbs":["POST"]}
	 */
	Description string `json:"description,omitempty"`

	/* {"description":"[Can only be changed if subscription is still in Provisioned state] ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should finish its first service period.","verbs":["POST"]}
	 */
	End *strfmt.DateTime `json:"end,omitempty"`

	/* {"default":"None","description":"The action that should be taken, should an invoice for some subscription to this rate plan remain unpaid despite the dunning period's being exceeded.<br><span class=\"label label-default\">CancelSubscription</span> &mdash; Demotes the subscription to the `Failed` state as soon as the dunning period is exceeded.<br><span class=\"label label-default\">None</span> &mdash; The subscription is allowed to continue in the `AwaitingPayment` state indefinitely even if the dunning period is exceeded.For slow payment cycles &mdash; or when manual invoice remediation is common &mdash; <span class=\"label label-default\">None</span> is recommended.<br>In a heavily-automated SaaS environment, automatic cancellation via <span class=\"label label-default\">CancelSubscription</span> is recommended.","verbs":["POST","PUT","GET"]}
	 */
	FailedPaymentBehaviour string `json:"failedPaymentBehaviour,omitempty"`

	/* {"description":"ID of the Subscription you wish to update.","verbs":["POST"]}

	Required: true
	*/
	ID *string `json:"id"`

	/* {"description":"New name to assign to the updated subscription. This is primarily for your benefit &mdash; for example, to enable you to identify subscriptions at a glance in the BillForward web interface (e.g. 'Customer 1425, guy@mail.com, Premium membership').","verbs":["POST"]}
	 */
	Name string `json:"name,omitempty"`

	/* {"description":"ID of a parent subscription which will collect the charges raised by this subscription. The parent becomes responsible for paying those charges.","verbs":["POST"]}
	 */
	ParentID string `json:"parentID,omitempty"`

	/* {"description":"[Can only be changed if subscription is still in Provisioned state] ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should enter its first service period.","verbs":["POST"]}
	 */
	Start *strfmt.DateTime `json:"start,omitempty"`

	/* {"description":"[Can only be changed if subscription is still in Provisioned state] The state into which you wish to move the updated subscription.<br><span class=\"label label-default\">AwaitingPayment</span> &mdash; The subscription is activated. After `start` time is surpassed, it will begin service and raise its first invoice.","verbs":["POST"]}
	 */
	State string `json:"state,omitempty"`
}

// Validate validates this update subscription request
func (m *UpdateSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailedPaymentBehaviour(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateSubscriptionRequestTypeFailedPaymentBehaviourPropEnum []interface{}

// prop value enum
func (m *UpdateSubscriptionRequest) validateFailedPaymentBehaviourEnum(path, location string, value string) error {
	if updateSubscriptionRequestTypeFailedPaymentBehaviourPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["CancelSubscription","None"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			updateSubscriptionRequestTypeFailedPaymentBehaviourPropEnum = append(updateSubscriptionRequestTypeFailedPaymentBehaviourPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, updateSubscriptionRequestTypeFailedPaymentBehaviourPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateSubscriptionRequest) validateFailedPaymentBehaviour(formats strfmt.Registry) error {

	if swag.IsZero(m.FailedPaymentBehaviour) { // not required
		return nil
	}

	// value enum
	if err := m.validateFailedPaymentBehaviourEnum("failedPaymentBehaviour", "body", m.FailedPaymentBehaviour); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSubscriptionRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var updateSubscriptionRequestTypeStatePropEnum []interface{}

// prop value enum
func (m *UpdateSubscriptionRequest) validateStateEnum(path, location string, value string) error {
	if updateSubscriptionRequestTypeStatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Trial","Provisioned","Paid","AwaitingPayment","Cancelled","Failed","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			updateSubscriptionRequestTypeStatePropEnum = append(updateSubscriptionRequestTypeStatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, updateSubscriptionRequestTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateSubscriptionRequest) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}
