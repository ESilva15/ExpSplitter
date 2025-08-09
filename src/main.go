package main

import (
	"expenses/cmd"
	"expenses/expenses"
	"log"
)

func main() {
	if err := expenses.StartApp(); err != nil {
		log.Fatal("failed to start app:", err.Error())
	}

	cmd.Execute()
}
