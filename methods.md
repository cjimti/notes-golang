[Golang Index](http://go.imti.co/)

# Methods

Methods in Go.

- A method is a function with a **receiver** argument
  - `func (recv <type>) fn()`
  - The type a method belongs to is known as the **receiver**
- The type must be delared in the same package as the function.
- Pointer receivers `func (recv *<type>)` should be used when the
function needs to modify the receiver.

## Example Methods

```go
package main

// Item is a thing
type Item struct {
    Name string
    Qty  int
    cfg  string
}

// GetCfg gets the configuration string of the Item
func (i *Item) GetCfg() string {
    return i.cfg
}

// SetCfg sets a configuration string on the Item
func (i *Item) SetCfg(cfg string) *Item {
    i.cfg = cfg
    return i
}

func main() {
}

```

## Constructors

[Allocation with new](https://golang.org/doc/effective_go.html#allocation_new)
allows you to construct a Type.

`item := new(Item)`

You can create constructor functions as follows:

```go
func NewItem() *Item {
	item := new(Item)
	// Perform operations on item here.
    return item
}
```
Construct a new item as follows:

`item := NewItem()`

## Example Code

- [Go Playground](https://play.golang.org/p/gVDJpm9_8Aa)
- [Source: methods.go](examples/methods.go)

## Resources

- [Effective Go: Methods](https://golang.org/doc/effective_go.html#methods)
- [Effective Go: Functions](https://golang.org/doc/effective_go.html#functions)
- [Go by Example: Functions](https://gobyexample.com/functions)
- Best Practices / Patterns
    - [Accept Interfaces and return concrete types](http://idiomaticgo.com/post/best-practice/accept-interfaces-return-structs/)