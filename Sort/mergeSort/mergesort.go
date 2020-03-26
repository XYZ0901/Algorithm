package mergeSort

func MergeSort(arr []int) []int {
	return __mergeSort(arr)
}

func __mergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	mid := n / 2
	arr1 := arr[:mid]
	arr2 := arr[mid:]
	arr1 = __mergeSort(arr1)
	arr2 = __mergeSort(arr2)
	if arr1[len(arr1)-1]<arr2[len(arr2)-1] {
		return append(arr1, arr2...)
	}
	return __merge(arr1, arr2)
}

func __merge(arr1, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)

	arr := []int{}
	i, j := 0, 0
	for i < n || j < m {
		if i < n && j < m {
			if arr1[i] < arr2[j] {
				arr = append(arr, arr1[i])
				i++
			} else {
				arr = append(arr, arr2[j])
				j++
			}
		} else if i < n {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	return arr
}