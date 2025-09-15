// Package luadec holds shopspring/decimal binds to be used on lua scripts
package luadec

import (
	"github.com/shopspring/decimal"
	lua "github.com/yuin/gopher-lua"
)

// AddDecimal adds to decimal numbers passed as strings
func AddDecimal(L *lua.LState) int {
	decA, _ := decimal.NewFromString(L.CheckString(1))
	decB, _ := decimal.NewFromString(L.CheckString(2))

	res := decA.Add(decB)

	L.Push(lua.LString(res.StringFixed(2)))
	return 1
}
