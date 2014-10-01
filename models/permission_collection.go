package models

type PermissionCollection struct {
}

func (this *PermissionCollection) AllNames() []string {
	return []string{
		"list_users",
		"list_groups",
	}
}

func (this *PermissionCollection) allGroupNamesMap() (allGroupNamesMap map[string]bool) {
	allNames := this.AllNames()
	allGroupNamesMap = make(map[string]bool, len(allNames))
	for _, name := range allNames {
		allGroupNamesMap[GroupNameFromPermissionName(name)] = false
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
		permission := Permission{GroupName: groupName}
		abilities[permission.Name()] = userHasPermission
	}
	return
}

func (this *PermissionCollection) Create() {
	for _, name := range this.AllNames() {
		Groups().Create(GroupNameFromPermissionName(name))
	}
}

func (this *PermissionCollection) Allow(name_or_uid string, permissionName string) {
	Users().AddUserToGroup(name_or_uid, GroupNameFromPermissionName(permissionName))
}

func Permissions() *PermissionCollection {
	return &PermissionCollection{}
}
