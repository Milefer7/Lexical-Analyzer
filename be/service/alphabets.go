package service

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dao/mysql"
)

type AlphabetService struct{}

func NewAlphabetService() *AlphabetService {
	return &AlphabetService{}
}

func (AlphabetService *AlphabetService) ReadAlphabets() (error, []model.Alphabet) {
	return mysql.NewAlphabetMySQL().ReadAlphabets()
}

func (AlphabetService *AlphabetService) UpdateAlphabets(Alphabets []model.Alphabet) error {
	// 遍历判断update的值是否存在
	// 如果存在,返回错误
	//err, alpDbs := mysql.NewAlphabetMySQL().ReadAlphabets()
	//if err != nil {
	//	return err
	//}
	//// 打印Alphabets的值
	//log.Printf("Alphabets: %v", Alphabets)
	//for _, alphabet := range Alphabets {
	//	for _, a := range alpDbs {
	//		if alphabet.Value == a.Value {
	//			err := errors.New("修改的值已存在")
	//			return err
	//		}
	//	}
	//}
	//// 不存在则插入
	return mysql.NewAlphabetMySQL().UpdateAlphabets(Alphabets)
}

func (AlphabetService *AlphabetService) CreateAlphabets(Alphabets []model.Alphabet) error {
	return mysql.NewAlphabetMySQL().CreateAlphabets(Alphabets)
}
