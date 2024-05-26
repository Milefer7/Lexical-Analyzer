package mysql

import (
	"github.com/Milefer7/compliation_exp/model"
	"log"
)

type AlphabetMySQL struct{}

func NewAlphabetMySQL() *AlphabetMySQL {
	return &AlphabetMySQL{}
}

func (AlphabetMySQL *AlphabetMySQL) ReadAlphabets() (error, []model.Alphabet) {
	alphabet := []model.Alphabet{}
	result := db.Find(&alphabet)
	if result.Error != nil {
		log.Printf("ReadAlphabet数据库查询错误: %s", result.Error)
	}
	return result.Error, alphabet
}

func (AlphabetMySQL *AlphabetMySQL) UpdateAlphabets(Alphabets []model.Alphabet) error {
	for _, alphabet := range Alphabets {
		result := db.Save(&alphabet)
		if result.Error != nil {
			log.Printf("UpdateAlphabet数据库更新错误: %s", result.Error)
			return result.Error
		}
	}
	return nil
}

func (AlphabetMySQL *AlphabetMySQL) CreateAlphabets(Alphabets []model.Alphabet) error {
	for _, alphabet := range Alphabets {
		result := db.Create(&alphabet)
		if result.Error != nil {
			log.Printf("CreateAlphabet数据库创建错误: %s", result.Error)
			return result.Error
		}
	}
	return nil
}
