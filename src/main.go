// main will start the expenses app
package main

import (
	"log"

	"github.com/ESilva15/expenses/cmd"
	"github.com/ESilva15/expenses/expenses"
)

func main() {
	if err := expenses.StartApp(); err != nil {
		log.Fatal("failed to start app:", err.Error())
	}

	cmd.Execute()
}
