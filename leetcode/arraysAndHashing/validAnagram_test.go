package arraysandhashing

import "testing"

func validAnagramGetTestCase(num int) []string {
    switch num {
    case 1:
        return []string{"anagram", "nagaram"}
    case 2:
        return []string{"rat", "car"}
    }
    return []string{"", ""}
}

func TestValidAnagram1(t *testing.T) {
    stringS := validAnagramGetTestCase(1)[0]
    stringT := validAnagramGetTestCase(1)[1]
    isValidAnagram := isAnagram1(stringS, stringT)

    if !isValidAnagram {
        t.Error("Test case 1 failed for approach 1.")
    }

    stringS = validAnagramGetTestCase(2)[0]
    stringT = validAnagramGetTestCase(2)[1]
    isValidAnagram = isAnagram1(stringS, stringT)

    if isValidAnagram {
        t.Error("Test case 2 failed for approach 1.")
    }
}

func TestValidAnagram22(t *testing.T) {
    stringS := validAnagramGetTestCase(1)[0]
    stringT := validAnagramGetTestCase(1)[1]
    isValidAnagram := isAnagram2(stringS, stringT)

    if !isValidAnagram {
        t.Error("Test case 1 failed for approach 2.")
    }

    stringS = validAnagramGetTestCase(2)[0]
    stringT = validAnagramGetTestCase(2)[1]
    isValidAnagram = isAnagram2(stringS, stringT)

    if isValidAnagram {
        t.Error("Test case 2 failed for approach 2.")
    }
}
