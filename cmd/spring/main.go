package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: spring init --name <name>")
		return
	}

	command := os.Args[1]
	if command != "init" {
		fmt.Printf("unknown command: %s\n", command)
		return
	}

	fmt.Println("Initializing Spring Boot project...")
	if len(os.Args) < 4 || os.Args[2] != "--name" {
		fmt.Println("usage: spring init --name <name>")
		return
	}

	name := os.Args[3]

	fmt.Printf("Creating Spring Boot project: %s\n", name)

}
