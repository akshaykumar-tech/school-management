package main

import (
	"fmt"

	"github.com/akshaykumar-tech/school-management/controller"
	"github.com/akshaykumar-tech/school-management/store"
)

func main() {
	server := controller.Server{}

	server.NewServer(store.Postgres{})

	fmt.Println("server is ", server)
}
