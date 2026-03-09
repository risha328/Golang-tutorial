package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 3.14
	var c string = "Hello, Go!"
	var d bool = true
	var e *int = &a
	var f []int = []int{1, 2, 3, 4, 5}
	var g map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println("int", a)
	fmt.Println("float64", b)
	fmt.Println("string", c)

	fmt.Println("bool", d)
	fmt.Println("pointer", e)
	fmt.Println("slice", f)
	fmt.Println("map", g)
}
