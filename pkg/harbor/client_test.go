package harbor

import (
	"context"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor/project"
	"testing"
)

var ctx = context.TODO()

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	err = client.Project.Update(ctx, "demo-project", project.OptPublicProject(false))
	if err != nil {
		t.Fatal(err)
	}
	projects, err := client.Project.List(ctx, project.OptListProjectsPublic(false))
	if err != nil {
		t.Fatal(err)
	}
	for _, project := range projects {
		t.Logf("project ID: %d\n", project.ProjectID)
	}
}
