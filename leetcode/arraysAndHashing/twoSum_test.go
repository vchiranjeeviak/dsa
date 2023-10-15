package arraysandhashing

import (
	"reflect"
	"testing"
)

type twoSumTestCase struct {
	array  []int
	target int
	result []int
}

func twoSumGetTestCases(num int) twoSumTestCase {
	switch num {
	case 1:
		return twoSumTestCase{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		}
	case 2:
		return twoSumTestCase{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		}
	case 3:
		return twoSumTestCase{
			[]int{3, 3},
			6,
			[]int{0, 1},
		}
	}
	return twoSumTestCase{}
}

func TestTwoSum1(t *testing.T) {
	testCase := twoSumGetTestCases(1)
	actualResult := twoSum1(testCase.array, testCase.target)

	if !reflect.DeepEqual(actualResult, testCase.result) {
		t.Error("Test case 1 failed for approach 1.")
	}

	testCase = twoSumGetTestCases(2)
	actualResult = twoSum1(testCase.array, testCase.target)

	if !reflect.DeepEqual(actualResult, testCase.result) {
		t.Error("Test case 2 failed for approach 1.")
	}

	testCase = twoSumGetTestCases(3)
	actualResult = twoSum1(testCase.array, testCase.target)

	if !reflect.DeepEqual(actualResult, testCase.result) {
		t.Error("Test case 3 failed for approach 1.")
	}
}

func TestTwoSum2(t *testing.T) {
	testCase := twoSumGetTestCases(1)
	actualResult := twoSum2(testCase.array, testCase.target)

	if !reflect.DeepEqual(actualResult, testCase.result) {
		t.Error("Test case 1 failed for approach 2.")
	}

	testCase = twoSumGetTestCases(2)
	actualResult = twoSum2(testCase.array, testCase.target)

	if !reflect.DeepEqual(actualResult, testCase.result) {
		t.Error("Test case 2 failed for approach 2.")
	}

	testCase = twoSumGetTestCases(3)
	actualResult = twoSum2(testCase.array, testCase.target)

	if !reflect.DeepEqual(actualResult, testCase.result) {
		t.Error("Test case 3 failed for approach 2.")
	}
}
