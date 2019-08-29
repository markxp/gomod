package main // import "github.com/markxp/gomod"

import (
	"fmt"

	a "github.com/markxp/gomod/sub"
	b "github.com/markxp/gomod/sub2"
)

func main() {
	fmt.Printf("hello world\n")
	x, y := 1, 2
	fmt.Printf("x = %d, y = %d, =>  x - y = %d\n", x, y, a.Sub(x, y))
	fmt.Println(b.Sub("hello world", "world", "gopher"))
}
