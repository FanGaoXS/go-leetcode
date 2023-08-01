package main

func main() {
	nums := []int{0, 1, 3}
	res := missingNumber2(nums)
	println(res)
}

func search(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}

// search2 双指针
func search2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] == target && nums[r] == target {
			return r - l + 1
		}
		if nums[l] != target {
			l++
		}
		if nums[r] != target {
			r--
		}
	}
	return 0
}

func missingNumber(nums []int) int {
	for i, num := range nums {
		if num != i {
			return i
		}
	}

	return len(nums)
}

// missingNumber2 hash表
func missingNumber2(nums []int) int {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for i := 0; i < len(nums)+1; i++ {
		if _, ok := set[i]; !ok {
			return i
		}
	}

	return -1
}
