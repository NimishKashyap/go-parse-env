package main

import (
	"log"

	"github.com/NimishKashyap/go-parse/parse_env"
)

func main() {
	log.Println("This is a test")

	parse_env.Parse(".environments")
}
