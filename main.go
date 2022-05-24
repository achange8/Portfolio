package main

import (
	"fmt"

	"github.com/achange8/Portfolio/router"
)

func main() {

	fmt.Println("Hello echo!")

	e := router.New()

	e.Logger.Fatal(e.Start(":8082"))
}
