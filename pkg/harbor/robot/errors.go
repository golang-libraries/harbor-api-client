package robot

import (
	"errors"
	harborErr "github.com/golang-libraries/harbor-api-client/pkg/harbor/errors"
)

var (
	ErrMultiplayPermissions = errors.New("multiplay permissions for project level robot account")
	ErrRobotNotFound        = errors.New("robot account with such name isn't found")
)

func catchRobotErr(err error, method harborErr.Method) harborErr.Error {
	const (
		service = harborErr.Service(ServiceName)
	)
	var newErr = harborErr.New(err, method, service)
	return newErr

}
