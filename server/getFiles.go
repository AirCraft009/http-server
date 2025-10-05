package server

import (
	"log"
	"os"
	"path/filepath"
	"slices"
)

func listFiles(path string) []string {

	//Reading contents of the directory
	files, err := os.ReadDir(path)

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	paths := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			subDirectoryPath := filepath.Join(path, file.Name())      // path of the subdirectory
			paths = slices.Concat(paths, listFiles(subDirectoryPath)) // calling listFiles() again for subdirectory
		} else {
			paths = append(paths, filepath.Join(path, file.Name()))
		}
	}
	return paths
}
