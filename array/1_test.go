package ds

import "testing"

/*
1. Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

	2 <= nums.length <= 104
	-109 <= nums[i] <= 109
	-109 <= target <= 109
	Only one valid answer exists.
*/

func twoSum(nums []int, target int) []int {
	// 2,7,11,15
	a := []int{0,0}
	for i := 0; i < len(nums); i ++ {
		for j := 1; j < len(nums); j ++ {
			x := nums[i] + nums[j]
			if (x == target && i != j) {
				a[0] = i
				a[1] = j
			}
		}
	}
	return a
}

func TestTwoSum(t *testing.T) {
	a1 := twoSum([]int{2,7,11,15}, 9)
	if a1[0] != 0 || a1[1] != 1 {
		t.Fatalf("expected [0 1] but got %v", a1)
	}

	a2 := twoSum([]int{3,2,4}, 6)
	if a2[0] != 1 || a2[1] != 2 {
		t.Fatalf("expected [1 2] but got %v", a2)
	}

	a3 := twoSum([]int{3,3}, 6)
	if a3[0] != 0 || a3[1] != 1 {
		t.Fatalf("expected [0 1] but got %v", a3)
	}
}
