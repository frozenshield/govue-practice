package models

type Users struct {
	ID           uint   `json:"users_id" gorm:"primary"`
	Microsoft_ID int    `json:"microsoft_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
}
