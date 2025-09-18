// main will start the expenses app
package main

import (
	"log"
	"os"

	"github.com/ESilva15/expenses/cmd"
	"github.com/ESilva15/expenses/expenses"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := expenses.StartApp(); err != nil {
		log.Fatal("failed to start app:", err.Error())
	}

	cmd.Execute()
}
