package main

import (
	"http-server/server"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	httpServer := server.NewServer(8080)
	httpServer.Listen()
	httpServer.AcceptConnections()
}
