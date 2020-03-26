package quickSort

import "math/rand"

func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	p := rand.Intn(n)
	arr[0],arr[p] = arr[p],arr[0]
	v := arr[0]
	x := 1
	for i := 1; i < n; i++ {
		if arr[i]<=v {
			arr[i],arr[x] = arr[x],arr[i]
			x++
		}
	}
	x--
	arr[x],arr[0] = arr[0],arr[x]
	res := QuickSort(arr[:x])
	res = append(res, v)
	res = append(res, QuickSort(arr[x+1:])...)
	return res
}
