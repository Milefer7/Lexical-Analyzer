package model

type Words struct {
	ID   int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name string `json:"name" binding:"required"`
	Word string `json:"word" binding:"required"`
}
