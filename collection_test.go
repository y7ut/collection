package collection

import (
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	type args[T item] struct {
		d []T
	}
	type testCase[T item] struct {
		name string
		args args[T]
		want []T
	}

	StringCases := []testCase[string]{
		{
			name: "String",
			args: args[string]{
				d: []string{"a", "b", "c"},
			},
			want: []string{
				"a", "b", "c",
			},
		},
	}

	for _, tt := range StringCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.d).Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

type Book struct {
	Name  string
	Price int
}

func TestFilter(t *testing.T) {
	type args[T item] struct {
		d []T
	}
	type testCase[T item] struct {
		name   string
		args   args[T]
		filter func(i Book) bool
		want   []string
	}

	IntCases := []testCase[Book]{
		{
			name: "book",
			args: args[Book]{
				d: []Book{
					{Name: "a", Price: 2},
					{Name: "b", Price: 3},
					{Name: "c", Price: 4},
				},
			},
			filter: func(a Book) bool {
				return a.Price >= 3
			},
			want: []string{
				"b", "c",
			},
		},
	}

	for _, tt := range IntCases {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.args.d).Filter(tt.filter).Value()
			wantt := make([]string, 0)
			for _, v := range result {
				wantt = append(wantt, v.Name)
			}
			if got := wantt; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPtr(t *testing.T) {
	type args[T item] struct {
		d []T
	}
	type testCase[T item] struct {
		name   string
		args   args[T]
		filter func(i *Book) bool
		want   []string
	}

	IntCases := []testCase[*Book]{
		{
			name: "bookPtr",
			args: args[*Book]{
				d: []*Book{
					{Name: "a", Price: 2},
					{Name: "b", Price: 3},
					{Name: "c", Price: 4},
				},
			},
			filter: func(a *Book) bool {
				return a.Price >= 3
			},
			want: []string{
				"b", "c",
			},
		},
	}

	for _, tt := range IntCases {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.args.d).Filter(tt.filter).Value()
			wantt := make([]string, 0)
			for _, v := range result {
				wantt = append(wantt, v.Name)
			}
			if got := wantt; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEachPtr(t *testing.T) {
	type args[T item] struct {
		d []T
	}
	type testCase[T item] struct {
		name string
		args args[T]
		each func(i *Book)
		want []int
	}

	IntCases := []testCase[*Book]{
		{
			name: "bookPtr",
			args: args[*Book]{
				d: []*Book{
					{Name: "a", Price: 2},
					{Name: "b", Price: 3},
					{Name: "c", Price: 4},
				},
			},
			each: func(a *Book) {
				a.Price = a.Price + 1
			},
			want: []int{
				3, 4, 5,
			},
		},
	}

	for _, tt := range IntCases {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.args.d).Each(tt.each).Value()
			wantt := make([]int, 0)
			for _, v := range result {
				wantt = append(wantt, v.Price)
			}
			if got := wantt; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestSort(t *testing.T) {
	type args[T item] struct {
		d []T
	}
	type testCase[T item] struct {
		name string
		args args[T]
		sort func(i, j Book) bool
		want []int
	}

	IntCases := []testCase[Book]{
		{
			name: "book",
			args: args[Book]{
				d: []Book{
					{Name: "a", Price: 6},
					{Name: "b", Price: 11},
					{Name: "c", Price: 9},
					{Name: "d", Price: 1},
				},
			},
			sort: func(i, j Book) bool {
				return i.Price < j.Price
			},
			want: []int{
				1, 6, 9, 11,
			},
		},
		{
			name: "book Desc",
			args: args[Book]{
				d: []Book{
					{Name: "a", Price: 9},
					{Name: "b", Price: 11},
					{Name: "c", Price: 9},
					{Name: "d", Price: 1},
				},
			},
			sort: func(i, j Book) bool {
				return i.Price >= j.Price
			},
			want: []int{
				11, 9, 9, 1,
			},
		},
		{
			name: "book 15",
			args: args[Book]{
				d: []Book{
					{Name: "a", Price: 1},
					{Name: "b", Price: 2},
					{Name: "c", Price: 3},
					{Name: "d", Price: 4},
					{Name: "e", Price: 5},
					{Name: "f", Price: 6},
					{Name: "g", Price: 7},
					{Name: "h", Price: 8},
					{Name: "i", Price: 9},
					{Name: "j", Price: 10},
					{Name: "k", Price: 11},
					{Name: "l", Price: 12},
					{Name: "m", Price: 13},
					{Name: "n", Price: 14},
					{Name: "o", Price: 15},
					{Name: "p", Price: 16},
				},
			},
			sort: func(i, j Book) bool {
				return i.Price > j.Price
			},
			want: []int{
				16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
			},
		},
		// TODO: Add some test cases with 15 books.
	}
	for _, tt := range IntCases {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.args.d).Sort(tt.sort).Value()
			wantt := make([]int, 0)
			for _, v := range result {
				wantt = append(wantt, v.Price)
			}
			if got := wantt; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPeekCollection(t *testing.T) {
	type args[T item] struct {
		d     []T
		index int
	}
	type testCase[T item] struct {
		name string
		args args[T]
		want T
	}

	cases := []testCase[string]{
		{
			name: "string first",
			args: args[string]{
				d: []string{"a", "b", "c"},
				index: 0,
			},
			want: "a",
		},
		{
			name: "string last",
			args: args[string]{
				d: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
				index: 9,
			},
			want: "j",
		},
		{
			name: "string negative",
			args: args[string]{
				d: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
				index: -2,
			},
			want: "o",
		},
		{
			name: "string out of range",
			args: args[string]{
				d: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
				index: 17,
			},
			want: "",
		},
		{
			name: "string out of range",
			args: args[string]{
				d: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
				index: -27,
			},
			want: "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.d).Peek(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error in %s, New() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
