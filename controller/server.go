package controller

import "github.com/akshaykumar-tech/school-management/store"

type Server struct {
	PostgressDb store.Postgres
}

func (akshay *Server) NewServer() {
	akshay.PostgressDb.NewStore()
}
