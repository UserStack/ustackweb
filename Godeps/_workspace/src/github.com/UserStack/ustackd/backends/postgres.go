package backends

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var PREPARE_POSTGRES = []string{
	fmt.Sprintf(`CREATE TABLE IF NOT EXISTS Users (
		uid INTEGER PRIMARY KEY DEFAULT nextval('UsersSeq'),
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		state INTEGER DEFAULT %d,
		CONSTRAINT UniqueUserNames UNIQUE (name)
	);`, STATUS_ACTIVE),
	`CREATE TABLE IF NOT EXISTS Groups (
		gid INTEGER PRIMARY KEY DEFAULT nextval('GroupsSeq'),
		name TEXT NOT NULL,
		CONSTRAINT UniqueGroupNames UNIQUE (name)
	);`,
	`CREATE TABLE IF NOT EXISTS UserGroups (
		uid INTEGER NOT NULL REFERENCES Users(uid) ON DELETE CASCADE,
		gid INTEGER NOT NULL REFERENCES Groups(gid) ON DELETE CASCADE,
		CONSTRAINT UniqueUidGidPairs UNIQUE (uid, gid)
	);`,
	`CREATE TABLE IF NOT EXISTS UserValues (
		uid INTEGER REFERENCES Users(uid) ON DELETE CASCADE,
		key TEXT NOT NULL,
		value TEXT NOT NULL,
		CONSTRAINT UniqueUidKeyPairs UNIQUE (uid, key)
	);`,
}

type PostgresBackend struct {
	SqlBackend
}

func NewPostgresBackend(url string) (PostgresBackend, error) {
	var backend PostgresBackend
	db, err := sql.Open("postgres", url)
	if err == nil {
		backend = PostgresBackend{SqlBackend{db: db}}
		return backend, backend.init(PREPARE_POSTGRES)
	} else {
		return backend, err
	}
}

func (backend *PostgresBackend) init(sqls []string) (err error) {
	// create sequences and ignore if this is failing
	backend.db.Exec(`CREATE SEQUENCE UsersSeq;`)
	backend.db.Exec(`CREATE SEQUENCE GroupsSeq;`)
	backend.db.Exec(`CREATE OR REPLACE FUNCTION convert_to_integer(v_input text)
	RETURNS INTEGER AS $$
	DECLARE v_int_value INTEGER DEFAULT NULL;
	BEGIN
	    BEGIN
	        v_int_value := v_input::INTEGER;
	    EXCEPTION WHEN OTHERS THEN
	        RAISE NOTICE 'Invalid integer value: "%".  Returning NULL.', v_input;
	        RETURN NULL;
	    END;
	RETURN v_int_value;
	END;
	$$ LANGUAGE plpgsql;`)
	err = backend.SqlBackend.init(sqls)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.createUserStmt, err = backend.db.Prepare(
		`INSERT INTO Users (name, password) VALUES ($1, $2) RETURNING uid;`)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.createGroupStmt, err = backend.db.Prepare(
		`INSERT INTO Groups (name) VALUES ($1) RETURNING gid;`)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.deleteUserStmt, err = backend.db.Prepare(
		`DELETE FROM Users
		WHERE uid = convert_to_integer($1) OR name = $2;`)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.setUserStateStmt, err = backend.db.Prepare(`UPDATE Users
		SET state = $1
		WHERE name = $2 OR uid = convert_to_integer($3);`)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.uidForNameUidStmt, err = backend.db.Prepare(`SELECT uid FROM Users
		WHERE name = $1 OR uid = convert_to_integer($2);`)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.deleteGroupStmt, err = backend.db.Prepare(`DELETE FROM Groups
		WHERE gid = convert_to_integer($1) OR name = $2;`)
	if err != nil {
		panic(err)
	}
	backend.SqlBackend.gidForNameGidStmt, err = backend.db.Prepare(`SELECT gid FROM Groups
		WHERE gid = convert_to_integer($1) OR name = $2;`)
	if err != nil {
		panic(err)
	}
	return
}

func (backend *PostgresBackend) CreateUser(name string, password string) (int64, *Error) {
	if name == "" || password == "" {
		return 0, &Error{"EINVAL", "User name and password can't be blank"}
	}
	var uid int64
	err := backend.SqlBackend.createUserStmt.QueryRow(name, password).Scan(&uid)
	if err != nil {
		return 0, &Error{"EFAULT", err.Error()}
	}
	return uid, nil
}

func (backend *PostgresBackend) CreateGroup(name string) (int64, *Error) {
	if name == "" {
		return 0, &Error{"EINVAL", "Invalid group name"}
	}
	var gid int64
	err := backend.SqlBackend.createGroupStmt.QueryRow(name).Scan(&gid)
	if err != nil {
		return 0, &Error{"EEXIST", err.Error()}
	}
	return gid, nil
}
