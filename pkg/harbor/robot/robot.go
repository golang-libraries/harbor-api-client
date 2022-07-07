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

//ect: , Resource: repository, Action: push
//Effect: , Resource: repository, Action: pull
//Effect: , Resource: artifact, Action: delete
//Effect: , Resource: helm-chart, Action: read
//Effect: , Resource: helm-chart-version, Action: create
//Effect: , Resource: helm-chart-version, Action: delete
//Effect: , Resource: tag, Action: create
//Effect: , Resource: tag, Action: delete
//Effect: , Resource: artifact-label, Action: create
//Effect: , Resource: scan, Action: createEff

const (
	ActionPull   = "pull"
	ActionPush   = "push"
	ActionDelete = "delete"
	ActionRead   = "read"
	ActionCreate = "create"

	LevelProject = "project"
	LevelSystem  = "system"

	ResourceRepository = "repository"
	ResourceArtifact   = "artifact"
	ResourceTag        = "tag"

	ScopeAllProjects       = "*"
	NeverExpires     int64 = -1
)

func ExpiresAfterDays(days int64) int64 {
	return days
}

func ExpiresAfterYears(years int64) int64 {
	return years * 365
}

func ExpiresAfterMonths(months int64) int64 {
	return months * 30
}

func ExpiresNever() int64 {
	return NeverExpires
}

func AccessTagCreate() *models.Access {
	return &models.Access{
		Action:   ActionCreate,
		Resource: ResourceTag,
	}
}

func AccessTagDelete() *models.Access {
	return &models.Access{
		Action:   ActionDelete,
		Resource: ResourceTag,
	}
}

func AccessArtifactDelete() *models.Access {
	return &models.Access{
		Action:   ActionDelete,
		Resource: ResourceArtifact,
	}
}

func AccessRepositoryPull() *models.Access {
	return &models.Access{
		Action:   ActionPull,
		Resource: ResourceRepository,
	}
}

func AccessRepositoryPush() *models.Access {
	return &models.Access{
		Action:   ActionPush,
		Resource: ResourceRepository,
	}
}

func NewPermission(projectName string, access []*models.Access) *models.RobotPermission {
	return &models.RobotPermission{
		Access:    access,
		Kind:      LevelProject,
		Namespace: projectName,
	}
}

func NewPermissionList(permissions ...*models.RobotPermission) []*models.RobotPermission {
	permissionList := make([]*models.RobotPermission, 0)
	if len(permissions) == 0 {
		permissionList[0] = &models.RobotPermission{}
	}
	permissionList = append(permissionList, permissions...)
	return permissionList
}

func NewAccessList(accesses ...*models.Access) []*models.Access {
	accessList := make([]*models.Access, 0)
	if len(accesses) == 0 {
		accessList[0] = &models.Access{}
	}
	accessList = append(accessList, accesses...)
	return accessList
}

func PullAndPushPermissions(projectName string) *models.RobotPermission {
	return NewPermission(projectName, NewAccessList(AccessRepositoryPush(), AccessRepositoryPull()))
}

type CreateOpt func(robotCreate *models.RobotCreate) error

func WithCreateDescription(description string) CreateOpt {
	return func(robotCreate *models.RobotCreate) error {
		robotCreate.Description = description
		return nil
	}
}

func WithCreateDuration(duration int64) CreateOpt {
	return func(robotCreate *models.RobotCreate) error {
		robotCreate.Duration = duration
		return nil
	}
}
func WithCreateSecret(secret string) CreateOpt {
	return func(robotCreate *models.RobotCreate) error {
		robotCreate.Secret = secret
		return nil
	}
}
func WithCreateDisabled() CreateOpt {
	return func(robotCreate *models.RobotCreate) error {
		robotCreate.Disable = true
		return nil
	}
}

type Service interface {
	Get(ctx context.Context, id int64) (*models.Robot, error)
	Create(ctx context.Context, robotName, level string, permissions []*models.RobotPermission, opts ...CreateOpt) (*models.RobotCreated, error)
	Delete(ctx context.Context, robotID int64) error
	List(ctx context.Context, query string, opts ...ListOpt) ([]*models.Robot, error)
	Update(ctx context.Context, robotID int64, robotModel *models.Robot) error
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
	const method = "Get"
	params := &robot.GetRobotByIDParams{
		RobotID: id,
		Context: ctx,
	}
	resp, err := svc.robot.GetRobotByID(params, nil)
	if err != nil {
		return nil, catchRobotErr(err, method)
	}
	return resp.GetPayload(), nil
}

func (svc *ServiceImpl) Create(ctx context.Context, robotName, level string, permissions []*models.RobotPermission, opts ...CreateOpt) (*models.RobotCreated, error) {
	const method = "Create"
	if level == LevelProject && len(permissions) > 1 {
		return nil, catchRobotErr(ErrMultiplayPermissions, method)
	}
	newRobot := &models.RobotCreate{
		Name:        robotName,
		Permissions: permissions,
		Level:       level,
	}
	for _, opt := range opts {
		if err := opt(newRobot); err != nil {
			return nil, err
		}
	}
	params := &robot.CreateRobotParams{
		Context: ctx,
		Robot:   newRobot,
	}
	resp, err := svc.robot.CreateRobot(params, nil)
	if err != nil {
		return nil, catchRobotErr(err, method)
	}
	return resp.GetPayload(), nil
}

func (svc *ServiceImpl) Delete(ctx context.Context, robotID int64) error {
	const method = "Delete"
	params := &robot.DeleteRobotParams{
		Context: ctx,
		RobotID: robotID,
	}
	resp, err := svc.robot.DeleteRobot(params, nil)
	if err != nil {
		return catchRobotErr(err, method)
	}
	_ = resp
	return nil
}

type ListOpt func(listParams *robot.ListRobotParams) error

func WithListOptions(page, pageSize int64, sort string) ListOpt {
	return func(listParams *robot.ListRobotParams) error {
		listParams.Sort = &sort
		listParams.Page = &page
		listParams.PageSize = &pageSize
		return nil
	}
}

func (svc *ServiceImpl) List(ctx context.Context, query string, opts ...ListOpt) ([]*models.Robot, error) {
	const method = "List"
	params := &robot.ListRobotParams{
		Context: ctx,
		Q:       &query,
	}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, catchRobotErr(err, method)
		}
	}
	resp, err := svc.robot.ListRobot(params, nil)
	if err != nil {
		return nil, catchRobotErr(err, method)
	}
	return resp.GetPayload(), nil

}

//Update Robot Account

type UpdateOpt func(robotModel *models.Robot) error

func WithUpdateDuration(duration int64) UpdateOpt {
	return func(robotModel *models.Robot) error {
		robotModel.Duration = duration
		return nil
	}
}

type robotUpdateBuilder struct {
	robotModel *models.Robot
}

func (ub *robotUpdateBuilder) Build() *models.Robot {
	return ub.robotModel
}

func (ub *robotUpdateBuilder) Duration(duration int64) *robotUpdateBuilder {
	ub.robotModel.Duration = duration
	return ub
}

func (ub *robotUpdateBuilder) Description(description string) *robotUpdateBuilder {
	ub.robotModel.Description = description
	return ub
}

func NewUpdateBuilder(existRobot *models.Robot) *robotUpdateBuilder {
	return &robotUpdateBuilder{
		robotModel: existRobot,
	}
}

func (svc *ServiceImpl) Update(ctx context.Context, robotID int64, robotModel *models.Robot) error {
	const method = "Update"
	params := &robot.UpdateRobotParams{
		Context: ctx,
		RobotID: robotID,
		Robot:   robotModel,
	}
	resp, err := svc.robot.UpdateRobot(params, nil)
	_ = resp
	if err != nil {
		return catchRobotErr(err, method)
	}
	return nil
}
