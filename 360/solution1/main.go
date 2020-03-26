package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		n, m := 0, 0
		_, err := fmt.Scanf("%d %d", &n, &m)
		if err == io.EOF {
			break
		}
		arrA := [][]float64{}
		for i := 0; i <= n; i++ {
			arrA = append(arrA, make([]float64, m+1))
		}
		for i := 1; i <= n; i++ {
			for j := 0; j <= m; j++ {
				fi, fj := float64(i), float64(j)
				if j == 0 {
					arrA[i][j] = 1
					continue
				}
				if j == 1 {
					arrA[i][j] = fi / (fi + fj)
					continue
				}
				arrA[i][j] = fi / (fi + fj)
				if i > 1 && j > 1 {
					arrA[i][j] += arrA[i-1][j-2] * fi * fj * (fj - 1) / (fi + fj) / (fi + fj - 1) / (fi + fj - 2)
				}
				if j > 2 {
					arrA[i][j] += arrA[i][j-3] * fj * (fj - 1) * (fj - 2) / (fi + fj) / (fi + fj - 1) / (fi + fj - 2)
				}
			}
		}
		fmt.Println(arrA)
	}
}
