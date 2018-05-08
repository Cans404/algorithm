// created by Cans, 20180508

package main

import (
	"fmt"
)

func insertSort(data []*int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && *data[j] < *data[j-1]; j-- {
			*data[j-1], *data[j] = *data[j], *data[j-1]
		}
	}
}

func shellSort(arr []int, gap int) {
	for gap >= 1 {
		for gid := 0; gid < gap; gid++ {
			var grp []int
			var pgrp []*int

			for i := gid; i < len(arr); i += gap {
				grp = append(grp, arr[i])
				pgrp = append(pgrp, &arr[i])
			}

			// debug 打印分组
			// fmt.Println(grp)

			insertSort(pgrp)
		}

		gap = gap / 2

		//fmt.Println()
	}

	fmt.Println(arr)
}

func main() {
	arr := []int{8, 9, 1, 7, 3, 2, 1, 5, 4, 6, 0}
	shellSort(arr, len(arr)/2)
}
