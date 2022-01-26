/**
 * % go test -run Reverse -v ./leetcode
 */
package main

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	execPalindromeNumber(t, 121, true)
	execPalindromeNumber(t, -121, false)
	execPalindromeNumber(t, 10, false)
	execPalindromeNumber(t, 0, true)
}

func execPalindromeNumber(t *testing.T, input int, expect bool) {
	ret := isPalindrome(input)
	if ret != expect {
		t.Fatalf("Wrong Answer. input: %v, result: %v, expect: %v\n", input, ret, expect)
	}
}
