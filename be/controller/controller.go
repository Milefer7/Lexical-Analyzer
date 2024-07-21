package controller

import (
	"github.com/Milefer7/compliation_exp/biz"
	"github.com/Milefer7/compliation_exp/code"
	model2 "github.com/Milefer7/compliation_exp/dal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 词法分析模块

//	关键词表

// 获取关键词表
func ReadKeywords(c *gin.Context) {
	keywords, err := biz.ReadKeywords(c)
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", keywords)
}

// 更新关键词表
func UpdateKeywords(c *gin.Context) {
	keywords := []*model2.Keywords{}
	err := c.ShouldBindJSON(&keywords)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.UpdateKeywords(keywords)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 创建关键词表
func CreateKeywords(c *gin.Context) {
	keywords := []*model2.Keywords{}
	err := c.ShouldBindJSON(&keywords)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.CreateKeywords(keywords)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 分界符表
// 获取分界符表
func ReadDelimiters(c *gin.Context) {
	delimiters, err := biz.ReadDelimiters()
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", delimiters)
}

// 更新分界符表
func UpdateDelimiters(c *gin.Context) {
	delimiters := []*model2.Delimiter{}
	err := c.ShouldBindJSON(&delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.UpdateDelimiters(delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 创建分界符表
func CreateDelimiters(c *gin.Context) {
	delimiters := []*model2.Delimiter{}
	err := c.ShouldBindJSON(&delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.CreateDelimiters(delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 词法分析
// 获取词法分析结果
func LexicalAnalysis(c *gin.Context) {
	// 接受json数据
	data := model2.LexicalCode{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	// 调用service层函数
	log.Println("data.Code: ", data.Code)
	LexicalAnalysis, err := biz.LexicalAnalysis(data.Code)
	log.Printf("词法分析结果：%v", LexicalAnalysis)
	// 将结果返回给前端
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", LexicalAnalysis)
}

// 系统维护模块

// 字符表
// 获取字符表
func ReadAlphabets(c *gin.Context) {
	err, alphabet := biz.ReadAlphabets()
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", alphabet)
}

// 更新字符表
func UpdateAlphabets(c *gin.Context) {
	//获取前端传的数据
	alphabet := []*model2.Alphabet{}
	err := c.ShouldBindJSON(&alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.UpdateAlphabets(alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

func CreateAlphabets(c *gin.Context) {
	alphabet := []*model2.Alphabet{}
	err := c.ShouldBindJSON(&alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.CreateAlphabets(alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 单词分类
// 更新单词
func UpdateWords(c *gin.Context) {
	// 获取前端传的数据
	words := []*model2.Words{}
	err := c.ShouldBindJSON(&words)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.UpdateWords(words)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

func CreateWords(c *gin.Context) {
	words := []*model2.Words{}
	err := c.ShouldBindJSON(&words)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = biz.CreateWords(words)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

func ReadWords(c *gin.Context) {
	err, words := biz.ReadWords()
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", words)
}
