package stacks

// Leetcode 22
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// Approach 1: Add each item if the current item is an opening bracket or there is no item already in stack or the last element is not the matching closing one of the current one. Pop if the latest one is the opening bracket of current closed one.
// Time complexity: O(n)
// Space complexity: O(n)
func validParanthesis1(s string) bool {
	if len(s) < 2 {
		return false
	}

	combinations := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' || len(stack) == 0 || stack[len(stack)-1] != combinations[v] {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
