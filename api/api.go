package api

import (
	"github.com/akshaykumar-tech/school-management/controller"
	"github.com/akshaykumar-tech/school-management/store"
)

type ApiRouts struct {
	Server controller.ServerOperation
}

func (api *ApiRouts) Startapp(server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgres{})

}
