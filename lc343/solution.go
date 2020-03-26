package main

import "fmt"

var memo []int

func integerBreak(n int) int {
	memo = make([]int,n+1)
	memo[1] = 1
	for i:=2;i<=n ;i++  {
		for j:=1;j<=i-1 ;j++  {
			memo[i] = max3(memo[i],j*(i-j),j*memo[i-j])
		}
	}
	return memo[n]
}

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	} else if c > b {
		return c
	}
	return b
}

func main() {
	fmt.Println(integerBreak(10))
}
