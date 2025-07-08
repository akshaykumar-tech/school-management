package api

import (
	"github.com/akshaykumar-tech/school-management/controller"
	"github.com/akshaykumar-tech/school-management/db"
	"github.com/akshaykumar-tech/school-management/model"
	"github.com/gin-gonic/gin"
)

type ApiRouts struct{}

func (api *ApiRouts) Startapp() {
	postgres := db.NewPostgres()
	postgres.DB.AutoMigrate(&model.User{})

	r := gin.Default()
	userController := controller.NewUserController(postgres)
	userController.RegisterRoutes(r)
	r.Run(":8080")
}
