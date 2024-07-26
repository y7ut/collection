// Package collection provide a collection type that can set generics type for any type or any pointer
package collection

import (
	"container/list"
	"reflect"
)

// item alias any for generics
type item interface {
	any
}

// Collection can set generics for any type or any pointer
// you must choose using pointer or not
type Collection[T item] struct {
	data      *list.List
	itemIsPtr bool
}

// New can be used to create a new collection
// you can use collection.New([]string{"a", "b", "c"})
// or collection.New([]int{1, 2, 3})
// or collection.New([]Book{{Name: "a", Price: 2}, {Name: "b", Price: 3}, {Name: "c", Price: 4}})
// or collection.New([]*Book{{Name: "a", Price: 2}, {Name: "b", Price: 3}, {Name: "c", Price: 4}})
func New[T item](d []T) *Collection[T] {

	t := reflect.TypeOf(*new(T))
	var isPtr = t.Kind() == reflect.Ptr

	l := list.New()
	for _, item := range d {
		t := item
		l.PushBack(t)
	}

	return &Collection[T]{data: l, itemIsPtr: isPtr}
}

// Each can be used to iterate over the collection
func (c *Collection[T]) Each(f func(i T)) *Collection[T] {
	for current := c.data.Front(); current != nil; current = current.Next() {
		var currentItem T = current.Value.(T)
		f(currentItem)
	}
	return c
}

// Map can be used to map the collection
func (c *Collection[T]) Map(f func(i T) T) *Collection[T] {
	for current := c.data.Front(); current != nil; current = current.Next() {
		var currentItem T = current.Value.(T)
		currentItem = f(currentItem)
		current.Value = currentItem
	}
	return c
}

// Filter can be used to filter the collection only when the function returns true
func (c *Collection[T]) Filter(f func(i T) bool) *Collection[T] {
	match := list.New()
	for current := c.data.Back(); current != nil; current = current.Prev() {
		var currentItem T = current.Value.(T)
		if f(currentItem) {
			match.PushFront(currentItem)
		}
	}
	c.data = match
	return c

}

// Len can be used to get the length
func (c *Collection[T]) Len() int {
	return c.data.Len()
}

// Value can be used to get the value slice
func (c *Collection[T]) Value() []T {
	lt := make([]T, 0)
	for current := c.data.Front(); current != nil; current = current.Next() {
		var currentItem T = current.Value.(T)
		lt = append(lt, currentItem)
	}
	return lt
}

// Merge can be used to merge two collections
func (c *Collection[T]) Merge(other *Collection[T]) {
	c.data.PushBackList(other.data)
}

// Sort sorts the elements in the Collection using the provided function.
func (c *Collection[T]) Sort(f func(i, j T) bool) *Collection[T] {
	c.data = mergeSort(c.data, f)
	return c
}

// Peek can be used to get the value
func (c *Collection[T]) Peek(i int) T {

	if i < 0 {
		i = c.data.Len() + i
	}

	if i > c.data.Len() || i < 0 {
		return *new(T)
	}
	var current *list.Element

	// if i is greater than half of the length, traverse from the end of the list
	if i > c.data.Len()/2 {
		current = c.data.Back()
		for i = c.data.Len() - 1 - i; i > 0; i-- {
			current = current.Prev()
		}
		return current.Value.(T)
	}

	current = c.data.Front()
	for ; i > 0; i-- {
		current = current.Next()
	}
	return current.Value.(T)
}
