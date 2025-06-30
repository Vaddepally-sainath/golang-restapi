package main

import (
	"new-ott-project/User-service/Database"
	"new-ott-project/User-service/Usercontroller"

	"github.com/gin-gonic/gin"
)

func main() {

	Database.InitDB()

	router := gin.Default()
	router.POST("/user",Usercontroller.Storeuserdata)
	// router.GET("/content",Controller.Getallvideos)
	router.GET("/user/:id",Usercontroller.Getuserbyid)
	// router.PUT("/content/:id",Controller.Updatevideobyid)
	// router.DELETE("/content/:id",Controller.Deletevideobyid)

	router.POST("/login", Usercontroller.Login)

	router.Run(":7000")
}