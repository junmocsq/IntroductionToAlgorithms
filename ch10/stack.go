package ch10

type Stack interface {
	Empty() bool
	Push(val interface{}) bool
	Pop() (val interface{}, ok bool)
}
type stackEle struct {
	val interface{}
}

type StackWithSlice struct {
	arr []*stackEle
}

func NewStackWithSlice() *StackWithSlice {
	return &StackWithSlice{}
}

func (s *StackWithSlice) Empty() bool {
	if len(s.arr) == 0 {
		return true
	}
	return false
}

func (s *StackWithSlice) Push(val interface{}) bool {
	ele := &stackEle{
		val: val,
	}
	s.arr = append(s.arr, ele)
	return true
}
func (s *StackWithSlice) Pop() (val interface{}, ok bool) {
	if s.Empty() {
		return
	}
	length := len(s.arr)
	e := s.arr[length-1]
	s.arr = s.arr[:length-1]
	return e.val, true
}

type StackWithArray struct {
	arr      []*stackEle
	length   int
	capacity int
}

func NewStackWithArray(capacity int) *StackWithArray {
	return &StackWithArray{
		arr:      make([]*stackEle, capacity),
		length:   0,
		capacity: capacity,
	}
}

func (s *StackWithArray) Empty() bool {
	if s.length == 0 {
		return true
	}
	return false
}

func (s *StackWithArray) full() bool {
	if s.length == s.capacity {
		return true
	}
	return false
}

func (s *StackWithArray) Push(val interface{}) bool {
	if s.full() {
		return false
	}
	ele := &stackEle{
		val: val,
	}
	s.arr[s.length] = ele
	s.length++
	return true
}
func (s *StackWithArray) Pop() (val interface{}, ok bool) {
	if s.Empty() {
		return
	}
	e := s.arr[s.length-1]
	s.length--
	return e.val, true
}

type stackLinkedEle struct {
	ele  *stackEle
	next *stackLinkedEle
	pre  *stackLinkedEle
}
type StackWithLinked struct {
	tail   *stackLinkedEle
	length int
}

func NewStackWithLinked() *StackWithLinked {
	return &StackWithLinked{}
}
func (s *StackWithLinked) addTail(val interface{}) bool {
	ele := &stackLinkedEle{
		ele: &stackEle{
			val,
		},
		next: nil,
		pre:  nil,
	}
	if !s.Empty() {
		s.tail.next = ele
		ele.pre = s.tail
	}
	s.tail = ele
	s.length++
	return true
}
func (s *StackWithLinked) deleteTail() (val interface{}, ok bool) {
	if s.Empty() {
		return
	}
	e := s.tail
	s.tail = e.pre
	if s.tail != nil {
		s.tail.next = nil
	}
	s.length--
	return e.ele.val, true
}

func (s *StackWithLinked) Empty() bool {
	if s.length == 0 {
		return true
	}
	return false
}

func (s *StackWithLinked) Push(val interface{}) bool {
	return s.addTail(val)
}
func (s *StackWithLinked) Pop() (val interface{}, ok bool) {
	return s.deleteTail()
}
