package main

import "fmt"

var res [100]int
var k = 0

func Select(a []int, n int, sum int) {
	if n <= 0 || sum <= 0 {
		return
	}
	if k > 0 {
		if sum == a[n-1] {
			for i := k - 1; i >= 0; i -= 1 {
				fmt.Print(res[i], "+")
			}
			fmt.Println(a[n-1])
		}

	}

	res[k] = a[n-1]
	k++
	Select(a, n-1, sum-a[n-1])
	k--
	Select(a, n-1, sum)
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	Select(data, 10, 15)
}
