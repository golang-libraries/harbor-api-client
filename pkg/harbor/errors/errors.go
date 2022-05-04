package errors

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	projects "github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/project"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/robot"
	"net/http"
	"strconv"
)

const (
	UnknownKind          = "UnknownErrorKind"
	NotFoundKind         = "NotFound"
	ConnectionFailedKind = "ConnectionFailed"
)

type Kind string

type StatusCode int

type Message string

type Method string

type Error struct {
	Kind
	Err error
	Method
	StatusCode
	Message
}

func (e Error) Error() string {
	var errStr = fmt.Sprintf("ErrKind: %s, Method: %s", e.Kind, e.Method)
	if e.StatusCode != 0 {
		errStr = fmt.Sprintf("ErrKind: %s, Method: %s, Status Code: %d", e.Kind, e.Method, e.StatusCode)
	}
	if e.Message != "" {
		errStr = fmt.Sprintf("ErrKind: %s, Method: %s, Status Code: %d, Description: %s", e.Kind, e.Method, e.StatusCode, e.Message)
	}
	return errStr
}

func (e Error) Unwrap() error {
	return e.Err
}

func ErrStatusCode(err error) int {
	harbErr, ok := err.(*Error)
	if ok {
		return int(harbErr.StatusCode)
	}
	return 0
}

func New(args ...interface{}) Error {
	newErr := Error{}
	for _, arg := range args {
		switch val := arg.(type) {
		case *runtime.APIError:
			newErr.StatusCode = StatusCode(val.Code)
			newErr.Err = val
		case Kind:
			newErr.Kind = val
		case StatusCode:
			newErr.StatusCode = val
		case Message:
			newErr.Message = val
		case error:
			newErr.Err = val
		case Method:
			newErr.Method = val
		}
	}
	return newErr
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

func StrToStatusCode(code string) StatusCode {
	i := StrToInt32(code)
	return StatusCode(i)
}

func StrToInt32(num string) int {
	i, err := strconv.Atoi(num)
	_ = err
	return i
}

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
