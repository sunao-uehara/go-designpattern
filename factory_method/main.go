package main

import (
	"fmt"
)

func main() {
	product, ok := productFactories["p1"]
	if !ok {
		panic("product not found")
	}

	fmt.Printf("%v\n", product)
	product.Name()
}