package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ThreeArray := [][]int{}
	if len(nums) < 3 {
		return ThreeArray
	}
	sort.Ints(nums)
	for k, _ := range nums {
		if nums[k] > 0 || k == (len(nums)-2) {
			break
		}
		if k != 0 && nums[k] == nums[k-1] {
			continue
		}

		l, r := k+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] == -nums[k] {
				ThreeArray = append(ThreeArray, []int{nums[k], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l]+nums[r] > -nums[k] {
				r--
			} else {
				l++
			}
		}
	}
	return ThreeArray
}

func main() {
	a := []int{-4,-1,-1,0,1,2}
	fmt.Println(threeSum(a))
}
