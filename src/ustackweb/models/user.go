package models

import (
	"github.com/UserStack/ustackd/backends"
)

type User struct {
	*backends.User
}

func (this *User) Name() string {
	return this.Email
}
