package api

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/srgklmv/comfortel/internal/middleware"
)

type controller interface {
	userController
}

type userController interface {
	CreateUser(*gin.Context)
	GetUser(*gin.Context)
	GetUsers(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

func SetRoutes(engine *gin.Engine, conn *sql.DB, controller controller) {
	engine.Use(cors.Default())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	api := engine.Group("api", middleware.Transaction(conn))

	user := api.Group("/user")
	user.POST("", controller.CreateUser)
	user.GET("/:id", controller.GetUser)
	user.GET("", controller.GetUsers)
	user.PATCH("/:id", controller.UpdateUser)
	user.DELETE("/:id", controller.DeleteUser)
}
