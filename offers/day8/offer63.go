package day8

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	res := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > res {
			res = prices[i] - min
		}
	}

	return res
}
