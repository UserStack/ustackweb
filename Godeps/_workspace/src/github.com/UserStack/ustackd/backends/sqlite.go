package backends

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var PREPARE_SQLITE = []string{
	`PRAGMA encoding = "UTF-8";`,
	"PRAGMA foreign_keys = ON;",
	"PRAGMA journal_mode;",
	"PRAGMA integrity_check;",
	"PRAGMA busy_timeout = 60000;",
	"PRAGMA auto_vacuum = INCREMENTAL;",
	fmt.Sprintf(`CREATE TABLE IF NOT EXISTS Users (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		state INTEGER DEFAULT %d,
		CONSTRAINT UniqueUserNames UNIQUE (name) ON CONFLICT ROLLBACK
	);`, STATUS_ACTIVE),
	`CREATE TABLE IF NOT EXISTS Groups (
		gid INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		CONSTRAINT UniqueGroupNames UNIQUE (name) ON CONFLICT ROLLBACK
	);`,
	`CREATE TABLE IF NOT EXISTS UserGroups (
		uid INTEGER NOT NULL REFERENCES Users(uid) ON DELETE CASCADE,
		gid INTEGER NOT NULL REFERENCES Groups(gid) ON DELETE CASCADE,
		CONSTRAINT UniqueUidGidPairs UNIQUE (uid, gid) ON CONFLICT IGNORE
	);`,
	`CREATE TABLE IF NOT EXISTS UserValues (
		uid INTEGER REFERENCES Users(uid) ON DELETE CASCADE,
		key TEXT NOT NULL,
		value BLOB NOT NULL,
		CONSTRAINT UniqueUidKeyPairs UNIQUE (uid, key) ON CONFLICT REPLACE
	);`,
}

type SqliteBackend struct {
	SqlBackend
}

func NewSqliteBackend(url string) (SqliteBackend, error) {
	var backend SqliteBackend
	db, err := sql.Open("sqlite3", url)
	if err == nil {
		backend = SqliteBackend{SqlBackend{db: db}}
		return backend, backend.init(PREPARE_SQLITE)
	} else {
		return backend, err
	}
}

func (backend *SqliteBackend) CreateUser(name string, password string) (uid int64, err *Error) {
	uid, err = backend.SqlBackend.CreateUser(name, password)
	if err != nil {
		// after error occured, the statement is broken, we need to recreate it
		backend.createUserStmt.Close()
		backend.createUserStmt, _ = backend.db.Prepare(`INSERT INTO Users
			(name, password) VALUES (?, ?);`)
	}
	return
}

func (backend *SqliteBackend) CreateGroup(name string) (gid int64, err *Error) {
	gid, err = backend.SqlBackend.CreateGroup(name)
	if err != nil {
		backend.createGroupStmt.Close()
		backend.createGroupStmt, _ = backend.db.Prepare(`INSERT INTO Groups
			(name) VALUES (?);`)
	}
	return
}
