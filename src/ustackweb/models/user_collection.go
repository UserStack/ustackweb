package models

import (
	"github.com/UserStack/ustackd/backends"
)

type UserCollection struct {
}

func (this *UserCollection) All() []User {
	return []User{
		User{&backends.User{Uid: 1, Name: "foo"}},
		User{&backends.User{Uid: 2, Name: "admin"}},
		User{&backends.User{Uid: 3, Name: "abc"}},
		User{&backends.User{Uid: 4, Name: "def"}},
		User{&backends.User{Uid: 5, Name: "hij"}},
		User{&backends.User{Uid: 6, Name: "glk"}},
		User{&backends.User{Uid: 7, Name: "uvw"}},
		User{&backends.User{Uid: 8, Name: "xyz"}}}
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
