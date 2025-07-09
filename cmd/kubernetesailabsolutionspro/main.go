// cmd/kubernetesailabsolutionspro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"kubernetesailabsolutionspro/internal/kubernetesailabsolutionspro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := kubernetesailabsolutionspro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
