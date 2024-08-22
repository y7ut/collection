# collection

Generic collection

## Each

```go
type Book struct {
    Name  string
    Price int
}
books := []*Book{
    {Name: "a", Price: 2},
    {Name: "b", Price: 3},
    {Name: "c", Price: 4},
}
collection.New(books).Each(func(b *Book){
    b.Price = b.Price + 1
}).Value()

for b := range bbbss {
    fmt.Println(b.Price) // 4，5
}

```

## Map

```go
bookss := []Book{
    {Name: "a", Price: 2},
    {Name: "b", Price: 3},
    {Name: "c", Price: 4},
}
bbbss := collection.New(bookss).Map(func(b Book) Book {
    b.Price = b.Price + 1
    return b
}).Filter(func(i Book) bool {
    return i.Price >= 4
}).Value()

for b := range bbbss {
    fmt.Println(b.Price) // 4，5
}
```

## Filter

```go
books := []*Book{
    {Name: "a", Price: 2},
    {Name: "b", Price: 3},
    {Name: "c", Price: 4},
}
bb := collection.New(books).Each(func(b *Book) {
    b.Price = b.Price + 1
}).Filter(func(i *Book) bool {
    return i.Price >= 4
}).Value()

for b := range bb {
    fmt.Println(b.Name) // 3，4
}

```

## Merge

```go
bb := collection.New([]*Book{
    {Name: "a", Price: 2},
    {Name: "b", Price: 3},
    {Name: "c", Price: 4},
})

bb.Merge(
    collection.New(
        []*Book{
            {Name: "d", Price: 5},
            {Name: "e", Price: 6},
        },
    ),
    collection.New(
        []*Book{
            {Name: "f", Price: 7},
            {Name: "g", Price: 8},
        },
    ),
)


for b := range bb.Value() {
    fmt.Println(b.Price) // 2,3,4,5,6,7,8
}

```

## Sort

```go
books := []*Book{
    {Name: "a", Price: 2},
    {Name: "b", Price: 3},
    {Name: "c", Price: 4},
}

bb := collection.New(books).Sort(func(i, j *Book) bool {
    return i.Price > j.Price
}).Value()

for b := range bb {
    fmt.Println(b.Price) // 4,3,2
}

```

## Peek

```go
books := []*Book{
    {Name: "a", Price: 2},
    {Name: "b", Price: 3},
    {Name: "c", Price: 4},
}

bb := collection.New(books).Sort(func(i, j *Book) bool {
    return i.Price > j.Price
})

fmt.Println(bb.Peek(0)) // 4
fmt.Println(bb.Peekz(-1)) // 2



```
