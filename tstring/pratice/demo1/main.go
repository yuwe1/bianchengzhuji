package main

import (
	"fmt"
	"strings"
)

// 对字符串按照空格进行切割,存放到数组中，对数组进行反转
func ResverString(str *string) {
	strs := strings.Split(*str, " ")
	i := 0
	j := len(strs) - 1
	for i < j {
		temp := strs[i]
		strs[i] = strs[j]
		i += 1
		strs[j] = temp
		j -= 1
	}
	*str = strings.Join(strs, " ")
}
func ResverString_one(str string) {
	strs := []rune(str)
	i := 0
	j := len(strs) - 1
	for i < j {
		temp := strs[i]
		strs[i] = strs[j]
		i++
		strs[j] = temp
		j--
	}
	fmt.Println(string(strs))
}
func main() {
	ResverString_one("July")
	// str := "I am a student."

	// ResverString(&str)
	// fmt.Println(len(str))
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// cou(a, 2)
}
