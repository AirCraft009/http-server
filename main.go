package main

import (
	"http-server/handlers"
	"http-server/server"
)

func main() {
	httpServer := server.NewServer(8080, true, true)
	httpServer.AddFileSystem("frontend")
	httpServer.Handle("GET", "/", handlers.Homehandler)
	httpServer.ListenAndServe()
}
