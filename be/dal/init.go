package dal

import (
	"github.com/Milefer7/compliation_exp/dal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func Init() {
	// 赋予dsn值
	err := godotenv.Load("../be/.env")
	if err != nil {
		log.Fatalf("加载 .env 失败: %v", err)
	}
	dsn := os.Getenv("DSN_LOCAL")
	// 连接数据库
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接出错: %s", err)
	}
	// 数据库表初始化
	initTable(d)
	// 将连接赋值给db
	Db = d
}

func initTable(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Alphabet{},
		&model.Delimiter{},
		&model.Words{},
		&model.Keywords{},
	)
	if err != nil {
		log.Fatalf("数据库表初始化错误: %s", err)
	}

}
