package models

import (
	"github.com/UserStack/ustackweb/backend"
)

type UserCollection struct {
}

func (this *UserCollection) All() (users []User, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendUsers, backendError := connection.Users()
	if backendError == nil {
		users = make([]User, len(backendUsers))
		for idx, backendUser := range backendUsers {
			users[idx] = User{backendUser}
		}
	}
	return
}

func (this *UserCollection) Find(uid int64) (user *User, err *backend.Error) {
	allUsers, err := this.All()
	if err != nil {
		return
	}
	for _, aUser := range allUsers {
		if aUser.Uid == uid {
			user = &aUser
			return
		}
	}
	return
}

func (this *UserCollection) Login(username string, password string) (loggedIn bool, id int64, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	id, backendError := connection.LoginUser(username, password)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		loggedIn = id > 0
	}
	return
}

func (this *UserCollection) Create(username string, password string) (created bool, id int64, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	id, backendError := connection.CreateUser(username, password)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		created = id > 0
	}
	return
}

func (this *UserCollection) UpdateUsername(name_or_uid string, password string, newUsername string) (updated bool, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendError := connection.ChangeUserName(name_or_uid, password, newUsername)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		updated = err == nil
	}
	return
}

func (this *UserCollection) UpdatePassword(name_or_uid string, password string, newPassword string) (updated bool, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendError := connection.ChangeUserPassword(name_or_uid, password, newPassword)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		updated = err == nil
	}
	return
}

func (this *UserCollection) Destroy(name_or_uid string) (deleted bool, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendError := connection.DeleteUser(name_or_uid)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		deleted = err == nil
	}
	return
}

func (this *UserCollection) Enable(name_or_uid string) (deleted bool, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendError := connection.EnableUser(name_or_uid)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		deleted = err == nil
	}
	return
}

func (this *UserCollection) Disable(name_or_uid string) (deleted bool, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendError := connection.DisableUser(name_or_uid)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		deleted = err == nil
	}
	return
}

func Users() *UserCollection {
	return &UserCollection{}
}
