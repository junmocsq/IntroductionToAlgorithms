package ch08

import (
	"fmt"
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"math"
	"testing"
)

// 比较排序：在排序的最终结果中，各元素的次序依赖于它们之间的比较。

// 计数排序：假设n个输入元素中每一个都是0到k区间的一个整数。基本思想是，对于每一个输入元素，确定小于x的元素个数。
// 它是稳定的，具有相同值的元素在输出数组中的相对次序与它们输入数组中的相对次序相同。
func CountingSort(arr []int) {
	min, max := math.MaxInt64, math.MinInt64
	// 寻找当前最小值和最大值
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// 数组元素计数，第i位元素的值为(i+min),元素个数为counting[i]
	counting := make([]int, max-min+1)
	for _, v := range arr {
		counting[v-min]++
	}
	// 累计计数
	for i := 1; i < len(counting); i++ {
		counting[i] += counting[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i] - min // counting的下标
		// counting[index]表示不大于值(index+min)即arr[i]的个数，数组下标是从0开始的，而计数是从1开始，所以需要减1
		result[counting[index]-1] = arr[i]
		counting[index]--
	}
	copy(arr, result)
}

func TestCountingSort(t *testing.T) {
	arr := tools.RandArr(50, 10)
	CountingSort(arr)
	if !tools.CheckIsSort(arr) {
		t.Errorf("Counting sort failed!")
	}
	arr = tools.RandArr(20, 10)
	arrs := make([]*CountingValue, len(arr))
	for k, v := range arr {
		arrs[k] = &CountingValue{
			key: v,
			val: k,
		}
	}
	//t.Log(arr)
	CountingSort2(arrs)
	for _, v := range arrs {
		fmt.Printf("%d %d|", v.key, v.val)
	}
	t.Log()
}

func BenchmarkCountingSort(b *testing.B) {
	b.StopTimer()
	arr := tools.RandArr(200, 1000)
	temp := make([]int, 200)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(temp, arr)
		b.StartTimer()
		CountingSort(temp)
	}
}

type CountingValue struct {
	key int
	val interface{}
}

func CountingSort2(arr []*CountingValue) {
	min, max := math.MaxInt64, math.MinInt64
	// 寻找当前最小值和最大值
	for _, v := range arr {
		if v.key < min {
			min = v.key
		}
		if v.key > max {
			max = v.key
		}
	}

	// 数组元素计数，第i位元素的值为(i+min),元素个数为counting[i]
	counting := make([]int, max-min+1)
	for _, v := range arr {
		counting[v.key-min]++
	}
	// 累计计数
	for i := 1; i < len(counting); i++ {
		counting[i] += counting[i-1]
	}

	result := make([]*CountingValue, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i].key - min // counting的下标
		// counting[index]表示不大于值(index+min)即arr[i].key的个数，数组下标是从0开始的，而计数是从1开始，所以需要减1
		result[counting[index]-1] = arr[i]
		counting[index]--
	}
	copy(arr, result)
}
