package main

import "fmt"

var memo = [][]int{}

func knapsack01(w, v []int, c int) int {
	n := len(w)
	if n == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		memo = append(memo, make([]int, c+1))
	}
	for j := 0; j <= c; j++ {
		if j >= w[0] {
			memo[0][j] = v[0]
		} else {
			memo[0][j] = 0
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= c; i++ {
			memo[i][j] = memo[i-1][j]
			if j >= w[i] {
				memo[i][j] = max(memo[i][j], v[i]+memo[i-1][j-w[i]])
			}
		}
	}
	return memo[n-1][c]
}

func knapsack01_memo(w, v []int, c int) int {
	n := len(w)
	for i := 0; i < n; i++ {
		memo = append(memo, make([]int, c+1))
	}
	return bestValue_memo(w, v, n-1, c)
}

func bestValue_memo(w []int, v []int, index, c int) int {
	if index < 0 || c <= 0 {
		return 0
	}
	if memo[index][c] != 0 {
		return memo[index][c]
	}
	res := bestValue_memo(w, v, index-1, c)
	if c >= w[index] {
		res = max(res, v[index]+bestValue_memo(w, v, index-1, c-w[index]))
	}
	memo[index][c] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	w := []int{3, 5}
	v := []int{6, 4}
	fmt.Println(w, v)
}
