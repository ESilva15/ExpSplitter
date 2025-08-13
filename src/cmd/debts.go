package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func debts(cmd *cobra.Command, args []string) {
	log.Println("Added the debts command!")	
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
