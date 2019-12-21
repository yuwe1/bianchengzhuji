package main

import "fmt"

// 利用去全排序进行寻找最小的k个数
func FindNumBySort(data []int, k int) (result []int) {
	if len(data) != 0 {
		// 对数组进行快速排序
		QuickSelect(data, 0, len(data)-1)
		// 返回前k个数
		return data[0:k]
	}
	return result
}
func QuickSelect(data []int, left, right int) {
	// 指向基准数
	p := left
	// i,j分别代表了需要进行一趟排 序的数组的首尾
	i := left
	j := right
	// 定义我们的基准数，每次将数组的第一个值当作基准数比较
	key := data[i]
	for i <= j {
		for key <= data[j] && j >= p {
			j--
		}
		if p <= j {
			data[i], data[j] = data[j], data[i]
			p = j
		}
		for i <= p && key >= data[i] {
			i++
		}
		if i <= p {
			data[i], data[j] = data[j], data[i]
			p = i
		}
	}
	// fmt.Println(data)
	if p-left > 1 {
		QuickSelect(data, left, p-1)
	}
	if right-p > 1 {
		QuickSelect(data, p+1, right)
	}

}

// 利用部分排序寻找最小的k个数
func FindNumByPartSort(data []int, k int) (result []int) {
	// 默认前k个数为最小
	result = data[0:k]

	for i := k; i < len(data); i++ {
		// 找到最大值的坐标
		max := select_sort(result)
		if result[max] > data[i] {
			result[max], data[i] = data[i], result[max]
		}
	}

	return result
}
func select_sort(data []int) (max int) {
	max = 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[max] {
			max = i
		}
	}
	return max
}

// 用堆代替数组

func main() {
	data := []int{5, 3, 7, 6, 4, 1, 0, 2, 9, 10, 8}
	QuickSelect(data, 0, 10)
	fmt.Println(data)
	// fmt.Println(FindNumBySort(data, 5))
	// fmt.Println(select_sort(data))
	// fmt.Println(FindNumByPartSort(data, 5))

}
