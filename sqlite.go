package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/glauth/glauth/v2/pkg/plugins"
	"github.com/glauth/glauth/v2/pkg/handler"
)

type SqliteBackend struct {
}

func NewSQLiteHandler(opts ...handler.Option) handler.Handler {
	backend := SqliteBackend{}
	return plugins.NewDatabaseHandler(backend, opts...)
}

func (b SqliteBackend) GetDriverName() string {
	return "sqlite3"
}

func (b SqliteBackend) GetPrepareSymbol() string {
	return "?"
}

// Create db/schema if necessary
func (b SqliteBackend) CreateSchema(db *sql.DB) {
	statement, _ := db.Prepare(`
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	uidnumber INTEGER NOT NULL,
	primarygroup INTEGER NOT NULL,
	othergroups TEXT DEFAULT '',
	givenname TEXT DEFAULT '',
	sn TEXT DEFAULT '',
	mail TEXT DEFAULT '',
	loginshell TYEXT DEFAULT '',
	homedirectory TEXT DEFAULT '',
	disabled SMALLINT  DEFAULT 0,
	passsha256 TEXT DEFAULT '',
	passbcrypt TEXT DEFAULT '',
	otpsecret TEXT DEFAULT '',
	yubikey TEXT DEFAULT '',
	sshkeys TEXT DEFAULT '',
	custattr TEXT DEFAULT '{}')
`)
	statement.Exec()
	statement, _ = db.Prepare("CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name on users(name)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS ldapgroups (id INTEGER PRIMARY KEY, name TEXT NOT NULL, gidnumber INTEGER NOT NULL)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE UNIQUE INDEX IF NOT EXISTS idx_group_name on ldapgroups(name)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS includegroups (id INTEGER PRIMARY KEY, parentgroupid INTEGER NOT NULL, includegroupid INTEGER NOT NULL)")
	statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS capabilities (id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, action TEXT NOT NULL, object TEXT NOT NULL)")
	statement.Exec()
}

// Migrate schema if necessary
func (b SqliteBackend) MigrateSchema(db *sql.DB, checker func(*sql.DB, string) bool) {
	if !checker(db, "sshkeys") {
		statement, _ := db.Prepare("ALTER TABLE users ADD COLUMN sshkeys TEXT DEFAULT ''")
		statement.Exec()
	}

	if TableExists(db, "groups") {
		// Drop the table created during schema creation
		statement, _ := db.Prepare("DROP TABLE ldapgroups")
		statement.Exec()

		statement, _ = db.Prepare("ALTER TABLE groups RENAME TO ldapgroups")
		statement.Exec()
	}
}

// Indicates whether the table exists or not
func TableExists(db *sql.DB, tableName string) bool {
	var found string
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(id) FROM %s", tableName)).Scan(
		&found)
	if err != nil {
		return false
	}
	return true
}

func main() {}
