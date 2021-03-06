// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PricingComponent pricing component
// swagger:discriminator PricingComponent @type
type PricingComponent interface {
	runtime.Validatable

	// { "description" : "", "default" : "", "verbs":["POST","GET"] }
	// Required: true
	AtType() string
	SetAtType(string)

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy() string
	SetChangedBy(string)

	// { "description" : "The charge model of the pricing-component.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ChargeModel() *string
	SetChargeModel(*string)

	// { "description" : "The charge type of the pricing-component.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ChargeType() *string
	SetChargeType(*string)

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created() strfmt.DateTime
	SetCreated(strfmt.DateTime)

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	CrmID() string
	SetCrmID(string)

	// { "description" : "The default quantity of the pricing-component. If no value is supplied on a subscription this value will be used. This is useful for setting an expected purchase level of for introducing new pricing components to existing subscriptions and not having to back-fill values", "verbs":["POST","PUT","GET"] }
	// Required: true
	DefaultQuantity() *int32
	SetDefaultQuantity(*int32)

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	Description() string
	SetDescription(string)

	// {"default":"<span class=\"label label-default\">delayed</span>","description":"Default behaviour when a value is downgraded using this pricing component, this behaviour can be overridden when changing the value.<br><span class=\"label label-default\">immediate</span> &mdash; Downgrade will apply at the time the request is made.<br><span class=\"label label-default\">delayed</span> &mdash; Downgrade will apply at the end of the current billing cycle.","verbs":["POST","GET"]}
	DowngradeMode() string
	SetDowngradeMode(string)

	// { "description" : "", "verbs":["GET"] } When associating a pricing component with a product rate plan, this ID should be used.
	// Required: true
	ID() *string
	SetID(*string)

	// { "default" : "Aggregated",  "description" : "For <span class=\"label label-default\">setup</span> pricing components <span class=\"label label-default\">Immediate</span> invoicing will result in an invoice being issued on subscription being set to the AwaitingPayment state, irrespective of the subscription start date. <span class=\"label label-default\">Aggregated</span> invoicing will add a charge to the first invoice of the subscription.", "verbs":["POST","PUT","GET"] }
	InvoicingType() string
	SetInvoicingType(string)

	// { "description" : "The maximum quantity of the pricing-component.", "verbs":[] }
	MaxQuantity() int32
	SetMaxQuantity(int32)

	// { "default" : "0", "description" : "The minimum quantity of the pricing-component.", "verbs":[] }
	MinQuantity() int32
	SetMinQuantity(int32)

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	// Required: true
	Name() *string
	SetName(*string)

	// { "description" : "", "verbs":[] }
	// Required: true
	OrganizationID() *string
	SetOrganizationID(*string)

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	// Required: true
	ProductRatePlanID() *string
	SetProductRatePlanID(*string)

	// {"description":"A friendly non-unique name used to identify this pricing-component","verbs":["POST","PUT","GET"]}
	PublicName() string
	SetPublicName(string)

	// {  "default" : "[]", "description" : "The pricing-component-tiers associated with the pricing-component.", "verbs":["POST","PUT","GET"] }
	Tiers() []*PricingComponentTier
	SetTiers([]*PricingComponentTier)

	// { "description" : "The unit-of-measure associated with the pricing-component.", "verbs":["POST","PUT","GET"] }
	UnitOfMeasure() *UnitOfMeasure
	SetUnitOfMeasure(*UnitOfMeasure)

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	// Required: true
	UnitOfMeasureID() *string
	SetUnitOfMeasureID(*string)

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated() strfmt.DateTime
	SetUpdated(strfmt.DateTime)

	// {"default":"<span class=\"label label-default\">immediate</span>","description":"Default behaviour when a value is upgraded using this pricing component, this behaviour can be overridden when changing the value.<br><span class=\"label label-default\">immediate</span> &mdash; Upgrade will apply at the time the request is made.<br><span class=\"label label-default\">delayed</span> &mdash; Upgrade will apply at the end of the current billing cycle.","verbs":["POST","GET"]}
	UpgradeMode() string
	SetUpgradeMode(string)

	// { "default" : "current time", "description" : "The DateTime specifying when the pricing-component is valid from.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ValidFrom() *strfmt.DateTime
	SetValidFrom(*strfmt.DateTime)

	// {  "default" : "null", "description" : "The UTC DateTime specifying when the pricing-component is valid till.", "verbs":["POST","PUT","GET"] }
	ValidTill() *strfmt.DateTime
	SetValidTill(*strfmt.DateTime)

	// { "description" : "", "verbs":["GET"] }
	// Required: true
	VersionID() *string
	SetVersionID(*string)
}

type pricingComponent struct {
	atTypeField string

	changedByField string

	chargeModelField *string

	chargeTypeField *string

	createdField strfmt.DateTime

	crmIdField string

	defaultQuantityField *int32

	descriptionField string

	downgradeModeField string

	idField *string

	invoicingTypeField string

	maxQuantityField int32

	minQuantityField int32

	nameField *string

	organizationIdField *string

	productRatePlanIdField *string

	publicNameField string

	tiersField []*PricingComponentTier

	unitOfMeasureField *UnitOfMeasure

	unitOfMeasureIdField *string

	updatedField strfmt.DateTime

	upgradeModeField string

	validFromField *strfmt.DateTime

	validTillField *strfmt.DateTime

	versionIdField *string
}

func (m *pricingComponent) AtType() string {
	return "PricingComponent"
}
func (m *pricingComponent) SetAtType(val string) {

}

func (m *pricingComponent) ChangedBy() string {
	return m.changedByField
}
func (m *pricingComponent) SetChangedBy(val string) {
	m.changedByField = val
}

func (m *pricingComponent) ChargeModel() *string {
	return m.chargeModelField
}
func (m *pricingComponent) SetChargeModel(val *string) {
	m.chargeModelField = val
}

func (m *pricingComponent) ChargeType() *string {
	return m.chargeTypeField
}
func (m *pricingComponent) SetChargeType(val *string) {
	m.chargeTypeField = val
}

func (m *pricingComponent) Created() strfmt.DateTime {
	return m.createdField
}
func (m *pricingComponent) SetCreated(val strfmt.DateTime) {
	m.createdField = val
}

func (m *pricingComponent) CrmID() string {
	return m.crmIdField
}
func (m *pricingComponent) SetCrmID(val string) {
	m.crmIdField = val
}

func (m *pricingComponent) DefaultQuantity() *int32 {
	return m.defaultQuantityField
}
func (m *pricingComponent) SetDefaultQuantity(val *int32) {
	m.defaultQuantityField = val
}

func (m *pricingComponent) Description() string {
	return m.descriptionField
}
func (m *pricingComponent) SetDescription(val string) {
	m.descriptionField = val
}

func (m *pricingComponent) DowngradeMode() string {
	return m.downgradeModeField
}
func (m *pricingComponent) SetDowngradeMode(val string) {
	m.downgradeModeField = val
}

func (m *pricingComponent) ID() *string {
	return m.idField
}
func (m *pricingComponent) SetID(val *string) {
	m.idField = val
}

func (m *pricingComponent) InvoicingType() string {
	return m.invoicingTypeField
}
func (m *pricingComponent) SetInvoicingType(val string) {
	m.invoicingTypeField = val
}

func (m *pricingComponent) MaxQuantity() int32 {
	return m.maxQuantityField
}
func (m *pricingComponent) SetMaxQuantity(val int32) {
	m.maxQuantityField = val
}

func (m *pricingComponent) MinQuantity() int32 {
	return m.minQuantityField
}
func (m *pricingComponent) SetMinQuantity(val int32) {
	m.minQuantityField = val
}

func (m *pricingComponent) Name() *string {
	return m.nameField
}
func (m *pricingComponent) SetName(val *string) {
	m.nameField = val
}

func (m *pricingComponent) OrganizationID() *string {
	return m.organizationIdField
}
func (m *pricingComponent) SetOrganizationID(val *string) {
	m.organizationIdField = val
}

func (m *pricingComponent) ProductRatePlanID() *string {
	return m.productRatePlanIdField
}
func (m *pricingComponent) SetProductRatePlanID(val *string) {
	m.productRatePlanIdField = val
}

func (m *pricingComponent) PublicName() string {
	return m.publicNameField
}
func (m *pricingComponent) SetPublicName(val string) {
	m.publicNameField = val
}

func (m *pricingComponent) Tiers() []*PricingComponentTier {
	return m.tiersField
}
func (m *pricingComponent) SetTiers(val []*PricingComponentTier) {
	m.tiersField = val
}

func (m *pricingComponent) UnitOfMeasure() *UnitOfMeasure {
	return m.unitOfMeasureField
}
func (m *pricingComponent) SetUnitOfMeasure(val *UnitOfMeasure) {
	m.unitOfMeasureField = val
}

func (m *pricingComponent) UnitOfMeasureID() *string {
	return m.unitOfMeasureIdField
}
func (m *pricingComponent) SetUnitOfMeasureID(val *string) {
	m.unitOfMeasureIdField = val
}

func (m *pricingComponent) Updated() strfmt.DateTime {
	return m.updatedField
}
func (m *pricingComponent) SetUpdated(val strfmt.DateTime) {
	m.updatedField = val
}

func (m *pricingComponent) UpgradeMode() string {
	return m.upgradeModeField
}
func (m *pricingComponent) SetUpgradeMode(val string) {
	m.upgradeModeField = val
}

func (m *pricingComponent) ValidFrom() *strfmt.DateTime {
	return m.validFromField
}
func (m *pricingComponent) SetValidFrom(val *strfmt.DateTime) {
	m.validFromField = val
}

func (m *pricingComponent) ValidTill() *strfmt.DateTime {
	return m.validTillField
}
func (m *pricingComponent) SetValidTill(val *strfmt.DateTime) {
	m.validTillField = val
}

func (m *pricingComponent) VersionID() *string {
	return m.versionIdField
}
func (m *pricingComponent) SetVersionID(val *string) {
	m.versionIdField = val
}

// Validate validates this pricing component
func (m *pricingComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChargeType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDefaultQuantity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDowngradeMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoicingType(formats); err != nil {
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

	if err := m.validateProductRatePlanID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTiers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasure(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasureID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpgradeMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pricingComponentTypeAtTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tieredPricingComponent","flatPricingComponent","tieredVolumePricingComponent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pricingComponentTypeAtTypePropEnum = append(pricingComponentTypeAtTypePropEnum, v)
	}
}

const (
	// PricingComponentAtTypeTieredPricingComponent captures enum value "tieredPricingComponent"
	PricingComponentAtTypeTieredPricingComponent string = "tieredPricingComponent"
	// PricingComponentAtTypeFlatPricingComponent captures enum value "flatPricingComponent"
	PricingComponentAtTypeFlatPricingComponent string = "flatPricingComponent"
	// PricingComponentAtTypeTieredVolumePricingComponent captures enum value "tieredVolumePricingComponent"
	PricingComponentAtTypeTieredVolumePricingComponent string = "tieredVolumePricingComponent"
)

// prop value enum
func (m *pricingComponent) validateAtTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pricingComponentTypeAtTypePropEnum); err != nil {
		return err
	}
	return nil
}

var pricingComponentTypeChargeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["flat","tiered","tiered_volume"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pricingComponentTypeChargeModelPropEnum = append(pricingComponentTypeChargeModelPropEnum, v)
	}
}

const (
	// PricingComponentChargeModelFlat captures enum value "flat"
	PricingComponentChargeModelFlat string = "flat"
	// PricingComponentChargeModelTiered captures enum value "tiered"
	PricingComponentChargeModelTiered string = "tiered"
	// PricingComponentChargeModelTieredVolume captures enum value "tiered_volume"
	PricingComponentChargeModelTieredVolume string = "tiered_volume"
)

// prop value enum
func (m *pricingComponent) validateChargeModelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pricingComponentTypeChargeModelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *pricingComponent) validateChargeModel(formats strfmt.Registry) error {

	if err := validate.Required("chargeModel", "body", m.ChargeModel()); err != nil {
		return err
	}

	// value enum
	if err := m.validateChargeModelEnum("chargeModel", "body", *m.ChargeModel()); err != nil {
		return err
	}

	return nil
}

var pricingComponentTypeChargeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["setup","subscription","arrears","usage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pricingComponentTypeChargeTypePropEnum = append(pricingComponentTypeChargeTypePropEnum, v)
	}
}

const (
	// PricingComponentChargeTypeSetup captures enum value "setup"
	PricingComponentChargeTypeSetup string = "setup"
	// PricingComponentChargeTypeSubscription captures enum value "subscription"
	PricingComponentChargeTypeSubscription string = "subscription"
	// PricingComponentChargeTypeArrears captures enum value "arrears"
	PricingComponentChargeTypeArrears string = "arrears"
	// PricingComponentChargeTypeUsage captures enum value "usage"
	PricingComponentChargeTypeUsage string = "usage"
)

// prop value enum
func (m *pricingComponent) validateChargeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pricingComponentTypeChargeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *pricingComponent) validateChargeType(formats strfmt.Registry) error {

	if err := validate.Required("chargeType", "body", m.ChargeType()); err != nil {
		return err
	}

	// value enum
	if err := m.validateChargeTypeEnum("chargeType", "body", *m.ChargeType()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateDefaultQuantity(formats strfmt.Registry) error {

	if err := validate.Required("defaultQuantity", "body", m.DefaultQuantity()); err != nil {
		return err
	}

	return nil
}

var pricingComponentTypeDowngradeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pricingComponentTypeDowngradeModePropEnum = append(pricingComponentTypeDowngradeModePropEnum, v)
	}
}

const (
	// PricingComponentDowngradeModeImmediate captures enum value "immediate"
	PricingComponentDowngradeModeImmediate string = "immediate"
	// PricingComponentDowngradeModeDelayed captures enum value "delayed"
	PricingComponentDowngradeModeDelayed string = "delayed"
)

// prop value enum
func (m *pricingComponent) validateDowngradeModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pricingComponentTypeDowngradeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *pricingComponent) validateDowngradeMode(formats strfmt.Registry) error {

	if swag.IsZero(m.DowngradeMode()) { // not required
		return nil
	}

	// value enum
	if err := m.validateDowngradeModeEnum("downgradeMode", "body", m.DowngradeMode()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

var pricingComponentTypeInvoicingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Immediate","Aggregated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pricingComponentTypeInvoicingTypePropEnum = append(pricingComponentTypeInvoicingTypePropEnum, v)
	}
}

const (
	// PricingComponentInvoicingTypeImmediate captures enum value "Immediate"
	PricingComponentInvoicingTypeImmediate string = "Immediate"
	// PricingComponentInvoicingTypeAggregated captures enum value "Aggregated"
	PricingComponentInvoicingTypeAggregated string = "Aggregated"
)

// prop value enum
func (m *pricingComponent) validateInvoicingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pricingComponentTypeInvoicingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *pricingComponent) validateInvoicingType(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoicingType()) { // not required
		return nil
	}

	// value enum
	if err := m.validateInvoicingTypeEnum("invoicingType", "body", m.InvoicingType()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateProductRatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("productRatePlanID", "body", m.ProductRatePlanID()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateTiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Tiers()) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiers()); i++ {

		if swag.IsZero(m.tiersField[i]) { // not required
			continue
		}

		if m.tiersField[i] != nil {

			if err := m.tiersField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *pricingComponent) validateUnitOfMeasure(formats strfmt.Registry) error {

	if swag.IsZero(m.UnitOfMeasure()) { // not required
		return nil
	}

	if m.UnitOfMeasure() != nil {

		if err := m.UnitOfMeasure().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unitOfMeasure")
			}
			return err
		}
	}

	return nil
}

func (m *pricingComponent) validateUnitOfMeasureID(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasureID", "body", m.UnitOfMeasureID()); err != nil {
		return err
	}

	return nil
}

var pricingComponentTypeUpgradeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pricingComponentTypeUpgradeModePropEnum = append(pricingComponentTypeUpgradeModePropEnum, v)
	}
}

const (
	// PricingComponentUpgradeModeImmediate captures enum value "immediate"
	PricingComponentUpgradeModeImmediate string = "immediate"
	// PricingComponentUpgradeModeDelayed captures enum value "delayed"
	PricingComponentUpgradeModeDelayed string = "delayed"
)

// prop value enum
func (m *pricingComponent) validateUpgradeModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pricingComponentTypeUpgradeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *pricingComponent) validateUpgradeMode(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradeMode()) { // not required
		return nil
	}

	// value enum
	if err := m.validateUpgradeModeEnum("upgradeMode", "body", m.UpgradeMode()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom()); err != nil {
		return err
	}

	return nil
}

func (m *pricingComponent) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("versionID", "body", m.VersionID()); err != nil {
		return err
	}

	return nil
}

// UnmarshalPricingComponentSlice unmarshals polymorphic slices of PricingComponent
func UnmarshalPricingComponentSlice(reader io.Reader, consumer runtime.Consumer) ([]PricingComponent, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PricingComponent
	for _, element := range elements {
		obj, err := unmarshalPricingComponent(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPricingComponent unmarshals polymorphic PricingComponent
func UnmarshalPricingComponent(reader io.Reader, consumer runtime.Consumer) (PricingComponent, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPricingComponent(data, consumer)
}

func unmarshalPricingComponent(data []byte, consumer runtime.Consumer) (PricingComponent, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		AtType string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.AtType); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.AtType {
	case "PricingComponent":
		var result pricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "flatPricingComponent":
		var result FlatPricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "tieredPricingComponent":
		var result TieredPricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "tieredVolumePricingComponent":
		var result TieredVolumePricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}
