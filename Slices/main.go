package main

import "fmt"

func main() {

	// Basic Slice Declaration

	names := []string{"Ahmad", "ALi", "Hassan"}

	fmt.Println("Names:", names)

	// Append inside the Slice Elements

	names = append(names, "Sara")

	fmt.Println(names)

	myslice1 := []string{}

	fmt.Println(len(myslice1)) // how many elements in a slice
	fmt.Println(cap(myslice1)) // how much slice can shrink or grow
	fmt.Println(myslice1)

	myslice2 := []string{"Batman", "Superman"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// You can also create a slice from an array by slicing the array

	myarray := [5]int{4, 5, 2, 6, 7}
	fmt.Println(myarray)

	myslice3 := myarray[2:4] // start : End  = 2:4 number dena for slicing
	fmt.Println(len(myslice3))
	fmt.Println(cap(myslice3))
	fmt.Println(myslice3)

	// Length and capcity are the built in functions in slicess

	myslice4 := [3]int{1, 2, 3}

	fmt.Println("Length:", len(myslice4))
	fmt.Println("Capacity:", cap(myslice4))

	num := []int{4, 5, 6}

	// Normal loop on slicess

	// Tradiotanl Loop

	for i := 0; i < len(num); i++ {

		fmt.Println("Index ", i, "Value: ", num[i])

	}

	// Another Method is Using Range Keyword ke thourgh looping most common maps strings slices arrays

	num2 := []int{15, 24, 45}

	for i, v := range num2 {

		fmt.Println("Index :", i, "Value:", v)

	}

	for _, v := range num2 { // _ means we dont need index

		fmt.Println("Value:", v)
	}

	// sirf index chaye to v hta do
	for i := range num2 {

		fmt.Println("Index:", i)
	}

	// Multi Dimentional useful for tubular data form in db quries or Json formating

	matrix := [][]int{

		{1, 2, 3}, // Row 0 per ye data para
		{4, 5, 6}, // Row 1 per ye data para
	}

	fmt.Println(matrix[0])
	fmt.Println(matrix[1])

	// Very useful like to append row in db e.g

	matrix = append(matrix, []int{7, 8, 9}) // new row append krdi using multi dimential slices in db
	fmt.Println(matrix)

	// Fetching user data from db users list backend example

	users := []string{"Sara", "Ahmad", "Ali"}

	fmt.Println("Current Users:", users)

	users = append(users, "Aqsa", "Asif") // new users

	fmt.Println("New Users Added:", users)

	for _, user := range users {

		fmt.Println("Names are: ", user)
	}

	// We can Also create a slice using make funtion

	myslice5 := make([]int, 5) //, 10) // 5= lenght or 10 = capacity if capacity is not given its equal to lenght

	fmt.Println(myslice5)
	fmt.Println(len(myslice5))
	fmt.Println(cap(myslice5))

	// How to replace or the current element in slice

	myslice6 := []int{4, 5, 2, 13}

	myslice6[1] = 12

	fmt.Println(myslice6)

	// How to append multiple slices in one slice in golang

	myslice7 := []int{1, 2, 3}
	myslice8 := []int{4, 5, 6}

	myslice9 := append(myslice7, myslice8...)

	fmt.Println(myslice9)

	// This is the Slices which we wanted to study

}
