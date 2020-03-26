package main

import "fmt"

func climbStairs(n int) int {
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func climbStairs2(n int) int {
	x, y := 1, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}

func main() {
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs2(10))
}
