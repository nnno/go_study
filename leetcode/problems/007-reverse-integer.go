package main

import "math"

const (
	MaxInt = math.MaxInt32
	MinInt = -MaxInt
)

func reverse(x int) int {
	var res int

	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}

	if res > MaxInt || res < MinInt {
		return 0
	}
	return res
}
