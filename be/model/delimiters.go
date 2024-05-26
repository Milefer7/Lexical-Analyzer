package model

type Delimiter struct {
	ID   int    `json:"id" gorm:"primaryKey" binding:"required"`
	Name string `json:"name" gorm:"column:name" binding:"required"`
}
