package models

type PermissionCollection struct {
}

func (this *PermissionCollection) AllGroupNamesMap() map[string]bool {
	return map[string]bool{
		"ustack-perm-users-list": false,
		"ustack-perm-users-read": false,
	}
}

func (this *PermissionCollection) AllGroupNames() (allGroupNames []string) {
	allGroupNamesMap := this.AllGroupNamesMap()
	allGroupNames = make([]string, len(allGroupNamesMap))
	i := 0
	for groupName, _ := range allGroupNamesMap {
		allGroupNames[i] = groupName
		i++
	}
	return
}

func (this *PermissionCollection) AllByUser(name_or_uid string) (permissions map[string]bool) {
	permissions = this.AllGroupNamesMap()
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
