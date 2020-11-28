package main

import "fmt"

func main() {

	// 一 、使用字面量创建map
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// 二、使用map声明一个map
	// colors2 := make(map[string]string, 100)
	// 迭代map
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	removeColor(colors, "Coral")

}

// 定义一个函数，用于删除map中的元素
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
