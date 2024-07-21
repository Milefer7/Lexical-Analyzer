package biz

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dal/query"
)

func UpdateAlphabets(Alphabets []*model.Alphabet) error {
	err := query.Alphabet.Save(Alphabets...)
	if err != nil {
		return err
	}
	return nil
}

func CreateAlphabets(Alphabets []*model.Alphabet) error {
	err := query.Q.Alphabet.Create(Alphabets...)
	if err != nil {
		return err
	}
	return nil
}

func ReadAlphabets() (error, []*model.Alphabet) {
	alphabets, err := query.Alphabet.Find()
	if err != nil {
		return err, nil
	}
	return nil, alphabets
}
