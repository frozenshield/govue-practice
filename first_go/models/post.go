package models

type Post struct {
	ID      uint   `json:"posts_id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`

	UserID uint  `json:"user_id"`
	User   Users `json:"-" gorm:"foreignKey:UserID"`
}
