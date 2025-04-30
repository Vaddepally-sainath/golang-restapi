package Database

import (
	"log"
      "new-ott-project/User-service/Usermodel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	dsn:="root:Sainath@2000@tcp(127.0.0.1:3306)/newuserdetails?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		log.Fatal("database connection error:",err)
		return
	}
	log.Println("database connected successfully")

	DB.AutoMigrate(&Usermodel.Users{})
}


