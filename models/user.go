package models

import (
	"github.com/UserStack/ustackd/backends"
	"github.com/UserStack/ustackweb/backend"
)

type User struct {
	backends.User
}

func (this *User) DataKeys() (keys []string, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	keys, backendError := connection.GetUserDataKeys(this.Name)
	backend.VerifyConnection(backendError)
	return
}

func (this *User) Data(key string) (data string, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	data, backendError := connection.GetUserData(this.Name, key)
	backend.VerifyConnection(backendError)
	return
}
