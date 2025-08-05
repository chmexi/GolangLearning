package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

func NewRectangle(width float32, height float32) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func (r *Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

type Circle struct {
	Radius float32
}

func NewCircle(radius float32) *Circle {
	return &Circle{
		Radius: radius,
	}
}

func (c *Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

// 通用的打印形状信息函数
func PrintShapeInfo(s Shape, description string) {
	fmt.Println("------------------")
	fmt.Println(description)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("------------------")
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("------------------")
	fmt.Println("name: ", e.Name)
	fmt.Println("age: ", e.Age)
	fmt.Println("employee_id: ", e.EmployeeID)
	fmt.Println("------------------")
}

func main() {
	c := NewCircle(3)
	PrintShapeInfo(c, "c的形状信息为:")

	r := NewRectangle(3, 3)
	PrintShapeInfo(r, "r的形状信息为:")

	var e Employee = Employee{
		Person: Person{
			Name: "张三",
			Age:  666,
		},
		EmployeeID: 996,
	}
	e.PrintInfo()
}
