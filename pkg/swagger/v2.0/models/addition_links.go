// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AdditionLinks addition links
//
// swagger:model AdditionLinks
type AdditionLinks map[string]AdditionLink

// Validate validates this addition links
func (m AdditionLinks) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(k)
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this addition links based on the context it is used
func (m AdditionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
