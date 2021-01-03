package main

import (
	"fmt"

	"github.com/jonioliveira/eq-json/pkg/app"
)

func main() {
	if err := app.Start(); err != nil {
		fmt.Printf("error while running app, %s", err)
	}
}
