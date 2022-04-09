package benchmark_action_poc

func Fib(n int) int {
	prevFibs := []int{0, 1}
	for i := 3; i <= n; i++ {
		nextFib := prevFibs[0] + prevFibs[1]
		prevFibs = []int{prevFibs[1], nextFib}
	}
	if n > 1 {
		return prevFibs[1]
	}
	return prevFibs[0]
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