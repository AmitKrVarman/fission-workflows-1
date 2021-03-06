package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TypedValue Copy of protobuf's Any, to avoid protobuf requirement of a protobuf-based type.
// swagger:model TypedValue
type TypedValue struct {

	// type
	Type string `json:"type,omitempty"`

	// value
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this typed value
func (m *TypedValue) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
