package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Topic struct {
	gorm.Model
	Title   string
	Content string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("数据库连接失败")
	}
	defer db.Close()

	db.AutoMigrate(&Topic{})

	// 创建一条记录
	db.Create(&Topic{Title: "welcome", Content: "jingwei.link"})

	// 查找一个topic
	// 这里的err值不为空，会报出"recode not found"的错误
	var topic Topic
	err = db.Where("id=?", 0).Find(&topic).Error
	fmt.Println("error:", err)
	fmt.Println("topic的id:", topic.ID)
	// First
	err = db.Where("id=?", 0).First(&topic).Error
	fmt.Println("error:", err)
	fmt.Println("topic的id:", topic.ID)
	// Last
	err = db.Where("id=?", 0).Last(&topic).Error
	fmt.Println("error:", err)
	fmt.Println("topic的id:", topic.ID)

	// 查找topic列表
	// 这里的err值为nil
	var topics []Topic
	err = db.Where("id=?", 0).Find(&topics).Error
	fmt.Println("error:", err)
	fmt.Println("topics的长度", len(topics))
	// First
	err = db.Where("id=?", 0).First(&topics).Error
	fmt.Println("error：", err)
	fmt.Println("topics的长度", len(topics))
	// Last
	err = db.Where("id=?", 0).Last(&topics).Error
	fmt.Println("error：", err)
	fmt.Println("topics的长度", len(topics))
}
