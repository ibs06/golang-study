package main

import (
	"container/list"
	"fmt"
)

type Value interface {
	Len() int
}

type Cache struct {
	maxBytes  int64
	nbytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

func main() {
	c := &Cache{
		maxBytes:  0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: nil,
	}

	fmt.Println(c)

	fmt.Println(*c)
}
