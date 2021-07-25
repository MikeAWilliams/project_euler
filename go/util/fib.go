package util

func Fib(n int) int {
	if 2 >= n {
		return 1
	}
	nm1 := 1
	nm2 := 1
	for index := 2; index < n; index++ {
		tmp := nm2 + nm1
		nm1 = nm2
		nm2 = tmp
	}

	return nm2
}
