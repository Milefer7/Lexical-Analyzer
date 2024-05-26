package mysql

import (
	"github.com/Milefer7/compliation_exp/model"
	"log"
)

type WordsMySQL struct{}

func NewWordsMySQL() *WordsMySQL {
	return &WordsMySQL{}
}

func (WordsMySQL *WordsMySQL) UpdateWord(words []model.Words) error {
	for _, word := range words {
		if err := db.Save(&word).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

func (WordsMySQL *WordsMySQL) CreatWords(words []model.Words) error {
	for _, word := range words {
		if err := db.Create(&word).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}
