package cmd

import (
	"expenses/expenses"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func verifyArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("requires at least two arguments: [source file] and dest")
	}

	return nil
}

func migrate(cmd *cobra.Command, args []string) {
	migID, err := strconv.ParseUint(args[0], 10, 0)
	if err != nil {
		log.Fatalf("%s is an invalid migration ID", args[1])
	}

	err = expenses.Serv.MigGoto(uint(migID))
	if err != nil {
		log.Fatalf("Failed to apply migration: %v", err)
	}
}

// shelf
var migCmd = &cobra.Command{
	Use:   "migrate",
	Short: "",
	Long:  ``,
	Args:  verifyArgs,
	Run:   migrate,
}

func init() {
	rootCmd.AddCommand(migCmd)
}
