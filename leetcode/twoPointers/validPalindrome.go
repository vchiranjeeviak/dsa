package twopointers

import "unicode"

func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func isPalindrome1(s string) bool {
	str := []rune(s)
	newStr := []rune{}

	for _, v := range str {
		if isAlphaNumeric(v) {
			newStr = append(newStr, v)
		}
	}

	if len(newStr) < 2 {
		return true
	}

	for i := 0; i < len(newStr)/2; i++ {
		if unicode.ToLower(newStr[i]) != unicode.ToLower(newStr[len(newStr)-i-1]) {
			return false
		}
	}

	return true
}

func isPalindrome2(s string) bool {
	leftPointer, rightPointer := 0, len(s)-1
	str := []rune(s)

	for {
		if leftPointer >= rightPointer {
			break
		}
		for {
			if leftPointer < rightPointer && !isAlphaNumeric(str[leftPointer]) {
				leftPointer++
			} else {
				break
			}
		}
		for {
			if leftPointer < rightPointer && !isAlphaNumeric(str[rightPointer]) {
				rightPointer--
			} else {
				break
			}
		}
		if unicode.ToLower(str[leftPointer]) != unicode.ToLower(str[rightPointer]) {
			return false
		}
		leftPointer++
		rightPointer--
	}
	return true
}
