package main

import "fmt"

func main() {
	// age := 16

	// //if else
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age >= 13 {
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a minor.")
	// }

	// var role = "admin"
	// var hasPermission = false
	// if role == "admin" || hasPermission {
	// 	fmt.Println("Access granted.")
	// } else {
	// 	fmt.Println("Access denied.")
	// }

	//we can declare a variable in the if statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	//Go does not support the ternary conditional operator.
	// Instead, conditional logic must be implemented using standard if–else statements
	// to improve readability and maintain simplicity in the language design.
}
