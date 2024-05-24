package router

import "github.com/gin-gonic/gin"

func InitRouter(e *gin.Engine) {
	keyword := e.Group("/keyword")
	delimiter := e.Group("/delimiter")
	analysis := e.Group("/analysis")
	alphabet := e.Group("/alphabet")
	{
		alphabet.POST("/read", controller.Read)
	}
	word := e.Group("/word")
}
