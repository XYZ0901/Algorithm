package main

import (
	"fmt"
	"math/rand"
	"time"
)

func singleStack(arr []int) {
	n := len(arr)
	stack := []int{}
	out := [][]int{}
	for i := 0; i < n; i++ {
		out = append(out, make([]int, 2))
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			stack = append(stack, i)
			continue
		}
		if arr[i] > arr[stack[len(stack)-1]] {
			x := i
			stack = append(stack, x)
			continue
		}
		m := len(stack)
		for j := m; j > 0; j-- {
			if arr[i] < arr[stack[len(stack)-1]] {
				x := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				out[x][1] = i
				if len(stack) == 0 {
					out[x][0] = -1
				} else {
					out[x][0] = stack[len(stack)-1]
				}
			}
		}
		stack = append(stack, i)
	}
	m := len(stack)
	for j := m; j > 0; j-- {
		x := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out[x][1] = -1
		if len(stack) == 0 {
			out[x][0] = -1
		} else {
			out[x][0] = stack[len(stack)-1]
		}
	}
	fmt.Println(out)
}

type myList struct {
	v    int
	next *myList
}

func singleStackS(arr []int) [][]int {
	n := len(arr)
	stack := []myList{}
	out := [][]int{}
	for i := 0; i < n; i++ {
		out = append(out, make([]int, 2))
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			stack = append(stack, myList{
				v:    i,
				next: nil,
			})
			continue
		}
		if arr[i] > arr[stack[len(stack)-1].v] {
			stack = append(stack, myList{
				v:    i,
				next: nil,
			})
			continue
		} else if arr[i] == arr[stack[len(stack)-1].v] {
			now := stack[len(stack)-1]
			next := myList{
				v:    i,
				next: &now,
			}
			stack[len(stack)-1] = next
			continue
		}
		if i == 7 {
			n = len(arr)
		}
		m := len(stack)
		for j := m; j > 0; j-- {
			if arr[i] >= arr[stack[len(stack)-1].v] {
				break
			}
			ml := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for {
				out[ml.v][1] = i
				if len(stack) == 0 {
					out[ml.v][0] = -1
				} else {
					out[ml.v][0] = stack[len(stack)-1].v
				}
				if ml.next == nil {
					break
				}
				ml = *ml.next
			}
		}
		if len(stack)>0&&arr[i] == arr[stack[len(stack)-1].v] {
			now := stack[len(stack)-1]
			next := myList{
				v:    i,
				next: &now,
			}
			stack[len(stack)-1] = next
			continue
		}
		stack = append(stack, myList{v: i})
	}
	m := len(stack)
	for i := m; i > 0; i-- {
		ml := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for {
			out[ml.v][1] = -1
			if len(stack) == 0 {
				out[ml.v][0] = -1
			} else {
				out[ml.v][0] = stack[len(stack)-1].v
			}
			if ml.next == nil {
				break
			}
			ml = *ml.next
		}
	}
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i:=0;i<1000000;i++{
		arr = append(arr, rand.Intn(1000000))
	}
	start := time.Now()
	singleStackS(arr)
	fmt.Println(time.Since(start))
}
