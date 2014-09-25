package utils

import (
	"strings"
)

func NewContext(controllerName string, actionName string) *Context {
	return &Context{controllerName: controllerName,
		actionName: actionName}

}

type Context struct {
	controllerName string
	actionName     string
}

func (this *Context) ControllerName() string {
	return strings.ToLower(strings.TrimSuffix(this.controllerName, "Controller"))
}

func (this *Context) ActionName() string {
	return strings.ToLower(this.actionName)
}

func (this *Context) Is(controllerAndAction string) bool {
	fullName := strings.ToLower(this.controllerName + "." + this.actionName)
	return fullName == strings.ToLower(controllerAndAction)
}
