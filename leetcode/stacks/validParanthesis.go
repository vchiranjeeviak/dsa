package stacks

func validParanthesis1(s string) bool {
    if len(s) < 2 {
        return false
    }

    combinations := map[rune]rune{')': '(', ']': '[', '}': '{'}
    stack := []rune{}

    for  _, v := range s {
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
