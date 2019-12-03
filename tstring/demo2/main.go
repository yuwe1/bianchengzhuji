/*
 * @Author: your name
 * @Date: 2019-12-03 20:09:14
 * @LastEditTime: 2019-12-03 20:34:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebf:\mayun\bianchengzhuji\tstring\demo2\main.go
 */
package main

import "fmt"

// m代表前m个字符
func LeftShiftOne(str *string, n int, m int) {
	// 将字符串根据要求进行反转为两个部分
	ReverseStringOne(str, 0, m-1)
	ReverseStringOne(str, m, n-1)
	ReverseStringOne(str, 0, n-1)

}
func ReverseStringOne(str *string, from int, to int) {
	ru := []rune(*str)
	for from < to {
		c := ru[from]
		ru[from] = ru[to]
		from += 1
		ru[to] = c
		to -= 1
	}
	*str = string(ru)
}

func main() {
	str := "asdfghjkl"
	LeftShiftOne(&str, len(str), 3)
	fmt.Println(str)
}
