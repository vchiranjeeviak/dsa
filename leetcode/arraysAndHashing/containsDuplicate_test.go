package arraysandhashing

import "testing"

type containsDuplicateTestCase struct {
	array  []int
	result bool
}

func containsDuplicateGetTestCase(num int) containsDuplicateTestCase {
	switch num {
	case 1:
		return containsDuplicateTestCase{
			[]int{1, 2, 3, 1},
			true,
		}
	case 2:
		return containsDuplicateTestCase{
			[]int{1, 2, 3, 4},
			false,
		}
	case 3:
		return containsDuplicateTestCase{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		}
	}
	return containsDuplicateTestCase{
		[]int{1},
		false,
	}
}

func TestContainsDuplicate1(t *testing.T) {
	testCase1 := containsDuplicateGetTestCase(1)
	result := containsDuplicate1(testCase1.array)
	if result != testCase1.result {
		t.Error("Test case 1 failed for approach 1.")
	}

	testCase2 := containsDuplicateGetTestCase(2)
	result = containsDuplicate1(testCase2.array)
	if result != testCase2.result {
		t.Error("Test case 2 failed for approach 1.")
	}

	testCase3 := containsDuplicateGetTestCase(3)
	result = containsDuplicate1(testCase3.array)
	if result != testCase3.result {
		t.Error("Test case 3 failed for approach 1.")
	}
}

func TestContainsDuplicate2(t *testing.T) {
	testCase1 := containsDuplicateGetTestCase(1)
	result := containsDuplicate2(testCase1.array)
	if result != testCase1.result {
		t.Error("Test case 1 failed for approach 2.")
	}

	testCase2 := containsDuplicateGetTestCase(2)
	result = containsDuplicate2(testCase2.array)
	if result != testCase2.result {
		t.Error("Test case 2 failed for approach 2.")
	}

	testCase3 := containsDuplicateGetTestCase(3)
	result = containsDuplicate2(testCase3.array)
	if result != testCase3.result {
		t.Error("Test case 3 failed for approach 2.")
	}
}

func TestContainsDuplicate3(t *testing.T) {
	testCase1 := containsDuplicateGetTestCase(1)
	result := containsDuplicate3(testCase1.array)
	if result != testCase1.result {
		t.Error("Test case 1 failed for approach 3.")
	}

	testCase2 := containsDuplicateGetTestCase(2)
	result = containsDuplicate3(testCase2.array)
	if result != testCase2.result {
		t.Error("Test case 2 failed for approach 3.")
	}

	testCase3 := containsDuplicateGetTestCase(3)
	result = containsDuplicate3(testCase3.array)
	if result != testCase3.result {
		t.Error("Test case 3 failed for approach 3.")
	}
}
