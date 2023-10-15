package arraysandhashing

import "testing"

type validAnagramTestCase struct {
	array  []string
	result bool
}

func validAnagramGetTestCase(num int) validAnagramTestCase {
	switch num {
	case 1:
		return validAnagramTestCase{
			[]string{"anagram", "nagaram"},
			true,
		}
	case 2:
		return validAnagramTestCase{
			[]string{"rat", "car"},
			false,
		}
	}
	return validAnagramTestCase{
		[]string{"", ""},
		true,
	}
}

func TestValidAnagram1(t *testing.T) {
	testCase1 := validAnagramGetTestCase(1)
	stringS := testCase1.array[0]
	stringT := testCase1.array[1]
	isValidAnagram := isAnagram1(stringS, stringT)

	if isValidAnagram != testCase1.result {
		t.Error("Test case 1 failed for approach 1.")
	}

	testCase2 := validAnagramGetTestCase(2)
	stringS = testCase2.array[0]
	stringT = testCase2.array[1]
	isValidAnagram = isAnagram1(stringS, stringT)

	if isValidAnagram != testCase2.result {
		t.Error("Test case 2 failed for approach 1.")
	}
}

func TestValidAnagram2(t *testing.T) {
	testCase1 := validAnagramGetTestCase(1)
	stringS := testCase1.array[0]
	stringT := testCase1.array[1]
	isValidAnagram := isAnagram2(stringS, stringT)

	if isValidAnagram != testCase1.result {
		t.Error("Test case 1 failed for approach 2.")
	}

	testCase2 := validAnagramGetTestCase(2)
	stringS = testCase2.array[0]
	stringT = testCase2.array[1]
	isValidAnagram = isAnagram2(stringS, stringT)

	if isValidAnagram != testCase2.result {
		t.Error("Test case 2 failed for approach 2.")
	}
}
