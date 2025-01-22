package main

import (
	"fmt"
)

func main() {

	i, j := 10, 20
	p := &i
	*p = 30
	fmt.Println(*p)
	fmt.Println(p)

	p = &j
	*p = 40
	fmt.Println(*p)
	fmt.Println(p)

}
