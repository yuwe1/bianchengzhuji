//背包问题
package main

import (
	"fmt"
)

// 01背包问题
// 假设有N件物品，容量为V
var (
	N int
	V int
	// 只看前i个物品，总体积为j的情况写，总价值为多少
	f [][1024]int
	// 第i个物品的体积
	w []int
	// 第i个物品的价值
	v     []int
	f_two []int
)
var f1 [][]int

type T []int

func Init() {
	fmt.Scanf("%d %d\n", &N, &V)
	// 初始化容量
	v = make([]int, N+1)
	w = make([]int, N+1)
	f_two = make([]int, 1024)
	f = make([][1024]int, N+1)
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func getMaxValue() {
	//对每个物品进行赋值体积和价值
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d %d\n", &w[i], &v[i])
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= V; j++ {
			// 判断是否太重
			f[i][j] = f[i-1][j]
			if j >= w[i] {
				f[i][j] = max(f[i][j], f[i-1][j-w[i]]+v[i])
			}

		}
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= V; j++ {
			fmt.Print(f[i][j], "  ")
		}
		fmt.Println()
	}
	fmt.Println(f[N][V])
}
func getMaxValue_two() {
	for i := 1; i <= N; i++ {
		for j := V; j >= w[i]; j-- {
			fmt.Println(j, "  ", j-w[i])
			f_two[j] = max(f_two[j], f_two[j-w[i]]+v[i])
		}
	}

	fmt.Println(f_two[V])
}
func main() {
	Init()
	getMaxValue()
	getMaxValue_two()
}
