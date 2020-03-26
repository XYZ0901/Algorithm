package main

import (
	"Algorithm/Heap/MaxHeap"
	"fmt"
	"math/rand"
	"time"
)

func heapSort1(arr []int) []int {
	n := len(arr)
	maxheap := MaxHeap.CreatMaxHeap_null(n)
	for i := 0; i < n; i++ {
		maxheap.Insert(arr[i])
	}
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxheap.ExtractMax()
	}
	return arr
}

func heapSort2(arr []int) []int {
	maxheap := MaxHeap.CreatMaxHeap_arr(arr)
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxheap.ExtractMax()
	}
	return arr
}

func heapSort3(arr []int) {
	n := len(arr)
	for i:=(n+1)/2-1;i>=0;i-- {
		shiftDown(arr,n,i)
	}
	for i := n-1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr,i,0)
	}
}

func shiftDown(arr []int, n, k int) {
	// n number for arr
	// from k
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1
		}
		if arr[k] >= arr[j] {
			break
		}
		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}

func main() {
	arr := []int{}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(100000))
	}
	start := time.Now()
	heapSort1(arr)
	fmt.Println("heapSort1:", time.Since(start))

	arr2 := []int{}
	for i := 0; i < 1000000; i++ {
		arr2 = append(arr2, rand.Intn(100000))
	}
	start = time.Now()
	heapSort2(arr2)
	fmt.Println("heapSort2:", time.Since(start))

	arr3 := []int{}
	for i := 0; i < 1000000; i++ {
		arr3 = append(arr3, rand.Intn(1000000))
	}
	start = time.Now()
	heapSort3(arr3)
	fmt.Println("heapSort3:", time.Since(start))
}
