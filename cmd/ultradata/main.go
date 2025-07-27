// cmd/ultradata/main.go
package main

import (
	"flag"
	"log"
	"os"

	"ultradata/internal/ultradata"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := ultradata.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
