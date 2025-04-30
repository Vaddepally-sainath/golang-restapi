package Userservice

import (
	"fmt"
	"new-ott-project/User-service/Authentication"
	"new-ott-project/User-service/Database"
	"new-ott-project/User-service/Usermodel"
)

func Createuser(user Usermodel.Users)(Usermodel.Users,error){
	if Database.DB == nil {
        return user, fmt.Errorf("database is not initialized")
    }
	result:=Database.DB.Create(&user)
	return user,result.Error
}

func Getdatabyid(id int)(Usermodel.Users,error){
	var user Usermodel.Users
	result:=Database.DB.First(&user,id)
	return user,result.Error
}

func LoginUser(email, password string) (string, error) {
	var user Usermodel.Users
	err := Database.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return "", err
	}

	token, err := Authentication.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}