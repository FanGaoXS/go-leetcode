package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{2,8},{1,3}}
	result := merge(intervals)
	fmt.Printf("result = %v",result)
}

func merge(intervals [][]int) [][]int {
	//fmt.Printf("intervals = %v\n",intervals)
	//对intervals进行排序（以每个区间的start进行排序）
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0];
	})
	//fmt.Printf("intervals = %v\n",intervals)
	merged := [][]int{intervals[0]}
	for i :=1; i < len(intervals); i++ {
		//当前区间的start和end
		start := intervals[i][0]
		end := intervals[i][1]
		//上个已经排好序的区间的start和end
		//lastMergedStart := merged[len(merged)-1][0]
		lastMergedEnd := merged[len(merged)-1][1]
		//如果当前区间的start比上次排好序的区间的end还要大
		if start > lastMergedEnd {
			//则代表他们两个区间不可能重叠
			merged = append(merged,intervals[i])
		} else { //否则有可能重叠
			if end > lastMergedEnd {
				//如果当前区间的end还要大，则重新规划区间
				merged[len(merged)-1][1] = end
			}
			//否则仍然是原有区间
		}
	}
	return merged
}
