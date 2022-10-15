// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenAPISpecScoreElementsList open Api spec score elements list
//
// swagger:model OpenApiSpecScoreElementsList
type OpenAPISpecScoreElementsList struct {

	// elements
	Elements []*OpenAPISpecScoreElement `json:"elements"`

	// severity
	Severity APISecurityRiskSeverity `json:"severity,omitempty"`
}

// Validate validates this open Api spec score elements list
func (m *OpenAPISpecScoreElementsList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenAPISpecScoreElementsList) validateElements(formats strfmt.Registry) error {
	if swag.IsZero(m.Elements) { // not required
		return nil
	}

	for i := 0; i < len(m.Elements); i++ {
		if swag.IsZero(m.Elements[i]) { // not required
			continue
		}

		if m.Elements[i] != nil {
			if err := m.Elements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenAPISpecScoreElementsList) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	if err := m.Severity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("severity")
		}
		return err
	}

	return nil
}

// ContextValidate validate this open Api spec score elements list based on the context it is used
func (m *OpenAPISpecScoreElementsList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenAPISpecScoreElementsList) contextValidateElements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Elements); i++ {

		if m.Elements[i] != nil {
			if err := m.Elements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenAPISpecScoreElementsList) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Severity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("severity")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenAPISpecScoreElementsList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenAPISpecScoreElementsList) UnmarshalBinary(b []byte) error {
	var res OpenAPISpecScoreElementsList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}