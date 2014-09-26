package models

import (
	"fmt"
	"github.com/UserStack/ustackd/backends"
	"github.com/UserStack/ustackd/client"
)

type UserCollection struct {
	backend backends.Abstract
}

func (this *UserCollection) All() []User {
	backendUsers, _ := this.backend.Users()
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
	id, _ := this.backend.LoginUser(username, password)
	loggedIn = id > 0
	return
}

func (this *UserCollection) Create(username string, password string) (created bool, id int64) {
	id, _ = this.backend.CreateUser(username, password)
	created = id > 0
	return
}

func (this *UserCollection) Destroy(name_or_uid string) (deleted bool) {
	err := this.backend.DeleteUser(name_or_uid)
	deleted = err == nil
	return
}

func Users() *UserCollection {
	backend, _ := client.Dial("127.0.0.1:7654")
	return &UserCollection{backend: backend}
}
