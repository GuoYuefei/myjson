package myjson

import (
	"fmt"
	"testing"
)

func TestStackAnaly(t *testing.T) {
	s := newStackAnaly()
	if !s.isEmpty() {
		t.Error("test err", i)
	}
	if s.pop() != 0 {
		t.Error("test err", i)
	}
	s.push(1)
	s.push(2)
	s.push(3)
	if s.size() != 3 {
		t.Error("test err", i)
	}
	s.push(44)
	if s.front() != 44 {
		t.Error("test err", i)
	}

	if s.pop() != 44 {
		t.Error("test err", i)
	}

	if s.size() != 3 {
		t.Error("test err", i)
	}

	s.clear()
	if s.size() != 0 {
		t.Error("err", i)
	}

	if !s.isEmpty() {
		t.Error("err", i)
	}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(111)
	for !s.isEmpty() {
		fmt.Println(s.pop())
	}
	if s.size() != 0 {
		t.Error("err", i)
	}
}
