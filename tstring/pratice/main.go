/*
 * @Author: your name
 * @Date: 2019-12-03 20:53:25
 * @LastEditTime: 2019-12-03 21:03:56
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebf:\mayun\bianchengzhuji\tstring\pratice\main.go
 */
package main

import "strings"

import "fmt"

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

func main() {
	str := "I am a student."

	ResverString(&str)
	fmt.Println(len(str))
}
