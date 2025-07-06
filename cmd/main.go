package main

import (
	"fmt"

	"github.com/akshaykumar-tech/school-management/controller"
)

func main() {
	server := controller.Server{}

	server.NewServer()

	fmt.Println("server is ", server)
}
