package main

import "fmt"

func decreasingInsertionSort(slc []int) {
	for index, ele := range slc {
		// Avoiding the first item
		if index != 0 {
			key := ele
			j := index - 1
			// Loop until the previous index goes out of index and key is greater than previous
			for (j >= 0) && (slc[j] < key) {
				// Keep switching until loop ends
				slc[j+1] = slc[j]
				j--
			}
			// Place the key in the right position
			slc[j+1] = key
		}
	}
}

func main() {
	slc := []int{1, 4, 7, 2, 9}
	decreasingInsertionSort(slc)
	fmt.Println(slc)
}
