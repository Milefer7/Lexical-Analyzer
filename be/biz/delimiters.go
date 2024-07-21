package biz

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dal/query"
)

func ReadDelimiters() ([]*model.Delimiter, error) {
	delimiter, err := query.Delimiter.Find()
	if err != nil {
		return nil, err
	}
	return delimiter, nil
}

func UpdateDelimiters(Delimiters []*model.Delimiter) error {
	err := query.Q.Delimiter.Save(Delimiters...)
	if err != nil {
		return err
	}
	return nil
}

func CreateDelimiters(Delimiters []*model.Delimiter) error {
	err := query.Q.Delimiter.Create(Delimiters...)
	if err != nil {
		return err
	}
	return nil
}
