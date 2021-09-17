package ch04

import (
	"math"
	"testing"
)

// 最大子数组
// 分治法,最大子数组分布
//	全在左[low,mid]
//	全在右，[mid+1,high]
//	跨越中点，[low,high]
func mid(low, high int) int {
	return (low + high) / 2
}

// 跨越中点求最大
func findCrossingMid(arr []int, low, high int) (int, int, int) {
	mid := mid(low, high)
	if low > mid || mid+1 > high {
		return -1, -1, math.MinInt64
	}
	//left [low,mid]
	leftSum := arr[mid]
	leftIndex := mid
	sum := leftSum
	for i := mid - 1; i >= low; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftIndex = i
			leftSum = sum
		}
	}

	//right [mid+1,high]
	rightSum := arr[mid+1]
	rightIndex := mid + 1
	sum = rightSum
	for i := mid + 2; i <= high; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightIndex = i
			rightSum = sum
		}
	}
	//fmt.Println(leftIndex, rightIndex, leftSum + rightSum)
	return leftIndex, rightIndex, leftSum + rightSum
}

func findMaximumSubarray(arr []int, low, high int) (int, int, int) {
	if low > high {
		return -1, -1, math.MinInt64
	}
	if low == high {
		return low, low, arr[low]
	}
	mid := mid(low, high)

	mlow, mhigh, midMax := findCrossingMid(arr, low, high)
	llow, lhigh, leftMax := findMaximumSubarray(arr, low, mid)
	rlow, rhigh, rightMax := findMaximumSubarray(arr, mid+1, high)
	if leftMax >= midMax && leftMax >= rightMax {
		return llow, lhigh, leftMax
	} else if rightMax > midMax && rightMax >= leftMax {
		return rlow, rhigh, rightMax
	} else {
		return mlow, mhigh, midMax
	}
}

func FindMaximumSubarray(arr []int) (int, int, int) {
	return findMaximumSubarray(arr, 0, len(arr)-1)
}

func TestFindMaximumSubarray(t *testing.T) {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	t.Log(FindMaximumSubarray(arr))
	t.Log(FindMaxSubarray2(arr))
	t.Log(FindMaximumSubarrayByViolence(arr))
	arr = []int{-5, -4, -3, -2, -99}
	t.Log(FindMaximumSubarray(arr))
}

// exercise 4.1-2
func FindMaximumSubarrayByViolence(arr []int) (int, int, int) {
	addArr := make([]int, len(arr))
	for k, v := range arr {
		if k >= 1 {
			addArr[k] = v + addArr[k-1]
		} else {
			addArr[k] = v
		}
	}
	maxTemp, l, r := math.MinInt64, -1, -1
	for m := 0; m < len(addArr)-1; m++ {
		if maxTemp < addArr[m] {
			maxTemp = addArr[m]
			l = m
			r = m
		}
		for n := m + 1; n < len(addArr); n++ {
			temp := addArr[n] - addArr[m]
			if temp > maxTemp {
				maxTemp = temp
				l = m
				r = n
			}
		}
	}
	// 求出的是相减的左边，需要加1为区间
	return l + 1, r, maxTemp
}

// exercise 4.1-5
// 从左到右，记录目前为止已经处理过的最大子数组。对于[0,m+1]的最大子数组，要么在[0,m]中，要么包含m+1这个元素
func FindMaxSubarray2(arr []int) (int, int, int) {
	if len(arr) < 1 {
		return -1, -1, math.MinInt64
	}
	left, right, max := 0, 0, arr[0]
	for i := 1; i < len(arr); i++ {
		left, right, max = findMaxSubarray2(arr, i, left, right, max)
	}
	return left, right, max
}

func findMaxSubarray2(arr []int, high, left, right, max int) (int, int, int) {
	sum := 0
	for i := high; i >= 0; i-- {
		sum += arr[i]
		if sum > max {
			max = sum
			right = high
			left = i
		}
	}
	return left, right, max
}
