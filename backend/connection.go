package backend

import (
	"fmt"
	"os"

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
			error = &Error{backends.Error{Message: anError.Error()}}
		}
		return connection, error
	} else {
		hostname := os.Getenv("USTACKD_1_PORT_7654_TCP_ADDR")
		if hostname == "" {
			hostname = "127.0.0.1"
		}
		connection, anError := client.Dial(hostname + ":7654")
		if anError != nil {
			error = &Error{backends.Error{Message: anError.Error()}}
		}
		return connection, error
	}
}

func Connection() (connection backends.Abstract, error *Error) {
	if sharedConnection == nil {
		connection, error = Reconnect()
	} else {
		connection = sharedConnection
	}
	return
}

func Reconnect() (connection backends.Abstract, error *Error) {
	connection, error = NewConnection()
	if error == nil {
		sharedConnection = connection
	}
	return
}
