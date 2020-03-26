package main

import "fmt"

func maxUncrossedLines( A []int ,  B []int ) int {
	// write code here
	suma,sumb,a,b := 0,0,0,0
	for _,v:=range A{
		for i:=a;i<len(B) ;i++  {
			if v == B[i] {
				suma++
				a = i+1
				break
			}
		}
	}
	for _,v:=range B{
		for i:=b;i<len(A) ;i++  {
			if v == A[i] {
				sumb++
				b = i+1
				break
			}
		}
	}
	sum := max(suma,sumb)
	return sum
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	a := []int{2,5,1,2,5}
	b := []int{10,5,2,1,5,2}
	fmt.Println(maxUncrossedLines(a,b))
}
