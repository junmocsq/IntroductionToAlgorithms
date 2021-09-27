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
		start:    0,
		end:      capacity - 1,
	}
}
func (q *QueueWithArray) Empty() bool {
	return q.length == 0
}
func (q *QueueWithArray) full() bool {
	return q.length == q.capacity
}
func (q *QueueWithArray) Enqueue(val interface{}) bool {
	if q.full() {
		return false
	}
	ele := &queueEle{val: val}
	q.end++
	if q.end == q.capacity {
		q.end = 0
	}
	q.arr[q.end] = ele
	q.length++
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
func (q *QueueWithArray) LPush(val interface{}) bool {
	if q.full() {
		return false
	}
	ele := &queueEle{val: val}
	q.start--
	if q.start == -1 {
		q.start = q.capacity - 1
	}
	q.arr[q.start] = ele
	q.length++
	return true
}
func (q *QueueWithArray) RPop() (val interface{}, ok bool) {
	if q.Empty() {
		return
	}
	ele := q.arr[q.end]
	q.end--
	if q.end == -1 {
		q.end = q.capacity - 1
	}
	q.length--
	return ele.val, true
}

func (q *QueueWithArray) RPush(val interface{}) bool {
	return q.Enqueue(val)
}
func (q *QueueWithArray) LPop() (val interface{}, ok bool) {
	return q.Dequeue()
}

type QueueWithSlice struct {
	arr []*queueEle
}

func NewQueueWithSlice() *QueueWithSlice {
	return &QueueWithSlice{}
}
func (q *QueueWithSlice) Empty() bool {
	return len(q.arr) == 0
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
	return q.length == 0
}

func (q *QueueWithLinked) Enqueue(val interface{}) bool {
	return q.addTail(val)
}

func (q *QueueWithLinked) Dequeue() (val interface{}, ok bool) {
	return q.deleteHead()
}
