package mysql

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"log"
)

type DelimiterMySQL struct{}

func NewDelimiterMySQL() *DelimiterMySQL {
	return &DelimiterMySQL{}
}

func (d *DelimiterMySQL) ReadDelimiters() ([]model.Delimiter, error) {
	deli := []model.Delimiter{}
	result := DB.Find(&deli)
	if result.Error != nil {
		log.Printf("ReadDelimiter数据库查询错误: %s", result.Error)
	}
	return deli, result.Error
}

func (d *DelimiterMySQL) UpdateDelimiters(Delimiters []model.Delimiter) error {
	for _, delimiter := range Delimiters {
		result := DB.Save(&delimiter)
		if result.Error != nil {
			log.Printf("UpdateDelimiter数据库更新错误: %s", result.Error)
			return result.Error
		}
	}
	return nil
}

func (d *DelimiterMySQL) CreateDelimiters(Delimiters []model.Delimiter) error {
	for _, delimiter := range Delimiters {
		result := DB.Create(&delimiter)
		if result.Error != nil {
			log.Printf("CreateDelimiter数据库创建错误: %s", result.Error)
			return result.Error
		}
	}
	return nil
}
