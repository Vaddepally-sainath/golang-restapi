package main

import (
	"new-ott-project/Content-service/Authorization"
	"new-ott-project/Content-service/Controller"
	"new-ott-project/Content-service/Database"

	"github.com/gin-gonic/gin"
)

func main() {

	Database.InitDB()

	router := gin.Default()
	router.Use(Authorization.AuthMiddleware())

// content := router.Group("/content")
// {
//     content.GET("/all",Controller.Getallvideos)
// }
// routes for Contenttable
	router.POST("/content",Controller.Storevideometadata)
	 router.GET("/content",Controller.Getallvideos)
	router.GET("/content/:id",Controller.Getvideobyid)
	router.PATCH("/content/:id",Controller.Updatevideobyid)
	router.DELETE("/content/:id",Controller.Deletevideobyid)
	// main.go or routes.go
router.GET("/videos/:id", Controller.GetVideoWithCast)


	// routes for cast table
	router.POST("/cast",Controller.Storecastmetadata)
	 router.GET("/cast/:id",Controller.Getcastdatabyid)

	 // routes for genre table
	 router.POST("/genre",Controller.Storegenre)
	 router.GET("/genre",Controller.GetVideosByGenresHandler)



	router.Run(":9080")
}


