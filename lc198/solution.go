package main

var memo = []int{}

//func rob(nums []int) int {
//	memo = make([]int, len(nums))
//	for k,_:=range memo{
//		memo[k]-=1
//	}
//	return tryRob(nums,0)
//}
//
//func tryRob(nums []int,index int) int {
//	if index>= len(nums) {
//		return 0
//	}
//	if memo[index] != -1 {
//		return memo[index]
//	}
//	res := 0
//	for i:=index;i< len(nums);i++  {
//		res = max(res,nums[i]+tryRob(nums,i+2))
//	}
//	memo[index] = res
//	return res
//}

func rob(nums []int) int {
	memo = make([]int, len(nums))
	n := len(nums)
	if n == 0 {
		return n
	}
	memo[n-1] = nums[n-1]
	for i:=n-2;i>=0;i--{
		for j:=i;j<n;j++{
			if j+2>=n {
				memo[i] = max(memo[i],nums[j])
			}else{
				memo[i] = max(memo[i],nums[j]+memo[j+2])
			}

		}
	}
	return memo[0]
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	
}
