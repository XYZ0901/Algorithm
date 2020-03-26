package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	//memo := make([]int, n)
	t := 0
	y := 0
	for i := 0; i < n; i++ {
		for j:=0;j<len(triangle[i]) ;j++  {
			if j==0 {
				y = y+triangle[i][j]
			}
			if t+triangle[i][j]<y {
				y = t+triangle[i][j]
			}
		}
		t = y
	}
	return y
}

func main() {
	t := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	fmt.Println(minimumTotal(t))
}
