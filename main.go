package main

import (
	"log"

	"github.com/JAANVI999/JWT_Auth/controllers"
	"github.com/JAANVI999/JWT_Auth/database"
	"github.com/JAANVI999/JWT_Auth/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")

	connectionString := "root:Anti@19icej@tcp(127.0.0.1:3306)/mydb?parseTime=true"
	database.Connect(connectionString)

	//log.Println("Running database migration...")
	database.Migrate()

	//log.Println("Application setup completed.")
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.Register)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
