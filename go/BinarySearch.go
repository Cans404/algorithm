// created by Cans, 20180506

package main

import (
	"fmt"
)

func binarySearch(arr []int, low, high, key int) int {
	if arr[low] > key || arr[high] < key {
		return -1
	}

	// 逐步缩小范围
	for high-low > 1 {
		mid := (high + low) / 2

		// fmt.Println(low, mid, high)

		switch true {
		case key > arr[mid]:
			low = mid
			continue
		case key < arr[mid]:
			high = mid
			continue
		case key == arr[mid]:
			return mid
		}
	}

	// 缩小到只剩2个数时逐个排除
	switch key {
	case arr[low]:
		return low
	case arr[high]:
		return high
	default:
		return -1
	}
}

func main() {
	key := 99
	arr := []int{4, 13, 22, 22, 23, 36, 89, 227}
	index := binarySearch(arr, 0, len(arr)-1, key)
	if index != -1 {
		fmt.Printf("%d found at index: %d\n", key, index)
	} else {
		fmt.Printf("%d not found\n", key)
	}
}
