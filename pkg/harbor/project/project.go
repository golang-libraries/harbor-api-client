package project

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	harborErr "github.com/golang-libraries/harbor-api-client/pkg/harbor/errors"
	"github.com/golang-libraries/harbor-api-client/pkg/helpers"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/client/project"
	"github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"
)

const (
	LoggerName = "Project"
)

type ProjectReqOpt func(req *models.ProjectReq) error

func OptPublicProject(public bool) ProjectReqOpt {
	return func(req *models.ProjectReq) error {
		req.Metadata.Public = helpers.BooToStr(public)
		return nil
	}
}

func newProjectReq() *models.ProjectReq {
	return &models.ProjectReq{
		Metadata: &models.ProjectMetadata{},
	}
}

type Service interface {
	Get(ctx context.Context, nameOrID string) (*models.Project, error)
	Delete(ctx context.Context, nameOrID string) error
	Create(ctx context.Context, name string, opts ...ProjectReqOpt) (int64, error)
	Update(ctx context.Context, nameOrID string, opts ...ProjectReqOpt) error
	List(ctx context.Context, opts ...OptListProjects) ([]*models.Project, error)
	Head(ctx context.Context, name string) (bool, error)
	Log(name string) ([]*models.AuditLog, error)
	GetScanner(nameOrID string) (*models.ScannerRegistration, error)
}

type ServiceImpl struct {
	log     logr.Logger
	project project.ClientService
}

func NewProjectSvc(transport *client.Runtime, log logr.Logger) *ServiceImpl {
	return &ServiceImpl{
		log:     log.WithName(LoggerName),
		project: project.New(transport, strfmt.Default),
	}
}

func (svc *ServiceImpl) Get(ctx context.Context, nameOrID string) (*models.Project, error) {
	var params = &project.GetProjectParams{
		ProjectNameOrID: nameOrID,
		Context:         ctx,
	}

	svc.log.Info(fmt.Sprintf("Get harbor project: %s", nameOrID))
	response, err := svc.project.GetProject(params, nil)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}

func (svc *ServiceImpl) Delete(ctx context.Context, nameOrID string) error {
	params := &project.DeleteProjectParams{
		Context:         ctx,
		ProjectNameOrID: nameOrID,
	}
	_, err := svc.project.DeleteProject(params, nil)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ServiceImpl) Create(ctx context.Context, name string, opts ...ProjectReqOpt) (int64, error) {
	req := &models.ProjectReq{
		ProjectName: name,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return -1, err
		}
	}
	params := &project.CreateProjectParams{
		Context: ctx,
		Project: req,
	}
	resp, err := svc.project.CreateProject(params, nil)
	if err != nil {
		return helpers.BadID, err
	}
	id, err := helpers.ParseResourceLocation(resp.Location)
	if err != nil {
		return helpers.BadID, err
	}
	return id, nil
}

func (svc *ServiceImpl) Update(ctx context.Context, nameOrID string, opts ...ProjectReqOpt) error {
	req := newProjectReq()
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return err
		}
	}
	params := &project.UpdateProjectParams{
		Context:         ctx,
		ProjectNameOrID: nameOrID,
		Project:         req,
	}
	_, err := svc.project.UpdateProject(params, nil)
	if err != nil {
		return err
	}
	return nil
}

type OptListProjects func(params *project.ListProjectsParams) error

func OptListProjectsPublic(public bool) OptListProjects {
	return func(params *project.ListProjectsParams) error {
		params.SetPublic(&public)
		return nil
	}
}

func OptListProjectsPage(page, pageSize int64) OptListProjects {
	return func(params *project.ListProjectsParams) error {
		params.SetPage(&page)
		params.SetPageSize(&pageSize)
		return nil
	}
}

func OptListProjectsOwner(owner string) OptListProjects {
	return func(params *project.ListProjectsParams) error {
		params.SetOwner(&owner)
		return nil
	}
}

func OptListProjectsName(name string) OptListProjects {
	return func(params *project.ListProjectsParams) error {
		params.SetName(&name)
		return nil
	}
}

func OptListProjectWithDetails() OptListProjects {
	return func(params *project.ListProjectsParams) error {
		var withDetail = true
		params.SetWithDetail(&withDetail)
		return nil
	}
}

func OptListProjectSort(sort string) OptListProjects {
	return func(params *project.ListProjectsParams) error {
		params.SetSort(&sort)
		return nil
	}
}

func (svc *ServiceImpl) List(ctx context.Context, opts ...OptListProjects) ([]*models.Project, error) {
	params := &project.ListProjectsParams{
		Context: ctx,
	}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}
	resp, err := svc.project.ListProjects(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

func (svc *ServiceImpl) Head(ctx context.Context, name string) (bool, error) {
	const method harborErr.Method = "Project.Head"
	resp, err := svc.project.HeadProject(&project.HeadProjectParams{Context: ctx, ProjectName: name}, nil)
	if err != nil {
		return false, catchProjectErr(err, method)
	}
	_ = resp
	return true, nil
}

func (svc *ServiceImpl) Log(name string) ([]*models.AuditLog, error) {
	params := &project.GetLogsParams{
		ProjectName: name,
	}
	resp, err := svc.project.GetLogs(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

func (svc *ServiceImpl) GetScanner(nameOrID string) (*models.ScannerRegistration, error) {
	params := &project.GetScannerOfProjectParams{
		ProjectNameOrID: nameOrID,
	}
	resp, err := svc.project.GetScannerOfProject(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}
