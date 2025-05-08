package main

import (
	"os"

	"bitmap/core"
)

func main() {
	app := core.Application{}
	err := app.Run()
	if err != nil {
		os.Exit(1)
	}
}
