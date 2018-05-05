// created by Cans, 20180418

package main

import (
	"fmt"
)

// 处理枢纽值
func dealPivot(arr []int, left, right int) int {
	mid := (left + right) / 2

	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}

	if arr[mid] > arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}

	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}

	arr[mid], arr[right-1] = arr[right-1], arr[mid]

	return right - 1
}

func quickSort(arr []int, start, end int) {
	gap := end - start

	if gap == 1 {
		if arr[start] > arr[end] {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}

	if gap >= 2 {
		key := dealPivot(arr, start, end)
		i, j := start, end-2

		for i <= j {
			// i 和 j 彼此靠近
			for arr[i] < arr[key] {
				i++
			}
			for arr[j] > arr[key] {
				j--
			}
			// 无法再靠近时互换
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		// 枢纽值分割出2个区间
		arr[i], arr[key] = arr[key], arr[i]
		i++

		// fmt.Println(arr, i, j , start, end)
		// fmt.Println(arr[i:end+1], arr[start:j+1])

		// 对子区间递归处理
		if i < end {
			quickSort(arr, i, end)
		}

		if start < j {
			quickSort(arr, start, j)
		}
	}
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
