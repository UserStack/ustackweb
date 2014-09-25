package models

import (
	"github.com/UserStack/ustackd/backends"
)

type UserCollection struct {
}

func (this *UserCollection) All() []User {
	return []User{
		User{&backends.User{Uid: 1, Email: "foo"}},
		User{&backends.User{Uid: 2, Email: "admin"}},
		User{&backends.User{Uid: 3, Email: "abc"}},
		User{&backends.User{Uid: 4, Email: "def"}},
		User{&backends.User{Uid: 5, Email: "hij"}},
		User{&backends.User{Uid: 6, Email: "glk"}},
		User{&backends.User{Uid: 7, Email: "uvw"}},
		User{&backends.User{Uid: 8, Email: "xyz"}}}
}

func (this *UserCollection) Find(uid int64) *User {
	for _, user := range this.All() {
		if user.Uid == uid {
			return &user
		}
	}
	return &User{}
}

func Users() *UserCollection {
	return &UserCollection{}
}
