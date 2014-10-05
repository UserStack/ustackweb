package models

import (
	"fmt"
	"strings"
)

type PermissionCollection struct {
}

func (this *PermissionCollection) AllNames() []string {
	return []string{
		"list_users",
		"create_users",
		"read_users",
		"update_users",
		"delete_users",
		"enable_users",
		"disable_users",

		"list_groups",
		"create_groups",
		"read_groups",
		"delete_groups",

		"list_permissions",
		"create_permissions",
		"delete_permissions",
		"grant_permissions",
		"revoke_permissions",

		"list_stats",
	}
}

func (this *PermissionCollection) AllInternal() (permissions []*Permission) {
	permissions = make([]*Permission, 0)
	for _, name := range this.AllNames() {
		permissions = append(permissions, &Permission{Name: name, Internal: true})
	}
	return
}

func (this *PermissionCollection) All() (permissions []*Permission) {
	permissions = make([]*Permission, 0)
	for _, permission := range this.AllInternal() {
		permissions = append(permissions, permission)
	}
	permissionGroups, _ := Groups().AllPermissions()
	for _, group := range permissionGroups {
		permissionName := this.Name(group.Name)
		permissionFound := false
		for _, permission := range permissions {
			if permission.Name == permissionName {
				permissionFound = true
				break
			}
		}
		if !permissionFound {
			permissions = append(permissions, &Permission{Name: permissionName})
		}
	}
	return
}

func (this *PermissionCollection) CreateAllInternal() {
	for _, permission := range this.AllInternal() {
		Groups().Create(permission.GroupName())
	}
}

func (this *PermissionCollection) Create(object string, verb string) {
	Groups().Create(this.GroupName(this.BuildName(object, verb)))
}

func (this *PermissionCollection) Destroy(name string) {
	Groups().Destroy(this.GroupName(name))
}

func (this *PermissionCollection) IsPermissionGroupName(groupName string) (isPermissionGroupName bool) {
	parts := strings.Split(groupName, ".")
	return len(parts) == 3 && parts[0] == "perm"
}

// e.g. users, list => users_list
func (this *PermissionCollection) BuildName(object string, verb string) string {
	return strings.Join([]string{object, verb}, "_")
}

// e.g. list_users => perm.users.list
func (this *PermissionCollection) GroupName(name string) (groupName string) {
	parts := strings.Split(name, "_")
	if len(parts) == 2 {
		groupName = fmt.Sprintf("perm.%s.%s", parts[1], parts[0])
	}
	return
}

// e.g. list_users
func (this *PermissionCollection) Name(groupName string) (name string) {
	parts := strings.Split(groupName, ".")
	if this.IsPermissionGroupName(groupName) {
		name = fmt.Sprintf("%s_%s", parts[2], parts[1])
	}
	return
}

func Permissions() *PermissionCollection {
	return &PermissionCollection{}
}
