package ch06

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"reflect"
	"testing"
)

func left(parent int) int {
	return 1 + parent*2
}
func right(parent int) int {
	return left(parent) + 1
}
func parent(child int) int {
	return (child - 1) / 2
}

// 下沉
func maxHeapify(arr []int, i int) {
	parent := i
	largest := i
	for {
		l, r := left(parent), right(parent)
		isChange := false
		if l < len(arr) && arr[largest] < arr[l] {
			largest = l
			isChange = true
		}
		if r < len(arr) && arr[largest] < arr[r] {
			largest = r
			isChange = true
		}
		if isChange {
			arr[parent], arr[largest] = arr[largest], arr[parent]
			parent = largest
		} else {
			break
		}
	}
}

func minHeapify(arr []int, i int) {
	l, r := left(i), right(i)
	min := i
	if l < len(arr) && arr[min] > arr[l] {
		min = l
	}
	if r < len(arr) && arr[min] > arr[r] {
		min = r
	}
	if min != i {
		arr[i], arr[min] = arr[min], arr[i]
		maxHeapify(arr, min)
	}
}

func buildMaxHeap(arr []int) {
	for i := parent(len(arr) - 1); i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

func TestExercise6_3_1(t *testing.T) {
	arr := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	buildMaxHeap(arr)
	if !reflect.DeepEqual(arr, []int{84, 22, 19, 10, 3, 17, 6, 5, 9}) {
		t.Errorf("Create MaxHeap   failed!")
	}
	HeapSort(arr)
	if !tools.CheckIsSort(arr) {
		t.Errorf("MaxHeap Sort failed!")
	}
}

func HeapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr[:i], 0)
	}
}
