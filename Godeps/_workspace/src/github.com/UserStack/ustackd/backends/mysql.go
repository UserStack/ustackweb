package backends

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var PREPARE_MYSQL = []string{
	fmt.Sprintf(`CREATE TABLE IF NOT EXISTS Users (
		uid INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		state INTEGER DEFAULT %d,
		CONSTRAINT SingleKeys UNIQUE (name)
	) ENGINE=InnoDB;`, STATUS_ACTIVE),
	`CREATE TABLE IF NOT EXISTS Groups (
		gid INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		CONSTRAINT SingleKeys UNIQUE (name)
	) ENGINE=InnoDB;`,
	`CREATE TABLE IF NOT EXISTS UserGroups (
		uid INTEGER NOT NULL,
		gid INTEGER NOT NULL,
		CONSTRAINT SingleKeys UNIQUE (uid, gid),
		CONSTRAINT FOREIGN KEY (uid) REFERENCES Users(uid) ON DELETE CASCADE,
		CONSTRAINT FOREIGN KEY (gid) REFERENCES Groups(gid) ON DELETE CASCADE
	) ENGINE=InnoDB;`,
	fmt.Sprintf(`CREATE TABLE IF NOT EXISTS UserValues (
		uid INTEGER,
		%s VARCHAR(255) NOT NULL,
		value LONGTEXT NOT NULL,
		CONSTRAINT SingleKeys UNIQUE (uid, %s),
		CONSTRAINT FOREIGN KEY (uid) REFERENCES Users(uid) ON DELETE CASCADE
	) ENGINE=InnoDB;`, "`key`", "`key`"),
}

type MysqlBackend struct {
	SqlBackend
}

func NewMysqlBackend(url string) (MysqlBackend, error) {
	var backend MysqlBackend
	db, err := sql.Open("mysql", url)
	if err == nil {
		backend = MysqlBackend{SqlBackend{db: db}}
		return backend, backend.init(PREPARE_MYSQL)
	} else {
		return backend, err
	}
}

func (backend *MysqlBackend) init(prepare []string) error {
	var err error
	// set the default encoding, enable foreign keys, enable journal mode,
	// check the integrity, set timeout to 60 sec and enable the auto vacuum
	// and initialize all tables
	for _, stmt := range prepare {
		_, err = backend.db.Exec(stmt)
		if err != nil {
			panic(err)
		}
	}
	backend.createUserStmt, err = backend.db.Prepare(`INSERT INTO Users
		(name, password) VALUES (?, ?);`)
	if err != nil {
		panic(err)
	}
	backend.usersStmt, err = backend.db.Prepare(`SELECT name, uid, state FROM Users`)
	if err != nil {
		panic(err)
	}
	backend.deleteUserStmt, err = backend.db.Prepare(`DELETE FROM Users WHERE uid = ? OR name = ?;`)
	if err != nil {
		panic(err)
	}
	backend.loginUserStmt, err = backend.db.Prepare(fmt.Sprintf(
		"SELECT uid FROM Users WHERE name = ? AND password = ? AND state = %d;",
		STATUS_ACTIVE))
	if err != nil {
		panic(err)
	}
	backend.setUserStateStmt, err = backend.db.Prepare(`UPDATE IGNORE Users
		SET state = ?
		WHERE name = ? OR uid = ?;`)
	if err != nil {
		panic(err)
	}
	backend.uidForNameUidStmt, err = backend.db.Prepare(`SELECT uid FROM Users
		WHERE name = ? OR uid = ?;`)
	if err != nil {
		panic(err)
	}
	backend.setUserDataStmt, err = backend.db.Prepare(
		"INSERT INTO UserValues (uid, `key`, value) VALUES (?, ?, ?) " +
			"ON DUPLICATE KEY UPDATE value=VALUES(value);")
	if err != nil {
		panic(err)
	}
	backend.getUserDataStmt, err = backend.db.Prepare(
		"SELECT value FROM UserValues WHERE uid = ? AND `key` = ?;")
	if err != nil {
		panic(err)
	}
	backend.changeUserPasswordStmt, err = backend.db.Prepare(`UPDATE Users
		SET password = ?
		WHERE uid = ? AND password = ?;`)
	if err != nil {
		panic(err)
	}
	backend.changeUserNameStmt, err = backend.db.Prepare(`UPDATE Users
		SET name = ?
		WHERE uid = ? AND password = ?;`)
	if err != nil {
		panic(err)
	}
	backend.userGroupsStmt, err = backend.db.Prepare(`SELECT g.name, g.gid
		FROM Groups g
		JOIN UserGroups ug ON (ug.gid = g.gid)
		WHERE ug.uid = ?`)
	if err != nil {
		panic(err)
	}
	backend.createGroupStmt, err = backend.db.Prepare(`INSERT INTO Groups (name)
		VALUES (?);`)
	if err != nil {
		panic(err)
	}
	backend.groupsStmt, err = backend.db.Prepare(`SELECT name, gid FROM Groups;`)
	if err != nil {
		panic(err)
	}
	backend.deleteGroupStmt, err = backend.db.Prepare(`DELETE FROM Groups
		WHERE gid = ? OR name = ?;`)
	if err != nil {
		panic(err)
	}
	backend.gidForNameGidStmt, err = backend.db.Prepare(`SELECT gid FROM Groups
		WHERE gid = ? OR name = ?;`)
	if err != nil {
		panic(err)
	}
	backend.addUserToGroupStmt, err = backend.db.Prepare(
		`INSERT INTO UserGroups (uid, gid) VALUES (?, ?);`)
	if err != nil {
		panic(err)
	}
	backend.removeUserFromGroupStmt, err = backend.db.Prepare(`DELETE FROM UserGroups
		WHERE uid = ? AND gid = ?`)
	if err != nil {
		return err
	}
	backend.groupUsersStmt, err = backend.db.Prepare(`SELECT u.name, u.uid, u.state
		FROM Users u
		JOIN UserGroups ug ON (ug.uid = u.uid)
		WHERE ug.gid = ?`)
	if err != nil {
		panic(err)
	}
	backend.statsStmt, err = backend.db.Prepare(`SELECT 'Users', COUNT(*) FROM Users
												UNION
												SELECT 'Groups', COUNT(*) FROM Groups`)
	if err != nil {
		panic(err)
	}
	return nil
}
