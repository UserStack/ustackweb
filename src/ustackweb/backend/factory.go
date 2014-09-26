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

func Factory() backends.Abstract {
	if Type == Memory {
		backend, _ := backends.NewSqliteBackend(":memory:")
		return &backend
	} else {
		backend, _ := client.Dial("127.0.0.1:7654")
		return backend
	}
}
