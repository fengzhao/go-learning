package main

import "fmt"

// 声明一个基于int的类型Integer。有点像java中的extend关键字（类继承）
type Integer int

// 声明一个函数 Less，可以作用于Integer类型的变量上。有点像面对对象中，给某个类声明一个方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	// 声明
	var a Integer = 1
	// 调用其方法
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
}
