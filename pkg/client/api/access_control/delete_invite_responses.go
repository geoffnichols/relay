// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/puppetlabs/relay/pkg/client/api/models"
)

// DeleteInviteReader is a Reader for the DeleteInvite structure.
type DeleteInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteInviteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteInviteOK creates a DeleteInviteOK with default headers values
func NewDeleteInviteOK() *DeleteInviteOK {
	return &DeleteInviteOK{}
}

/*DeleteInviteOK handles this case with default header values.

The resource was deleted
*/
type DeleteInviteOK struct {
	Payload *DeleteInviteOKBody
}

func (o *DeleteInviteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/invites/{inviteId}][%d] deleteInviteOK  %+v", 200, o.Payload)
}

func (o *DeleteInviteOK) GetPayload() *DeleteInviteOKBody {
	return o.Payload
}

func (o *DeleteInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteInviteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInviteDefault creates a DeleteInviteDefault with default headers values
func NewDeleteInviteDefault(code int) *DeleteInviteDefault {
	return &DeleteInviteDefault{
		_statusCode: code,
	}
}

/*DeleteInviteDefault handles this case with default header values.

An error occurred
*/
type DeleteInviteDefault struct {
	_statusCode int

	Payload *DeleteInviteDefaultBody
}

// Code gets the status code for the delete invite default response
func (o *DeleteInviteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInviteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/invites/{inviteId}][%d] deleteInvite default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteInviteDefault) GetPayload() *DeleteInviteDefaultBody {
	return o.Payload
}

func (o *DeleteInviteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteInviteDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteInviteDefaultBody Error response
swagger:model DeleteInviteDefaultBody
*/
type DeleteInviteDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this delete invite default body
func (o *DeleteInviteDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInviteDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteInvite default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInviteDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInviteDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteInviteDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteInviteOKBody Information about the deleted resource
swagger:model DeleteInviteOKBody
*/
type DeleteInviteOKBody struct {

	// Deleted resource id
	// Required: true
	ResourceID *string `json:"resource_id"`

	// Was the resource successfully deleted?
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this delete invite o k body
func (o *DeleteInviteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteInviteOKBody) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("deleteInviteOK"+"."+"resource_id", "body", o.ResourceID); err != nil {
		return err
	}

	return nil
}

func (o *DeleteInviteOKBody) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("deleteInviteOK"+"."+"success", "body", o.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteInviteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInviteOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteInviteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
