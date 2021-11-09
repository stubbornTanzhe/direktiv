// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetInstanceListReader is a Reader for the GetInstanceList structure.
type GetInstanceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstanceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstanceListOK creates a GetInstanceListOK with default headers values
func NewGetInstanceListOK() *GetInstanceListOK {
	return &GetInstanceListOK{}
}

/* GetInstanceListOK describes a response with status code 200, with default header values.

successfully got namespace instances
*/
type GetInstanceListOK struct {
}

func (o *GetInstanceListOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/instances][%d] getInstanceListOK ", 200)
}

func (o *GetInstanceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}