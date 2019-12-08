package main

import "fmt"

func CalcALLPermutation(str *string, from int, to int) {
	strr := []rune(*str)
	if to <= 1 {
		return
	}
	if from == to {
		fmt.Println("最终结果：", string(strr))
	} else {
		for j := from; j <= to; j++ {
			strr[j], strr[from] = strr[from], strr[j]
			*str = string(strr)
			CalcALLPermutation(str, from+1, to)
		}
	}
}

func main() {
	str := "abcd"
	CalcALLPermutation(&str, 0, len(str)-1)
}
