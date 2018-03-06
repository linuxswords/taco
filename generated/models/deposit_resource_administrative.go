// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DepositResourceAdministrative Administrative metadata for the SDR resource.
// swagger:model depositResourceAdministrative
type DepositResourceAdministrative struct {

	// If this resource should be sent to Preservation.
	// Required: true
	SdrPreserve *bool `json:"sdrPreserve"`
}

// Validate validates this deposit resource administrative
func (m *DepositResourceAdministrative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSdrPreserve(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepositResourceAdministrative) validateSdrPreserve(formats strfmt.Registry) error {

	if err := validate.Required("sdrPreserve", "body", m.SdrPreserve); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DepositResourceAdministrative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepositResourceAdministrative) UnmarshalBinary(b []byte) error {
	var res DepositResourceAdministrative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
