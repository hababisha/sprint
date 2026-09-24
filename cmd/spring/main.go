package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hababisha/sprint/internal/initializr"
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
	config := initializr.Config{
		Name:        name,
		BootVersion: "4.1.1",
		JavaVersion: "21",
		Build:       "maven",
		Packaging:   "jar",
		Dependencies: []string{
			"web",
			"data-jpa",
			"postgresql",
		},
	}

	client := initializr.NewClient()
	response, err := client.Generate(config)
	if err != nil {
		fmt.Println("error creating client: ", err)
		return
	}

	zipPath := name + ".zip"

	err = initializr.Save(response, zipPath)
	if err != nil {
		fmt.Println("Error saving the zip: ", err)
		return
	}
	fmt.Println("project downloaded successfully")

	// Todo -modify this to use go's archive to make it portable

	cmd := exec.Command("unzip", zipPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error extracting project:", err)
		return
	}

	//delete the zip
	err = os.Remove(zipPath)
	if err != nil {
		fmt.Println("Error removing ZIP:", err)
		return
	}
}
