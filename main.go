package main

import "fmt"

// Item 2way
type Item struct {
	Next  *Item
	Prev  *Item
	Value interface{}
}

// List 2way
type List struct {
	FirstItem *Item
	LastItem  *Item
}

// NewList cool

func main() {
	var newList List
	newList.PushFront(1)
	newList.PushFront(2)
	newList.PushBack(4)
	newList.PushFront(3)
	newList.PushBack(5)
	newList.Remove(*newList.Last())
	a := newList.First()
	b := newList.Last()
	x := newList.Len()
	fmt.Printf("%d %d %d", a.Value, b.Value, x)
}

//PushFront in first
func (l *List) PushFront(v int) {
	var newItem Item
	newItem.Value = v
	if l.FirstItem == nil {
		l.FirstItem = &newItem
		l.LastItem = &newItem
	} else {
		l.FirstItem.Prev = &newItem
		newItem.Next = l.FirstItem
		newItem.Prev = nil
		l.FirstItem = &newItem
	}
}

//PushBack back
func (l *List) PushBack(v int) {
	var newItem Item
	newItem.Value = v
	if l.FirstItem == nil {
		l.FirstItem = &newItem
		l.LastItem = &newItem
	} else {
		l.LastItem.Next = &newItem
		newItem.Prev = l.LastItem
		newItem.Next = nil
		l.LastItem = &newItem
	}
}

//Remove delete
func (l *List) Remove(i Item) {
	if i.Prev == nil {
		l.FirstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.LastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
}

//Len asdf
func (l *List) Len() int {

	if l.FirstItem == nil {
		return 0
	}
	count := 1
	next := l.FirstItem.Next
	for next != nil {
		next = next.Next
		count++
	}
	return count

}

//First asdf
func (l *List) First() *Item {

	return l.FirstItem

}

//Last asdf
func (l *List) Last() *Item {

	return l.LastItem

}
