package ch10

type queueEle struct {
	val interface{}
}

type Queue interface {
	Empty() bool
	Enqueue(val interface{}) bool
	Dequeue() (val interface{}, ok bool)
}

type QueueWithArray struct {
	arr        []*queueEle
	length     int
	capacity   int
	start, end int // 元素开始和结束，环
}

func NewQueueWithArray(capacity int) *QueueWithArray {
	return &QueueWithArray{
		arr:      make([]*queueEle, capacity),
		length:   0,
		capacity: capacity,
	}
}
func (q *QueueWithArray) Empty() bool {
	if q.length == 0 {
		return true
	}
	return false
}
func (q *QueueWithArray) full() bool {
	if q.length == q.capacity {
		return true
	}
	return false
}
func (q *QueueWithArray) Enqueue(val interface{}) bool {
	if q.full() {
		return false
	}
	ele := &queueEle{val: val}
	next := q.end + 1
	if next == q.capacity {
		next = 0
	}
	q.arr[next] = ele
	q.length++
	q.end = next
	return true
}

func (q *QueueWithArray) Dequeue() (val interface{}, ok bool) {
	if q.Empty() {
		return
	}
	ele := q.arr[q.start]
	q.start++
	if q.start == q.capacity {
		q.start = 0
	}
	q.length--
	return ele.val, true
}

type QueueWithSlice struct {
	arr []*queueEle
}

func NewQueueWithSlice() *QueueWithSlice {
	return &QueueWithSlice{}
}
func (q *QueueWithSlice) Empty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

func (q *QueueWithSlice) Enqueue(val interface{}) bool {

	ele := &queueEle{val: val}
	q.arr = append(q.arr, ele)
	return true
}

func (q *QueueWithSlice) Dequeue() (val interface{}, ok bool) {
	if q.Empty() {
		return
	}
	ele := q.arr[0]
	q.arr = q.arr[1:]
	return ele.val, true
}

type queueLinkedEle struct {
	ele       *queueEle
	next, pre *queueLinkedEle
}

type QueueWithLinked struct {
	length     int
	root, tail *queueLinkedEle
}

func NewQueueWithLinked() *QueueWithLinked {
	return &QueueWithLinked{}
}

func (q *QueueWithLinked) addTail(val interface{}) bool {
	ele := &queueLinkedEle{
		ele: &queueEle{
			val,
		},
		next: nil,
		pre:  nil,
	}
	if q.Empty() {
		q.root = ele
	} else {
		q.tail.next = ele
		ele.pre = q.tail
	}
	q.tail = ele
	q.length++
	return true
}
func (q *QueueWithLinked) deleteHead() (val interface{}, ok bool) {
	if q.Empty() {
		return
	}
	e := q.root
	q.root = e.next
	if q.root != nil {
		q.root.pre = nil
	} else {
		q.tail = nil
	}
	q.length--
	return e.ele.val, true
}

func (q *QueueWithLinked) Empty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

func (q *QueueWithLinked) Enqueue(val interface{}) bool {
	return q.addTail(val)
}

func (q *QueueWithLinked) Dequeue() (val interface{}, ok bool) {
	return q.deleteHead()
}
