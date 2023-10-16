package stacks

import "testing"

type validParanthesisTestCase struct {
    input string
    result bool
}

func validParanthesisGetTestCases(num int) validParanthesisTestCase {
    switch num {
    case 1:
        return validParanthesisTestCase{
            input: "()",
            result: true,
        }
    case 2:
        return validParanthesisTestCase{
            input: "()[]{}",
            result: true,
        }
    case 3:
        return validParanthesisTestCase{
            input: "](",
            result: false,
        }
    }
    return validParanthesisTestCase{
        input: "",
        result: false,
    }
}

func TestValidParanthesis1(t *testing.T) {
    testCase1 := validParanthesisGetTestCases(1)
    isValid1 := validParanthesis1(testCase1.input)

    if isValid1 != testCase1.result {
        t.Error("Test case 1 failed for approach 1.")
    }

    testCase2 := validParanthesisGetTestCases(2)
    isValid2 := validParanthesis1(testCase2.input)

    if isValid2 != testCase2.result {
        t.Error("Test case 2 failed for approach 1.")
    }

    testCase3 := validParanthesisGetTestCases(3)
    isValid3 := validParanthesis1(testCase3.input)

    if isValid3 != testCase3.result {
        t.Error("Test case 3 failed for approach 1.")
    }
}
