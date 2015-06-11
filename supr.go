package main

import (
	"github.com/JakeClarke/supr/command"
	"log"
)

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
