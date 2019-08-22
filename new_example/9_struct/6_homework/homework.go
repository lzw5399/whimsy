package main

import (
	"fmt"
)

// 问题: 如果匿名字段和外层结构有同名字段，应该如何进行操作
func main() {
	// 单独测试父类
	//f := Father{
	//	Name: "父类外层name",
	//	Contact: struct {
	//		Name string
	//	}{
	//		Name: "父类里层Name",
	//	},
	//}
	// fmt.Println(f, f.Name, f.Contact.Name)

	// 测试子类
	s :=Son{
		Father: Father{
			Name: "父类外层Name",
			Contact: struct {
				Name string
			}{
				Name: "父类里层Name",
			},
		},
		Name:   "子类外层Name",
		//Contact: struct {
		//	Name string
		//}{
		//	Name: "子类里层Name",
		//},
	}

	fmt.Println(s)
}

type Father struct {
	Name    string
	FatherName string
	Contact struct {
		Name string
	}
}

type Son struct {
	Father
	Hi int
	Name string
	//Contact struct{
	//	Name string
	//}
}
