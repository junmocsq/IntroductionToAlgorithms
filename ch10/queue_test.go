package ch10

import "testing"

func TestQueue(t *testing.T) {
	var queue Queue

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	f := func(s Queue) {
		_, ok := s.Dequeue()
		if ok {
			t.Error("Dequeue failed", s)
		}
		for _, v := range arr {
			ok := s.Enqueue(v)
			if !ok {
				t.Error("Enqueue failed", s)
				t.FailNow()
			}
		}
		for i := 0; i < len(arr); i++ {
			v, ok := s.Dequeue()
			if !ok {
				t.Error("Dequeue failed", s)
				t.FailNow()
			}
			if v.(int) != arr[i] {
				t.Error("Enqueue Dequeue failed", s)
			}
		}
	}
	queue = NewQueueWithSlice()
	f(queue)
	queue = NewQueueWithArray(len(arr))
	f(queue)
	queue = NewQueueWithLinked()
	f(queue)
}
