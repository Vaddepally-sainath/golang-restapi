package Models

type Genre struct {
	ID     int      `gorm:"primaryKey" json:"id"`
	Name   string   `json:"name"`
	Videos []Videos `gorm:"many2many:video-genres" json:"videos"`
}
