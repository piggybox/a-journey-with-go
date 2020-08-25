package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "a journey with go", "help")
	flag.StringVar(&name, "n", "a journey with go", "help")
	flag.Parse()

	log.Printf("name: %s", name)
}
