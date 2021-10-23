package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	num1 := "101"
	num2 := "1009"
	result := addStrings(num1, num2)
	fmt.Println("result =",result)
}

func addStrings(num1 string, num2 string) string {
	var buffer bytes.Buffer
	var ints = []int{}
	add := false
	i := len(num1)-1	//指向num1的下标
	j := len(num2)-1	//指向num2的下标
	for !(i<0&&j<0&&!add) {
		var numI int = 0	//指向num1的下标的值
		var numJ int = 0
		var sum int = 0		//和
		if i>=0 {
			numI = int(num1[i])-48
			i--
		}
		if j>=0 {
			numJ = int(num2[j])-48
			j--
		}
		if add {
			sum = numI + numJ + 1
		} else {
			sum = numI + numJ
		}
		if sum >= 10 {
			add = true
			sum = sum%10
		} else {
			add  = false
		}
		ints = append(ints,sum)
	}
	//依次读每位的和结果，然后追加到buffer字符串中（倒序）
	for k := len(ints)-1; k >= 0; k-- {
		//将int类型转为string类型
		buffer.WriteString(strconv.Itoa(ints[k]))
	}
	return buffer.String()
}

func addStrings2(num1 string, num2 string) string {
	var s string
	add := false
	i := len(num1)-1	//num1的下标
	j := len(num2)-1	//num2的下标
	for !(i<0&&j<0&&!add) {
		var numI int = 0
		var numJ int = 0
		var sum int = 0
		if i>=0 {
			numI = int(num1[i])-48
			i--
		}
		if j>=0 {
			numJ = int(num2[j])-48
			j--
		}
		if add {
			sum = numI + numJ + 1
		} else {
			sum = numI + numJ
		}
		if sum >= 10 {
			add = true
			sum = sum%10
		} else {
			add  = false
		}
		//直接追加
		s = strconv.Itoa(sum) + s
	}
	return s
}
