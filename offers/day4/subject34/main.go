package main

func main() {}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res[0] = i
			break
		}
	}

	if res[0] == -1 {
		return []int{-1, -1}
	}

	for i := res[0]; i < len(nums); i++ {
		if nums[i] == target {
			res[1] = i
		}
	}

	return res
}

// searchRange2 双指针
func searchRange2(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] == target && nums[r] == target {
			return []int{l, r}
		}
		if nums[l] != target {
			l++
		}
		if nums[r] != target {
			r--
		}
	}
	return []int{-1, -1}
}
