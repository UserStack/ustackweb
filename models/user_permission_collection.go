package models


type UserPermissionCollection struct {
}

func (this *UserPermissionCollection) All(name_or_uid string) (userPermissions []*UserPermission) {
  permissions := Permissions().All()
  abilities := this.Abilities(name_or_uid)
  userPermissions = make([]*UserPermission, len(permissions))
  for idx, permission := range permissions {
    granted := abilities[permission.Name]
    userPermission := &UserPermission{Permission: permission, Granted: granted}
    userPermissions[idx] = userPermission
  }
  return
}

func (this *UserPermissionCollection) NoAbilities() (allGroupNamesMap map[string]bool) {
  allNames := Permissions().AllNames()
  allGroupNamesMap = make(map[string]bool, len(allNames))
  for _, name := range allNames {
    allGroupNamesMap[name] = false
  }
  return
}

func (this *UserPermissionCollection) Abilities(name_or_uid string) (abilities map[string]bool) {
  abilities = this.NoAbilities()
  groups, _ := Groups().AllByUser(name_or_uid)
  for _, group := range groups {
    name := Permissions().Name(group.Name)
    if _, isInternalPermissionGroup := abilities[name]; isInternalPermissionGroup {
      abilities[name] = true
    }
  }
  return
}

func (this *UserPermissionCollection) AllowAll(name_or_uid string) {
  for _, permission := range Permissions().All() {
    this.Allow(name_or_uid, permission.Name)
  }
}

func (this *UserPermissionCollection) Allow(name_or_uid string, permissionName string) {
  Users().AddUserToGroup(name_or_uid, Permissions().GroupName(permissionName))
}

func (this *UserPermissionCollection) Deny(name_or_uid string, permissionName string) {
  Users().RemoveUserFromGroup(name_or_uid, Permissions().GroupName(permissionName))
}

func UserPermissions() *UserPermissionCollection {
  return &UserPermissionCollection{}
}
