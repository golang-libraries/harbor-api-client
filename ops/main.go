package main

import (
	"fmt"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	apiext "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apiextensions"
	"os"
	//appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	defaultNamespace          = "harbor"
	defaultHarborInstanseName = "harbor"
	helmCRApiVersion          = "helm.cattle.io/v1"
	helmCRKind                = "HelmChart"
)

func readValues(filePath string) (string, error) {
	var values string
	valuesFile, err := os.ReadFile(filePath)
	if err != nil {
		return values, err
	}
	values = fmt.Sprintf("%s", valuesFile)
	return values, nil
}

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		pulumi.Printf("Install Harbor registry, by creating HelmChart resource\n")
		meta := metav1.ObjectMetaArgs{}
		ns, err := corev1.NewNamespace(ctx, defaultNamespace, &corev1.NamespaceArgs{Metadata: meta})
		if err != nil {
			return err
		}
		values, err := readValues("./values.yaml")
		chartMeta := metav1.ObjectMetaArgs{
			Namespace: ns.Metadata.Name(),
			Name:      pulumi.StringPtr(defaultHarborInstanseName),
		}
		spec := kubernetes.UntypedArgs{
			"spec": map[string]interface{}{
				"chart":           "harbor",
				"repo":            "https://helm.goharbor.io",
				"targetNamespace": ns.Metadata.Name(),
				"valuesContent":   values,
			},
		}
		cr := apiext.CustomResourceArgs{ApiVersion: pulumi.String(helmCRApiVersion),
			Kind:        pulumi.String(helmCRKind),
			Metadata:    chartMeta,
			OtherFields: spec}
		helmChart, err := apiext.NewCustomResource(ctx, defaultHarborInstanseName, &cr)
		if err != nil {
			return err
		}
		_ = helmChart
		return nil
	})
}
