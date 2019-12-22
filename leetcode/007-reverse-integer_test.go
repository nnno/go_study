/**
 * % go test -run Reverse -v ./leetcode
 */
package main

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	execReverseInteger(t, 123, 321)
	execReverseInteger(t, -123, -321)
	execReverseInteger(t, 120, 21)
	execReverseInteger(t, MaxInt, 7463847412)  // 2147483647
	execReverseInteger(t, MinInt, -7463847412) // -2147483647
	execReverseInteger(t, 1534236469, 0)
}

func execReverseInteger(t *testing.T, input int, expect int) {
	result := reverse(input)
	if result != expect {
		t.Fatalf("Wrong Answer. result: %v, expect: %v\n", result, expect)
	}
}
