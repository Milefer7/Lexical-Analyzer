package main

import (
	"github.com/Milefer7/compliation_exp/dal"
	"github.com/Milefer7/compliation_exp/dal/query"

	"github.com/Milefer7/compliation_exp/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	dal.Init()
	// 初始化路由
	e := gin.Default()
	// 初始化路由
	router.InitRouter(e)
	// 定义
	query.SetDefault(dal.Db)
	// 运行
	err := e.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
