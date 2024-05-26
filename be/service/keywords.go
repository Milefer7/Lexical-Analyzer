package service

import (
	"github.com/Milefer7/compliation_exp/dao/mysql"
	"github.com/Milefer7/compliation_exp/model"
)

type KeywordsService struct {
}

func NewKeywordsService() *KeywordsService {
	return &KeywordsService{}
}

func (ws *KeywordsService) ReadKeywords() ([]model.Keywords, error) {
	return mysql.NewKeywordsMySQL().ReadKeywords()
}

func (ws *KeywordsService) CreateKeywords(keywords []model.Keywords) error {
	return mysql.NewKeywordsMySQL().CreateKeywords(keywords)
}

func (ws *KeywordsService) UpdateKeywords(keywords []model.Keywords) error {
	return mysql.NewKeywordsMySQL().UpdateKeywords(keywords)
}
