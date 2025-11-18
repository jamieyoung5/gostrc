package gostrc_test

import (
	"github.com/jamieyoung5/gostrc"
	"testing"
)

func TestCircularQueue_EnqueueDequeue(t *testing.T) {
	cq := gostrc.NewCircularQueue[int](3)

	if !cq.IsEmpty() {
		t.Error("new queue should be empty")
	}

	if !cq.Enqueue(1) || !cq.Enqueue(2) || !cq.Enqueue(3) {
		t.Error("failed to enqueue items into empty queue")
	}

	if !cq.IsFull() {
		t.Error("queue should be full")
	}

	if cq.Enqueue(4) {
		t.Error("should not be able to enqueue to full queue")
	}

	val, ok := cq.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %d (ok: %v)", val, ok)
	}

	if !cq.Enqueue(4) {
		t.Error("should be able to enqueue after dequeueing from full queue")
	}

	expected := []int{2, 3, 4}
	for _, exp := range expected {
		val, ok := cq.Dequeue()
		if !ok || val != exp {
			t.Errorf("expected %d, got %d", exp, val)
		}
	}

	if !cq.IsEmpty() {
		t.Error("queue should be empty")
	}
}

func TestCircularQueue_Peek(t *testing.T) {
	cq := gostrc.NewCircularQueue[string](2)
	cq.Enqueue("first")
	cq.Enqueue("second")

	val, ok := cq.Peek()
	if !ok || val != "first" {
		t.Errorf("expected 'first', got %s", val)
	}

	if cq.Size() != 2 {
		t.Error("Peek should not remove element")
	}
}

func TestCircularQueue_EmptyDequeue(t *testing.T) {
	cq := gostrc.NewCircularQueue[int](1)
	_, ok := cq.Dequeue()
	if ok {
		t.Error("Dequeue on empty queue should return false")
	}
}
