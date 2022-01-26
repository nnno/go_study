/**
 * % go test -run TwoSum -v ./leetcode
 */
package main

import "testing"

func TestTwoSumSuccess(t *testing.T) {
	execTwoSum(t, []int{2, 7, 11, 15}, 9, []int{0, 1})
	execTwoSum(t, []int{3, 2, 4}, 6, []int{1, 2})
}

func execTwoSum(t *testing.T, nums []int, target int, correct []int) {
	result := twoSum(nums, target)
	if len(result) == 0 {
		t.Fatal("function twoSum do not return values")
		return
	}
	if result[0] != correct[0] || result[1] != correct[1] {
		t.Fatal("function twoSum returns incorrect values.")
	}
}
