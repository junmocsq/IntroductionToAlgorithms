package ch09

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"math"
	"math/rand"
	"testing"
)

func Maximum(arr []int) int {
	if len(arr) == 0 {
		return math.MinInt64
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func Minimum(arr []int) int {
	if len(arr) == 0 {
		return math.MaxInt64
	}
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxMin 同时获取最小值和最大值 3n/2
func MaxMin(arr []int) (min int, max int) {
	length := len(arr)
	if length == 0 {
		return math.MaxInt64, math.MinInt64
	}
	var i int
	if length%2 == 0 {
		if arr[0] > arr[1] {
			max = arr[0]
			min = arr[1]
		} else {
			max = arr[1]
			min = arr[0]
		}
		i = 2
	} else {
		min, max = arr[0], arr[0]
		i = 1
	}
	for ; i < length; i += 2 {
		if arr[i] > arr[i+1] {
			if arr[i] > max {
				max = arr[i]
			}
			if arr[i+1] < min {
				min = arr[i+1]
			}
		} else {
			if arr[i+1] > max {
				max = arr[i+1]
			}
			if arr[i] < min {
				min = arr[i]
			}
		}
	}
	return min, max

}

func FindK(arr []int, k int) (result []int, ok bool) {
	if len(arr) < k {
		return
	}
	findK(arr, 0, len(arr)-1, k-1)
	return arr[:k], true
}

func findK(arr []int, l, r, k int) {
	divide := partition(arr, l, r)
	if divide < 0 {
		return
	}
	if divide == k {
		return
	} else if divide > k {
		findK(arr, l, divide-1, k)
	} else {
		findK(arr, divide+1, r, k)
	}
}

func randArr(arr []int, l, r int) {
	random := rand.Intn(r - l + 1)
	arr[random+l], arr[l] = arr[l], arr[random+l]
}

// 从左端开始，
func partition(arr []int, l, r int) int {
	if l > r {
		return -1
	}
	if l == r {
		return l
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

func TestFindK(t *testing.T) {
	arr := tools.RandArr(100, 10000)
	t.Log(arr)
	t.Log(FindK(arr, 10))
}
