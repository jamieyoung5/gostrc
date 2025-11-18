package gostrc_test

import (
	"github.com/jamieyoung5/gostrc"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	s := gostrc.NewStack[int]()
	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}

	s.Push(10)
	s.Push(20)

	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}

	if val := s.Pop(); val != 20 {
		t.Errorf("expected 20, got %d", val)
	}

	if val := s.Pop(); val != 10 {
		t.Errorf("expected 10, got %d", val)
	}

	if !s.IsEmpty() {
		t.Error("stack should be empty after popping all elements")
	}
}

func TestStack_Peek(t *testing.T) {
	s := gostrc.NewStack[string]()
	s.Push("A")
	s.Push("B")

	if val := s.Peek(); val != "B" {
		t.Errorf("expected peek 'B', got %s", val)
	}

	// check that Peek didnt remove the item
	if s.Size() != 2 {
		t.Errorf("expected size 2 after peek, got %d", s.Size())
	}
}

func TestStack_EmptyPop(t *testing.T) {
	s := gostrc.NewStack[int]()
	val := s.Pop()
	if val != 0 {
		t.Errorf("expected zero value 0, got %d", val)
	}

	sStr := gostrc.NewStack[string]()
	valStr := sStr.Pop()
	if valStr != "" {
		t.Errorf("expected empty string, got %s", valStr)
	}
}
