package biz

import (
	"github.com/Milefer7/compliation_exp/dal"
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dal/query"

	"github.com/gin-gonic/gin"
)

func ReadKeywords(c *gin.Context) ([]*model.Keywords, error) {
	var q = query.Use(dal.Db)
	Q := q.Keywords.WithContext(c.Request.Context())
	keywords, err := Q.Find()
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

func UpdateKeywords(u []*model.Keywords) error {
	err := query.Keywords.Save(u...)
	if err != nil {
		return err
	}
	return nil
}

func CreateKeywords(keywords []*model.Keywords) error {
	err := query.Q.Keywords.Create(keywords...)
	if err != nil {
		return err
	}
	return nil
}
