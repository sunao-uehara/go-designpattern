package main

import (
	"fmt"
)

var productFactories = make(map[string]Product)

func init() {
	Register("p1", NewP1())
	Register("p2", NewP2())
}

type Product interface {
	Name() string
}

func Register(name string, factory Product) {
	if factory == nil {
		panic(fmt.Sprintf("Product factory %s does not exist.", name))
	}
	_, registered := productFactories[name]
	if registered {
		fmt.Printf("Product factory %s already registered. Ignoring.", name)
	}
	productFactories[name] = factory
}
