package main

import "fmt"

type BinaryHeap struct {
	items []int
	size  int
}

func (b *BinaryHeap) len() int {
	return len(b.items)
}

func (b *BinaryHeap) isEmpty() bool {
	return b.len() == 0
}

func (b *BinaryHeap) push(item int) error {
	if b.len() == b.size {
		return fmt.Errorf("Куча переполнена")
	}
	b.items = append(b.items, item)
	b.heapify()
	return nil
}

func (b *BinaryHeap) pop() (int, error) {
	if b.isEmpty() {
		return 0, fmt.Errorf("Куча пуста")
	}
	item := b.items[0]
	if b.len() == 1 {
		b.items = nil
		return item, nil
	}
	b.items[0], b.items[b.len()-1] = b.items[b.len()-1], b.items[0]
	b.heapify()
	return item, nil
}

func (b *BinaryHeap) heapify() {
	for i := 1; i <= b.len()/2; {
		maxId := i
		if b.items[i] < b.items[i*2] {
			maxId = i * 2
		}
		if i*2+1 <= b.len() && b.items[maxId] < b.items[i*2+1] {
			maxId = i*2 + 1
		}
		if maxId == i {
			break
		}
		b.items[i], b.items[maxId] = b.items[maxId], b.items[i]
		i = maxId
	}
}
