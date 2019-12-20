package main

import "fmt"

func BuildMaxHeap(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		adjustDown(data, i)
	}
}
func adjustDown(data []int, i int) {

	for largest := 2*i + 1; largest < len(data); largest = 2*largest + 1 {
		if largest == len(data)-1 {
			if data[largest] > data[i] {
				// fmt.Println(data[largest], "和", data[i], "进行交换")
				data[largest], data[i] = data[i], data[largest]
				i = largest
			} else {
				// fmt.Println(data[largest], "不和", data[i], "进行交换")
			}
		} else {
			if data[largest+1] > data[largest] && data[largest+1] > data[i] {
				// fmt.Println(data[largest+1], "和", data[i], "进行交换")
				data[largest+1], data[i] = data[i], data[largest+1]
				i = largest
			} else if data[largest+1] < data[largest] && data[largest] > data[i] {
				// fmt.Println(data[largest], "和", data[i], "进行交换")
				data[largest], data[i] = data[i], data[largest]
				i = largest
			} else {
				// fmt.Println(data[largest], "或", data[largest+1], "不和", data[i], "进行交换")
			}
		}
	}
}

// 维护最大堆
func topK(data []int, k int) {
	// 建立前K个数的最大堆
	BuildMaxHeap(data[0:k])
	for i := k; i < len(data); i++ {
		if data[i] < data[0] { //如果剩余的数中有小的数则替换
			data[i], data[0] = data[0], data[i]
			adjustDown(data[0:k], 0)
		}
	}
}
func main() {
	data := []int{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
	// BuildMaxHeap(data)
	topK(data, 5)
	fmt.Println(data)
}
