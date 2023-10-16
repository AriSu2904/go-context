package main

import (
	"context"
	"fmt"
)

func main() {

	bg := context.Background()
	fmt.Println(bg)

	todo := context.TODO()
	fmt.Println(todo)
}
