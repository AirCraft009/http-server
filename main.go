package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

/*
*

	func main() {
		httpServer := server.NewServer(8080)
		httpServer.Handle("GET", "/", handlers.Homehandler)
		httpServer.Listen()
		httpServer.AcceptConnections()
	}
*/
func main() {
	fmt.Println(listFiles("http-server"))
}

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
