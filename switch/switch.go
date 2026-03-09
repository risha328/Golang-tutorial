package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// default:
	// 	fmt.Println("Unknown")

	// }

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I'm an int and my value is %d\n", v)
		case string:
			fmt.Printf("I'm a string and my value is %s\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}

	}
	whoAmI(42)
}
