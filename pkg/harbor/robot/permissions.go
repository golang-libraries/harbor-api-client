package robot

import "github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"

//Build complete permissions list
type permissionListBuilder struct {
	permissions []*models.RobotPermission
}

type permissionListBuilderSystem struct {
	pbl *permissionListBuilder
}

func (pbs *permissionListBuilderSystem) Add(perm *models.RobotPermission) *permissionListBuilderSystem {
	pbs.pbl.permissions = append(pbs.pbl.permissions, perm)
	return pbs
}

func (pbs *permissionListBuilderSystem) Build() []*models.RobotPermission {
	return pbs.pbl.permissions
}
func NewPermissionListBuilder() *permissionListBuilder {
	return &permissionListBuilder{
		permissions: make([]*models.RobotPermission, 0),
	}
}

func (pbl *permissionListBuilder) ProjectPermissions(perm *models.RobotPermission) []*models.RobotPermission {
	pbl.permissions = append(pbl.permissions, perm)
	return pbl.permissions
}

func (pbl *permissionListBuilder) SystemLevel() *permissionListBuilderSystem {
	return &permissionListBuilderSystem{pbl: pbl}
}

//Single permission object builder
type permissionBuilder struct {
	permission *models.RobotPermission
}

type permissionBuilderComplete struct {
	pb *permissionBuilder
}

type accessAppender struct {
	pb *permissionBuilder
}

func NewPermissionsBuilder() *permissionBuilder {
	return &permissionBuilder{
		permission: &models.RobotPermission{
			Kind: LevelProject,
		},
	}
}

func (ap *accessAppender) AddAccess(access *models.Access) *permissionBuilderComplete {
	ap.pb.permission.Access = append(ap.pb.permission.Access, access)
	return &permissionBuilderComplete{pb: ap.pb}
}

func (pbc *permissionBuilderComplete) Build() *models.RobotPermission {
	return pbc.pb.permission
}

func (pb *permissionBuilder) ProjectName(projectName string) *accessAppender {
	pb.permission.Namespace = projectName
	return &accessAppender{pb: pb}
}
