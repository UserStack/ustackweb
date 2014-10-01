package models

import (
	"fmt"
	"strings"
)

type Permission struct {
	GroupName string // e.g. ustack.perm.users.list
}

// e.g. list_users
func (this *Permission) Name() (name string) {
	parts := strings.Split(this.GroupName, ".")
	if len(parts) == 4 {
		name = fmt.Sprintf("%s_%s", parts[3], parts[2])
	}
	return
}
