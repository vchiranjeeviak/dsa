package arraysandhashing

import "sort"

// Leetcode 217
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Approach 1: We can loop through each item in the array and by pivoting through that point, we loop through the remaining elements of that array starting from that point. We will see if there is same element in the remaining array.
// Time complexity: O(n^2)
// Space complexity: O(1)
func containsDuplicate1(nums []int) bool {
	lenOfNums := len(nums)
	if lenOfNums == 1 {
		return false
	}
	for i := 0; i < lenOfNums-1; i++ {
		for j := i + 1; j < lenOfNums; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Approach 2: We can sort the array so that all the similar or duplicate items come together side by side. Now, we can loop through the array normally and see if any adjacent items are equal.
// Time complexity: O(n lg n)
// Space complexity: O(1) if used space optimised sorting algorithm
func containsDuplicate2(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// Approach 3: While iterating through the array, we can memorise the items and identify the duplicate when it appears depending on the memory. To use efficient memory, we use Hashmap.
// Time complexity: O(n)
// Space complexity: O(n)
func containsDuplicate3(nums []int) bool {
	numsMap := map[int]bool{}
	for _, val := range nums {
		if numsMap[val] {
			return true
		}
		numsMap[val] = true
	}
	return false
}
