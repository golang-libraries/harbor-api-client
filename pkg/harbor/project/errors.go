package project

import (
	"errors"
	harborErr "github.com/golang-libraries/harbor-api-client/pkg/harbor/errors"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/project"
)

func catchProjectErr(err error, method harborErr.Method) harborErr.Error {
	const (
		service = harborErr.Service("Project")
	)
	var newErr = harborErr.New(err, method, service)

	return newErr

}

func IsNotFound(err error) bool {
	var headProjectNotFound *project.HeadProjectNotFound
	if errors.As(err, &headProjectNotFound) {
		return true
	}
	return false
}
