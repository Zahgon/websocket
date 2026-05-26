package main

import (
	"log"
)

func main() {
	log.SetFlags(0)

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

// run initializes the chatServer and then
// starts a http.Server for the passed in address.
func run() error { _ = "STUB: not implemented"; return nil }
