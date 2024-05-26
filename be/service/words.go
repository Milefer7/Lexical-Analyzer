package service

import (
	"github.com/Milefer7/compliation_exp/dao/mysql"
	"github.com/Milefer7/compliation_exp/model"
)

type WordsService struct{}

func NewWordsService() *WordsService {
	return &WordsService{}
}

func (ws *WordsService) UpdateWords(word []model.Words) error {
	return mysql.NewWordsMySQL().UpdateWord(word)
}

func (ws *WordsService) CreateWords(word []model.Words) error {
	return mysql.NewWordsMySQL().CreatWords(word)
}

func (ws *WordsService) ReadWords() (error, []model.Words) {
	return mysql.NewWordsMySQL().ReadWords()
}
