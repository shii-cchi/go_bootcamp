package present_heap

import (
	"container/heap"
	"errors"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h *PresentHeap) Len() int { return len(*h) }

func (h *PresentHeap) Less(i, j int) bool {
	if (*h)[i].Value == (*h)[j].Value {
		return (*h)[i].Size < (*h)[j].Size
	}
	return (*h)[i].Value > (*h)[j].Value
}

func (h *PresentHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *PresentHeap) Push(x interface{}) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getNCoolestPresents(n int, presents []Present) ([]Present, error) {
	if n > len(presents) || n < 0 {
		return nil, errors.New("invalid value of n")
	}

	h := &PresentHeap{}

	for _, p := range presents {
		heap.Push(h, p)
	}

	coolestPresents := make([]Present, 0, n)
	for i := 0; i < n; i++ {
		if h.Len() == 0 {
			break
		}
		coolestPresents = append(coolestPresents, heap.Pop(h).(Present))
	}

	return coolestPresents, nil
}
