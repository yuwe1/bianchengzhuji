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
func cou(a []int, n int) {
	if n < 0 {
		return
	}
	fmt.Println("执行次数", a[n])
	if n > 1 {
		cou(a, n-1)
	}
}

func bsearch1(a []int, e int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if a[mid] == e {
		return mid
	} else if a[mid] > e {
		return bsearch1(a, e, left, mid-1)
	} else {
		return bsearch1(a, e, mid+1, right)
	}
	return -1
}
func main() {
	// str := "I am a student."

	// ResverString(&str)
	// fmt.Println(len(str))
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// cou(a, 2)
}
