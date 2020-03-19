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

func BenchmarkPushPop(b *testing.B) {
	var s []int
	sh := NewSliceHeap(&s, func(i, j int) bool { return s[i] < s[j] })
	for i := 0; i < b.N; i++ {
		for j := 0; j < 3; j++ {
			heap.Push(sh, j)
		}
		for j := 0; j < 3; j++ {
			if heap.Pop(sh) != j {
				b.Fail()
			}
		}
	}
	if sh.Len() != 0 {
		b.Fail()
	}
}

// Copied from https://golang.org/src/container/heap/example_intheap_test.go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkPushPopWithExplicitImplementation(b *testing.B) {
	sh := &IntHeap{}
	heap.Init(sh)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 3; j++ {
			heap.Push(sh, j)
		}
		for j := 0; j < 3; j++ {
			if heap.Pop(sh) != j {
				b.Fail()
			}
		}
	}
	if sh.Len() != 0 {
		b.Fail()
	}
}
