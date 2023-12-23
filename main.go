package main

import (
	"afm/internal/app"

	"os"
)

func main() {
	a := app.New()

	if err := a.Run(os.Args); err != nil {
		panic(err)
	}
}
