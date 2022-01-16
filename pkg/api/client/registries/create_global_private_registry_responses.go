// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateGlobalPrivateRegistryReader is a Reader for the CreateGlobalPrivateRegistry structure.
type CreateGlobalPrivateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGlobalPrivateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGlobalPrivateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGlobalPrivateRegistryOK creates a CreateGlobalPrivateRegistryOK with default headers values
func NewCreateGlobalPrivateRegistryOK() *CreateGlobalPrivateRegistryOK {
	return &CreateGlobalPrivateRegistryOK{}
}

/* CreateGlobalPrivateRegistryOK describes a response with status code 200, with default header values.

successfully created global private registry
*/
type CreateGlobalPrivateRegistryOK struct {
}

func (o *CreateGlobalPrivateRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/functions/registries/private][%d] createGlobalPrivateRegistryOK ", 200)
}

func (o *CreateGlobalPrivateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateGlobalPrivateRegistryBody create global private registry body
// Example: {"data":"admin:8QwFLg%D$qg*","reg":"https://prod.customreg.io"}
swagger:model CreateGlobalPrivateRegistryBody
*/
type CreateGlobalPrivateRegistryBody struct {

	// Target registry connection data containing the user and token.
	// Required: true
	Data *string `json:"data"`

	// Target registry URL
	// Required: true
	Reg *string `json:"reg"`
}

// Validate validates this create global private registry body
func (o *CreateGlobalPrivateRegistryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateGlobalPrivateRegistryBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Registry Payload"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

func (o *CreateGlobalPrivateRegistryBody) validateReg(formats strfmt.Registry) error {

	if err := validate.Required("Registry Payload"+"."+"reg", "body", o.Reg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create global private registry body based on context it is used
func (o *CreateGlobalPrivateRegistryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateGlobalPrivateRegistryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateGlobalPrivateRegistryBody) UnmarshalBinary(b []byte) error {
	var res CreateGlobalPrivateRegistryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}