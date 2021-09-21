package ch08

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"testing"
)

// 分桶
// 插入排序排序桶内元素
func BucketSort(arr []int, start, end, period int) {
	var periodArr []int
	for i := start; i <= end; i += period {
		periodArr = append(periodArr, i)
	}
	tempArr := make([][]int, len(periodArr))
	var index int
	// 分桶
	for _, v := range arr {
		index = -1
		for _, _v := range periodArr {
			if v >= _v {
				index++
			} else {
				break
			}
		}
		tempArr[index] = append(tempArr[index], v)
	}
	index = 0
	for _, items := range tempArr {
		insertionSort(items) // 插入排序
		for _, v := range items {
			arr[index] = v
			index++
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
}

func TestBucketSort(t *testing.T) {
	arr := tools.RandArr(20, 50)
	t.Log(arr)
	BucketSort(arr, 0, 50, 5)
	t.Log(arr)
}

func BenchmarkBucketSort(b *testing.B) {
	b.StopTimer()
	arr := tools.RandArr(200, 10000)
	temp := make([]int, 200)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(temp, arr)
		b.StartTimer()
		BucketSort(temp, 0, 10000, 1000)
	}
}
