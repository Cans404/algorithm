// created by Cans, 20180414

package main

import "fmt"

func main() {
	data := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}
	// fmt.Println(data)
	// fmt.Println()

	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}

		// fmt.Println(data)
	}

        fmt.Println(data)
}
