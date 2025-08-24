package main

import "fmt"

func main() {

	// Loops in Golang
	// for loop only but we use in different ways

	// Classic Loop

	// Specific number ya condition us tak chlta

	for i := 1; i <= 5; i++ {

		fmt.Println("Number :", i)

	}

	// while like loop

	count := 1
	for count <= 5 {

		fmt.Println(count)
		count++
	}

	// Infinite loop

	//	for {
	//		fmt.Println("Press c to avoid infinite loop")
	//	}

	// Loops for specific ranges used in slices maps arrays strings

	//	nums := []int{34, 52, 32, 31}

	//	for index, value = range nums {

	//	print(index, value)

	//}
}
