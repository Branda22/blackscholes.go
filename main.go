package main

import (
	"fmt"

	"github.com/branda22/blackscholes/blackscholes"
)

func main() {
	bs := blackscholes.NewBlackScholes(50.0, 45.0, .1, .3, 60)
	fmt.Println(bs)
	fmt.Println("d1", bs.D1())
	fmt.Println("d2", bs.D2())
	fmt.Println("delta", bs.Delta())
}
