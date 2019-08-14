// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifier":       schema_pkg_apis_sbtech_v1_K8SWatcherNotifier(ref),
		"github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifierSpec":   schema_pkg_apis_sbtech_v1_K8SWatcherNotifierSpec(ref),
		"github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifierStatus": schema_pkg_apis_sbtech_v1_K8SWatcherNotifierStatus(ref),
	}
}

func schema_pkg_apis_sbtech_v1_K8SWatcherNotifier(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "K8SWatcherNotifier is the Schema for the k8swatchernotifiers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifierSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifierStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifierSpec", "github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1.K8SWatcherNotifierStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sbtech_v1_K8SWatcherNotifierSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "K8SWatcherNotifierSpec defines the desired state of K8SWatcherNotifier",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sbtech_v1_K8SWatcherNotifierStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "K8SWatcherNotifierStatus defines the observed state of K8SWatcherNotifier",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
