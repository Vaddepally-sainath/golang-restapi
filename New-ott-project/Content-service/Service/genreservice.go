package Service

import (
	"new-ott-project/Content-service/Database"
	"new-ott-project/Content-service/Models"
)

func Creategenres(genre Models.Genre) (Models.Genre,error){
  
  result:=Database.DB.Create(&genre)
  return genre,result.Error
}

