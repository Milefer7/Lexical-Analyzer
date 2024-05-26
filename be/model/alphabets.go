package model

// Alphabet 是一个结构体，表示键值对数据
type Alphabet struct {
	Key   int    `gorm:"primaryKey" json:"key" binding:"required" gorm:"column:key"`
	Value string `json:"value" binding:"required" gorm:"column:value"`
}
