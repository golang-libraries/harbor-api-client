package project

import (
	harborErr "github.com/golang-libraries/harbor-api-client/pkg/harbor/errors"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/project"
)

func catchProjectErr(err error, method harborErr.Method) harborErr.Error {
	var kind harborErr.Kind
	var originErr error
	var statusCode harborErr.StatusCode
	var msg harborErr.Message
	switch e := err.(type) {
	case *project.HeadProjectNotFound:
		kind = harborErr.NotFoundKind
		originErr = e
		statusCode = harborErr.StatusCode(404)
		msg = "Harbor project with such name or ID doesn't exist."
		errStack := e.GetPayload().Errors
		_ = errStack

	default:
		return harborErr.New(err.Error(), method)

	}
	return harborErr.New(kind, originErr, statusCode, msg, method)
}
