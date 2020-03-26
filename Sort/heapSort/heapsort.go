package heapSort

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
