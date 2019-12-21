package main

import (
	"fmt"
)

// 解法一：散列映射
func SelectNum(data []int, sum int) [][]int {
	// 构建一个空间为n的散列表即map,bool值用来标记是否已经被使用，下次不需要再进行使用
	var m map[int]bool
	m = make(map[int]bool, len(data))
	// 定义一个存放结果的散列
	var result [][]int

	for i := 0; i < len(data); i++ {
		m[data[i]] = false
	}
	// 循环遍历我们的数组
	for i := 0; i < len(data); i++ {

		if _, ok := m[sum-data[i]]; ok && m[data[i]] == false {
			// 定义一个临时数组
			temp := []int{data[i], sum - data[i]}
			// append函数会自动扩充数组
			result = append(result, temp)
			m[sum-data[i]] = true
		}
	}
	return result
}

// 解法二：排序夹逼
func SelectNum_Two(data []int, sum int) [][]int {
	var result [][]int
	// 先排序数组
	Qiuck_Sort(data, 0, len(data)-1)
	// 定义两个前后指针指向数组的首和尾
	begin, end := 0, len(data)-1
	for begin < end {
		if data[begin]+data[end] > sum {
			// 缩小范围
			end--
		} else if data[begin]+data[end] < sum {
			begin++
		} else {
			temp := []int{data[begin], data[end]}
			result = append(result, temp)
			// 继续缩小范围进行寻找
			begin++
			end--
		}
	}

	return result
}

// 快速排序对数组进行排序
func Qiuck_Sort(data []int, left, right int) {
	// 定义一个p指向基准数
	p := left
	// 定义i，j分别指向首和尾
	i, j := left, right
	key := data[i]
	for i <= j {
		for key <= data[j] && p <= j {
			j--
		}
		if p <= j {
			data[i], data[j] = data[j], data[i]
			p = j
		}
		for key >= data[i] && p >= i {
			i++
		}
		if p >= i {
			data[i], data[j] = data[j], data[i]
			p = i
		}
	}
	if p-left > 1 {
		Qiuck_Sort(data, left, p-1)
	}
	if right-p > 1 {
		Qiuck_Sort(data, p+1, right)
	}

}
func main() {
	data := []int{5, 3, 7, 6, 4, 1, 0, 2, 9, 10, 8}
	fmt.Println(SelectNum(data, 5))
	// Qiuck_Sort(data, 0, len(data)-1)
	// fmt.Println(data)
	fmt.Println(SelectNum_Two(data, 5))
}
