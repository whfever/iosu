package main

import "fmt"

// User 结构体用来演示指针的使用
type User struct {
	name string
	age  int
}

// 使用指针接收者的方法可以修改结构体字段
func (u *User) setAge(age int) {
	u.age = age // 直接修改原始数据
}

// 使用值接收者的方法不会修改原始数据
func (u User) getAgeDoubled() int {
	u.age *= 2 // 只修改副本
	return u.age
}

func main() {
	// 1. 基本类型指针
	num := 42
	ptr := &num // 取地址
	fmt.Printf("指针的值：%v\n", ptr)
	fmt.Printf("指针指向的值：%v\n", *ptr)

	*ptr = 24 // 通过指针修改值
	fmt.Printf("修改后的num：%v\n", num)

	// 2. 结构体指针
	user := &User{ // 创建结构体并获取指针
		name: "张三",
		age:  20,
	}

	user.setAge(21) // 通过指针方法修改年龄
	fmt.Printf("新年龄：%v\n", user.age)

	doubled := user.getAgeDoubled() // 值方法不会修改原始数据
	fmt.Printf("翻倍后的年龄：%v\n", doubled)
	fmt.Printf("原始年龄：%v\n", user.age) // 原始年龄未变
}
