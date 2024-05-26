package service

import (
	"github.com/Milefer7/compliation_exp/dao/mysql"
	"github.com/Milefer7/compliation_exp/model"
)

type AlphabetService struct{}

func NewAlphabetService() *AlphabetService {
	return &AlphabetService{}
}

func (AlphabetService *AlphabetService) ReadAlphabets() (error, []model.Alphabet) {
	return mysql.NewAlphabetMySQL().ReadAlphabets()
}

func (AlphabetService *AlphabetService) UpdateAlphabets(Alphabets []model.Alphabet) error {
	return mysql.NewAlphabetMySQL().UpdateAlphabets(Alphabets)
}

func (AlphabetService *AlphabetService) CreateAlphabets(Alphabets []model.Alphabet) error {
	return mysql.NewAlphabetMySQL().CreateAlphabets(Alphabets)
}
