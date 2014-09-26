package models

import (
	"fmt"
	"github.com/UserStack/ustackd/backends"
)

type UserCollection struct {
	backend backends.Abstract
}

func (this *UserCollection) All() []User {
	var err *backends.Error
	fmt.Println(err)
	_, err = this.backend.CreateUser("pauluxxxks", "barssx")
	fmt.Println(err)
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

func (this *UserCollection) Destroy(username string) (deleted bool) {
	err := this.backend.DeleteUser(username)
	deleted = err == nil
	return
}

func Users() *UserCollection {
	backend, _ := backends.NewSqliteBackend("./tmp.db")
	return &UserCollection{backend: &backend}
}
