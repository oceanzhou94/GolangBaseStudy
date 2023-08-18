package main

import "fmt"

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 将数组拆分成两部分
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	// 递归地对左右两部分进行排序
	left = mergeSort(left)
	right = mergeSort(right)

	// 合并左右两部分
	return merge(left, right)
}

// 合并两个有序数组
func merge(left []int, right []int) []int {
	result := make([]int, 0)

	// 取出两个数组的最小值，加入结果数组中
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// 将剩余的元素加入结果数组中
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}

// 希尔排序
func shellSort(arr []int) []int {
	n := len(arr)
	gap := n / 2

	// 缩小增量排序
	for gap > 0 {
		for i := gap; i < n; i++ {
			j := i
			for j >= gap && arr[j-gap] > arr[j] {
				arr[j-gap], arr[j] = arr[j], arr[j-gap]
				j -= gap
			}
		}
		gap /= 2
	}

	return arr
}
func main() {
	arr := []int{10, 5, 2, 6, 3, 8, 7, 1, 4, 9}
	fmt.Println("原始数组：", arr)

	arr = mergeSort(arr)
	fmt.Println("归并排序后：", arr)
}
