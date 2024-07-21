package model

type Keywords struct {
	Id      int    `json:"id" gorm:"primaryKey" binding:"required"`
	Keyword string `json:"keyword" `
}
