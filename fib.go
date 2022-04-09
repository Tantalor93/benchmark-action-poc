package benchmark_action_poc

func Fib(n int) int {
	// test
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Reverse(s string) string {
	rs := []rune(s)
	rs2 := []rune(s)
	j := 0
	for i := len(s)-1; i >= 0; i-- {
		rs2[j] = rs[i]
		j++
	}
	return string(rs2)
}