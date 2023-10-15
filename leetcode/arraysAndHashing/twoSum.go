package arraysandhashing

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
