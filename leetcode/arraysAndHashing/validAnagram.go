package arraysandhashing

import "sort"

func isAnagram1(s string, t string) bool {
	// if length of both strings are not same, they cant be anagrams
	if len(s) != len(t) {
		return false
	}

	// convert strings into rune arrays to sort
	arrayS := []rune(s)
	arrayT := []rune(t)

	// sort the arrays
	sort.Slice(arrayS, func(i, j int) bool {
		return arrayS[i] < arrayS[j]
	})
	sort.Slice(arrayT, func(i, j int) bool {
		return arrayT[i] < arrayT[j]
	})

	// check if any pair of items at same index are mismatching
	for i := 0; i < len(s); i++ {
		if arrayS[i] != arrayT[i] {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	frequencyMap := map[rune]int{}

	// Increment the frequency of a letter if found in string s
	for _, v := range s {
		frequencyMap[v]++
	}

	// Decrement the frequency of a letter if found in string t
	for _, v := range t {
		frequencyMap[v]--
	}

	// It is not anagram if frequency of any letter is not 0
	for _, v := range frequencyMap {
		if v != 0 {
			return false
		}
	}

	return true
}
