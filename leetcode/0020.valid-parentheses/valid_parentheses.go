package leetcode

func IsValid(s string) bool {

	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	//建立map去存那三對主角

	parenMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, parenMap[c])
		} else if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}
