package Controller

import (
	"net/http"
	"new-ott-project/Content-service/Database"
	"new-ott-project/Content-service/Models"
	"new-ott-project/Content-service/Service"
	"strconv"

	"github.com/gin-gonic/gin"
)
func Storevideometadata(c *gin.Context){
	var Videos Models.Videos
	err:=c.ShouldBindJSON(&Videos)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	

	video,err:=Service.Createvideos(Videos)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,video)
}

func Getallvideos(c *gin.Context){
	videos,err:=Service.Getalldata()
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,videos)
}

func Getvideobyid(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	video,err:=Service.Getdatabyid(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,video)
}

func Updatevideobyid(c *gin.Context){
	var updateddata Models.Videos
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	err=c.ShouldBindJSON(&updateddata)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	
	video,err:=Service.Updatedatabyid(id,updateddata)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,video)
}

func Deletevideobyid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format: " + err.Error(),
		})
		return
	}

	video, err := Service.Deletebyid(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Video not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    video,
	})
}


// controllers/video_controller.go
func GetVideoWithCast(c *gin.Context) {
    videoID := c.Param("id")

    var video Models.Videos
    if err := Database.DB.Preload("Cast").First(&video, videoID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching video with cast"})
        return
    }

    c.JSON(http.StatusOK, video)
}
type GenreRequest struct {
	Genres []string `json:"genres" binding:"required"`
}

func GetVideosByGenresHandler(c *gin.Context) {
	var req GenreRequest

	// Step 1: Bind JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: genres list required"})
		return
	}

	// Step 2: Call the service layer
	videos, err := Service.GetVideosByGenres(req.Genres)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
		return
	}

	// Step 3: Return response
	c.JSON(http.StatusOK, gin.H{"videos": videos})
}