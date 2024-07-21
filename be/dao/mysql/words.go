package mysql

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"log"
)

type WordsMySQL struct{}

func NewWordsMySQL() *WordsMySQL {
	return &WordsMySQL{}
}

func (WordsMySQL *WordsMySQL) UpdateWord(words []model.Words) error {
	for _, word := range words {
		if err := DB.Save(&word).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

func (WordsMySQL *WordsMySQL) CreatWords(words []model.Words) error {
	for _, word := range words {
		if err := DB.Create(&word).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

func (WordsMySQL *WordsMySQL) ReadWords() (error, []model.Words) {
	words := []model.Words{}
	if err := DB.Find(&words).Error; err != nil {
		log.Println(err.Error())
		return err, nil
	}
	return nil, words
}
