package twopointers

import "testing"

type validPalindromeTestCase struct {
	str    string
	result bool
}

func getValidPalindromeTestCase(num int) validPalindromeTestCase {
	switch num {
	case 1:
		return validPalindromeTestCase{
			"A man, a plan, a canal: Panama",
			true,
		}
	case 2:
		return validPalindromeTestCase{
			"race a car",
			false,
		}
	case 3:
		return validPalindromeTestCase{
			" ",
			true,
		}
	}
	return validPalindromeTestCase{
		"",
		true,
	}
}

func TestValidPalindrome1(t *testing.T) {
	for i := 0; i < 3; i++ {
		testCase := getValidPalindromeTestCase(i + 1)
		if isPalindrome1(testCase.str) != testCase.result {
			t.Errorf("Testcase %d failed for approach 1.", i+1)
		}
	}
}

func TestValidPalindrome2(t *testing.T) {
	for i := 0; i < 3; i++ {
		testCase := getValidPalindromeTestCase(i + 1)
		if isPalindrome2(testCase.str) != testCase.result {
			t.Errorf("Testcase %d failed for approach 2.", i+1)
		}
	}
}
