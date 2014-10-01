package models

import (
	"github.com/UserStack/ustackd/backends"
	"github.com/UserStack/ustackweb/backend"
	"strings"
)

type GroupCollection struct {
}

func (this *GroupCollection) collect(backendGroups []backends.Group) (groups []Group) {
	groups = make([]Group, len(backendGroups))
	for idx, backendGroup := range backendGroups {
		groups[idx] = Group{backendGroup}
	}
	return
}

func (this *GroupCollection) All() (groups []Group, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendGroups, backendError := connection.Groups()
	backend.VerifyConnection(backendError)
	if backendError == nil {
		groups = this.collect(backendGroups)
	}
	return
}

func (this *GroupCollection) allWithPrefix(prefix string, with bool) (groups []Group, err *backend.Error) {
	allGroups, err := this.All()
	if err != nil {
		return
	}
	groups = make([]Group, 0)
	for _, group := range allGroups {
		hasPrefix := strings.HasPrefix(group.Name, prefix)
		if (hasPrefix && with) || (!hasPrefix && !with) {
			groups = append(groups, group)
		}
	}
	return
}

func (this *GroupCollection) AllWithPrefix(prefix string) (groups []Group, err *backend.Error) {
	return this.allWithPrefix(prefix, true)
}

func (this *GroupCollection) AllWithoutPrefix(prefix string) (groups []Group, err *backend.Error) {
	return this.allWithPrefix(prefix, false)
}

func (this *GroupCollection) AllByUser(name_or_uid string) (groups []Group, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendGroups, backendError := connection.UserGroups(name_or_uid)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		groups = this.collect(backendGroups)
	}
	return
}

func (this *GroupCollection) Create(name string) (created bool, id int64, err *backend.Error) {
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

func (this *GroupCollection) Destroy(name_or_uid string) (deleted bool, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendError := connection.DeleteGroup(name_or_uid)
	backend.VerifyConnection(backendError)
	if backendError == nil {
		deleted = err == nil
	}
	return
}

func Groups() *GroupCollection {
	return &GroupCollection{}
}
