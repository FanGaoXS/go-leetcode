package main

func main() {

}

func trap(height []int) int {
	var v int = 0	//返回值体积
	var maxHeightIndex int = 0	//最高处的下标
	//优化方案：i起始从1开始（没有必要再去比较height[0]和height[0]）
	for i := 0; i < len(height); i++ {
		if height[i] > height[maxHeightIndex] {
			maxHeightIndex = i
		}
	}
	var moreHeightIndex int = 0
	//优化方案：i起始从1开始（没有必要再去比较height[0]和height[0]）
	for i := 0; i < maxHeightIndex; i++ {
		if height[i] > height[moreHeightIndex] {
			moreHeightIndex = i
			continue
		}
		v += height[moreHeightIndex] - height[i]
	}
	moreHeightIndex = len(height) - 1
	//优化方案：i起始从倒数第二开始（没有必要再去比较height[尾]和height[尾]）
	for i := len(height)-1; i > maxHeightIndex ; i-- {
		if height[i] > height[moreHeightIndex] {
			moreHeightIndex = i
			continue
		}
		v += height[moreHeightIndex] - height[i]
	}
	return v
}
