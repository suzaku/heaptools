package heaptools

import (
	"container/heap"
	"testing"
)

type Item struct {
	value string
}

func TestNewSliceHeap(t *testing.T) {
	t.Parallel()
	t.Run("Should work with slices of primitive type", func(t *testing.T) {
		s := []int{1, 9, 8, 4}
		sh := NewSliceHeap(&s, func(i, j int) bool { return s[i] < s[j] })
		assertEq(t, 1, heap.Pop(sh))
		heap.Push(sh, 10)
		assertEq(t, 4, heap.Pop(sh))
		heap.Push(sh, 2)
		heap.Remove(sh, 1)
		assertEq(t, 2, heap.Pop(sh))
		assertEq(t, 9, heap.Pop(sh))
		assertEq(t, 10, heap.Pop(sh))
		assertEq(t, 0, sh.Len())
	})
	t.Run("Should also work with slices of structs", func(t *testing.T) {
		s := []Item{
			{value: "hello"},
			{value: "world"},
			{value: "one"},
			{value: "four"},
		}
		sh := NewSliceHeap(&s, func(i, j int) bool {
			vi, vj := s[i].value, s[j].value
			if len(vi) == len(vj) {
				return vi < vj
			}
			return len(vi) < len(vj)
		})
		assertContainValue(t, "one", heap.Pop(sh))
		assertContainValue(t, "four", heap.Pop(sh))
		assertContainValue(t, "hello", heap.Pop(sh))
		assertContainValue(t, "world", heap.Pop(sh))
		assertEq(t, 0, sh.Len())
	})
}

func assertEq(t *testing.T, expect int, got interface{}) {
	if gotV, ok := got.(int); !ok {
		t.Fatalf("Got non-int value: %v", got)
	} else {
		if expect != gotV {
			t.Fatalf("Expect %d, got %d", expect, gotV)
		}
	}
}

func assertContainValue(t *testing.T, expect string, got interface{}) {
	if gotVal, ok := got.(Item); !ok {
		t.Fatalf("Unknown popped value: %v", gotVal)
	} else {
		if expect != gotVal.value {
			t.Fatalf("Expect %s, got %s", expect, gotVal.value)
		}
	}
}
