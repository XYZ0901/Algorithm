package main

import (
	"math/rand"
	"testing"
)

func Test_heapSort1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Benchmark_heapSort(b *testing.B)  {
	arr :=[]int{}
	n := 1000000
	for i:=0;i<n;i++{
		arr = append(arr, rand.Intn(n))
	}
	for i:=0;i<b.N;i++{
		heapSort2(arr)
	}
}