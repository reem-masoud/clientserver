package main

import (
	"fmt"

	"github.com/clientserver/server"
)

//"strconv"

func main() {
	//routes.init()

	fmt.Println("Running...")
	//	e := routes.Router()

	//e.Logger.Fatal(e.Start(":8080"))
	server.StartServer()
	//logger.Fatal(server.StartServer())
}
