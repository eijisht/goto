package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	HOME, _ := os.UserHomeDir()
	f, err := os.ReadFile(HOME + "/goto/index.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]string

	if err := yaml.Unmarshal(f, &m); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		for key := range m {
			fmt.Printf("%s ", key)
		}

		fmt.Printf("\n")
		os.Exit(0)
	}

	project := os.Args[1]
	path, exists := m[project]

	if project == "help" {
		os.Exit(0)
	}

	if project == "add" {
		_, exists = m[os.Args[2]]

		if len(os.Args) != 4 {
			log.Fatalf("Not enough arguments. Usage: goto add <alias> <absolute path to directory>")
			os.Exit(1)
		}
		if exists {
			fmt.Printf("Project already exists\n")
			os.Exit(0)
		}

		alias := os.Args[2]
		path := os.Args[3]
		m[alias] = path

		yamlfile, err := yaml.Marshal(m)
		if err != nil {
			log.Fatalf("couldnt convert to yaml")
		}

		os.WriteFile("index.yaml", yamlfile, 0666)
		os.Exit(0)
	}

	if exists {
		fmt.Printf("%s\n", path)
		os.Exit(0)
	}

	os.Exit(1)
}
