package service

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dao/mysql"
)

type DelimitersService struct{}

func NewDelimitersService() *DelimitersService {
	return &DelimitersService{}
}

func (d *DelimitersService) ReadDelimiters() ([]model.Delimiter, error) {
	return mysql.NewDelimiterMySQL().ReadDelimiters()
}

func (d *DelimitersService) UpdateDelimiters(Delimiters []model.Delimiter) error {
	return mysql.NewDelimiterMySQL().UpdateDelimiters(Delimiters)
}

func (d *DelimitersService) CreateDelimiters(Delimiters []model.Delimiter) error {
	return mysql.NewDelimiterMySQL().CreateDelimiters(Delimiters)
}
