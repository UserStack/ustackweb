package models

import (
	"fmt"
	"strings"
)

type PermissionCollection struct {
}

func (this *PermissionCollection) allNames() []string {
	return []string{
		"list_users",
		"list_groups",
	}
}

func (this *PermissionCollection) All() (permissions []*Permission) {
	names := this.allNames()
	permissions = make([]*Permission, len(names))
	for idx, name := range names {
		permissions[idx] = &Permission{Name: name}
	}
	return
}

func (this *PermissionCollection) allGroupNamesMap() (allGroupNamesMap map[string]bool) {
	allNames := this.allNames()
	allGroupNamesMap = make(map[string]bool, len(allNames))
	for _, name := range allNames {
		allGroupNamesMap[this.GroupName(name)] = false
	}
	return
}

func (this *PermissionCollection) allGroupNamesMapByUser(name_or_uid string) (groupNamesMapByUser map[string]bool) {
	groupNamesMapByUser = this.allGroupNamesMap()
	groups, _ := Groups().AllByUser(name_or_uid)
	for _, group := range groups {
		if _, isPermissionGroup := groupNamesMapByUser[group.Name]; isPermissionGroup {
			groupNamesMapByUser[group.Name] = true
		}
	}
	return
}

func (this *PermissionCollection) Abilities(name_or_uid string) (abilities map[string]bool) {
	groupNames := this.allGroupNamesMapByUser(name_or_uid)
	abilities = make(map[string]bool, len(groupNames))
	for groupName, userHasPermission := range groupNames {
		abilities[this.Name(groupName)] = userHasPermission
	}
	return
}

func (this *PermissionCollection) Create() {
	for _, name := range this.allNames() {
		Groups().Create(this.GroupName(name))
	}
}

func (this *PermissionCollection) Allow(name_or_uid string, permissionName string) {
	Users().AddUserToGroup(name_or_uid, this.GroupName(permissionName))
}

func (this *PermissionCollection) Deny(name_or_uid string, permissionName string) {
	Users().RemoveUserFromGroup(name_or_uid, this.GroupName(permissionName))
}

func (this *PermissionCollection) IsPermissionGroupName(groupName string) (isPermissionGroupName bool) {
	parts := strings.Split(groupName, ".")
	return len(parts) == 3 && parts[0] == "perm"
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
