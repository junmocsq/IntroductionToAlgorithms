package ch10

import (
	"fmt"
	"reflect"
)

// 带哨兵循环列表
type doubleEle struct {
	val       interface{}
	next, pre *doubleEle
}

type DoubleLinked struct {
	root *doubleEle
	length int
}


func NewDoubleLinked() *DoubleLinked {
	sentinel := &doubleEle{
		val: nil,
	}
	sentinel.next = sentinel
	sentinel.pre = sentinel
	return &DoubleLinked{
		root:   sentinel,
		length: 0,
	}
}

func (l *DoubleLinked) Search(k interface{}) (*doubleEle, bool) {
	x := l.root.next
	for x != l.root && !reflect.DeepEqual(x.val, k) {
		x = x.next
	}
	return x, reflect.DeepEqual(x.val, k)
}

func (l *DoubleLinked) Insert(k interface{}) {
	x := l.root
	ele := &doubleEle{
		val:  k,
		next: x.next,
		pre:  x,
	}
	ele.next.pre = ele
	ele.pre.next = ele
	l.length++
}
func (l *DoubleLinked) Delete(k interface{}) *doubleEle {
	v, ok := l.Search(k)
	if !ok {
		return nil
	}
	v.pre.next = v.next
	v.next.pre = v.pre
	l.length--
	return v
}

func (l *DoubleLinked) Print() {
	x := l.root.next
	for x != l.root {
		fmt.Printf("%v ", x.val)
		x = x.next
	}
}

