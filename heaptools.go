package heaptools

import (
	"container/heap"
	"reflect"
)

var _ heap.Interface = &sliceHeap{}

type sliceHeap struct {
	slice   reflect.Value
	less    func(i, j int) bool
	swapper func(i, j int)
}

func (h *sliceHeap) Len() int {
	return h.slice.Elem().Len()
}

func (h *sliceHeap) Less(i, j int) bool {
	return h.less(i, j)
}

func (h *sliceHeap) Swap(i, j int) {
	if i == j {
		return
	}
	h.swapper(i, j)
}

func (h *sliceHeap) Push(x interface{}) {
	e := h.slice.Elem()
	e.Set(reflect.Append(e, reflect.ValueOf(x)))
	h.swapper = reflect.Swapper(e.Interface())
}

func (h *sliceHeap) Pop() interface{} {
	e := h.slice.Elem()
	last := e.Index(e.Len() - 1)
	e.SetLen(e.Len() - 1)
	return last.Interface()
}

func NewSliceHeap(slice interface{}, less func(i, j int) bool) heap.Interface {
	v := reflect.ValueOf(slice)
	sh := &sliceHeap{
		slice:   v,
		less:    less,
		swapper: reflect.Swapper(v.Elem().Interface()),
	}
	heap.Init(sh)
	return sh
}
