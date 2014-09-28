package models

import (
	"fmt"
	"github.com/UserStack/ustackweb/backend"
)

type UserCollection struct {
}

func (this *UserCollection) All() []User {
	backendUsers, _ := backend.Current().Users()
	users := make([]User, len(backendUsers))
	for idx, backendUser := range backendUsers {
		users[idx] = User{backendUser}
	}
	fmt.Println(users)
	return users
}

func (this *UserCollection) Find(uid int64) *User {
	for _, user := range this.All() {
		if user.Uid == uid {
			return &user
		}
	}
	return &User{}
}

func (this *UserCollection) Login(username string, password string) (loggedIn bool) {
	id, _ := backend.Current().LoginUser(username, password)
	loggedIn = id > 0
	return
}

func (this *UserCollection) Create(username string, password string) (created bool, id int64) {
	id, _ = backend.Current().CreateUser(username, password)
	created = id > 0
	return
}

func (this *UserCollection) UpdateUsername(name_or_uid string, password string, newUsername string) (updated bool) {
	err := backend.Current().ChangeUserName(name_or_uid, password, newUsername)
	updated = err == nil
	return
}

func (this *UserCollection) UpdatePassword(name_or_uid string, password string, newPassword string) (updated bool) {
	err := backend.Current().ChangeUserPassword(name_or_uid, password, newPassword)
	updated = err == nil
	return
}

func (this *UserCollection) Destroy(name_or_uid string) (deleted bool) {
	err := backend.Current().DeleteUser(name_or_uid)
	deleted = err == nil
	return
}

func Users() *UserCollection {
	return &UserCollection{}
}
