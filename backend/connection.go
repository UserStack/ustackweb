package backend

import (
	"fmt"
	"github.com/UserStack/ustackd/backends"
	"github.com/UserStack/ustackd/client"
)

const (
	Memory = iota
	Remote
)

var Type int
var sharedConnection backends.Abstract

func VerifyConnection(error *backends.Error) {
	if error != nil && error.Message == "EOF" {
		fmt.Println("Connection lost", error)
		sharedConnection = nil
	}
}

func NewConnection() (connection backends.Abstract, error *Error) {
	if Type == Memory {
		aConnection, anError := backends.NewSqliteBackend(":memory:")
		if anError == nil {
			connection = &aConnection
		} else {
			error = &Error{Raw: fmt.Sprintf("%s", anError)}
		}
		return connection, error
	} else {
		connection, anError := client.Dial("127.0.0.1:7654")
		if anError != nil {
			error = &Error{Raw: fmt.Sprintf("%s", anError)}
		}
		return connection, error
	}
}

func Connection() (connection backends.Abstract, error *Error) {
	if sharedConnection == nil {
		connection, error = NewConnection()
		if error == nil {
			sharedConnection = connection
		}
	} else {
		connection = sharedConnection
	}
	return
}
