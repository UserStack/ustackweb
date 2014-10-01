package models

import (
	"fmt"
	"strings"
)

type Permission struct {
	GroupName string // e.g. ustack.perm.user.list
}

// e.g. list_user
func (this *Permission) Name() string {
	parts := strings.Split(this.GroupName, ".")
	return fmt.Sprintf("%s_%s", parts[3], parts[2])
}
