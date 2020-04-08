// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/puppetlabs/relay/pkg/client/api/models"
)

// GetWorkflowsReader is a Reader for the GetWorkflows structure.
type GetWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWorkflowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowsOK creates a GetWorkflowsOK with default headers values
func NewGetWorkflowsOK() *GetWorkflowsOK {
	return &GetWorkflowsOK{}
}

/*GetWorkflowsOK handles this case with default header values.

A list of workflows
*/
type GetWorkflowsOK struct {
	Payload *GetWorkflowsOKBody
}

func (o *GetWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflowsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowsOK) GetPayload() *GetWorkflowsOKBody {
	return o.Payload
}

func (o *GetWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsDefault creates a GetWorkflowsDefault with default headers values
func NewGetWorkflowsDefault(code int) *GetWorkflowsDefault {
	return &GetWorkflowsDefault{
		_statusCode: code,
	}
}

/*GetWorkflowsDefault handles this case with default header values.

An error occurred
*/
type GetWorkflowsDefault struct {
	_statusCode int

	Payload *GetWorkflowsDefaultBody
}

// Code gets the status code for the get workflows default response
func (o *GetWorkflowsDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowsDefault) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflows default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowsDefault) GetPayload() *GetWorkflowsDefaultBody {
	return o.Payload
}

func (o *GetWorkflowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWorkflowsDefaultBody Error response
swagger:model GetWorkflowsDefaultBody
*/
type GetWorkflowsDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get workflows default body
func (o *GetWorkflowsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflows default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBody An object containing an array of workflows
swagger:model GetWorkflowsOKBody
*/
type GetWorkflowsOKBody struct {

	// A list of workflows
	Workflows []*models.Workflow `json:"workflows"`
}

// Validate validates this get workflows o k body
func (o *GetWorkflowsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBody) validateWorkflows(formats strfmt.Registry) error {

	if swag.IsZero(o.Workflows) { // not required
		return nil
	}

	for i := 0; i < len(o.Workflows); i++ {
		if swag.IsZero(o.Workflows[i]) { // not required
			continue
		}

		if o.Workflows[i] != nil {
			if err := o.Workflows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowsOK" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
