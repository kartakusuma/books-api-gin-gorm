package models

type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"type:varchar(50)"`
	Author      string `json:"author" gorm:"type:varchar(50)"`
	Description string `json:"description" gorm:"type:varchar(50)"`
}
