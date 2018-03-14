package main

import "fmt"

// Item is a thing
type Item struct {
	Name string
	Qty  int
	cfg  string
}

// NewItem creates an Item
func NewItem() *Item {
	item := new(Item)
	// Perform operations on item here.
	return item
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

	item := NewItem()
	item.SetCfg("Example Configuration")

	fmt.Printf("Item Configuration: %s\n", item.GetCfg())
}