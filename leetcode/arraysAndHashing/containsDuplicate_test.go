package arraysandhashing

import "testing"

func getTestCase(num int) []int {
    switch num {
    case 1:
        return []int{1,2,3,1}
    case 2:
        return []int{1,2,3,4}
    case 3:
        return []int{1,1,1,3,3,4,3,2,4,2}
    }
    return []int{1}
}

func TestContainsDuplicate1(t *testing.T) {
    result := containsDuplicate1(getTestCase(1))
    if result != true {
        t.Error("Test case 1 failed for approach 1.")
    }

    result = containsDuplicate1(getTestCase(2))
    if result != false {
        t.Error("Test case 2 failed for approach 1.")
    }

    result = containsDuplicate1(getTestCase(3))
    if result != true {
        t.Error("Test case 3 failed for approach 1.")
    }
}

func TestContainsDuplicate2(t *testing.T) {
    result := containsDuplicate2(getTestCase(1))
    if result != true {
        t.Error("Test case 1 failed for approach 2.")
    }

    result = containsDuplicate2(getTestCase(2))
    if result != false {
        t.Error("Test case 2 failed for approach 2.")
    }

    result = containsDuplicate2(getTestCase(3))
    if result != true {
        t.Error("Test case 3 failed for approach 2.")
    }
}

func TestContainsDuplicate3(t *testing.T) {
    result := containsDuplicate3(getTestCase(1))
    if result != true {
        t.Error("Test case 1 failed for approach 3.")
    }

    result = containsDuplicate3(getTestCase(2))
    if result != false {
        t.Error("Test case 2 failed for approach 3.")
    }

    result = containsDuplicate3(getTestCase(3))
    if result != true {
        t.Error("Test case 3 failed for approach 3.")
    }
}
