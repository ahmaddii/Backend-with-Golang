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

}
