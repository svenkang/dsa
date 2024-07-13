package main

import (
	"fmt"
	"math"
)

/**
You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.
Return the score of s.

Example 1:
Input: s = "hello"
Output: 13
Explanation:
The ASCII values of the characters in s are: 'h' = 104, 'e' = 101, 'l' = 108, 'o' = 111. So, the score of s would be |104 - 101| + |101 - 108| + |108 - 108| + |108 - 111| = 3 + 7 + 0 + 3 = 13.
Example 2:
Input: s = "zaz"
z=122
a=97
25
|z-a|+|a-z| = |122-97|+|97-122| = 25+25 = 50
Output: 50
Explanation:
The ASCII values of the characters in s are: 'z' = 122, 'a' = 97. So, the score of s would be |122 - 97| + |97 - 122| = 25 + 25 = 50.
*/
func main() {
	// 1. go through examples

	// 2. make more examples with edge cases
		// what if the string is empty=0
		// what if string is even/odd = fine
		// what if string is all same = 0
		// string must have ascii values

	// 3. think through the problem and solution at high level
		// looop through the string
		// if next is there compute, otherwise return total
		// for compute, get curr ascii code and next ascii code do abs diff
		// save the result to total

	// 4. code the solution
	input := "hello"
	// len 3
	// stop at 1

	s := []rune(input)
	var total float64
	total = 0
	for i, r := range s {
		if i < len(s) - 1 {
			total = total + math.Abs(float64(int(r) - int(s[i+1])))
		}
	}
	fmt.Println(int(total))
	
	// 5. walk through the solution
	// 6. clean up the code
}
