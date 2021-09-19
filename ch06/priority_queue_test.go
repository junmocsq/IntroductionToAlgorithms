package ch06

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"testing"
)

// 优先队列
type priQueue struct {
	arr    []int
	length int
	isMax  bool // 最大堆为true 最小堆为false
}

func NewPriorityQueue(isMax bool) *priQueue {
	return &priQueue{isMax: isMax}
}

func (p *priQueue) Insert(val int) {
	if len(p.arr) == p.length {
		p.arr = append(p.arr, val)
	} else {
		p.arr[p.length] = val
	}
	p.length++
	p.up(p.length - 1)
}

func (p *priQueue) compare(i, j int) bool {
	res := i > j
	if p.isMax {
		return res
	} else {
		return !res
	}
}
func (p *priQueue) up(i int) {
	for i > 0 {
		parent := p.parent(i)
		if p.compare(p.arr[i], p.arr[parent]) {
			p.arr[parent], p.arr[i] = p.arr[i], p.arr[parent]
			i = parent
		} else {
			break
		}
	}
}

// 6.5-6 p93
func (p *priQueue) up2(i int) {
	temp := p.arr[i]
	for i > 0 {
		parent := p.parent(i)
		if p.compare(temp, p.arr[parent]) {
			p.arr[i] = p.arr[parent]
			i = parent
		} else {
			break
		}
	}
	p.arr[i] = temp
}

func (p *priQueue) down(i int) {
	for {
		largest := i
		l, r := p.left(i), p.right(i)
		isChange := false
		if l < p.length && p.compare(p.arr[l], p.arr[largest]) {
			largest = l
			isChange = true
		}
		if r < p.length && p.compare(p.arr[r], p.arr[largest]) {
			largest = r
			isChange = true
		}
		if isChange {
			p.arr[i], p.arr[largest] = p.arr[largest], p.arr[i]
			i = largest
		} else {
			break
		}
	}
}
func (p *priQueue) left(i int) int {
	return 2*i + 1
}
func (p *priQueue) parent(i int) int {
	return (i - 1) / 2
}
func (p *priQueue) right(i int) int {
	return p.left(i) + 1
}
func (p *priQueue) empty() bool {
	return p.length == 0
}
func (p *priQueue) First() (val int, ok bool) {
	if !p.empty() {
		val = p.arr[0]
		ok = true
	}
	return
}
func (p *priQueue) ExtractFirst() (v int, ok bool) {
	if p.empty() {
		return
	}
	v = p.arr[0]
	ok = true
	p.length--
	p.arr[0] = p.arr[p.length]
	p.down(0)
	return
}
func (p *priQueue) Increase(index, incr int) {
	if p.length <= index {
		return
	}
	p.arr[index] += incr
	if incr == 0 {
		return
	} else if incr > 0 {
		p.up(index)
	} else {
		p.down(index)
	}
}

// 6.5-8 p93
func (p *priQueue) Delete(index int) (val int, ok bool) {
	if p.length <= index {
		return
	}
	val = p.arr[index]
	ok = true
	p.length--
	p.arr[index] = p.arr[p.length]
	parent := p.parent(index)
	if p.compare(p.arr[index], p.arr[parent]) {
		p.up(index)
	} else {
		p.down(index)
	}
	return
}

func TestPriorityQueue(t *testing.T) {
	arr := []int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}
	p := NewPriorityQueue(true)
	for _, v := range arr {
		p.Insert(v)
	}
	var res []int
	for {
		r, ok := p.ExtractFirst()
		if !ok {
			break
		}
		res = append(res, r)
	}
	if !tools.CheckIsSort(res) {
		t.Errorf("Priority Queue Sort failed!")
	}
	for _, v := range arr {
		p.Insert(v)
	}
	t.Log(p.arr[:p.length])
	p.Delete(1)
	t.Log(p.arr[:p.length])
}
