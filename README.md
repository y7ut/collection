# collection

基于泛型的集合操作，可以多补充些

## Each

### 仅作用于基础类型指针类型的数据集

```go
type Book struct {
    Name  string
    Price int
}
books := []*Book{
    &{Name: "a", Price: 2},
    &{Name: "b", Price: 3},
    &{Name: "c", Price: 4},
}
collection.New(books).Each(func(b *Book){
    b.Price = b.Price + 1
}).Value()

```

## Map

### 比较万能

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
    return i.Price >= 3
}).Value()

for _, b := range bbbss {
    fmt.Println(b.Price) // 3，4，5
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

for _, b := range bb {
    fmt.Println(b.Name) // 3，4
}

```
