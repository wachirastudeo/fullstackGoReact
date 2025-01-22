package main

import "fmt"

type Persion struct {
	Name string
	Age  int
}

func main() {
	persion1 := Persion{"wachira", 25}
	persion2 := Persion{Name: "pae", Age: 30}
	fmt.Println("Name:", persion1.Name, "Age:", persion1.Age)
	fmt.Println(persion2)
}
