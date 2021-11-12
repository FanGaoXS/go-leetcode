package subject78

func subsets(nums []int) [][]int {
	lists := [][]int{}
	lists = append(lists,[]int{})	//加入一个空集
	for i := 0; i < len(nums); i++ {
		size := len(lists)
		for j := 0; j < size; j++ {
			integers := make([]int, len(lists[j]))
			copy(integers,lists[j])
			integers = append(integers,nums[i])
			lists = append(lists, integers)
		}
	}
	return lists
}