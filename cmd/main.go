package main

import (
	"fmt"

	"github.com/akshaykumar-tech/school-management/api"
	"github.com/akshaykumar-tech/school-management/controller"
)

func main() {

	api := api.ApiRouts{}
	api.Startapp(controller.Server{})

	fmt.Println("server is ", api)
}
