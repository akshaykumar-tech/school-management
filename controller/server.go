package controller

import "github.com/akshaykumar-tech/school-management/store"

type Server struct {
	PostgressDb store.StoreOperation
}

func (akshay *Server) NewServer(pgstore store.Postgres) {
	akshay.PostgressDb = &pgstore
	akshay.PostgressDb.NewStore()
}
