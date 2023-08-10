package day10

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	// len(s)>=2

	max := 1 // 只有两个字符的时候，最长不重复子串至少为1
	l := 0   // 左指针指向s[0]
	r := 1   // 右指针指向s[1]
	set := make(map[byte]struct{})
	set[s[l]] = struct{}{} // 将s[0]加入到set中
	for r < len(s) {
		if _, ok := set[s[r]]; !ok {
			set[s[r]] = struct{}{}
			r++
			max = Max(max, len(set))
		} else {
			delete(set, s[l])
			l++
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
