package main

import "fmt"

func main() {
	nums := []int{1,2,3,4}
	result := permute(nums)
	for i := range result {
		fmt.Printf("result[i] = %v\n",result[i])
	}
}

var result = [][]int{}

func permute(nums []int) [][]int {
	// ps：在leetcode中需要先清空全局变量result，不然会获取到上次的result（bug）
	result = [][]int{}	//结果集
	ints := []int{}		//单次排列
	dfs(nums,ints)
	return result
}

func dfs(nums []int, ints []int)  {
	if len(ints) == len(nums) {
		cInts := make([]int, len(ints))
		copy(cInts,ints)
		//将ints copy到cInts中再添加到result中
		result = append(result,cInts)
		return
	}
	for i := 0; i < len(nums); i++ {
		//如果ints中包含nums[i]则不添加此数
		if isArrayContains(ints,nums[i]) {
			continue
		}
		//否则将该数添加到排列中
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
