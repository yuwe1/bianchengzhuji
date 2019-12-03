/*
 * @Author: your name
 * @Date: 2019-12-02 20:45:31
 * @LastEditTime: 2019-12-03 19:58:41
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebf:\mayun\bianchengzhuji\tstring\main.go
 */
package main

import (
	"fmt"
)

func LeftShiftOne(str *string, n int) {
	// 保存第一个字符
	// 将字符传存放到rune数组
	// string转换成rune
	strs := []rune(*str)
	// 存放第一个字符
	f := strs[0]
	*str = ""
	// 循环数组
	for i := 1; i < n; i++ {
		strs[i-1] = strs[i]
	}
	strs[n-1] = f

	*str += string(strs[n-1])
	*str = string(strs)
}

func main() {
	str := "asdfghjkl"
	LeftShiftOne(&str, len(str))
	fmt.Println(str)
}
