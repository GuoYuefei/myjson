package myjson

import (
	"fmt"
	"testing"
)

func TestStackAnaly(t *testing.T) {
	s := newStackAnaly()
	if !s.isEmpty() {
		t.Error("test err", s.size())
	}
	s.push(1)
	if s.pop() != 1 {
		t.Error("test err", s.size())
	}
	s.push(2)
	s.push(3)
	if s.size() != 2 {
		t.Error("test err", s.size())
	}
	s.push(44)
	if s.front() != 44 {
		t.Error("test err", s.size())
	}

	if s.pop() != 44 {
		t.Error("test err", s.size())
	}

	if s.size() != 2 {
		t.Error("test err", s.size())
	}

	s.clear()
	if s.size() != 0 {
		t.Error("err", s.size())
	}

	if !s.isEmpty() {
		t.Error("err", s.size())
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
