package biz

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dal/query"
)

func UpdateWords(word []*model.Words) error {
	err := query.Words.Save(word...)
	if err != nil {
		return err
	}
	return nil
}

func CreateWords(word []*model.Words) error {
	err := query.Words.Create(word...)
	if err != nil {
		return err
	}
	return nil
}

func ReadWords() (error, []*model.Words) {
	words, err := query.Words.Find()
	if err != nil {
		return err, nil
	}
	return nil, words
}
