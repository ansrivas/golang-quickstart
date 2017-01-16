package main

import (
	"fmt"

	pack1 "github.com/user/package1"
	pack2 "github.com/user/package2"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(pack1.Average(2, 3))
	fmt.Println(pack2.Add(2, 3))
}
