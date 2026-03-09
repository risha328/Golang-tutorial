package main

import "fmt"

//for => only construct in go for looping
func main() {
	//while loop
	// i := 1

	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//for loop
	// for j := 1; j <= 5; j++ {
	// 	fmt.Println(j)
	// }
	//infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// //classic for loop
	// for i := 0; i <= 3; i++ {
	// 	//break => exit the loop
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	//continue => skip the current iteration
	// 	// if i == 1 {
	// 	// 	continue
	// 	// }
	// 	fmt.Println(i)
	// }

	for i := range 10 {
		fmt.Println(i)
	}

}
