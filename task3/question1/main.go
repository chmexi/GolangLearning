package main

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint
	Name  string
	Age   int
	Grade string
}

func (s *Student) BeforeSave(db *gorm.DB) error {
	students := []Student{}
	db.Find(&students)
	fmt.Println(students)
	return nil
}

func main() {
	dsn := "root:meng0987612345@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error cannot open database error: ", err)
	}
	db.AutoMigrate(&Student{})
	zhnagsan := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	zhnagsi := Student{
		Name:  "张四",
		Age:   10,
		Grade: "三年级",
	}
	fmt.Println("插入")
	db.Create(&zhnagsan)
	db.Create(&zhnagsi)
	students := []Student{}
	fmt.Println("查找")
	db.Where("age > ?", 18).Find(&students)
	fmt.Println(students)
	fmt.Println("修改")
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	db.Find(&students)
	fmt.Println(students)
	fmt.Println("删除")
	db.Model(&Student{}).Where("age < ?", 15).Delete(&Student{})
	db.Find(&students)
	fmt.Println(students)
}
