package main

import "fmt"

func main() {
	// var nums [5]int

	// //array length
	// fmt.Println("Length of array:", len(nums))
	// //array indexing
	// nums[0] = 10
	// nums[1] = 20
	// nums[2] = 30
	// nums[3] = 40
	// nums[4] = 50
	// fmt.Println("Array:", nums)
	// fmt.Println("First element:", nums[0])
	// fmt.Println("Last element:", nums[len(nums)-1])

	// var vals [3]bool
	// vals[2] = true
	// fmt.Println("Array of booleans:", vals)

	// var names [3]string
	// names[0] = "Alice"
	// names[1] = "Bob"

	// fmt.Println("Array of strings:", names)

	//2D array
	nums := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2D Array:", nums)

	//fixed size array
	//memory optimization
	//contant time access
	//Go finds the element using this formula:

	// Address = base_address + index * element_size

	// Example:

	// arr[2] = base + (2 × size_of_int)
}
