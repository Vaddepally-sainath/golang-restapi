package Service

import (
	"fmt"
	"new-ott-project/Content-service/Database"
	"new-ott-project/Content-service/Models"
)

func Createcast(Videos Models.Cast)( Models.Cast,error){ 
	if Database.DB == nil {
		return Videos, fmt.Errorf("database is not initialized")
	}
	result := Database.DB.Create(&Videos)
	return Videos, result.Error
}

func Getcastbyid(id int)(Models.Cast,error){
	var video Models.Cast
	result:=Database.DB.First(&video,id)
	return video,result.Error
}