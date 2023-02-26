package main

import (
	"fmt"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/router"
)

// CURD board, modify user PW, ID //done
// todos : google oauth login
// may be file upload, downlard
func main() {

	fmt.Println("Hello echo!")
	database.Connect()
	e := router.New()

	e.Logger.Fatal(e.Start(":8082"))
}
