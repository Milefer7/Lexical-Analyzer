package router

import (
	"github.com/Milefer7/compliation_exp/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	keyword := e.Group("/keywords")
	{
		keyword.GET("", controller.ReadKeywords)
		keyword.PUT("", controller.UpdateKeywords)
		keyword.POST("", controller.CreateKeywords)
	}
	delimiter := e.Group("/delimiters")
	{
		delimiter.GET("", controller.ReadDelimiters)
		delimiter.PUT("", controller.UpdateDelimiters)
		delimiter.POST("", controller.CreateDelimiters)
	}
	analysis := e.Group("/analysis")
	{
		analysis.GET("/lexical", controller.LexicalAnalysis)
	}
	alphabet := e.Group("/alphabets")
	{
		alphabet.GET("", controller.ReadAlphabets)
		alphabet.PUT("", controller.UpdateAlphabets)
		alphabet.POST("", controller.CreateAlphabets)
	}
	word := e.Group("/words")
	{
		word.PUT("", controller.UpdateWords)
		word.POST("", controller.CreateWords)
	}
}
