package project_test

import (
	"context"
	"fmt"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor"
	nanoid "github.com/matoous/go-nanoid/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	"testing"
)

func TestProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Project Suite")
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
	debugLog = log.New(GinkgoWriter, "testProjectSvc", log.Flags())
})

var _ = Describe("Tests for Project Service", func() {
	const (
		notExistProjectName = "not-exists-project"
		emptyProjectName    = ""
	)

	var (
		projectName string
	)
	BeforeEach(func() {
		projectName = prepareTestProject()
	})

	Describe("Head projects", func() {

		//var (
		//	projectName string
		//)
		//BeforeEach(func() {
		//	projectName = prepareTestProject()
		//})

		Context("Head project which is not existed or project with empty name", func() {
			It("Nil error and as False result, just get know that project doesn't exist", func() {
				exist, err := client.Project.Head(ctx, notExistProjectName)
				Expect(err).Should(Succeed())
				Expect(exist).Should(BeFalse())
			})
			It("Try tp head project with empty name", func() {
				exist, err := client.Project.Head(ctx, emptyProjectName)
				debugLog.Printf("Test error: %s\n", err.Error())
				Expect(err).ShouldNot(Succeed())
				Expect(exist).Should(BeFalse())
			})
			Context("Try to head existed project", func() {
				It("Should get project successfully", func() {
					exist, err := client.Project.Head(ctx, projectName)
					Expect(err).Should(Succeed())
					Expect(exist).Should(BeTrue())
				})
			})

		})
		//AfterEach(func() {
		//	deleteTestProject(projectName)
		//})
	})
	Describe("Test .Delete() project", func() {
		//var (
		//	projectName string
		//)
		//BeforeEach(func() {
		//	projectName = prepareTestProject()
		//})
		Context("Delete already exists project", func() {
			It("Should correctly delete existing project without errors.", func() {
				err := client.Project.Delete(ctx, projectName)
				Expect(err).Should(Succeed())
			})
		})
		//AfterEach(func() {
		//	deleteTestProject(projectName)
		//})
	})
	Describe("Test .Get() project", func() {
		//var (
		//	projectName string
		//)
		//BeforeEach(func() {
		//	projectName = prepareTestProject()
		//})
		Context("Get projects with wrong parameters", func() {
			It("Try to get not exists project", func() {
				project, err := client.Project.Get(ctx, notExistProjectName)
				debugLog.Printf("ERR FROM NOT EXISTS PROJECT: %s\n", err)
				Expect(err).ShouldNot(Succeed())
				Expect(project).Should(BeNil())

			})
			It("Get exists project", func() {
				project, err := client.Project.Get(ctx, projectName)
				Expect(err).Should(Succeed())
				Expect(project).ShouldNot(BeNil())
			})
		})

	})
	Describe("Test getting project audit log", func() {
		It("Should get at least one audit log event", func() {
			event, err := client.Project.Log(ctx, projectName)
			Expect(err).Should(Succeed())
			Expect(event).ShouldNot(BeNil())
		})
	})
	AfterEach(func() {
		deleteTestProject(projectName)
	})
})

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
