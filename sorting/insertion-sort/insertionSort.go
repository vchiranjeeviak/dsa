package main

import "fmt"

// Keep cards placing example in mind
// How you place cards in sorted order in one hand while taking each one from the pile
func insertSort(slc []int) {
	// Using for - range loop to avoid array index out of bound errors
	for index, ele := range slc {
		key := ele
		// Avoiding the first item
		if index != 0 {
			j := index - 1
			// Loop until the previous index goes out of index and key is smaller than previous
			for (j >= 0) && (key < slc[j]) {
				// Keep switching until loop ends
				slc[j+1] = slc[j]
				j--
			}
			// Place the key in the right position
			slc[j+1] = ele
		}
	}
}

func main() {
	slc := []int{1, 7, 2, 4, 9}
	insertSort(slc)
	fmt.Println(slc)
}
