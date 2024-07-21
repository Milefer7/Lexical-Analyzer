package mysql

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"log"
)

type KeywordsMySQL struct {
}

func NewKeywordsMySQL() *KeywordsMySQL {
	return &KeywordsMySQL{}
}

func (kw *KeywordsMySQL) ReadKeywords() ([]model.Keywords, error) {
	keywords := []model.Keywords{}
	result := DB.Find(&keywords)
	if result.Error != nil {
		log.Printf("readKeyword数据库查询错误: %s", result.Error)
	}
	return keywords, result.Error
}

func (kw *KeywordsMySQL) CreateKeywords(keywords []model.Keywords) error {
	for _, keyword := range keywords {
		result := DB.Create(&keyword)
		if result.Error != nil {
			log.Printf("CreateKeyword数据库插入错误: %s", result.Error)
			return result.Error
		}
	}
	return nil
}

func (kw *KeywordsMySQL) UpdateKeywords(keywords []model.Keywords) error {
	for _, keyword := range keywords {
		result := DB.Save(&keyword)
		if result.Error != nil {
			log.Printf("UpdateKeyword数据库更新错误: %s", result.Error)
			return result.Error
		}
	}
	return nil
}
