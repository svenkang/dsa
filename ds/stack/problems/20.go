package main

import "fmt"
import "github.com/svenkang/dsa/ds/stack"
/**

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.
Example 1:
Input: s = "()"
Output: true
Example 2:
Input: s = "()[]{}"
Output: true
Example 3:
Input: s = "(]"
Output: false

Constraints:
    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.
*/

func main() {
	if !isValid("()") {
		fmt.Println("FAIL: case 1: should be valid")
	} else {
		fmt.Println("PASS: case 1: should be valid")
	}
	if !isValid("()[]{}") {
		fmt.Println("FAIL: case 2: should be valid")
	} else {
		fmt.Println("PASS: case 2: should be valid")
	}
	if isValid("(]") {
		fmt.Println("FAIL: case 3: should not be valid")
	} else {
		fmt.Println("PASS: case 3: should not be valid")
	}
	if !isValid("([])") {
		fmt.Println("FAIL: case 4: should be valid")
	} else {
		fmt.Println("PASS: case 4: should be valid")
	}
	if isValid("([{]])") {
		fmt.Println("FAIL: case 5: should be invalid")
	} else {
		fmt.Println("PASS: case 5: should be invalid")
	}
	return 
}

/** implement this function **/
func isValid(s string) bool {
	if len(s) == 0 || len(s) % 2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := stack.Stack[rune]{}
	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack.Push(r)
		} else if v, err := stack.Peek(); err == nil && r != pairs[v] {
			return false
		} else {
			stack.Pop()
		}
	}
	return stack.Len() == 0
}
