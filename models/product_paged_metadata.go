package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ProductPagedMetadata product paged metadata

swagger:model ProductPagedMetadata
*/
type ProductPagedMetadata struct {

	/* {"description":"Paging parameter. 0-indexed. Describes your current location within a pageable list of query results.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	CurrentOffset *int32 `json:"currentOffset"`

	/* {"description":"Paging parameter. 0-indexed. Describes which page (given a page size of `recordsRequested`) of the result set you are viewing.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	CurrentPage *int32 `json:"currentPage"`

	/* {"description":"Number of milliseconds taken by API to calculate response.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	ExecutionTime *int64 `json:"executionTime"`

	/* {"description":"Paging parameter. URL fragment that can be used to fetch next page of results.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	NextPage *string `json:"nextPage"`

	/* {"default":10,"description":"Paging parameter. Describes how many records you requested.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	RecordsRequested *int32 `json:"recordsRequested"`

	/* {"description":"Describes how many records were returned by your query.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	RecordsReturned *int32 `json:"recordsReturned"`

	/* {"description":"The results returned by your query.","verbs":["GET","PUT","POST"]}

	Required: true
	*/
	Results []*Product `json:"results"`
}

// Validate validates this product paged metadata
func (m *ProductPagedMetadata) Validate(formats strfmt.Registry) error {
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

func (m *ProductPagedMetadata) validateCurrentOffset(formats strfmt.Registry) error {

	if err := validate.Required("currentOffset", "body", m.CurrentOffset); err != nil {
		return err
	}

	return nil
}

func (m *ProductPagedMetadata) validateCurrentPage(formats strfmt.Registry) error {

	if err := validate.Required("currentPage", "body", m.CurrentPage); err != nil {
		return err
	}

	return nil
}

func (m *ProductPagedMetadata) validateExecutionTime(formats strfmt.Registry) error {

	if err := validate.Required("executionTime", "body", m.ExecutionTime); err != nil {
		return err
	}

	return nil
}

func (m *ProductPagedMetadata) validateNextPage(formats strfmt.Registry) error {

	if err := validate.Required("nextPage", "body", m.NextPage); err != nil {
		return err
	}

	return nil
}

func (m *ProductPagedMetadata) validateRecordsRequested(formats strfmt.Registry) error {

	if err := validate.Required("recordsRequested", "body", m.RecordsRequested); err != nil {
		return err
	}

	return nil
}

func (m *ProductPagedMetadata) validateRecordsReturned(formats strfmt.Registry) error {

	if err := validate.Required("recordsReturned", "body", m.RecordsReturned); err != nil {
		return err
	}

	return nil
}

func (m *ProductPagedMetadata) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
	}

	return nil
}
