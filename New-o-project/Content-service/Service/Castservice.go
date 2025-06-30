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

func updatacastbyid(id int, updateddata Models.Cast) (Models.Cast, error) {
	var video Models.Cast
	if err := Database.DB.First(&video, id).Error; err != nil {
		return video, err
	}
	video.Name = updateddata.Name
	video.Role = updateddata.Role
	result := Database.DB.Updates(&video)
	return updateddata, result.Error
}