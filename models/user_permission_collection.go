package models


type UserPermissionCollection struct {
}

func (this *UserPermissionCollection) All(name_or_uid string) (userPermissions []*UserPermission) {
  permissions := Permissions().All()
  abilities := Permissions().Abilities(name_or_uid)
  userPermissions = make([]*UserPermission, len(permissions))
  for idx, permission := range permissions {
    assigned := abilities[permission.Name]
    userPermission := &UserPermission{Permission: permission, Assigned: assigned}
    userPermissions[idx] = userPermission
  }
  return
}

func UserPermissions() *UserPermissionCollection {
  return &UserPermissionCollection{}
}
