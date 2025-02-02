package main

import (
	"fmt"
	"web-fiber/internal/app"
)

func main() {
	if err := app.Startup(); err != nil {
		panic(fmt.Sprintf("Application startup error: %s", err))
	}
}
