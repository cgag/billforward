package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ProductRatePlanPagedMetadata ProductRatePlanPagedMetadata product rate plan paged metadata

swagger:model ProductRatePlanPagedMetadata
*/
type ProductRatePlanPagedMetadata struct {

	/* {"description":"Paging parameter. 0-indexed. Describes your current location within a pageable list of query results.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	CurrentOffset int32 `json:"currentOffset,omitempty"`

	/* {"description":"Paging parameter. 0-indexed. Describes which page (given a page size of `recordsRequested`) of the result set you are viewing.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	CurrentPage int32 `json:"currentPage,omitempty"`

	/* {"description":"Number of milliseconds taken by API to calculate response.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	ExecutionTime int64 `json:"executionTime,omitempty"`

	/* {"description":"Paging parameter. URL fragment that can be used to fetch next page of results.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	NextPage string `json:"nextPage,omitempty"`

	/* {"default":10,"description":"Paging parameter. Describes how many records you requested.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	RecordsRequested int32 `json:"recordsRequested,omitempty"`

	/* {"description":"Describes how many records were returned by your query.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	RecordsReturned int32 `json:"recordsReturned,omitempty"`

	/* {"description":"The results returned by your query.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	Results []*ProductRatePlan `json:"results,omitempty"`
}

// Validate validates this product rate plan paged metadata
func (m *ProductRatePlanPagedMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentOffset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentPage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExecutionTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNextPage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordsRequested(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordsReturned(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductRatePlanPagedMetadata) validateCurrentOffset(formats strfmt.Registry) error {

	if err := validate.Required("currentOffset", "body", int32(m.CurrentOffset)); err != nil {
		return err
	}

	return nil
}

func (m *ProductRatePlanPagedMetadata) validateCurrentPage(formats strfmt.Registry) error {

	if err := validate.Required("currentPage", "body", int32(m.CurrentPage)); err != nil {
		return err
	}

	return nil
}

func (m *ProductRatePlanPagedMetadata) validateExecutionTime(formats strfmt.Registry) error {

	if err := validate.Required("executionTime", "body", int64(m.ExecutionTime)); err != nil {
		return err
	}

	return nil
}

func (m *ProductRatePlanPagedMetadata) validateNextPage(formats strfmt.Registry) error {

	if err := validate.RequiredString("nextPage", "body", string(m.NextPage)); err != nil {
		return err
	}

	return nil
}

func (m *ProductRatePlanPagedMetadata) validateRecordsRequested(formats strfmt.Registry) error {

	if err := validate.Required("recordsRequested", "body", int32(m.RecordsRequested)); err != nil {
		return err
	}

	return nil
}

func (m *ProductRatePlanPagedMetadata) validateRecordsReturned(formats strfmt.Registry) error {

	if err := validate.Required("recordsReturned", "body", int32(m.RecordsReturned)); err != nil {
		return err
	}

	return nil
}

func (m *ProductRatePlanPagedMetadata) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
	}

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {

			if err := m.Results[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
