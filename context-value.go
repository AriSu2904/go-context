package main

import (
	"context"
	"fmt"
)

func main() {

	parent := context.Background()

	childA := context.WithValue(parent, "a", "Value Child A")

	subChildA := context.WithValue(childA, "sub-a", "Value dari Sub Child A")
	fmt.Println(subChildA.Value("a"))
	fmt.Println(subChildA.Value("sub-a"))
}
