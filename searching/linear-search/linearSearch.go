package main

import "fmt"

func linearSearch(slc []int, target int) int {
	for index, ele := range slc {
		if ele == target {
			return index
		}
	}

	return -1
}

func main() {
	slc := []int{3, 4, 1, 6, 9, 5}
	fmt.Println(linearSearch(slc, 6))
}
