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
// 		)
//
//		books.Sort(func(i, j *Book) bool {
//			return i.Price > j.Price
//		})
//		fmt.Println(books.Peek(0))
//	}
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
	for current := c.data.Front(); current != nil; current = current.Next() {
		var currentItem T = current.Value.(T)
		f(currentItem)
	}
	return c
}

// Map applies the given function to each item in the collection and replaces the item with the result.
func (c *Collection[T]) Map(f func(i T) T) *Collection[T] {
	for current := c.data.Front(); current != nil; current = current.Next() {
		var currentItem T = current.Value.(T)
		currentItem = f(currentItem)
		current.Value = currentItem
	}
	return c
}

// Filter filters the collection using the given function and retains only the items for which the function returns true.
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

// Len returns the length of the collection.
func (c *Collection[T]) Len() int {
	return c.data.Len()
}

// Value returns a slice containing all the items in the collection.
func (c *Collection[T]) Value() []T {
	lt := make([]T, 0)
	for current := c.data.Front(); current != nil; current = current.Next() {
		var currentItem T = current.Value.(T)
		lt = append(lt, currentItem)
	}
	return lt
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
