package Models

type Cast struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`                  // e.g., "Lead", "Supporting"
	VideoID int    `json:"video_id"` // Foreign key
}
