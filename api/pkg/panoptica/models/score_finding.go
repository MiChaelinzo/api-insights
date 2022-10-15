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
	"github.com/go-openapi/validate"
)

// ScoreFinding ScoreFinding
//
// swagger:model ScoreFinding
type ScoreFinding struct {

	// Additional Info
	//
	// Finding instance additional attributes
	AdditionalInfo []*AdditionalInfo `json:"additional_info"`

	// Data
	//
	// Raw data filled in only for unclassified findings, when allowed
	Data interface{} `json:"data,omitempty"`

	// Description
	//
	// Finding description if finding has been classified, None otherwise
	Description string `json:"description,omitempty"`

	// Mitigation
	//
	// Finding mitigation if finding has been classified, None otherwise
	Mitigation string `json:"mitigation,omitempty"`

	// Name
	//
	// Finding name
	// Required: true
	Name *string `json:"name"`

	// Occurrences
	//
	// Number of findings of the same type
	Occurrences int64 `json:"occurrences,omitempty"`

	// Raw Finding Id
	//
	// Identifier of the raw finding if available
	// Format: uuid
	RawFindingID strfmt.UUID `json:"raw_finding_id,omitempty"`

	// Source
	//
	// Finding source or filled in with 'Undisclosed' if source cannot be revealed
	// Required: true
	Source *string `json:"source"`

	// Type Id
	//
	// Finding type identifier if finding has been classified, None otherwise
	// Format: uuid
	TypeID strfmt.UUID `json:"type_id,omitempty"`
}

// Validate validates this score finding
func (m *ScoreFinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRawFindingID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScoreFinding) validateAdditionalInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalInfo); i++ {
		if swag.IsZero(m.AdditionalInfo[i]) { // not required
			continue
		}

		if m.AdditionalInfo[i] != nil {
			if err := m.AdditionalInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additional_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScoreFinding) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ScoreFinding) validateRawFindingID(formats strfmt.Registry) error {
	if swag.IsZero(m.RawFindingID) { // not required
		return nil
	}

	if err := validate.FormatOf("raw_finding_id", "body", "uuid", m.RawFindingID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScoreFinding) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *ScoreFinding) validateTypeID(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeID) { // not required
		return nil
	}

	if err := validate.FormatOf("type_id", "body", "uuid", m.TypeID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this score finding based on the context it is used
func (m *ScoreFinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScoreFinding) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalInfo); i++ {

		if m.AdditionalInfo[i] != nil {
			if err := m.AdditionalInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additional_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScoreFinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScoreFinding) UnmarshalBinary(b []byte) error {
	var res ScoreFinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
