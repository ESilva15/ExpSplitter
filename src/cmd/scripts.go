package cmd

import (
	exp "expenses/expenses"
	"log"

	"github.com/spf13/cobra"
)

func scripts(cmd *cobra.Command, args []string) {
	lua := exp.App.Lua

	if err := lua.DoFile("./lua-scripts/test.lua"); err != nil {
		log.Printf("Failed to run lua script: %+v", err)
	}
}

// shelf
var scriptsCmd = &cobra.Command{
	Use:   "debts",
	Short: "",
	Long:  ``,
	Args:  nil,
	Run:   scripts,
}

func init() {
	rootCmd.AddCommand(scriptsCmd)
}
