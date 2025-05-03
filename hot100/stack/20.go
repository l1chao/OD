package main

func isValid(s string) bool {
	stack := []byte{}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, byte(v))
		} else {
			if len(stack) == 0 {
				return false
			}

			if v == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if v == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if v == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
