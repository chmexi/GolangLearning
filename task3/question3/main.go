package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
Sqlx入门
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	ID     uint    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func CreateTable() {
	dsn := "root:meng0987612345@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error cannot open database error: ", err)
	}

	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Book{})

	db.Create([]Employee{
		Employee{
			Name:       "小桌子难怪",
			Department: "销售部",
			Salary:     0.91,
		},
		Employee{
			Name:       "小张",
			Department: "技术部",
			Salary:     9.1,
		},
		Employee{
			Name:       "那艘i发",
			Department: "技术部",
			Salary:     9.91,
		}})

	// 插入 Book 表的测试数据
	db.Create([]Book{
		{
			Title:  "Go 语言实战",
			Author: "John Doe",
			Price:  69.99,
		},
		{
			Title:  "数据库系统概念",
			Author: "Jane Smith",
			Price:  30.99,
		},
		{
			Title:  "算法导论",
			Author: "Donald Knuth",
			Price:  129.99,
		},
	})
}

func main() {
	// CreateTable()
	db, err := sqlx.Open("mysql", "root:meng0987612345@tcp(127.0.0.1:3306)/learning")
	if err != nil {
		fmt.Println("连接数据库失败 error:", err)
		return
	}
	defer db.Close()

	techEmployees := []Employee{}
	err = db.Select(&techEmployees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		fmt.Println("查询技术部员工失败 error:", err)
		return
	}
	fmt.Println("技术部员工为:", techEmployees)

	highestSalaryEmployee := Employee{}
	err = db.Get(&highestSalaryEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		fmt.Println("查询工资最高员工失败 error:", err)
		return
	}
	fmt.Println("工资最高员工为:", highestSalaryEmployee)

	higherThan50Books := []Book{}
	err = db.Select(&higherThan50Books, "SELECT id, title, price, author FROM books WHERE price > ?", 50)
	if err != nil {
		fmt.Println("查询所有售价>50的书 error:", err)
		return
	}
	fmt.Println("售价大于50的书为:", higherThan50Books)
}
