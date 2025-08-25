package cmd

import (
	"log"

	exp "expenses/expenses"
	"expenses/luadec"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

func debts(cmd *cobra.Command, args []string) {
	log.Println("Added the debts command!")

	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("GetAllExpenses", L.NewFunction(exp.Serv.LuaGetAllExpenses))
	L.SetGlobal("AddDecimal", L.NewFunction(luadec.AddDecimal))

	if err := L.DoFile("./lua-scripts/test.lua"); err != nil {
		log.Printf("Failed to run lua script: %+v", err)
	}
}

// shelf
var debtsCmd = &cobra.Command{
	Use:   "debts",
	Short: "",
	Long:  ``,
	Args:  nil,
	Run:   debts,
}

func init() {
	rootCmd.AddCommand(debtsCmd)
}
