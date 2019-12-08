package main

import "fmt"

// 第一个步骤，找到排列中可以升序的首位置 a[i]
// 找到在位置i+i 到 n的比a[i]大的最后一个位置 b[j]
// 交换
// i+1 到 n进行反转
func CalAllPermutation(str string) string {

	// 记录首位置
	i := 0
	// 记录末位置
	j := 0
	for i = len(str) - 2; i >= 0; i-- {
		if str[i] <= str[i+1] {
			break
		}
	}
	if i < 0 {
		return str
	}
	//
	for j = len(str) - 1; j > i; j-- {
		if str[i] < str[j] {
			break
		}
	}
	str1 := []rune(str)
	// 交换
	str1[i], str1[j] = str1[j], str1[i]
	// 赋值
	str = string(str1)
	// 固定前面的字符串
	s := str[:(i + 1)]
	// 反转后面的字符串进行拼接
	s += ResverString(str[(i + 1):])
	return s

}
func ResverString(str string) string {
	r := []rune(str)
	j := len(r) - 1
	for i := 0; i < len(r)/2; i++ {
		r[i], r[j] = r[j], r[i]
		j--
	}
	return string(r)
}
func main() {
	i := 0
	s := "21345"
	for i < 90 {
		s = CalAllPermutation(s)
		fmt.Println(s)
		i++
	}
	// CalAllPermutation("21543")
	// fmt.Println(ResverString("abcd"))
	// fmt.Println(ResverString("a"))
	// fmt.Println(ResverString(""))
}
