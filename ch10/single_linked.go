package ch10

import (
	"fmt"
	"reflect"
)

type singleEle struct {
	val  interface{}
	next *singleEle
}

type SingleLinked struct {
	root   *singleEle
	tail   *singleEle
	length int
}

func NewSingleLinked() *SingleLinked {
	return &SingleLinked{}
}

func (s *SingleLinked) Empty() bool {
	return s.length == 0
}

func (s *SingleLinked) Clear() {
	s.root = nil
	s.tail = nil
	s.length = 0
}

func (s *SingleLinked) Add(val interface{}) bool {
	ele := &singleEle{val: val}
	if s.Empty() {
		s.root = ele
	} else {
		s.tail.next = ele
	}
	s.tail = ele
	s.length++
	return true
}

func (s *SingleLinked) Insert(index int, val interface{}) bool {
	if index == 0 {
		return s.InsertHead(val)
	} else if index == s.length {
		return s.InsertTail(val)
	} else if index > s.length {
		return false
	}
	temp := s.root
	for i := 1; i < index; i++ {
		temp = temp.next
	}
	next := temp.next
	temp.next = &singleEle{
		val:  val,
		next: next,
	}
	s.length++
	return true
}

func (s *SingleLinked) InsertHead(val interface{}) bool {
	ele := &singleEle{
		val:  val,
		next: nil,
	}
	if s.Empty() {
		s.tail = ele
	} else {
		ele.next = s.root
	}
	s.root = ele
	s.length++
	return true
}

func (s *SingleLinked) InsertTail(val interface{}) bool {
	return s.Add(val)
}

func (s *SingleLinked) Find(val interface{}) (index int) {
	index = -1
	temp := s.root
	for temp != nil {
		index++
		if reflect.DeepEqual(temp.val, val) {
			break
		}
		temp = temp.next
	}
	return index
}

func (s *SingleLinked) FindByIndex(index int) (val interface{}) {
	if index >= s.length {
		return nil
	}
	temp := s.root
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (s *SingleLinked) Delete(val interface{}) (index int) {
	index = -1
	if s.Empty() {
		return
	}
	if reflect.DeepEqual(s.root.val, val) {
		s.root = nil
		s.tail = nil
		s.length--
		index = 0
		return
	}
	if s.root.next != nil {
		index = 0
		temp := s.root

		for temp.next != nil {
			index++
			if reflect.DeepEqual(temp.next.val, val) {
				temp.next = temp.next.next
				if temp.next == nil {
					s.tail = temp
				}
				s.length--
				return index
			}
			temp = temp.next
		}
	}
	return -1
}

func (s *SingleLinked) DeleteTail() (interface{}, bool) {
	return s.DeleteByIndex(s.length - 1)
}

func (s *SingleLinked) DeleteHead() (interface{}, bool) {
	if s.Empty() {
		return nil, false
	}
	res := s.root
	next := s.root.next
	if next == nil {
		s.tail = nil
	}
	s.root = next
	s.length--
	return res.val, true
}

func (s *SingleLinked) DeleteByIndex(index int) (interface{}, bool) {
	if index >= s.length {
		return nil, false
	}
	if index == 0 {
		return s.DeleteHead()
	}
	temp := s.root
	for i := 1; i < index; i++ {
		temp = temp.next
	}
	res := temp.next
	temp.next = temp.next.next
	if temp.next == nil {
		s.tail = temp
	}
	s.length--
	return res.val, true
}

func (s *SingleLinked) Print() {
	fmt.Println("length:", s.length)
	temp := s.root
	for i := 0; i < s.length; i++ {
		fmt.Printf("%d:%v \t", i, temp.val)
		if (i+1)%5 == 0 {
			fmt.Println()
		}
		temp = temp.next
	}
	fmt.Println()
}

func (s *SingleLinked) Elements() []interface{} {
	var result = make([]interface{}, s.length)
	temp := s.root
	for i := 0; i < s.length; i++ {
		//fmt.Println(i,temp)
		result[i] = temp.val
		temp = temp.next
	}
	//fmt.Println("length:", s.length,s.root,s.tail)
	return result
}
