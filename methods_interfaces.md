# Methods and Interfaces

## Methods

- A method is a function with a **receiver** argument
  - `func (recv <type>) fn()`
  - The type a method belongs to is known as the **receiver**
- The type must be delared in the same package as the function.
- Pointer receivers `func (recv *<type>)` should be used when the
function needs to modify the receiver.

### Example Methods

```go
package example

// Item is a thing
type Item struct {
    Name string
    Qty  int
    cfg  string
}

// GetCfg gets the configuration string of the Item
func (i *Item) GetCfg() string {
    return i.Name
}

// SetCfg sets a configuration string on the Item
func (i *Item) SetCfg(cfg string) *Item {
    i.cfg = cfg
    return i
}

```

## Constructors

[Allocation with new](https://golang.org/doc/effective_go.html#allocation_new)
allows you to construct a Type.

`item := new(example.Item)`

You can create constructor functions as follows:

```go
func NewItem() *Item {
	item := new(Item)
	// Perform operations on item here.
    return item
}
```
Construct a new item as follows:

`item := example.NewItem()`

