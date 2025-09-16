package repo

import lua "github.com/yuin/gopher-lua"

// Migrator defines the interface for the migration procedures of different
// data storage.
type Migrator interface {
	Goto(id uint) error
	RunMigrations(lua *lua.LState) error
	Close()
}
