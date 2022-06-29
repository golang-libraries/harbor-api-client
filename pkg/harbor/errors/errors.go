package errors

import (
	"errors"
	"fmt"
	openApiErrs "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

const (
	UnauthorizedKind        = "Unauthorized"
	InternalServerErrorKind = "InternalServerError"
	UnknownKind             = "UnknownErrorKind"
	NotFoundKind            = "NotFound"
	ConnectionFailedKind    = "ConnectionFailed"
	UnprocessableEntityKind = "UnprocessableEntity"
	KindUnknownAPIErr       = "APIError"
	KindKnownAPIErr         = "HTTPRespError"

	StatusCodeUnknown             = 0
	StatusCodeInternalServerErr   = 500
	StatusCodeNotFound            = 404
	StatusCodeUnauthorized        = 401
	StatusCodeUnprocessableEntity = 422
	StatusCodeStatusCode          = 409
	StatusCodeForbidden           = 403
)

var (
	ErrFailedToParseResourceLocation = errors.New("failed to parse resource location")
)

type ResponseErr interface {
	GetPayload() *models.Errors
}

type Kind string

type StatusCode int

type Description string

type Method string

type Service string

type Error struct {
	Kind
	Err error
	Method
	Service
	StatusCode
	Description
}

type ModErr func(err Error)

func WithDescription(msg string) ModErr {
	return func(err Error) {
		err.Description = Description(msg)
	}
}
func WithMethod(method Method) ModErr {
	return func(err Error) {
		err.Method = method
	}
}

func WithErr(e error) ModErr {
	return func(err Error) {
		err.Err = e
	}
}

func WithSvc(svc Service) ModErr {
	return func(err Error) {
		err.Service = svc
	}
}

func WithModelErrors(modelErrors *models.Errors, harborErr *Error) {
	err := modelErrors.Validate(strfmt.Default)
	var compositeErr *openApiErrs.CompositeError
	if errors.As(err, &compositeErr) {
		harborErr.StatusCode = StatusCode(compositeErr.Code())
		harborErr.Err = compositeErr
	}
}

func (e Error) Error() string {
	var errStr = fmt.Sprintf("KIND: %s, METHOD: %s, UPSTREAM_ERR: %s", e.Kind, e.Method, e.Err.Error())
	if e.StatusCode != 0 {
		errStr = fmt.Sprintf("%s, STATUS_CODE: %d", errStr, e.StatusCode)
	}
	if e.Description != "" {
		errStr = fmt.Sprintf("%s, DESCRIPTION: %s", errStr, e.Description)
	}
	return errStr
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) IsNotFound() bool {
	return e.StatusCode == StatusCodeNotFound
}

func New(args ...interface{}) Error {
	newErr := Error{
		Kind: UnknownKind,
	}
	for _, arg := range args {
		switch val := arg.(type) {
		case Kind:
			newErr.Kind = val
		case StatusCode:
			newErr.StatusCode = val
		case Description:
			newErr.Description = val
		case error:
			newErr.Err = val
			fromUpstreamErr(val, &newErr)
		case Method:
			newErr.Method = val
		case Service:
			newErr.Service = val
		}
	}

	return newErr
}

func fromUpstreamErr(err error, harborErr *Error) {
	switch upstreamErr := err.(type) {
	case *runtime.APIError:
		harborErr.StatusCode = StatusCode(upstreamErr.Code)
		harborErr.Kind = KindUnknownAPIErr
		harborErr.Err = err
	case ResponseErr:
		harborErr.Kind = KindKnownAPIErr
	}

}
