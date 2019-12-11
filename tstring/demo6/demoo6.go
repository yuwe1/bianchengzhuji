package main

import "fmt"

func strToIntOne(str string) int32 {
	if len(str) == 0 {
		return 0
	}
	// 转换成字节数组
	r := []rune(str)
	//代表最终的结果，刚开始为0
	n := int32(0)
	for i := 0; i < len(r); i++ {
		// 表示当前的数字
		c := int32(r[i] - '0')
		n = n*10 + c
	}
	return n
}

func strToIntTwo(str string) int32 {
	if len(str) == 0 {
		return 0
	}
	r := []rune(str)
	n := int32(0)
	Max_int := int32(1<<31 - 1) //math.MaxInt32
	Min_int := int32(-1 << 31)  //math.MinInt32
	if str[0] == '-' {
		for i := 1; i < len(r); i++ {
			c := int32(r[i] - '0')

			if c-1 > (Max_int - n*10) {
				n = Min_int
				break
			}
			n = n*10 + c
		}
	}

	return n
}

func main() {
	a := "-2147483649"
	// fmt.Println(strToIntOne(a))
	fmt.Println(strToIntTwo(a))
}
