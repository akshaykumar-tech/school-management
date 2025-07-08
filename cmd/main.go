package main

import (
	"fmt"

	"github.com/akshaykumar-tech/school-management/api"
)

func main() {

	api := api.ApiRouts{}
	api.Startapp()

	fmt.Println("server is ", api)

}
