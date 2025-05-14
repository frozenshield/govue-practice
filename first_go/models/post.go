package models

type Post struct {
	ID     uint   `json:"posts_id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Conent string `json:"content"`
}
