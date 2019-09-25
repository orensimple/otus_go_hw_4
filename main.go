package main

import "fmt"

type Item struct {
	Next  *Item
	Prev  *Item
	Value interface{}
}

type List struct {
	FirstItem *Item
	LastItem  *Item
	Lenght    int
}

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
	l.Lenght++
}

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
	l.Lenght++
}

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
	l.Lenght--
}

func (l *List) Len() int {
	return l.Lenght
}

func (l *List) First() *Item {

	return l.FirstItem

}

func (l *List) Last() *Item {

	return l.LastItem

}
