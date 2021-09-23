package ch10

type doubleEle struct {
	val       interface{}
	next, pre *doubleEle
}

type DoubleLinked struct {
	root   *doubleEle
	tail   *doubleEle
	length int
}
