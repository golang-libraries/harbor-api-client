package robot

import harborErr "github.com/golang-libraries/harbor-api-client/pkg/harbor/errors"

func catchRobotErr(err error, method harborErr.Method) harborErr.Error {
	const (
		service = harborErr.Service(ServiceName)
	)
	var newErr = harborErr.New(err, method, service)
	return newErr

}
