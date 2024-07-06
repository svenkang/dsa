package main

import (
	"fmt"
	"reflect"
)

/**
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days
you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
[1,1,4,]
3
					   0  1  2  3  4  5  6  7
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]

Input: temperatures = []
Output: []

Input: temperatures = [10]
Output: [0]

Input: temperatures = [10, 9, 8, 7, 11]
Output: [4,3,2,1,0]
*/


func main() {
	q := []int{73,74,75,71,69,72,76,73}
	r := dailyTemperatures(q)
	a := []int{1,1,4,2,1,1,0,0}
	if !reflect.DeepEqual(r, a) {
		fmt.Printf("FAIL: Expected %v, Actual: %v", a, r)
	}
	return
}

// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]


func dailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	stack := []int{}

	for i, t := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < t {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}

	return result
}
