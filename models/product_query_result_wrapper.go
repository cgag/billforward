package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
ProductQueryResultWrapper product query result wrapper

swagger:model ProductQueryResultWrapper
*/
type ProductQueryResultWrapper struct {

	/* CurrentOffset current offset
	 */
	CurrentOffset int32 `json:"currentOffset,omitempty"`

	/* CurrentPage current page
	 */
	CurrentPage int32 `json:"currentPage,omitempty"`

	/* RecordsRequested records requested
	 */
	RecordsRequested int32 `json:"recordsRequested,omitempty"`

	/* RecordsReturned records returned
	 */
	RecordsReturned int32 `json:"recordsReturned,omitempty"`

	/* Results results
	 */
	Results []*Product `json:"results,omitempty"`
}

// Validate validates this product query result wrapper
func (m *ProductQueryResultWrapper) Validate(formats strfmt.Registry) error {
	return nil
}
