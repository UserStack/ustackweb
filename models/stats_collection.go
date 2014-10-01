package models

import (
	"github.com/UserStack/ustackweb/backend"
)

type StatsCollection struct {
}

func (this *StatsCollection) All() (stats map[string]int64, err *backend.Error) {
	connection, err := backend.Connection()
	if err != nil {
		return
	}
	backendStats, backendError := connection.Stats()
	backend.VerifyConnection(backendError)
	if backendError == nil {
		stats = backendStats
	}
	return
}

func Stats() *StatsCollection {
	return &StatsCollection{}
}
