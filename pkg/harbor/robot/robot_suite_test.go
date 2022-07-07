package robot_test

import (
	"context"
	"fmt"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor/robot"
	nanoid "github.com/matoous/go-nanoid/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"testing"
)

func TestProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RobotAccount Suite")
}

var (
	client   *harbor.Client
	ctx      = context.Background()
	debugLog *log.Logger
)

var _ = BeforeSuite(func() {
	var err error
	client, err = harbor.NewClient()
	Expect(err).Should(Succeed())
	Expect(client).ShouldNot(BeNil())
	debugLog = log.New(GinkgoWriter, "testRobotSvc", log.Flags())
})

var _ = Describe("Test RobotAccount Svc", func() {
	const (
		baseRobotName = "test-robot"
	)

	var (
		robotID     int64
		robotName   string
		projectName string
	)
	BeforeEach(func() {
		projectName = prepareTestProject()
		robotName = genUniqueName(baseRobotName)
		robotID = prepareTestSystemtRobot(projectName)

	})
	It("Get robot account", func() {
		robot, err := client.Robot.Get(ctx, robotID)
		Expect(err).Should(Succeed())
		Expect(robot).ShouldNot(BeNil())
	})
	It("Create system robot account for certain project", func() {
		permissions := robot.NewPermissionListBuilder().SystemLevel().Add(robot.NewPermissionsBuilder().ProjectName(projectName).AddAccess(robot.AccessRepositoryPull()).Build()).Build()
		newRobot, err := client.Robot.Create(ctx, robotName, robot.LevelSystem, permissions)
		Expect(err).Should(Succeed())
		Expect(newRobot).ShouldNot(BeNil())
		deleteTestRobotAccount(newRobot.ID)

	})
	It("Create project scope robot account", func() {
		permissions := robot.NewPermissionListBuilder().ProjectPermissions(robot.NewPermissionsBuilder().ProjectName(projectName).AddAccess(robot.AccessRepositoryPush()).Build())
		newRobot, err := client.Robot.Create(ctx, robotName, robot.LevelProject, permissions)
		Expect(err).Should(Succeed())
		Expect(newRobot).ShouldNot(BeNil())
	})
	It("Update robot account", func() {
		const (
			description = "Test robot account"
		)
		getRobot, err := client.Robot.Get(ctx, robotID)
		Expect(err).Should(Succeed())
		Expect(getRobot).ShouldNot(BeNil())
		err = client.Robot.Update(ctx, robotID, robot.NewUpdateBuilder(getRobot).Description(description).Duration(robot.NeverExpires).Build())
		Expect(err).Should(Succeed())
	})
	AfterEach(func() {
		deleteTestRobotAccount(robotID)
		deleteTestProject(projectName)

	})

})

func genUniqueName(name string) string {
	id, _ := nanoid.Generate("abcdefg", 12)
	return fmt.Sprintf("%s-%s", name, id)
}

func prepareTestSystemtRobot(projectName string) int64 {
	_ = projectName
	id, err := nanoid.Generate("abcdefg", 12)
	robotName := fmt.Sprintf("test-robot-%s", id)
	newRobot, err := client.Robot.Create(ctx, robotName, robot.LevelSystem, robot.NewPermissionList(robot.PullAndPushPermissions(projectName)))
	Expect(err).Should(Succeed())
	Expect(newRobot).ShouldNot(BeNil())
	return newRobot.ID
}

func prepareTestProject() string {
	id, err := nanoid.Generate("abcdefg", 12)
	ExpectWithOffset(1, err).Should(Succeed())
	projectName := fmt.Sprintf("test-project-%s", id)
	projectId, err := client.Project.Create(ctx, projectName)
	Expect(err).Should(Succeed())
	Expect(id).ShouldNot(BeZero())
	debugLog.Printf("Test project prepared, name: %s,  ID: %d", projectName, projectId)
	return projectName
}

func deleteTestProject(nameOrID string) {
	projectExist, err := client.Project.Head(ctx, nameOrID)
	ExpectWithOffset(1, err).Should(Succeed())
	if projectExist {
		debugLog.Printf("Delete test harbor project: %s\n", nameOrID)
		err := client.Project.Delete(ctx, nameOrID)
		ExpectWithOffset(1, err).Should(Succeed())
	}
}

func deleteTestRobotAccount(id int64) {
	err := client.Robot.Delete(ctx, id)
	Expect(err).Should(Succeed())
}
