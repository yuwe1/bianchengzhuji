package main

import "fmt"

import "sort"

// 蛮力轮询
func StringContain(a string, b string) bool {
	ar := []rune(a)
	br := []rune(b)
	// 用短字符串去寻找长字符串进行比较，因此主循环用短字符串进行控制
	for i := 0; i < len(br); i++ {
		// 定义一个 j 用来每次循环都要保证从 0 开始,j 代表 ar的下表
		j := 0
		// 循环长字符串
		for j = 0; j < len(ar) && br[i] != ar[j]; j++ {

		}
		// 如果在遍历的过程中遇到不符合情况的字符，则上一个循环会比较到 ar 最后，可用 j 作为判断
		// 如果在循环的过程中找到相同字符，第二个循环会提前结束
		if j >= len(ar) {
			return false
		}
	}
	return true
}

// 排序后轮询
type RuneSclice []rune

// 使用官方的sort.Sort()需实现相关接口
func (p RuneSclice) Len() int {
	return len(p)
}
func (p RuneSclice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p RuneSclice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func StringSortContain(a string, b string) bool {
	ar := []rune(a)
	var ars RuneSclice
	ars = ar
	br := []rune(b)
	var brs RuneSclice
	brs = br
	// 排序
	sort.Sort(ars)
	sort.Sort(brs)
	// 定义两个指针分别指向ars和brs
	pa := 0
	pb := 0
	for pb < brs.Len() {
		// 移动pa指针寻找相同的字符，pb暂时不移动
		for (pa < ars.Len()) && (ars[pa] < brs[pb]) {
			pa += 1
		}
		// 说明pa遍历结束仍然未找到，存在不相同的字符，返回false
		if pa >= ars.Len() || ars[pa] > brs[pb] {
			return false
		}
		// 说明找到了相同的，开始寻找下一个，pb向后移动
		pb += 1
	}
	return true
}
func main() {
	fmt.Println(StringContain("ABCD", "BAE"))
	fmt.Println(StringSortContain("adhfsf", "f"))
}
