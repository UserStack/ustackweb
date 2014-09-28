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

func build() (backend backends.Abstract) {
	if Type == Memory {
		b, err := backends.NewSqliteBackend(":memory:")
		if err == nil {
			backend = &b
		}
	} else {
		backend, _ = client.Dial("127.0.0.1:7654")
	}
	return
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
