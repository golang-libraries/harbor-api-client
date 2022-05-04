package harbor

import (
	"errors"
	"github.com/go-openapi/runtime"
	projects "github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/project"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/robot"
	"net/http"
)

const (
	UnknownKind = "UnknownErrorKind"
)

type Kind string

type Error struct {
	Kind        string
	Path        string
	Err         error
	RawResponse interface{}
	StatusCode  int
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func NewError(err error) *Error {
	var apiErr *runtime.APIError
	if ok := errors.As(err, &apiErr); ok {
		return &Error{
			Err:         err,
			Path:        apiErr.OperationName,
			RawResponse: apiErr.Response,
		}
	}
	return &Error{
		Err: err,
	}
}

var (
	ErrProjectNotFound              = errors.New("project with such name isn't found")
	ErrWrongProjectName             = errors.New("wrong project name")
	ErrFailedToParseProjectLocation = errors.New("failed to parse project resource location")
	ErrRobotNotFound                = errors.New("robot account with such name isn't found")
	ErrProjectWithoutRobots         = errors.New("there're no robot accounts in project")
	ErrMoreThanOneItemInResponse    = errors.New("more than one item in the response for given resource name")
	ErrLdapGroupNotFound            = errors.New("ldap group is not found")
	//	ErrGroupIsNotFound           = errors.New("there's  no group with such name")
	ErrFailedToParseResourceLocation = errors.New("failed to parse resource location")
	ErrThereIsNoGroupsAtAll          = errors.New("there's no groups at this harbor registry instance")
	ErrFailedToGetSystemInfo         = errors.New("failed to get full information scope about Harbor instance")
	//ErrCreateProjectBadRequest = errors.New("create project bad request")
	ErrCreateRobotBadRequest = errors.New("create robot bad request")
)

func IsForbiddenApiErr(err error) bool {
	var projectDeleteForbiddenErr *projects.DeleteProjectForbidden
	if errors.As(err, &projectDeleteForbiddenErr) {
		return true
	}
	var forbiddenApiErr *runtime.APIError
	if errors.As(err, &forbiddenApiErr) && forbiddenApiErr.Code == http.StatusForbidden {
		return true
	}
	return false
}

func IsNotFoundApiErr(err error) bool {
	var headProjectNotFoundErr *projects.HeadProjectNotFound
	if errors.As(err, &headProjectNotFoundErr) {
		return true
	}
	return false
}

func IsBadRequestApiErr(err error) bool {
	var robotCreateBadRequest *robot.CreateRobotBadRequest
	if errors.As(err, &robotCreateBadRequest) {
		return true
	}
	return false
}
