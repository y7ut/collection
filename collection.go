// Package collection provides a generic collection type that can hold any type or any pointer type.
//
// This package allows the creation of a collection of any type, including custom types and pointers to custom types.
//
// Example usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/yourusername/collection"
//	)
//
//	type Book struct {
//	    Name  string
//	    Price int
//	}
//
//	func main() {
//	    c := collection.New([]string{"a", "b", "c"})
//	    c.Each(func(i string) {
//	        fmt.Println(i)
//	    })
//
//	    books := collection.New([]Book{{Name: "a", Price: 2}, {Name: "b", Price: 3}, {Name: "c", Price: 4}})
//	    books.Each(func(b Book) {
//	        fmt.Printf("Book: %s, Price: %d\n", b.Name, b.Price)
//	    })
//		books.Merge(
//			collection.New(
//				[]*Book{
//					{Name: "d", Price: 5},
//					{Name: "e", Price: 6},
//				},
//			),
//			collection.New(
//				[]*Book{
//					{Name: "f", Price: 7},
//					{Name: "g", Price: 8},
//				},
//			),
//		)
//
//		books.Sort(func(i, j *Book) bool {
//			return i.Price > j.Price
//		})
//		fmt.Println(books.Peek(0))
//	}
package collection

import (
	"container/list"
	"iter"
	"reflect"
)

// item alias any for generics
type item interface {
	any
}

// Collection can set generics for any type or any pointer
// you must choose using pointer or not
type Collection[T item] struct {
	// data is a list of any type or any pointer type
	data *list.List

	// itemIsPtr is true if T is a pointer
	itemIsPtr bool
}

// New creates a new collection from a slice of any type or pointer type.
//
// Example usage:
//
//	collection.New([]string{"a", "b", "c"})
//	collection.New([]int{1, 2, 3})
//	collection.New([]Book{{Name: "a", Price: 2}, {Name: "b", Price: 3}, {Name: "c", Price: 4}})
//	collection.New([]*Book{{Name: "a", Price: 2}, {Name: "b", Price: 3}, {Name: "c", Price: 4}})
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

// Each iterates over the collection and applies the given function to each item.
func (c *Collection[T]) Each(f func(i T)) *Collection[T] {
	c.Value()(func(i T) bool {
		f(i)
		return true
	})
	return c
}

// Map applies the given function to each item in the collection and replaces the item with the result.
func (c *Collection[T]) Map(f func(i T) T) *Collection[T] {
	c.all()(func(item *list.Element, value T) bool {
		value = f(value)
		item.Value = value
		return true
	})
	return c
}

// Filter filters the collection using the given function and retains only the items for which the function returns true.
func (c *Collection[T]) Filter(f func(i T) bool) *Collection[T] {
	match := list.New()
	c.Value()(func(value T) bool {
		if f(value) {
			match.PushBack(value)
		}
		return true
	})
	c.data = match
	return c
}

// Len returns the length of the collection.
func (c *Collection[T]) Len() int {
	return c.data.Len()
}

// Merge can be used to merge two collections
func (c *Collection[T]) Merge(other ...*Collection[T]) {
	for _, collection := range other {
		c.data.PushBackList(collection.data)
	}
}

// Sort sorts the elements in the Collection using the provided function.
func (c *Collection[T]) Sort(f func(i, j T) bool) *Collection[T] {
	c.data = mergeSort(c.data, f)
	return c
}

// Peek returns the item at the specified index. If the index is negative, it returns the item from the end of the collection.
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

// Clone returns a new collection that is a copy of the current collection.
func (c *Collection[T]) Clone() *Collection[T] {
	clone := list.New()
	c.all()(func(item *list.Element, value T) bool {
		clone.PushBack(value)
		return true
	})
	return &Collection[T]{data: clone, itemIsPtr: c.itemIsPtr}
}

// Reverse reverses the order of the collection.
func (c *Collection[T]) Reverse() *Collection[T] {
	reverse := list.New()
	c.Value()(func(value T) bool {
		reverse.PushFront(value)
		return true
	})
	c.data = reverse
	return c
}

// Value returns a interator that yields all the items value in the collection.
func (c *Collection[T]) Value() iter.Seq[T] {
	return func(yield func(T) bool) {
		for current := c.data.Front(); current != nil; current = current.Next() {
			if !yield(current.Value.(T)) {
				return
			}
		}
	}
}

// all returns a interator that yields all the items with their values in the collection.
func (c *Collection[T]) all() iter.Seq2[*list.Element, T] {
	return func(yield func(*list.Element, T) bool) {
		for current := c.data.Front(); current != nil; current = current.Next() {
			if !yield(current, current.Value.(T)) {
				return
			}
		}
	}
}

// All returns a interator that yields index their values in the collection.
func (c *Collection[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		index := 0
		for value := range c.Value() {
			if !yield(index, value) {
				return
			}
			index++
		}
	}
}
