package main

import (
	"new-ott-project/Content-service/Controller"
	"new-ott-project/Content-service/Database"

	"github.com/gin-gonic/gin"
	_ "new-ott-project/Content-service/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)
// @title Example API
// @version 1.0
// @description This is a sample server with two GET APIs using Gin and Swagger.
// @host localhost:9080
// @BasePath /


func main() {

	Database.InitDB()

	r := gin.Default()
	// r.Use(Authorization.AuthMiddleware())

// content := router.Group("/content")
// {
//     content.GET("/all",Controller.Getallvideos)
// }
// routes for Contenttable

 r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
v1 := r.Group("/api/v1")
{
	v1.POST("/content",Controller.Storevideometadata)
	 v1.GET("/content",Controller.Getallvideos)
	v1.GET("/content/:id",Controller.Getvideobyid)
	v1.PATCH("/content/:id",Controller.Updatevideobyid)
	v1.DELETE("/content/:id",Controller.Deletevideobyid)
	// main.go or routes.go
v1.GET("/videos/:id", Controller.GetVideoWithCast)


	// routes for cast table
	v1.POST("/cast",Controller.Storecastmetadata)
	 v1.GET("/cast/:id",Controller.Getcastdatabyid)

	 // routes for genre table
	 v1.POST("/genre",Controller.Storegenre)
	 v1.GET("/genre",Controller.GetVideosByGenresHandler)
}


	


	r.Run(":9080")
}


