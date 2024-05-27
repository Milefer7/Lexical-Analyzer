package controller

import (
	"github.com/Milefer7/compliation_exp/code"
	"github.com/Milefer7/compliation_exp/model"
	"github.com/Milefer7/compliation_exp/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 词法分析模块

//	关键词表

// 获取关键词表
func ReadKeywords(c *gin.Context) {
	keywords, err := service.NewKeywordsService().ReadKeywords()
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", keywords)
}

// 更新关键词表
func UpdateKeywords(c *gin.Context) {
	keywords := []model.Keywords{}
	err := c.ShouldBindJSON(&keywords)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewKeywordsService().UpdateKeywords(keywords)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 创建关键词表
func CreateKeywords(c *gin.Context) {
	keywords := []model.Keywords{}
	err := c.ShouldBindJSON(&keywords)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewKeywordsService().CreateKeywords(keywords)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 分界符表
// 获取分界符表
func ReadDelimiters(c *gin.Context) {
	delimiters, err := service.NewDelimitersService().ReadDelimiters()
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", delimiters)
}

// 更新分界符表
func UpdateDelimiters(c *gin.Context) {
	delimiters := []model.Delimiter{}
	err := c.ShouldBindJSON(&delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewDelimitersService().UpdateDelimiters(delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

// 创建分界符表
func CreateDelimiters(c *gin.Context) {
	delimiters := []model.Delimiter{}
	err := c.ShouldBindJSON(&delimiters)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewDelimitersService().CreateDelimiters(delimiters)
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
	data := model.LexicalCode{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	// 调用service层函数
	log.Println("data.Code: ", data.Code)
	LexicalAnalysis, err := service.LexicalAnalysis(data.Code)
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
	err, alphabet := service.NewAlphabetService().ReadAlphabets()
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
	alphabet := []model.Alphabet{}
	err := c.ShouldBindJSON(&alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewAlphabetService().UpdateAlphabets(alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

func CreateAlphabets(c *gin.Context) {
	alphabet := []model.Alphabet{}
	err := c.ShouldBindJSON(&alphabet)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewAlphabetService().CreateAlphabets(alphabet)
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
	words := []model.Words{}
	err := c.ShouldBindJSON(&words)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewWordsService().UpdateWords(words)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

func CreateWords(c *gin.Context) {
	words := []model.Words{}
	err := c.ShouldBindJSON(&words)
	if err != nil {
		code.CommonResp(c, http.StatusBadRequest, code.Fail, err.Error(), code.EmptyData)
		return
	}
	err = service.NewWordsService().CreateWords(words)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", code.EmptyData)
}

func ReadWords(c *gin.Context) {
	err, words := service.NewWordsService().ReadWords()
	if err != nil {
		println(err.Error())
		code.CommonResp(c, http.StatusInternalServerError, code.Fail, err.Error(), code.EmptyData)
		return
	}
	code.CommonResp(c, http.StatusOK, code.Success, "", words)
}
