package main

import (
	"math"
	"strconv"
)

func myAtoi(s string) int {
	len := len(s)
	//初始化结果
	res := 0
	if len == 0 {
		return res
	}
	//是否正数（默认正数）
	sign := 1
	//下标
	i := 0
	//结果集字符串
	str := ""
	if s[i] == ' ' {
		for i<len && s[i] == ' ' {
			i++
		}
	}	//读取空格字符
	if i<len && s[i] == '+' {
		sign = 1
		i++
	} else if i<len && s[i] == '-' {
		sign = -1
		i++
	}	//读取正负符号
	if i<len && s[i] >= '0' && s[i]<= '9' {
		for i<len && s[i] >= '0' && s[i]<= '9' {
			str = str + string(s[i])
			i++
		}
	}	//读取数字字符
	atoi, _ := strconv.Atoi(str)	//将字符字符串转为int类型
	if sign * atoi < math.MinInt32 {
		return math.MinInt32
	} else if sign * atoi > math.MaxInt32 {
		return math.MaxInt32
	}
	return sign * atoi
}
