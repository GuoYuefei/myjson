package myjson

import (
	"fmt"
	"testing"
)

func TestStackAnaly(t *testing.T) {
	s := NewStackAnaly()
	if !s.IsEmpty() {
		t.Error("test err", i)
	}
	if s.Pop() != 0 {
		t.Error("test err", i)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Size() != 3 {
		t.Error("test err", i)
	}
	s.Push(44)
	if s.Top() != 44 {
		t.Error("test err", i)
	}

	if s.Pop() != 44 {
		t.Error("test err", i)
	}

	if s.Size() != 3 {
		t.Error("test err", i)
	}

	s.Clear()
	if s.Size() != 0 {
		t.Error("err", i)
	}

	if !s.IsEmpty() {
		t.Error("err", i)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(111)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
	if s.Size() != 0 {
		t.Error("err", i)
	}
}
