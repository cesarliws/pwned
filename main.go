package main

import (
	"fmt"
	"os"

	hibp "github.com/mattevans/pwned-passwords"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		v := args[0]

		client := hibp.NewClient()

		compromised, err := client.Pwned.Compromised(v)

		if err != nil {
			fmt.Printf("Unexpected error running client.Pwned.Compromised(): %s", err)
		}

		if compromised {
			fmt.Printf("Compromised hash for: '%v'", v)
		}
	}

}
