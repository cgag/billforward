package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*PendingComponentValueChange PendingComponentValueChange pending component value change

swagger:model PendingComponentValueChange
*/
type PendingComponentValueChange struct {

	/* At at
	 */
	At *strfmt.DateTime `json:"at,omitempty"`

	/* DiscardHTTPVerb discard http verb
	 */
	DiscardHTTPVerb *string `json:"discardHttpVerb,omitempty"`

	/* DiscardURL discard url
	 */
	DiscardURL *string `json:"discardUrl,omitempty"`

	/* Value value
	 */
	Value *int32 `json:"value,omitempty"`
}

// Validate validates this pending component value change
func (m *PendingComponentValueChange) Validate(formats strfmt.Registry) error {
	return nil
}
