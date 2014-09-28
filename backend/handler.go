package backend

import (
	"github.com/UserStack/ustackd/backends"
	"github.com/UserStack/ustackd/client"
)

const (
	Memory = iota
	Remote
)

var Type int
var backend backends.Abstract

func build() backends.Abstract {
	if Type == Memory {
		backend, _ := backends.NewSqliteBackend(":memory:")
		return &backend
	} else {
		backend, _ := client.Dial("127.0.0.1:7654")
		return backend
	}
}

func Current() backends.Abstract {
	if backend == nil {
		Reconnect()
	}
	return backend
}

func Reconnect() {
	backend = build()
}
