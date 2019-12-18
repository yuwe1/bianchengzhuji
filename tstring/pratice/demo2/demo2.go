package main

import "fmt"

// 字符串的左右移动
// 字符串为*和26个字母的任意组合，现在需要将*都移动到左侧，字母都移动到右侧，保持相对顺序不变
func movestringrorf(str string) {
	// 定义 两个字符串变量用于存放字母和*
	left := ""
	right := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "*" {
			left += "*"
		} else {
			right += string(str[i])
		}
	}
	// 完成字符串的拼接
	result := left + right
	fmt.Println(result)
}

// 倒叙进行赋值，后面没有用到的空间都赋值为"*"
func movestring_two(str string) {
	strs := []rune(str)
	// 求出字符串的长度
	now := len(str)
	for i := now - 1; i >= 0; i-- {
		if string(strs[i]) != "*" {
			now = now - 1
			strs[now] = strs[i]
		}
	}
	for now = now - 1; now >= 0; now-- {
		strs[now] = '*'
	}
	fmt.Println(string(strs))
}
func main() {
	str := "**ab*b*c*a**"
	// fmt.Println(reflect.TypeOf(str[0]))
	// fmt.Println(str[0])
	// fmt.Println(string(str[0]))
	movestringrorf(str)
	movestring_two(str)
}
