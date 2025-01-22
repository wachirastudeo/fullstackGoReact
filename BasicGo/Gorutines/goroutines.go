package main

import (
	"fmt"
	"time"
)

func Printnumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)

	}
}
func main() {
	go Printnumbers()

	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
		time.Sleep(150 * time.Millisecond)
	}
}
