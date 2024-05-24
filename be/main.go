package main

import (
	"github.com/Milefer7/compliation_exp/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	model.Init()
	// 初始化路由
	e := gin.Default()

	// 运行
	err := e.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
