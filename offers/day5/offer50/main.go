package offer50

func firstUniqChar(s string) byte {
	uniqueBytes := make(map[rune]int)

	for _, b := range s {
		if _, ok := uniqueBytes[b]; ok {
			uniqueBytes[b] += 1
			continue
		}
		uniqueBytes[b] = 1
	}

	for _, b := range s {
		if uniqueBytes[b] == 1 {
			return byte(b)
		}
	}

	return ' '
}

func firstUniqChar2(s string) byte {
	chars := [26]int{}

	for _, c := range s {
		chars[c-97]++
	}

	for _, c := range s {
		if chars[c-97] == 1 {
			return byte(c)
		}
	}

	return ' '
}
