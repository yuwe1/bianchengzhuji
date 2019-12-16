package main

import "fmt"

// 常规解法
func LookPanlind_one(str string) string {
	// 找到该字符串的所有字串存放到字符串数组
	re := getAllSubstring(str)
	max := 0
	sign := 0
	// 遍历 数组
	for i := 0; i < len(re); i++ {
		if IsPalindrome_one(re[i]) {
			if max < len(re[i]) {
				max = len(re[i])
				sign = i
			}
		}
	}
	return re[sign]
}

// 获取所有的字串
func getAllSubstring(str string) (re []string) {
	for i := 0; i < len(str); i++ {
		j := len(str)
		for j = len(str); j > i; j-- {
			re = append(re, str[i:j])
		}
	}
	return re
}

// 判断是否是回文字符串
func IsPalindrome_one(str string) bool {
	// 非法输入
	if len(str) == 0 {
		return false
	}
	// 定义两个指针
	right := len(str) - 1
	left := 0
	for i := left; i < len(str)/2; i++ {
		if str[i] != str[right] {
			return false
		}
		right--
	}
	return true
}

// 中心扩展法
func LookPanlind_two(str string) (s string) {

	i, j, max, c := 0, 0, 0, 0
	for i = 0; i < len(str); i++ {
		// 回文长度为奇数
		for j = 0; (i+j) < len(str) && (i-j) >= 0; j++ {
			if str[i-j] != str[i+j] {
				break
			}

			c = 2*j + 1
		}
		if max < c {
			s = str[(i - j + 1):(i + j)]
			println("奇数", s)
			max = c
		}
		// 回文长度为偶数
		for j = 0; (i+j+1) < len(str) && (i-j) >= 0; j++ {
			if str[i-j] != str[i+j+1] {
				break
			}

			c = 2*j + 2
		}
		if max < c {
			s = str[(i - j + 1):(i + j + 1)]
			println("偶数", s)
			max = c
		}
	}
	return s
}

func main() {
	str := "12212321"
	// fmt.Println(str[0:0])
	// fmt.Println(str[0:1])
	// fmt.Println(str[0:3])
	// fmt.Println(str[2:2])
	// fmt.Println(LookPanlind_one(str))
	fmt.Println(LookPanlind_two(str))
}
