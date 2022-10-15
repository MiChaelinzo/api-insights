// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIEndpoint ApiEndpoint
//
// swagger:model ApiEndpoint
type APIEndpoint struct {

	// Api Id
	//
	// API service this endpoint belongs to. Empty if still undetermined.
	// Format: uuid
	APIID strfmt.UUID `json:"api_id,omitempty"`

	// compliance
	Compliance *APIServiceComplianceSimple `json:"compliance,omitempty"`

	// Host
	//
	// IP v4/v6 address of the API endpoint
	// Required: true
	Host *string `json:"host"`

	// Hostname
	//
	// Hostname of the API endpoint if known
	Hostname string `json:"hostname,omitempty"`

	// id
	//
	// Unique id of the Endpoint
	// Required: true
	// Format: uuid
	Identifier *strfmt.UUID `json:"identifier"`

	// Location
	Location string `json:"location,omitempty"`

	// Port
	//
	// Port of the API endpoint
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port"`

	// proto
	// Required: true
	Proto *IPProtoEnum `json:"proto"`

	// scheme
	Scheme URLSchemeEnum `json:"scheme,omitempty"`
}

// Validate validates this Api endpoint
func (m *APIEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIEndpoint) validateAPIID(formats strfmt.Registry) error {
	if swag.IsZero(m.APIID) { // not required
		return nil
	}

	if err := validate.FormatOf("api_id", "body", "uuid", m.APIID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIEndpoint) validateCompliance(formats strfmt.Registry) error {
	if swag.IsZero(m.Compliance) { // not required
		return nil
	}

	if m.Compliance != nil {
		if err := m.Compliance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compliance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compliance")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpoint) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *APIEndpoint) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	if err := validate.FormatOf("identifier", "body", "uuid", m.Identifier.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIEndpoint) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", *m.Port, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *APIEndpoint) validateProto(formats strfmt.Registry) error {

	if err := validate.Required("proto", "body", m.Proto); err != nil {
		return err
	}

	if err := validate.Required("proto", "body", m.Proto); err != nil {
		return err
	}

	if m.Proto != nil {
		if err := m.Proto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proto")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpoint) validateScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheme) { // not required
		return nil
	}

	if err := m.Scheme.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scheme")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scheme")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Api endpoint based on the context it is used
func (m *APIEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompliance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIEndpoint) contextValidateCompliance(ctx context.Context, formats strfmt.Registry) error {

	if m.Compliance != nil {
		if err := m.Compliance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compliance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compliance")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpoint) contextValidateProto(ctx context.Context, formats strfmt.Registry) error {

	if m.Proto != nil {
		if err := m.Proto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proto")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpoint) contextValidateScheme(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Scheme.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scheme")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scheme")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIEndpoint) UnmarshalBinary(b []byte) error {
	var res APIEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
