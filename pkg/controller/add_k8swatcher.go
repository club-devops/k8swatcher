package controller

import (
	"github.com/operator-framework/k8s-watcher/pkg/controller/k8swatcher"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, k8swatcher.Add)
}
