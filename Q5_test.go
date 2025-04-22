package main

func isUgly(n int) bool {
	if n == 0 {
		return false
	} // 边界条件

	for n%2 == 0 { // 把所有的2出干净
		n /= 2 // n = n / 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	if n != 1 {
		return false
	} else {
		return true
	}
}
