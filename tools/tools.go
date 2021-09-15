package tools

import (
	"math/rand"
	"time"
)

func RandArr(num, maxValue int) []int {
	var arr []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(maxValue))
	}
	return arr
}

func CheckIsSort(arr []int) bool {
	length := len(arr)
	if length < 3 {
		return true
	}
	result := true
	// 判断是否为递增
	for i := 0; i < length-1; i++ {
		if arr[i] > arr[i+1] {
			result = false
			break
		}
	}
	if result {
		return result
	}
	result = true
	// 判断是否为递减
	for i := 0; i < length-1; i++ {
		if arr[i] < arr[i+1] {
			break
		}
	}
	return result
}
