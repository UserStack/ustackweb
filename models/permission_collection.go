package models

type PermissionCollection struct {
}

func (this *PermissionCollection) AllGroupNames() []string {
	return []string{
		"ustack-perm-users-list",
		"ustack-perm-users-read",
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

func (this *PermissionCollection) AllByUser(name_or_uid string) (permissions map[string]bool) {
	permissions = this.allGroupNamesMap()
	groups, _ := Groups().AllByUser(name_or_uid)
	for _, group := range groups {
		if _, isPermissionGroup := permissions[group.Name]; isPermissionGroup {
			permissions[group.Name] = true
		}
	}
	return
}

func Permissions() *PermissionCollection {
	return &PermissionCollection{}
}
