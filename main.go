package main

import "fmt"

type Result struct {
	profit int
	t, p, c int
}

func max(a, b Result) Result {
	if a.profit > b.profit {
		return a
	}
	return b
}


var dp = make(map[int]Result)

func solve(n int) Result {
	if n <= 0 {
		return Result{0, 0, 0, 0}
	}

	if val, ok := dp[n]; ok {
		return val
	}

	best := Result{0, 0, 0, 0}

	if n >= 5 {
		res := solve(n - 5)
		res.profit += 1500
		res.t++
		best = max(best, res)
	}

	if n >= 4 {
		res := solve(n - 4)
		res.profit += 1000
		res.p++
		best = max(best, res)
	}

	if n >= 10 {
		res := solve(n - 10)
		res.profit += 2000
		res.c++
		best = max(best, res)
	}

	dp[n] = best
	return best
}

func main() {
	var n int
	fmt.Print("Enter Time Units: ")
	fmt.Scan(&n)

	result := solve(n)

	fmt.Println("Maximum Profit:", result.profit)
	fmt.Printf("T:%d P:%d C:%d\n", result.t, result.p, result.c)
}