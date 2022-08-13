package server

import (
	"github.com/gin-gonic/gin"

	ping "manyu/collect/pkg/ping/v1"
	user "manyu/collect/pkg/user/v1"
	form "manyu/collect/pkg/form/v1"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())

	pingGroup := router.Group("ping")
	pingGroup.GET("", ping.NewModule().GetController().Ping)

	v1 := router.Group("api/v1")
	{
		userGroup := v1.Group("user")
		{
			userController := user.NewUserModule().GetController()
			userGroup.GET("/user", userController.Ping)
		}

		formGroup := v1.Group("form")
		{
			formController := form.NewFormModule().GetController()
			formGroup.GET("/form", formController.Ping)
		}

	}

	return router
}