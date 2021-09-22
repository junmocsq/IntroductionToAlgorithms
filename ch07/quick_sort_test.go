package ch07

import (
	"math/rand"
	"testing"
	"time"
)

func QuickSort(arr []int) {
	rand.Seed(time.Now().Unix())
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	divide := partition(arr, l, r)
	if divide < 0 {
		return
	}
	quickSort(arr, l, divide-1)
	quickSort(arr, divide+1, r)
}

// 随机化一个[l,r]中的元素和l交换
func randArr(arr []int, l, r int) {
	random := rand.Intn(r - l + 1)
	arr[random+l], arr[l] = arr[l], arr[random+l]
}

// 从左端开始，
func partition(arr []int, l, r int) int {
	if l >= r {
		return -1
	}
	i := l - 1

	// 随机化
	randArr(arr, l, r)

	for j := l; j < r; j++ {
		if arr[j] <= arr[r] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

// 两端同时开始 i为左节点 j为右节点
func partition2(arr []int, l, r int) int {
	if l >= r {
		return -1
	}
	i := l - 1
	j := r - 1
	// 小于pivot的[l,i] 大于等于pivot [j+1,r-1]
	for i < j {
		if arr[j] >= arr[r] {
			j--
		} else {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	//arr[i+1], arr[r] = arr[r], arr[i+1]
	arr[j+1], arr[r] = arr[r], arr[j+1]
	return i + 1
}

func TestQuickSort(t *testing.T) {
	arr := []int{2, 8, 7, 1, 3, 5, 6, 4}
	QuickSort(arr)
	t.Log(arr)
}
