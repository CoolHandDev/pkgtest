package main

import (
	"fmt"

	"pkgtest/pkg/chicken"
	"pkgtest/pkg/pizza"
)

func main() {
	fmt.Println("hello")
	fmt.Println(chicken.Sound())
	fmt.Println(pizza.Shape())
}
