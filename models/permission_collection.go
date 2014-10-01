package models

type PermissionCollection struct {
}

func (this *PermissionCollection) AllGroupNames() []string {
	return []string{
		"ustack.perm.users.list",
		"ustack.perm.groups.list",
	}
}

func (this *PermissionCollection) allGroupNamesMap() (allGroupNamesMap map[string]bool) {
	allGroupNames := this.AllGroupNames()
	allGroupNamesMap = make(map[string]bool, len(allGroupNamesMap))
	for _, groupName := range allGroupNames {
		allGroupNamesMap[groupName] = false
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

func Permissions() *PermissionCollection {
	return &PermissionCollection{}
}
