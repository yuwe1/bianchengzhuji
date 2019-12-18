package main

import "fmt"

// 字符个数的统计
func calculateTotalNum(str string) {
	// strs := []rune(str)
	// 定义一个map[rune]int
	var m map[rune]int
	m = make(map[rune]int, len(str))
	// fmt.Println("map：", len(m))
	for _, v := range str {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	// fmt.Println("map：", len(m))
	for i, v := range m {
		fmt.Println("字符：", string(i), "    个数：", v)
	}
}
func main() {
	calculateTotalNum("aaabbbcccddda")
}
