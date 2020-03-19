# heaptools

heaptools provide helper functions to make initializing slice heaps easier.

## Install

`go get github.com/suzaku/heaptools`

## Example

```go
package main

import (
    "fmt"
    "container/heap"


    "github.com/suzaku/heaptools"
)

type Item struct {
	value string
}

func main() {
    s := []Item{
        {value: "hello"},
        {value: "world"},
        {value: "one"},
        {value: "four"},
    }
    sh := heaptools.NewSliceHeap(&s, func(i, j int) bool {
        vi, vj := s[i].value, s[j].value
        if len(vi) == len(vj) {
            return vi < vj
        }
        return len(vi) < len(vj)
    })
    for sh.Len() > 0 {
        it, _ := heap.Pop(sh).(Item)
        fmt.Println(it.value)
    }
}
```


