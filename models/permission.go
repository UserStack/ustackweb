package models

import (
	"fmt"
	"strings"
)

type Permission struct {
	GroupName string // e.g. perm.users.list
}

// e.g. list_users
func (this *Permission) Name() (name string) {
	parts := strings.Split(this.GroupName, ".")
	if len(parts) == 3 {
		name = fmt.Sprintf("%s_%s", parts[2], parts[1])
	}
	return
}

// e.g. list_users => perm.users.list
func GroupNameFromPermissionName(name string) (groupName string) {
	parts := strings.Split(name, "_")
	if len(parts) == 2 {
		groupName = fmt.Sprintf("perm.%s.%s", parts[1], parts[0])
	}
	return
}
