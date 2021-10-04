package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("No input provided\n")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	pass, hasPassword := url.User.Password()

	if !hasPassword {
		pass = ""
	}

	port := url.Port()

	if len(port) == 0 {
		port = "5432"
	}

	fmt.Printf("PGUSER=%s\n", url.User.Username())
	fmt.Printf("PGPASSWORD=%s\n", pass)
	fmt.Printf("PGHOST=%s\n", url.Hostname())
	fmt.Printf("PGPORT=%s\n", port)
	fmt.Printf("PGDATABASE=%s\n", url.Path[1:])
}
