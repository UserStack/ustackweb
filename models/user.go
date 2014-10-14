package models

import (
	"strconv"
	"time"

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

func (this *User) DataAsTime(key string) (data time.Time, err *backend.Error) {
	value, err := this.Data(key)
	if err != nil {
		return
	}
	i, backendError := strconv.ParseInt(value, 10, 0)
	if backendError != nil {
		return
	}
	data = time.Unix(i, 0)
	return
}

func (this *User) AllData() (allData map[string]interface{}, err *backend.Error) {
	keys, err := this.DataKeys()
	if err != nil {
		return
	}
	allData = make(map[string]interface{})
	for _, key := range keys {
		if key == "currentlogin" || key == "lastlogin" {
			value, err := this.DataAsTime(key)
			if err == nil {
				allData[key] = value
			}
		} else {
			value, err := this.Data(key)
			if err == nil {
				allData[key] = value
			}
		}
	}
	return
}
