package ch01

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"testing"
)

func TestMergeSort(t *testing.T) {
	length := 1000
	arr := tools.RandArr(length, 20000)
	mergeSort(arr, 0, length-1)
	if !tools.CheckIsSort(arr) {
		t.Errorf("merge sort failed!")
	}

	arr = []int{3,41,26,52,38,57,9,49}
	t.Log(mergeSort(arr,0,7))
}

func BenchmarkMerge(b *testing.B) {
	b.StopTimer()
	length := 1000
	arr := tools.RandArr(length, 20000)
	sortArr := make([]int, length)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(sortArr, arr)
		b.StartTimer()
		mergeSort(sortArr, 0, length-1)
	}
}
func mergeSort(arr []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) / 2
	inversionNumLeft := mergeSort(arr, l, mid)
	inversionNumRight := mergeSort(arr, mid+1, r)
	inversionNum := merge(arr, l, r)
	return inversionNumLeft + inversionNumRight + inversionNum
}

// inversionNum 逆序对数目 p2.4
func merge(arr []int, l, r int) (inversionNum int) {
	mid := (l + r) / 2
	temp := make([]int, r-l+1)
	m, n := l, mid+1
	i := 0
	for m <= mid && n <= r {
		if arr[m] > arr[n] {
			temp[i] = arr[n]
			n++
		} else {
			temp[i] = arr[m]
			inversionNum += n - mid - 1
			m++
		}
		i++
	}
	if m <= mid {
		for _, v := range arr[m : mid+1] {
			inversionNum += n - mid - 1
			temp[i] = v
			i++
		}
	}
	if n <= r {
		for _, v := range arr[n : r+1] {
			temp[i] = v
			i++
		}
	}
	for k, v := range temp {
		arr[l+k] = v
	}
	return
}
