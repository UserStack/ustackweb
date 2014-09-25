package models

import (
	"github.com/UserStack/ustackd/backends"
)

type UserCollection struct {
}

func (this UserCollection) All() []backends.User {
	return []backends.User{
		backends.User{Uid: 1, Email: "foo"},
		backends.User{Uid: 2, Email: "admin"},
		backends.User{Uid: 3, Email: "abc"},
		backends.User{Uid: 4, Email: "def"},
		backends.User{Uid: 5, Email: "hij"},
		backends.User{Uid: 6, Email: "glk"},
		backends.User{Uid: 7, Email: "uvw"},
		backends.User{Uid: 8, Email: "xyz"}}
}

func (this UserCollection) Find(uid int) backends.User {
	for _, user := range this.All() {
		if user.Uid == uid {
			return user
		}
	}
	return backends.User{}
}

func Users() UserCollection {
	return UserCollection{}
}
