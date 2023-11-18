package main

import "fmt"

type Deque struct {
	items []int
	size  int
}

func (d *Deque) len() int {
	return len(d.items)
}

func (d *Deque) isEmpty() bool {
	return d.len() == 0
}

func (d *Deque) pushFront(item int) error {
	if d.size == d.len() {
		return fmt.Errorf("Дек переполнен")
	}
	d.items = append([]int{item}, d.items...)
	return nil
}

func (d *Deque) popFront() (int, error) {
	if d.isEmpty() {
		return 0, fmt.Errorf("Дек пуст")
	}
	item := d.items[0]
	if d.len() == 1 {
		d.items = nil
		return item, nil
	}
	d.items = d.items[1:]
	return item, nil
}

func (d *Deque) pushBack(item int) error {
	if d.size == d.len() {
		return fmt.Errorf("Дек переполнен")
	}
	d.items = append(d.items, item)
	return nil
}

func (d *Deque) popBack() (int, error) {
	if d.isEmpty() {
		return 0, fmt.Errorf("Дек пуст")
	}
	item := d.items[d.len()-1]
	if d.len() == 1 {
		d.items = nil
		return item, nil
	}
	d.items = d.items[:d.len()-1]
	return item, nil
}
