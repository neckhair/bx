package mocks

import (
	"embed"
	"log"
)

//go:embed data
var dataDir embed.FS

func readFromFile(filename string) string {
	content, err := dataDir.ReadFile("data/" + filename)
	if err != nil {
		log.Fatalf("could not open test data file: %v", err)
	}
	return string(content)
}

func ContactList() string {
	return readFromFile("list_contacts.json")
}
