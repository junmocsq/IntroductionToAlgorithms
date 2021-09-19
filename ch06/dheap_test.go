package ch06

import "testing"

// d叉堆 p93 6-2

type dheap struct {
	arr   []int
	d     int
	isMax bool
}

func NewDHeap(d int, isMax bool) *dheap {
	return &dheap{d: d, isMax: isMax}
}
func (dh *dheap) childLeft(i int) int {
	return dh.d*i + 1
}

func (dh *dheap) childRight(i int) int {
	return dh.d*(i+1) - 1
}

func (dh *dheap) parent(i int) int {
	return (i - 1) / dh.d
}
func (dh *dheap) compare(i, j int) bool {
	if dh.isMax {
		return i > j
	} else {
		return i < j
	}
}
func (dh *dheap) Insert(val int) {
	dh.arr = append(dh.arr, val)
	dh.up(len(dh.arr) - 1)
}

func (dh *dheap) ExtractFirst() (val int, ok bool) {
	if len(dh.arr) == 0 {
		return
	}
	val = dh.arr[0]
	ok = true
	dh.arr[0] = dh.arr[len(dh.arr)-1]
	dh.arr = dh.arr[:len(dh.arr)-1]
	dh.down(0)
	return
}

func (dh *dheap) IncreaseKey(i, incr int) {
	dh.arr[i] += incr
	if incr == 0 {
		return
	} else if incr > 0 {
		dh.up(i)
	} else {
		dh.down(i)
	}
	return
}

func (dh *dheap) up(i int) {
	for i > 0 {
		p := dh.parent(i)

		if dh.compare(dh.arr[i], dh.arr[p]) {
			dh.arr[i], dh.arr[p] = dh.arr[p], dh.arr[i]
			i = p
		} else {
			break
		}
	}
}

func (dh *dheap) down(i int) {
	for {
		l, r := dh.childLeft(i), dh.childRight(i)
		largest := i
		isChange := false
		for k := l; k <= r && k < len(dh.arr); k++ {
			if dh.compare(dh.arr[k], dh.arr[largest]) {
				largest = k
				isChange = true
			}
		}
		if isChange {
			dh.arr[i], dh.arr[largest] = dh.arr[largest], dh.arr[i]
			i = largest
		} else {
			break
		}
	}
}

func TestDHeap(t *testing.T) {
	arr := []int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}
	p := NewDHeap(4, false)
	for _, v := range arr {
		p.Insert(v)
	}
	t.Log(p.arr)
}
