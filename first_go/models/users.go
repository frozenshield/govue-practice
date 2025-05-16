package models

type Users struct {
	ID           uint   `json:"users_id" gorm:"primaryKey;column:id"`
	Microsoft_ID int    `json:"microsoft_id"`
	Username     string `json:"username" gorm:"type:VARCHAR(100)"`
	Password     string `json:"password" gorm:"type:VARCHAR(255)"`
	Post         []Post `gorm:"foreignKey:UserID"`
}
