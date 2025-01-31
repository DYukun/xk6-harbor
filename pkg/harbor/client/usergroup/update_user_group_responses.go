// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// UpdateUserGroupReader is a Reader for the UpdateUserGroup structure.
type UpdateUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserGroupOK creates a UpdateUserGroupOK with default headers values
func NewUpdateUserGroupOK() *UpdateUserGroupOK {
	return &UpdateUserGroupOK{}
}

/* UpdateUserGroupOK describes a response with status code 200, with default header values.

Success
*/
type UpdateUserGroupOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateUserGroupOK) Error() string {
	return fmt.Sprintf("[PUT /usergroups/{group_id}][%d] updateUserGroupOK ", 200)
}

func (o *UpdateUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateUserGroupBadRequest creates a UpdateUserGroupBadRequest with default headers values
func NewUpdateUserGroupBadRequest() *UpdateUserGroupBadRequest {
	return &UpdateUserGroupBadRequest{}
}

/* UpdateUserGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateUserGroupBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /usergroups/{group_id}][%d] updateUserGroupBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateUserGroupBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupUnauthorized creates a UpdateUserGroupUnauthorized with default headers values
func NewUpdateUserGroupUnauthorized() *UpdateUserGroupUnauthorized {
	return &UpdateUserGroupUnauthorized{}
}

/* UpdateUserGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateUserGroupUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /usergroups/{group_id}][%d] updateUserGroupUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateUserGroupUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupForbidden creates a UpdateUserGroupForbidden with default headers values
func NewUpdateUserGroupForbidden() *UpdateUserGroupForbidden {
	return &UpdateUserGroupForbidden{}
}

/* UpdateUserGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUserGroupForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /usergroups/{group_id}][%d] updateUserGroupForbidden  %+v", 403, o.Payload)
}
func (o *UpdateUserGroupForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupNotFound creates a UpdateUserGroupNotFound with default headers values
func NewUpdateUserGroupNotFound() *UpdateUserGroupNotFound {
	return &UpdateUserGroupNotFound{}
}

/* UpdateUserGroupNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateUserGroupNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateUserGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /usergroups/{group_id}][%d] updateUserGroupNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUserGroupNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupInternalServerError creates a UpdateUserGroupInternalServerError with default headers values
func NewUpdateUserGroupInternalServerError() *UpdateUserGroupInternalServerError {
	return &UpdateUserGroupInternalServerError{}
}

/* UpdateUserGroupInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateUserGroupInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateUserGroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /usergroups/{group_id}][%d] updateUserGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateUserGroupInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
