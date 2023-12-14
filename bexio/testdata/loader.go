package testdata

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func readFromFile(filename string) string {
	path := filepath.Join("testdata", filename)
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open test data file: %v", err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("could not read test data: %v", err)
	}
	return string(data)
}

func ContactList() string {
	return readFromFile("list_contacts.json")
}
