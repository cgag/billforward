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

/*CancelSubscriptionRequest CancelSubscriptionRequest

swagger:model CancelSubscriptionRequest
*/
type CancelSubscriptionRequest struct {

	/* CancelChildren cancel children
	 */
	CancelChildren *bool `json:"cancelChildren,omitempty"`

	/* CancelEmptyParent cancel empty parent
	 */
	CancelEmptyParent *bool `json:"cancelEmptyParent,omitempty"`

	/* Specifies whether the service will end immediately on cancellation or if it will continue until the end of the current period. Default: AtPeriodEnd
	 */
	CancellationCredit *string `json:"cancellationCredit,omitempty"`

	/* ServiceEnd service end
	 */
	ServiceEnd *string `json:"serviceEnd,omitempty"`

	/* Source source

	Required: true
	*/
	Source string `json:"source,omitempty"`

	/* State state
	 */
	State *string `json:"state,omitempty"`

	/* SubscriptionID subscription ID

	Required: true
	*/
	SubscriptionID string `json:"subscriptionID,omitempty"`
}

// Validate validates this cancel subscription request
func (m *CancelSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancellationCredit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
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

var cancelSubscriptionRequestCancellationCreditEnum []interface{}

func (m *CancelSubscriptionRequest) validateCancellationCreditEnum(path, location string, value string) error {
	if cancelSubscriptionRequestCancellationCreditEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Credit","None"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			cancelSubscriptionRequestCancellationCreditEnum = append(cancelSubscriptionRequestCancellationCreditEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, cancelSubscriptionRequestCancellationCreditEnum); err != nil {
		return err
	}
	return nil
}

func (m *CancelSubscriptionRequest) validateCancellationCredit(formats strfmt.Registry) error {

	if swag.IsZero(m.CancellationCredit) { // not required
		return nil
	}

	if err := m.validateCancellationCreditEnum("cancellationCredit", "body", *m.CancellationCredit); err != nil {
		return err
	}

	return nil
}

var cancelSubscriptionRequestServiceEndEnum []interface{}

func (m *CancelSubscriptionRequest) validateServiceEndEnum(path, location string, value string) error {
	if cancelSubscriptionRequestServiceEndEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Immediate","AtPeriodEnd"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			cancelSubscriptionRequestServiceEndEnum = append(cancelSubscriptionRequestServiceEndEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, cancelSubscriptionRequestServiceEndEnum); err != nil {
		return err
	}
	return nil
}

func (m *CancelSubscriptionRequest) validateServiceEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEnd) { // not required
		return nil
	}

	if err := m.validateServiceEndEnum("serviceEnd", "body", *m.ServiceEnd); err != nil {
		return err
	}

	return nil
}

func (m *CancelSubscriptionRequest) validateSource(formats strfmt.Registry) error {

	if err := validate.RequiredString("source", "body", string(m.Source)); err != nil {
		return err
	}

	return nil
}

func (m *CancelSubscriptionRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.RequiredString("subscriptionID", "body", string(m.SubscriptionID)); err != nil {
		return err
	}

	return nil
}
