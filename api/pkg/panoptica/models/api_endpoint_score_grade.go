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

// APIEndpointScoreGrade ApiEndpointScoreGrade
//
// swagger:model ApiEndpointScoreGrade
type APIEndpointScoreGrade struct {

	// Additional Info
	AdditionalInfo []*AdditionalInfo `json:"additional_info"`

	// Categories
	// Required: true
	Categories map[string]CategoryScoreGrade `json:"categories"`

	// confidence
	Confidence RiskConfidenceEnum `json:"confidence,omitempty"`

	// counters history
	CountersHistory *CountersHistory `json:"counters_history,omitempty"`

	// endpoint
	// Required: true
	Endpoint *APIEndpoint `json:"endpoint"`

	// endpoint id
	// Format: uuid
	EndpointID strfmt.UUID `json:"endpoint_id,omitempty"`

	// risk
	// Required: true
	Risk *APISecurityRiskSeverity `json:"risk"`

	// Scorer Version
	// Required: true
	ScorerVersion *int64 `json:"scorer_version"`

	// trend
	Trend RiskTrendEnum `json:"trend,omitempty"`
}

// Validate validates this Api endpoint score grade
func (m *APIEndpointScoreGrade) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountersHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScorerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrend(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIEndpointScoreGrade) validateAdditionalInfo(formats strfmt.Registry) error {
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

func (m *APIEndpointScoreGrade) validateCategories(formats strfmt.Registry) error {

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
	}

	for k := range m.Categories {

		if err := validate.Required("categories"+"."+k, "body", m.Categories[k]); err != nil {
			return err
		}
		if val, ok := m.Categories[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("categories" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIEndpointScoreGrade) validateConfidence(formats strfmt.Registry) error {
	if swag.IsZero(m.Confidence) { // not required
		return nil
	}

	if err := m.Confidence.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("confidence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("confidence")
		}
		return err
	}

	return nil
}

func (m *APIEndpointScoreGrade) validateCountersHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.CountersHistory) { // not required
		return nil
	}

	if m.CountersHistory != nil {
		if err := m.CountersHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters_history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counters_history")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpointScoreGrade) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpointScoreGrade) validateEndpointID(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointID) { // not required
		return nil
	}

	if err := validate.FormatOf("endpoint_id", "body", "uuid", m.EndpointID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIEndpointScoreGrade) validateRisk(formats strfmt.Registry) error {

	if err := validate.Required("risk", "body", m.Risk); err != nil {
		return err
	}

	if err := validate.Required("risk", "body", m.Risk); err != nil {
		return err
	}

	if m.Risk != nil {
		if err := m.Risk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("risk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("risk")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpointScoreGrade) validateScorerVersion(formats strfmt.Registry) error {

	if err := validate.Required("scorer_version", "body", m.ScorerVersion); err != nil {
		return err
	}

	return nil
}

func (m *APIEndpointScoreGrade) validateTrend(formats strfmt.Registry) error {
	if swag.IsZero(m.Trend) { // not required
		return nil
	}

	if err := m.Trend.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trend")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("trend")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Api endpoint score grade based on the context it is used
func (m *APIEndpointScoreGrade) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfidence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountersHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIEndpointScoreGrade) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *APIEndpointScoreGrade) contextValidateCategories(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
	}

	for k := range m.Categories {

		if val, ok := m.Categories[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *APIEndpointScoreGrade) contextValidateConfidence(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Confidence.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("confidence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("confidence")
		}
		return err
	}

	return nil
}

func (m *APIEndpointScoreGrade) contextValidateCountersHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.CountersHistory != nil {
		if err := m.CountersHistory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters_history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counters_history")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpointScoreGrade) contextValidateEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.Endpoint != nil {
		if err := m.Endpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpointScoreGrade) contextValidateRisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Risk != nil {
		if err := m.Risk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("risk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("risk")
			}
			return err
		}
	}

	return nil
}

func (m *APIEndpointScoreGrade) contextValidateTrend(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Trend.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trend")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("trend")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIEndpointScoreGrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIEndpointScoreGrade) UnmarshalBinary(b []byte) error {
	var res APIEndpointScoreGrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
