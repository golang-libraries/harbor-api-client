package robot

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/robot"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

const (
	ServiceName = "Robot"
)

type Service interface {
	Get(ctx context.Context, id int64) (*models.Robot, error)
}

type ServiceImpl struct {
	log   logr.Logger
	robot robot.ClientService
}

func NewRobotSvc(transport *client.Runtime, log logr.Logger) *ServiceImpl {
	return &ServiceImpl{
		log:   log.WithName(ServiceName),
		robot: robot.New(transport, strfmt.Default),
	}
}

func (svc ServiceImpl) Get(ctx context.Context, id int64) (*models.Robot, error) {
	const method = "Get()"
	params := &robot.GetRobotByIDParams{
		RobotID: id,
	}
	resp, err := svc.robot.GetRobotByID(params, nil)
	if err != nil {
		return nil, catchRobotErr(err, method)
	}
	return resp.GetPayload(), nil
}
