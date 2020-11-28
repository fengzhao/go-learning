package main

import (
	"fmt"
	"reflect"
)

func main() {

	var p = new(int)
	// p的值：0xc000072090 （指针是内存地址）
	fmt.Println(p)
	// p的数据类型：p是指针
	fmt.Println("the type of x is :", reflect.TypeOf(p))

	//
	fmt.Println(*p)

	// 切片声明
	// make(元素类型,切片长度,切片容量)
	mySlice1 := make([]int, 5, 10)
	// 当长度和容量相同时，可以省略容量，make(元素类型,切片长度)
	//mySlice1 := make([]int, 5)

	fmt.Println(mySlice1)
	// [0 0 0 0 0]
	fmt.Println("the type of mySlice1 is :", reflect.TypeOf(mySlice1))
	// the type of mySlice1 is : []int

	// make短声明一个空map
	stringMap := make(map[string][]string, 100)
	fmt.Println(stringMap)
	// map[]

	// 添加键值对，值是切片字面量
	stringMap["Red"] = []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	var slice3 = []string{"1", "2", "3", "4"}
	stringMap["slice3"] = slice3
	fmt.Println(stringMap)
	// map[]

}
