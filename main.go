package main

import (
	"log"
	"os"

	"github.com/romulodm/go-main-cli/app"
)

func main() {
	aplication := app.Generate()
	err := aplication.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
