package Controller

import (
	"net/http"
	"new-ott-project/Content-service/Models"
	"new-ott-project/Content-service/Service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Storecastmetadata(c *gin.Context) {
	var Videos Models.Cast
	err := c.ShouldBindJSON(&Videos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	video, err := Service.Createcast(Videos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, video)
}

func Getcastdatabyid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	video, err := Service.Getdatabyid(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, video)
}