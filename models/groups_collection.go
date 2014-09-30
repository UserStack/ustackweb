package models

import (
	"github.com/UserStack/ustackweb/backend"
)

type GroupsCollection struct {
}

func (this *GroupsCollection) All() (groups []Group, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendGroups, backendError := connection.Groups()
	if backendError == nil {
		groups = make([]Group, len(backendGroups))
		for idx, backendGroup := range backendGroups {
			groups[idx] = Group{backendGroup}
		}
	}
	return
}

func (this *GroupsCollection) Create(name string) (created bool, id int64, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	id, backendError := connection.CreateGroup(name)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		created = id > 0
	}
	return
}

func Groups() *GroupsCollection {
	return &GroupsCollection{}
}
