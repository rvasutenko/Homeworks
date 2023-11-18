package main

import "fmt"

func main() {
	return
}

type Queue struct {
	items []int
	size  int
}

func (q *Queue) len() int {
	return len(q.items)
}

func (q *Queue) isEmpty() bool {
	return q.len() == 0
}

func (q *Queue) enqueue(item int) error {
	if q.size == q.len() {
		return fmt.Errorf("Очередь переполнена")
	}
	q.items = append(q.items, item)
	return nil
}

func (q *Queue) dequeue() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Очередь пуста")
	}
	item := q.items[0]
	if q.len() == 1 {
		q.items = nil
		return item, nil
	}
	q.items = q.items[1:]
	return item, nil
}
