package repo

import lua "github.com/yuin/gopher-lua"

type Migrator interface {
	Goto(id uint) error
	RunMigrations(lua *lua.LState) error
	Close()
}
