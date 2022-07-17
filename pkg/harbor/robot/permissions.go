package robot

import "github.com/golang-libraries/harbor-api-client/pkg/swagger/v2.0/models"

// Actions and resource constants are borrowed from github.com/goharbor/harbor/src/common/rbac/const.go
const (
	ActionAll  = "*"    // action match any other actions
	ActionPull = "pull" // pull repository tag
	ActionPush = "push" // push repository tag

	// create, read, update, delete, list actions compatible with restful api methods
	ActionCreate = "create"
	ActionRead   = "read"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionList   = "list"

	ActionOperate     = "operate"
	ActionScannerPull = "scanner-pull" // for robot account created by scanner to pull image, bypass the policy check
	ActionStop        = "stop"         // for stop scan/scan-all execution

	LevelProject = "project"
	LevelSystem  = "system"
)

// const resource variables
const (
	ResourceAll                   = "*"             // resource match any other resources
	ResourceConfiguration         = "configuration" // project configuration compatible for portal only
	ResourceHelmChart             = "helm-chart"
	ResourceHelmChartVersion      = "helm-chart-version"
	ResourceHelmChartVersionLabel = "helm-chart-version-label"
	ResourceLabel                 = "label"
	ResourceLog                   = "log"
	ResourceLdapUser              = "ldap-user"
	ResourceMember                = "member"
	ResourceMetadata              = "metadata"
	ResourceQuota                 = "quota"
	ResourceRepository            = "repository"
	ResourceTagRetention          = "tag-retention"
	ResourceImmutableTag          = "immutable-tag"
	ResourceRobot                 = "robot"
	ResourceNotificationPolicy    = "notification-policy"
	ResourceScan                  = "scan"
	ResourceScanner               = "scanner"
	ResourceArtifact              = "artifact"
	ResourceTag                   = "tag"
	ResourceAccessory             = "accessory"
	ResourceArtifactAddition      = "artifact-addition"
	ResourceArtifactLabel         = "artifact-label"
	ResourcePreatPolicy           = "preheat-policy"
	ResourcePreatInstance         = "preheat-instance"
	ResourceSelf                  = "" // subresource for self

	ResourceAuditLog           = "audit-log"
	ResourceCatalog            = "catalog"
	ResourceProject            = "project"
	ResourceUser               = "user"
	ResourceUserGroup          = "user-group"
	ResourceRegistry           = "registry"
	ResourceReplication        = "replication"
	ResourceDistribution       = "distribution"
	ResourceGarbageCollection  = "garbage-collection"
	ResourceReplicationAdapter = "replication-adapter"
	ResourceReplicationPolicy  = "replication-policy"
	ResourceScanAll            = "scan-all"
	ResourceSystemVolumes      = "system-volumes"
)

var (
	FullAccess = []*models.Access{
		AccessTagCreate(),
		AccessTagDelete(),
		AccessRepositoryPull(),
		AccessRepositoryPush(),
		AccessArtifactDelete(),
	}
)

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

func (ap *accessAppender) AddRepositoryPull() *accessAppender {
	ap.pb.permission.Access = append(ap.pb.permission.Access, AccessRepositoryPull())
	return ap
}
func (ap *accessAppender) AddRepositoryPush() *accessAppender {
	ap.pb.permission.Access = append(ap.pb.permission.Access, AccessRepositoryPush())
	return ap
}

func (ap *accessAppender) AddFullAccess() *accessAppender {
	ap.pb.permission.Access = FullAccess
	return ap
}

func (pbc *permissionBuilderComplete) Build() *models.RobotPermission {
	return pbc.pb.permission
}

func (pb *permissionBuilder) ProjectName(projectName string) *accessAppender {
	pb.permission.Namespace = projectName
	return &accessAppender{pb: pb}
}
