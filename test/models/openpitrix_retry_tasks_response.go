// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRetryTasksResponse openpitrix retry tasks response
// swagger:model openpitrixRetryTasksResponse
type OpenpitrixRetryTasksResponse struct {

	// task set
	TaskSet OpenpitrixRetryTasksResponseTaskSet `json:"task_set"`
}

// Validate validates this openpitrix retry tasks response
func (m *OpenpitrixRetryTasksResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRetryTasksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRetryTasksResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRetryTasksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
