package main

import (
	"log"
	"os"

	"github.com/mikaelmicheline/goroutines-study/app"
)

func main() {
	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
