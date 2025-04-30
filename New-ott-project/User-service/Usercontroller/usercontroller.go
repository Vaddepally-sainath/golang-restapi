package Usercontroller

import (
	"net/http"
	"new-ott-project/User-service/Usermodel"
	"new-ott-project/User-service/Userservice"

	"strconv"

	"github.com/gin-gonic/gin"
)

func Storeuserdata(c *gin.Context){
	var users Usermodel.Users
	err:=c.ShouldBindJSON(&users)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	user,err:=Userservice.Createuser(users)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,user)
}

func Getuserbyid(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	user,err:=Userservice.Getdatabyid(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,user)
}
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := Userservice.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}