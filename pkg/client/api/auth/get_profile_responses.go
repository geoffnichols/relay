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

// GetProfileReader is a Reader for the GetProfile structure.
type GetProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProfileOK creates a GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {
	return &GetProfileOK{}
}

/*GetProfileOK handles this case with default header values.

The current user's profile
*/
type GetProfileOK struct {
	Payload *GetProfileOKBody
}

func (o *GetProfileOK) Error() string {
	return fmt.Sprintf("[GET /auth/profile][%d] getProfileOK  %+v", 200, o.Payload)
}

func (o *GetProfileOK) GetPayload() *GetProfileOKBody {
	return o.Payload
}

func (o *GetProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProfileOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfileDefault creates a GetProfileDefault with default headers values
func NewGetProfileDefault(code int) *GetProfileDefault {
	return &GetProfileDefault{
		_statusCode: code,
	}
}

/*GetProfileDefault handles this case with default header values.

An error occurred
*/
type GetProfileDefault struct {
	_statusCode int

	Payload *GetProfileDefaultBody
}

// Code gets the status code for the get profile default response
func (o *GetProfileDefault) Code() int {
	return o._statusCode
}

func (o *GetProfileDefault) Error() string {
	return fmt.Sprintf("[GET /auth/profile][%d] getProfile default  %+v", o._statusCode, o.Payload)
}

func (o *GetProfileDefault) GetPayload() *GetProfileDefaultBody {
	return o.Payload
}

func (o *GetProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProfileDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetProfileDefaultBody Error response
swagger:model GetProfileDefaultBody
*/
type GetProfileDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get profile default body
func (o *GetProfileDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProfileDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProfile default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProfileDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProfileDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetProfileDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetProfileOKBody get profile o k body
swagger:model GetProfileOKBody
*/
type GetProfileOKBody struct {
	models.ProfileEntity

	// user
	User *models.UserProfile `json:"user,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetProfileOKBody) UnmarshalJSON(raw []byte) error {
	// GetProfileOKBodyAO0
	var getProfileOKBodyAO0 models.ProfileEntity
	if err := swag.ReadJSON(raw, &getProfileOKBodyAO0); err != nil {
		return err
	}
	o.ProfileEntity = getProfileOKBodyAO0

	// GetProfileOKBodyAO1
	var dataGetProfileOKBodyAO1 struct {
		User *models.UserProfile `json:"user,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetProfileOKBodyAO1); err != nil {
		return err
	}

	o.User = dataGetProfileOKBodyAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetProfileOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getProfileOKBodyAO0, err := swag.WriteJSON(o.ProfileEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getProfileOKBodyAO0)
	var dataGetProfileOKBodyAO1 struct {
		User *models.UserProfile `json:"user,omitempty"`
	}

	dataGetProfileOKBodyAO1.User = o.User

	jsonDataGetProfileOKBodyAO1, errGetProfileOKBodyAO1 := swag.WriteJSON(dataGetProfileOKBodyAO1)
	if errGetProfileOKBodyAO1 != nil {
		return nil, errGetProfileOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetProfileOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get profile o k body
func (o *GetProfileOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ProfileEntity
	if err := o.ProfileEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProfileOKBody) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProfileOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProfileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProfileOKBody) UnmarshalBinary(b []byte) error {
	var res GetProfileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
