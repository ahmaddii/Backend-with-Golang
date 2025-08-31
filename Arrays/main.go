package main

import "fmt"

// Array in is like other programming lang
// store fixed size of elements of same data type

func main() {

	var numbers [5]int // int type ka data

	numbers2 := [4]int{10, 22, 33, 12} // short declaration

	//[5] = {0,1,2,3,4}

	numbers[0] = 32 // har aik index per value lagadi.
	numbers[1] = 15
	numbers[2] = 20
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Arrays", numbers)
	fmt.Println("Arrays2:", numbers2)

	// Array with Inferred Size

	numbers3 := [...]string{"Apple", "Banna", "Grapes"}
	fmt.Println("Inferred Array:", numbers3)

	// How to access a array Element

	fmt.Println(numbers3[1])
	fmt.Println(numbers3[2])

	// how to loop through an array

	for i := 0; i < len(numbers); i++ { // len built in function to count lenght of an array

		fmt.Println(numbers[i]) // print inside elements of array after every loop

	}

	// Multi Dimentional Array ke name ki cheez bhi hoti ha

	matrix := [2][3]int{

		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {

		fmt.Println(matrix[i])

		for j := 0; j < len(matrix[i]); j++ {

			fmt.Println(matrix[i][j])

		}

	}

	// So what to chose array vs slices

	// Array mostly use where we know fixed size or used at low hardwere devices
	// But mostly in backend where data changes dynamically

	// But slices which changes dynamically shrink or grow with whene user increase

	// User Info

	users := [3]string{"ahmad", "ali", "ayeza"}
	fmt.Println("Users:", users)

	// Isliye hum slices user kre ge

	userss := []string{"ahmad", "ali", "ayeza"}
	userss = append(userss, "Sara") // see we can append or grow our data dynammically which we cant handle in arrays
	fmt.Println("Userss", userss)
}
