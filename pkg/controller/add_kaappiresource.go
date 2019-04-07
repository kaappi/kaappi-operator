package controller

import (
	"github.com/kaappi/kaappi-operator/pkg/controller/kaappiresource"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kaappiresource.Add)
}
