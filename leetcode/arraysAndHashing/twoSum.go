package arraysandhashing

// Leetcode 1
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// Approach 1: While looping through the array, select an element and loop through the remaining items in the array and see if the same item appears again.
// Time complexity: O(n^2)
// Space complexity: O(1)
func twoSum1(nums []int, target int) []int {
	lenOfNums := len(nums)

	for i := 0; i < lenOfNums-1; i++ {
		for j := i + 1; j < lenOfNums; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

// Approach 2: Loop through the array and see if the difference between the target and current item is already there in the map. If already there, we got the result, else add the current item in the map.
// Time complexity: O(n)
// Space complexity: O(n)
func twoSum2(nums []int, target int) []int {
	trackingMap := map[int]int{}

	for i, v := range nums {
		indexOfDifference := trackingMap[target-v]
		if indexOfDifference > 0 {
			return []int{indexOfDifference - 1, i}
		}
		trackingMap[v] = i + 1
	}
	return []int{-1, -1}
}
