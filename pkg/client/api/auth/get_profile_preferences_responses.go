// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/puppetlabs/relay/pkg/client/api/models"
)

// GetProfilePreferencesReader is a Reader for the GetProfilePreferences structure.
type GetProfilePreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfilePreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProfilePreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProfilePreferencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProfilePreferencesOK creates a GetProfilePreferencesOK with default headers values
func NewGetProfilePreferencesOK() *GetProfilePreferencesOK {
	return &GetProfilePreferencesOK{}
}

/*GetProfilePreferencesOK handles this case with default header values.

The current user's public preferences
*/
type GetProfilePreferencesOK struct {
	Payload models.UserPreferences
}

func (o *GetProfilePreferencesOK) Error() string {
	return fmt.Sprintf("[GET /auth/profile/preferences][%d] getProfilePreferencesOK  %+v", 200, o.Payload)
}

func (o *GetProfilePreferencesOK) GetPayload() models.UserPreferences {
	return o.Payload
}

func (o *GetProfilePreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilePreferencesDefault creates a GetProfilePreferencesDefault with default headers values
func NewGetProfilePreferencesDefault(code int) *GetProfilePreferencesDefault {
	return &GetProfilePreferencesDefault{
		_statusCode: code,
	}
}

/*GetProfilePreferencesDefault handles this case with default header values.

An error occurred
*/
type GetProfilePreferencesDefault struct {
	_statusCode int

	Payload *GetProfilePreferencesDefaultBody
}

// Code gets the status code for the get profile preferences default response
func (o *GetProfilePreferencesDefault) Code() int {
	return o._statusCode
}

func (o *GetProfilePreferencesDefault) Error() string {
	return fmt.Sprintf("[GET /auth/profile/preferences][%d] getProfilePreferences default  %+v", o._statusCode, o.Payload)
}

func (o *GetProfilePreferencesDefault) GetPayload() *GetProfilePreferencesDefaultBody {
	return o.Payload
}

func (o *GetProfilePreferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProfilePreferencesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetProfilePreferencesDefaultBody Error response
swagger:model GetProfilePreferencesDefaultBody
*/
type GetProfilePreferencesDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get profile preferences default body
func (o *GetProfilePreferencesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProfilePreferencesDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProfilePreferences default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProfilePreferencesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProfilePreferencesDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetProfilePreferencesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
