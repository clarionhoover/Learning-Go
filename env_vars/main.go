package main

import (
	"fmt"
	"os"
)

var (
	user string
)

func main() {
	user = os.Getenv("USER")
	if user == "root" {
		fmt.Println("You're running this with sudo or root")

	} else {
		fmt.Println("This program needs to be run as root")
		os.Exit(1)
	}
}
