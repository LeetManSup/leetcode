'''
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
'''

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        def add(a, b: int) -> int:
            return b + a
        def sub(a, b: int) -> int:
            return b - a
        def mul(a, b: int) -> int:
            return b * a
        def div(a, b: int) -> int:
            return int(b / a)   # using '//' is wrong, there's no truncating toward zero for negatives.
        ops = {"+": add, "-": sub, "*": mul, "/": div}
        stack = []
        for token in tokens:
            try:
                stack.append(int(token))
            except:
                a, b = stack.pop(), stack.pop()
                stack.append(ops[token](a, b))
        return stack[0]
