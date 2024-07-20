package main

import (
	"github.com/Milefer7/compliation_exp/dao/mysql"
	"github.com/Milefer7/compliation_exp/model"
	"gorm.io/gen"
)

func main() {

	// 创建一个新的生成器，配置输出路径和模式
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
	})

	// 使用已经初始化的数据库
	g.UseDB(mysql.db)

	// 这个是从数据库表生成模型的模板
	//userTpl := g.GenerateModel("inv_lists")

	// 这个是应用自己写的模型
	g.ApplyBasic(model.Alphabet{}, model.LexicalAnalysis) // Associations
	g.ApplyBasic(userTpl)

	// 执行生成操作
	g.Execute()
}
