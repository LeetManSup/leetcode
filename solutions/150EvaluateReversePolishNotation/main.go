/*
You are given an array of strings tokens that represents
an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents
the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward
zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a
reverse polish notation.
The answer and all the intermediate calculations can be
represented in a 32-bit integer.
*/

func evalRPN(tokens []string) int {
	stack := []int{}
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return b + a },
		"-": func(a, b int) int { return b - a },
		"*": func(a, b int) int { return b * a },
		"/": func(a, b int) int { return b / a },
	}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
			continue
		}
		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, ops[token](a, b))
	}
	return stack[0]
}