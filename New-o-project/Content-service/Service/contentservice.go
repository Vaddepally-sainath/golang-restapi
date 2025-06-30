package Service

import (
	"fmt"
	"log"
	"new-ott-project/Content-service/Database"
	"new-ott-project/Content-service/Models"
)




func Createvideos(Videos Models.Videos) (Models.Videos, error) {
	if Database.DB == nil {	
		return Videos, fmt.Errorf("database is not initialized")
	}
	result := Database.DB.Create(&Videos)
	if result.Error != nil {
		log.Println("Error creating video:", result.Error)
		return Videos, result.Error
	}
	return Videos, nil
}

func Getalldata() ([]Models.Videos, error) {
	var video []Models.Videos
	result := Database.DB.Find(&video)

log.Println("Data:",video)
log.Println("working")

	return video, result.Error
}


func Getdatabyid(id int) (Models.Videos, error) {
	var video Models.Videos
	if err := Database.DB.Preload("Genres").First(&video, id).Error; err != nil {
		log.Println("Error fetching video by ID:", err)
		return Models.Videos{}, err
	}
	return video, nil
}

func Updatedatabyid(id int, updateddata Models.Videos) (Models.Videos, error) {
	var video Models.Videos
	if err := Database.DB.First(&video, id).Error; err != nil {
		return video, err
	}
	video.Genre = updateddata.Genre
	video.Ispremium = updateddata.Ispremium
	video.Releaseyear = updateddata.Releaseyear
	video.Title = updateddata.Title
	result := Database.DB.Updates(&video)
	return updateddata,result.Error
}

func Deletebyid(id int) (string, error) {
	var video Models.Videos

	// Check if the record exists before deleting
	if err :=Database.DB.First(&video, id).Error; err != nil {
		return "", fmt.Errorf("video with ID %d not found", id)
	}

	// Perform delete
	result:= Database.DB.Delete(&video) 
	

	return "video deleted successfully", result.Error
}

func GetVideosByGenres(Name []string) ([]Models.Videos, error) {
	var videos []Models.Videos
	var genre_id []int
	var video_id []int

	// Step 1: Get genre IDs from genre names
	if err := Database.DB.Model(&Models.Genre{}).
		Where("name IN ?", Name).
		Pluck("id", &genre_id).Error; err != nil {
		return nil, err
	}

	if len(genre_id) == 0 {
		return []Models.Videos{}, nil // No matching genres
	}

	// Step 2: Get video IDs that are linked to all selected genres
	if err := Database.DB.Table("video_genres").
		Select("video_id").
		Where("genre_id IN ?", genre_id).
		Group("video_id").
		Having("COUNT(DISTINCT genre_id) = ?", len(genre_id)).
		Pluck("video_id", &video_id).Error; err != nil {
		return nil, err
	}

	// Step 3: Fetch videos with preloaded Genres
	if err := Database.DB.Preload("Genres").
		Where("id IN ?", video_id).
		Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func Getalldatas()([]Models.Videos,error){
	var vide []Models.Videos
	result:=Database.DB.Find(&vide)
	return vide,result.Error
}



