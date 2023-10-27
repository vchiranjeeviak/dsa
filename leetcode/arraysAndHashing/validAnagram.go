package arraysandhashing

import "sort"

// Leetcode 242
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// Approach 1: Sort the both the strings, compare each letter one by one and return false if found even one non-matching letter pair.
// Time complexity: O(n lg n)
// Space complexity: O(1) if used space optimised sorting algorithm
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

// Appraoch 2: Two strings are anagrams if there are same set of letters in both strings with same frequency. Increment the frequency of each letter every time it is appeared in one string and decrement the same if found in other string. If frequency of all letters remain 0 after traversing through two string, then they are anagrams.
// Time complexity: O(n)
// Space complexity: O(n)
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
