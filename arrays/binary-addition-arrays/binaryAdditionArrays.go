package main

import "fmt"

func reverse(slc []int) []int {
	for i := 0; i < len(slc)/2; i++ {
		slc[i], slc[len(slc)-i-1] = slc[len(slc)-i-1], slc[i]
	}
	return slc
}

func binaryAdditionArrays(slc1 []int, slc2 []int) []int {
	len1 := len(slc1)
	len2 := len(slc2)

	res := []int{}

	if len1 > len2 {
		len1, len2 = len2, len1
		slc1, slc2 = slc2, slc1
	}

	carry := 0

	for i := 0; i < len2; i++ {
		sum := slc2[len2-i-1] + carry
		if len1-i > 0 {
			sum += slc1[len1-i-1]
		}
		if sum > 1 {
			carry = 1
			sum = sum % 2
		}
		res = append(res, sum)
	}
	res = append(res, carry)

	return reverse(res)
}

func main() {
	slc1 := []int{1, 1, 1}
	slc2 := []int{1, 0, 1}
	fmt.Println(binaryAdditionArrays(slc1, slc2))
}
