package main

import "fmt"

// const age = 30
// const name = "John Doe"
// const pi = 3.14
// const isGoFun = true

// func main() {
// 	fmt.Println("Name:", name)
// 	fmt.Println("Age:", age)
// 	fmt.Println("Pi:", pi)
// 	fmt.Println("Is Go Fun?", isGoFun)
// }

func main() {
	const (
		port = 8080
		host = "localhost"
	)
	fmt.Printf("Server running at http://%s:%d\n", host, port)
}
