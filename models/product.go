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

// Product Product
// swagger:model Product
type Product struct {

	// {"description":"","verbs":[]}
	AccountID string `json:"accountID,omitempty"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// {"description":"Customer-relationship-management ID of the product.","verbs":["GET","PUT","POST"]}
	CrmID string `json:"crmID,omitempty"`

	// {"description":"","verbs":["GET"]}
	// Required: true
	Deleted bool `json:"deleted"`

	// {"description":"A description &mdash; for your benefit &mdash; of the product. For example: you could explain what service this product entitles a customer to.","verbs":["POST","PUT","GET"]}
	// Required: true
	Description *string `json:"description"`

	// {"description":"Number of length-measures which constitute the product's period.","verbs":["POST","PUT","GET"]}
	// Required: true
	Duration *int32 `json:"duration"`

	// {"description":"Measure describing the magnitude of the product's period.","verbs":["POST","PUT","GET"]}
	// Required: true
	DurationPeriod *string `json:"durationPeriod"`

	// {"description":"ID uniquely identifying this product.","verbs":["GET"]}
	ID string `json:"id,omitempty"`

	// metadata
	Metadata DynamicMetadata `json:"metadata,omitempty"`

	// {"description":"A unique name &mdash; for your benefit &mdash; used to identify this product within BillForward. It should reflect the fact that this product confers some service to a customer (e.g. \"Gold membership\").<br>The product can also be defined by the frequency with which it recurs (e.g. \"Monthly Gold membership\").<br>Remember also that rate plans can override the timing prescribed by their product. If you intend to override that timing, you may consider the product's period duration to be an unimportant factor when it comes to naming it.","verbs":["POST","PUT","GET"]}
	// Required: true
	Name *string `json:"name"`

	// payment terms
	PaymentTerms int64 `json:"paymentTerms,omitempty"`

	// {"default":"recurring","description":"The frequency of the product &mdash; either recurring or non-recurring.","verbs":["POST","PUT","GET"]}
	// Required: true
	ProductType *string `json:"productType"`

	// {"description":"A friendly non-unique name used to identify this product","verbs":["POST","PUT","GET"]}
	PublicName string `json:"publicName,omitempty"`

	// {"description":"","verbs":[]}
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// {"default":0,"description":"Number of trial-length-measures which constitute the product's trial period","verbs":["POST","PUT","GET"]}
	// Required: true
	Trial *int32 `json:"trial"`

	// {"default":"none","description":"Measure describing the magnitude of the product's trial period.","verbs":["POST","PUT","GET"]}
	// Required: true
	TrialPeriod *string `json:"trialPeriod"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDurationPeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrial(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrialPeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

var productTypeDurationPeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["minutes","days","months","years"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productTypeDurationPeriodPropEnum = append(productTypeDurationPeriodPropEnum, v)
	}
}

const (
	// ProductDurationPeriodMinutes captures enum value "minutes"
	ProductDurationPeriodMinutes string = "minutes"
	// ProductDurationPeriodDays captures enum value "days"
	ProductDurationPeriodDays string = "days"
	// ProductDurationPeriodMonths captures enum value "months"
	ProductDurationPeriodMonths string = "months"
	// ProductDurationPeriodYears captures enum value "years"
	ProductDurationPeriodYears string = "years"
)

// prop value enum
func (m *Product) validateDurationPeriodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productTypeDurationPeriodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Product) validateDurationPeriod(formats strfmt.Registry) error {

	if err := validate.Required("durationPeriod", "body", m.DurationPeriod); err != nil {
		return err
	}

	// value enum
	if err := m.validateDurationPeriodEnum("durationPeriod", "body", *m.DurationPeriod); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var productTypeProductTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nonrecurring","recurring"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productTypeProductTypePropEnum = append(productTypeProductTypePropEnum, v)
	}
}

const (
	// ProductProductTypeNonrecurring captures enum value "nonrecurring"
	ProductProductTypeNonrecurring string = "nonrecurring"
	// ProductProductTypeRecurring captures enum value "recurring"
	ProductProductTypeRecurring string = "recurring"
)

// prop value enum
func (m *Product) validateProductTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productTypeProductTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Product) validateProductType(formats strfmt.Registry) error {

	if err := validate.Required("productType", "body", m.ProductType); err != nil {
		return err
	}

	// value enum
	if err := m.validateProductTypeEnum("productType", "body", *m.ProductType); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateTrial(formats strfmt.Registry) error {

	if err := validate.Required("trial", "body", m.Trial); err != nil {
		return err
	}

	return nil
}

var productTypeTrialPeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","minutes","days","months"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productTypeTrialPeriodPropEnum = append(productTypeTrialPeriodPropEnum, v)
	}
}

const (
	// ProductTrialPeriodNone captures enum value "none"
	ProductTrialPeriodNone string = "none"
	// ProductTrialPeriodMinutes captures enum value "minutes"
	ProductTrialPeriodMinutes string = "minutes"
	// ProductTrialPeriodDays captures enum value "days"
	ProductTrialPeriodDays string = "days"
	// ProductTrialPeriodMonths captures enum value "months"
	ProductTrialPeriodMonths string = "months"
)

// prop value enum
func (m *Product) validateTrialPeriodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productTypeTrialPeriodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Product) validateTrialPeriod(formats strfmt.Registry) error {

	if err := validate.Required("trialPeriod", "body", m.TrialPeriod); err != nil {
		return err
	}

	// value enum
	if err := m.validateTrialPeriodEnum("trialPeriod", "body", *m.TrialPeriod); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
