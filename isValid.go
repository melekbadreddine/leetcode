func isValid(s string) bool {
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    stack := []rune{}

    for _, char := range s {
        if opening, isClosing := pairs[char]; isClosing {
            if len(stack) == 0 || stack[len(stack)-1] != opening {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, char)
        }
    }

    return len(stack) == 0
}
