package main

import (
	"Algorithm/Sort/mergeSort"
	"Algorithm/Sort/quickSort"
	"fmt"
	"math/rand"
	"time"
)

func GetRandArr(n int) []int {
	arrRand := []int{}
	for i := 0; i < n; i++ {
		arrRand = append(arrRand, rand.Intn(n))
	}
	return arrRand
}

func GetSmalRandArr(n, m int) []int {
	arrRand := []int{}
	for i := 0; i < n; i++ {
		arrRand = append(arrRand, i)
	}
	for i := 0; i < m; i++ {
		x, y := rand.Intn(n), rand.Intn(n)
		arrRand[x], arrRand[y] = arrRand[y], arrRand[x]
	}
	return arrRand
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr1 := []int{}
	arr2 := []int{}
	start := time.Time{}
	var mergeTime time.Duration
	mergeTime = 0
	n,m := 1000000,10
	fmt.Println("数组长度为",n,"有序中更换",m,"个数字")
	for i := 0; i < 10; i++ {
		arr1 = GetRandArr(n)
		start = time.Now()
		mergeSort.MergeSort(arr1)
		mergeTime += time.Since(start)
	}
	fmt.Println("无序归并：", mergeTime/10)
	mergeTime = 0
	for i := 0; i < 10; i++ {
		arr2 = GetSmalRandArr(n, m)
		start = time.Now()
		mergeSort.MergeSort(arr2)
		mergeTime += time.Since(start)
	}
	fmt.Println("有序归并：", mergeTime/10)

	mergeTime = 0
	for i := 0; i < 10; i++ {
		arr1 = GetRandArr(n)
		start = time.Now()
		quickSort.QuickSort(arr2)
		mergeTime += time.Since(start)
	}
	fmt.Println("无序快排：", mergeTime/10)
	mergeTime = 0
	for i := 0; i < 10; i++ {
		arr2 = GetSmalRandArr(n, m)
		start = time.Now()
		quickSort.QuickSort(arr2)
		mergeTime += time.Since(start)
	}
	fmt.Println("有序快排：", mergeTime/10)
}
