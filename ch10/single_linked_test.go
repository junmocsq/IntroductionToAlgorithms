package ch10

import (
	"github.com/junmocsq/IntroductionToAlgorithms/tools"
	"testing"
)

func TestNewSingleLinked(t *testing.T) {
	linked := NewSingleLinked()
	arr := tools.RandArr(20, 1000)
	for _, v := range arr {
		linked.Add(v)
	}
	for k, v := range linked.Elements() {
		if arr[k] != v.(int) {
			t.Error("single linked Add failed!")
		}
	}
	if v, ok := linked.DeleteTail(); !ok || v.(int) != arr[19] {
		t.Error("single linked Delete Tail failed!")
	}
	if v, ok := linked.DeleteHead(); !ok || v.(int) != arr[0] {
		t.Error("single linked Delete Head failed!")
	}

	v := 999999
	if !linked.InsertHead(v) {
		t.Error("single linked Insert Head failed!")
	}
	if !linked.Insert(1, v) {
		t.Error("single linked Insert failed!")
	}
	if !linked.InsertTail(v) {
		t.Error("single linked Insert Tail failed!")
	}
	eles := linked.Elements()
	if eles[0] != v || eles[1] != v || eles[linked.length-1] != v {
		t.Error("single linked Insert failed!")
	}

	ele := linked.FindByIndex(10)
	if ele == nil {
		t.Error("single linked FindByIndex failed!")
	}
	if linked.Find(ele) < 0 {
		t.Error("single linked Find failed!")
	}

	if linked.Delete(ele) < 0 {
		t.Error("single linked Delete failed!")
	}

	eles = linked.Elements()
	if v, ok := linked.DeleteTail(); !ok || v.(int) != eles[len(eles)-1].(int) {
		t.Error("single linked DeleteTail failed!")
	}
	if v, ok := linked.DeleteHead(); !ok || v.(int) != eles[0].(int) {
		t.Error("single linked DeleteHead failed!")
	}
	eles = linked.Elements()
	if v, ok := linked.DeleteByIndex(10); !ok || v.(int) != eles[10].(int) {
		t.Error("single linked DeleteByIndex failed!")
	}
}
