package ch10

import "testing"

func TestStack(t *testing.T) {
	var stack Stack

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	f := func(s Stack) {
		_, ok := s.Pop()
		if ok {
			t.Error("Pop failed", s)
		}
		for _, v := range arr {
			ok := s.Push(v)
			if !ok {
				t.Error("Push failed", s)
				t.FailNow()
			}
		}
		for i := len(arr) - 1; i >= 0; i-- {
			v, ok := s.Pop()
			if !ok {
				t.Error("Pop failed", s)
				t.FailNow()
			}
			if v.(int) != arr[i] {
				t.Error("Pop Push failed", s)
			}
		}
	}

	stack = NewStackWithSlice()
	f(stack)
	stack = NewStackWithArray(len(arr))
	f(stack)
	stack = NewStackWithLinked()
	f(stack)
}
