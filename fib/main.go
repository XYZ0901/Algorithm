package main

import (
	"fmt"
	"time"
)

// dp
func fib(n int) int {
	memo := make([]int,n+1)
	memo[0]=0
	memo[1] =1
	for i:=2;i<=n ;i++  {
		memo[i]=memo[i-1]+memo[i-2]
	}
	return memo[n]
}

func main() {
	n := 1000
	startTime := time.Now()
	res := fib(n)
	endTime := time.Since(startTime)
	fmt.Println("fib(", n, ") = ", res)
	fmt.Println("time : ", endTime)
}
