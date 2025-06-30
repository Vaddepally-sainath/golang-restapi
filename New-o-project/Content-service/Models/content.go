package Models

type Videos struct {
	ID          int     `gorm:"primaryKey" json:"ID"`
	Title       string  `json:"Title" binding:"required"`
	Releaseyear int     `json:"Releaseyear" binding:"required,lte=2025,gte=1900"`
	Ispremium   bool    `json:"Ispremium"`
	Cast        []Cast  `gorm:"foreignKey:VideoID"`
	Genre       []Genre `gorm:"many2many:video-genres" json:"genre"`
}
