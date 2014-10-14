package backends

type NilBackend struct {
}

func (backend *NilBackend) CreateUser(name string, password string) (int64, *Error) {
	return 0, nil
}

func (backend *NilBackend) DisableUser(nameuid string) *Error {
	return nil
}

func (backend *NilBackend) EnableUser(nameuid string) *Error {
	return nil
}

func (backend *NilBackend) SetUserData(nameuid string, key string, value string) *Error {
	return nil
}

func (backend *NilBackend) GetUserData(nameuid string, key string) (string, *Error) {
	return "", nil
}

func (backend *NilBackend) GetUserDataKeys(nameuid string) (keys []string, err *Error) {
	return
}

func (backend *NilBackend) LoginUser(name string, password string) (int64, *Error) {
	return 0, nil
}

func (backend *NilBackend) ChangeUserPassword(nameuid string, password string, newpassword string) *Error {
	return nil
}

func (backend *NilBackend) ChangeUserName(nameuid string, password string, newname string) *Error {
	return nil
}

func (backend *NilBackend) UserGroups(nameuid string) ([]Group, *Error) {
	return nil, nil
}

func (backend *NilBackend) DeleteUser(nameuid string) *Error {
	return nil
}

func (backend *NilBackend) Users() ([]User, *Error) {
	return nil, nil
}

func (backend *NilBackend) CreateGroup(name string) (int64, *Error) {
	return 0, nil
}

func (backend *NilBackend) AddUserToGroup(nameuid string, groupgid string) *Error {
	return nil
}

func (backend *NilBackend) RemoveUserFromGroup(nameuid string, groupgid string) *Error {
	return nil
}

func (backend *NilBackend) DeleteGroup(groupgid string) *Error {
	return nil
}

func (backend *NilBackend) Groups() ([]Group, *Error) {
	return nil, nil
}

func (backend *NilBackend) GroupUsers(groupgid string) ([]User, *Error) {
	return nil, nil
}

func (backend *NilBackend) Stats() (map[string]int64, *Error) {
	return nil, nil
}

func (backend *NilBackend) Close() {

}
