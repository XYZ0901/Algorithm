package MaxHeap

import (
	"fmt"
)

type MaxHeap struct {
	data     []int
	count    int
	capacity int
}

func (maxheap MaxHeap) Print() {
	count := maxheap.count
	numList := []int{1}
	nub := 1
	sum := maxheap.count - 1
	for count > 0 {
		for i := 1; i <= nub; i++ {
			k := numList[i-1]
			fmt.Print(maxheap.data[k], " ")
			numList = append(numList, k*2)
			numList = append(numList, k*2+1)
		}
		fmt.Println()
		numList = numList[nub:]
		nub *= 2
		if sum < nub {
			nub = sum
		} else {
			sum = sum - nub
		}
		count /= 2
	}
}

func (maxheap *MaxHeap) shiftDown(i int) {
	k := i
	for 2*k <= maxheap.count {
		j := 2 * k
		if j+1 <= maxheap.count && maxheap.data[j+1] > maxheap.data[j] {
			j += 1
		}
		if maxheap.data[k] >= maxheap.data[j] {
			break
		}
		maxheap.data[k], maxheap.data[j] = maxheap.data[j], maxheap.data[k]
		k = j
	}
}

func (maxheap *MaxHeap) shiftUp() {
	count := maxheap.count
	for count > 1 && maxheap.data[count/2] < maxheap.data[count] {
		maxheap.data[count], maxheap.data[count/2] = maxheap.data[count/2], maxheap.data[count]
		count /= 2
	}
}

func (maxheap *MaxHeap) ExtractMax() (ret int) {
	if maxheap.count > 0 {
		ret = maxheap.data[1]
		maxheap.data[1], maxheap.data[maxheap.count] = maxheap.data[maxheap.count], maxheap.data[1]
		maxheap.count--
		maxheap.shiftDown(1)
	}
	return
}

func (maxheap *MaxHeap) Insert(item int) {
	if maxheap.count+1 <= maxheap.capacity {
		maxheap.data[maxheap.count+1] = item
		maxheap.count++
		maxheap.shiftUp()
	}
}

func (maxheap *MaxHeap) Size() int {
	return maxheap.count
}

func (maxheap *MaxHeap) IsEmpty() bool {
	return maxheap.count == 0
}

func CreatMaxHeap_null(capacity int) MaxHeap {
	datas := make([]int, capacity+1)
	maxheap := MaxHeap{
		data:     datas,
		count:    0,
		capacity: capacity+1,
	}
	return maxheap
}

func CreatMaxHeap_arr(arr []int) MaxHeap {
	datas := make([]int, 1)
	datas = append(datas,arr...)
	maxheap := MaxHeap{
		data:     datas,
		count:    len(arr),
		capacity: len(arr)+1,
	}
	for i:=maxheap.count/2;i>=1;i--{
		maxheap.shiftDown(i)
	}
	return maxheap
}