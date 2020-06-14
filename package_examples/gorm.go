package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	logrus.SetReportCaller(true)
}

type Product struct {
	gorm.Model
	Title   string
	Content string
	Created time.Time
}

func main() {
	db, err := gorm.Open("sqlite3", "package_examples/data/test.db")
	if err != nil {
		logrus.Error(err)
		panic("Connect Sqlite3 err ")
	}

	defer db.Close()

	//自动迁移
	db.AutoMigrate(&Product{})

	//创建
	db.Create(&Product{Title: "XiaoMi", Content: "New Product", Created: time.Now()})

	//读取
	var product Product
	first := db.First(&product, 1)                 // 查询id为1的product
	row := db.First(&product, "Title = ?", "XiaoMi") // 查询code为l1212的product
	fmt.Println(first)
	fmt.Println(*row)
}
