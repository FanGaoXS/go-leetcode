package main

func main() {}

func findRepeatNumber(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return n
		}
		set[n] = struct{}{}
	}
	return 0
}
