package main

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	var reverseNum = 0
	for x > reverseNum {
		reverseNum = reverseNum*10 + (x % 10)
		x /= 10
	}
	return x == reverseNum || x == reverseNum/10
}
