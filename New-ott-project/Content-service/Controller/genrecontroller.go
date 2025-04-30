package Controller

import (
	"net/http"
	"new-ott-project/Content-service/Models"
	"new-ott-project/Content-service/Service"

	"github.com/gin-gonic/gin"
)

func Storegenre(c *gin.Context) {
	var genres Models.Genre
	err := c.ShouldBindJSON(&genres)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
		return

	}
	Genres,err:=Service.Creategenres(genres)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":err.Error()})
		return

	}
	c.JSON(http.StatusOK,Genres)
}
