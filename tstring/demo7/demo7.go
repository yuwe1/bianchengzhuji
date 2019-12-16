package main

import "fmt"

// 两头往中间扫
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

// 从中间往两头进行扫描比较
func IsPalindrome_two(str string) bool {
	if len(str) == 0 {
		return false
	}
	right := 0
	left := 0
	// 判断字符串的长度尾偶数or奇数
	if len(str)%2 == 0 {
		right = (len(str) / 2) - 1
		left = (len(str) / 2)
	} else {
		right = (len(str) / 2) - 1
		left = (len(str) / 2) + 1
	}
	// 遍历字符串
	for {
		if right == -1 {
			break
		}
		if str[right] != str[left] {
			return false
		}
		right--
		left++
	}
	return true
}
func main() {
	// fmt.Println(IsPalindrome_one("我爱我"))
	fmt.Println(IsPalindrome_one("aba"))
	fmt.Println(IsPalindrome_one("madam"))
	fmt.Println(IsPalindrome_two("ada"))
	fmt.Println(IsPalindrome_two("madam"))
}
