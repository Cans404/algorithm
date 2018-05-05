// created by Cans, 20180414

package main

import "fmt"

func main() {
	data := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}
	// fmt.Println(data)
	// fmt.Println()

	for i := 0; i < len(data)-1; i++ {
		min := i

		for j := i; j < len(data)-1; j++ {
			if data[min] > data[j+1] {
				min = j + 1
			}
		}

		if min != i {
			data[i], data[min] = data[min], data[i]
		}

		// fmt.Println(data)
	}

        fmt.Println(data)
}
