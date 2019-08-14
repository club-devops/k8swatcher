package apis

import (
	v1 "github.com/operator-framework/k8s-watcher/pkg/apis/sbtech/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
