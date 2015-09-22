package main

import (
	"fmt"
	"log"
	"path/filepath"
)

// Convert a `/` sepearated path to a valid Windows path using os.PathSeparator
func convertToWindowsPath(path string) string {
	return filepath.FromSlash(path)
}

// Catch-all function for error checking
func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
