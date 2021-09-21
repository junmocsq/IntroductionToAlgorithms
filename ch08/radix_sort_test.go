package ch08

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"testing"
)

// 基数排序：按最低位有效位排序解决问题

func RadixSort(arr []int, num int) {
	arrs := make([]*CountingValue, len(arr))
	for i := 0; i < num; i++ {
		div := 1
		for j := 0; j < i; j++ {
			div *= 10
		}
		for k, v := range arr {
			if arrs[k] == nil {
				arrs[k] = &CountingValue{}
			}
			arrs[k].key = v / div % 10
			arrs[k].val = v
		}
		CountingSort2(arrs)
		for k, v := range arrs {
			arr[k] = v.val.(int)
		}
	}
}

func TestRadixSort(t *testing.T) {
	arr := tools.RandArr(200, 100000000)
	RadixSort(arr, 8)
	if !tools.CheckIsSort(arr) {
		t.Errorf("Radix sort failed!")
	}
}

func BenchmarkRadixSort(b *testing.B) {
	b.StopTimer()
	arr := tools.RandArr(200, 10000)
	temp := make([]int, 200)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(temp, arr)
		b.StartTimer()
		RadixSort(temp, 4)
	}
}
