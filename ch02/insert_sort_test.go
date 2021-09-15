package ch02

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"testing"
)

func InsertSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for ; i >= 0 && arr[i] > key; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = key
	}
}

func TestInsertSort(t *testing.T) {
	arr := tools.RandArr(100,20000)
	InsertSort(arr)
	if !tools.CheckIsSort(arr){
		t.Errorf("insert sort failed!")
	}
}
