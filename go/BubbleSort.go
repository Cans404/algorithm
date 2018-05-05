// created by Cans, 20180413

package main

import "fmt"

func main() {
	data := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}
	// fmt.Println(data)
	// fmt.Println()

	for i := 0; i < len(data); i++ {
		var counter int

		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				counter++
			}

			// fmt.Println(counter)
		}

		if counter == 0 {
			break
		}

		// fmt.Println(data)
	}

        fmt.Println(data)
}
