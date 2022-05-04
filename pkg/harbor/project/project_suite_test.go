package project_test

import (
	"context"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor"
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
	Describe("Head projects", func() {
		Context("Head project which not exists", func() {
			It("Should throw err, and False result", func() {
				const (
					projectName = "not-exists-project"
				)
				exist, err := client.Project.Head(ctx, projectName)
				debugLog.Printf("ERROR STR: %s\n", err.Error())
				Expect(err).ShouldNot(Succeed())
				Expect(exist).Should(BeFalse())
			})

		})
	})
})
