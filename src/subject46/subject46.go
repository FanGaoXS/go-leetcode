package main

import "fmt"

func main() {
	nums := []int{0,1}
	result := permute(nums)
	for i := range result {
		fmt.Printf("result[i] = %v\n",result[i])
	}
}

var (
	result = [][]int{}
)

func permute(nums []int) [][]int {
	ints := []int{}
	dfs(nums,ints)
	return result
}

func dfs(nums []int,ints []int)  {
	if len(ints) == len(nums) {
		result = append(result,ints)
		return
	}
	for i := 0; i < len(nums); i++ {
		//如果ints中包含nums[i]则不添加此数
		if isArrayContains(ints,nums[i]) {
			continue
		}
		ints = append(ints,nums[i])
		dfs(nums,ints)
		ints = ints[:len(ints)-1]
	}
}

//判断ints中是否包含nums[i]
func isArrayContains(ints []int,num int) (ok bool) {
	for i := 0; i < len(ints); i++ {
		if num == ints[i] {
			ok = true
			break
		}
	}
	return ok
}
