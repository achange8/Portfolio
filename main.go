package main

import (
	"fmt"

	"github.com/achange8/Portfolio/router"
)

//CURD board, modify user PW, ID //done
//todos : google oauth login
//may be file upload, downlard
//배포, apache, aws service study
func main() {

	fmt.Println("Hello echo!")

	e := router.New()

	e.Logger.Fatal(e.Start(":8082"))
}
